package main

func mergeMap(old, new map[string]interface{}) map[string]interface{} {
	for k, v := range old {
		new[k] = v
	}
	return new
}
