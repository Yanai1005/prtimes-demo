package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Input struct {
	IndustryId       string   `json:"industry_id"`
	ImportantAspects []string `json:"important_aspects"`
}

func main() {
	http.HandleFunc("/industry_ids", getIndustryIDs)
	http.HandleFunc("/check", checkApiRequest)
	http.ListenAndServe(":8080", nil)
}

// 指定可能な業種とそのid
func getIndustryIDs(w http.ResponseWriter, r *http.Request) {
	industryMap := map[string]string{
		"0": "その他",
		"1": "水産・農林業",
		"2": "鉱業",
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(industryMap); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Apiに対するリクエストについて、とりあえず構造とアクセス方法が正しければOK, ダメなら原因を表示する
func checkApiRequest(w http.ResponseWriter, r *http.Request) {
	var requiredParams []string = []string{"industry_id", "important_aspects"}
	if r.Method != http.MethodPost {
		fmt.Fprintln(w, "POSTメソッドでアクセスしてください")
		return
	}

	for _, v := range requiredParams {
		if r.FormValue(v) == "" {
			fmt.Fprintf(w, "%sは必須項目です", v)
			return
		}
	}

	var input Input
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		fmt.Fprintf(w, "%s", err)
	}

	fmt.Fprintln(w, "OK")
}
