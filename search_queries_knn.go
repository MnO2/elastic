package elastic

type KnnQuery struct {
	vector    []float64
	K  		  int
	queryName string
}


func NewKnnQuery(K int) *KnnQuery {
	return &KnnQuery{
		vector: make([]float64, 0),
		K: K,
	}
}

func (q *KnnQuery) Vector(v ...float64) *KnnQuery {
	q.vector = append(q.vector, v...)
	return q
}

func (q *KnnQuery) QueryName(queryName string) *KnnQuery {
	q.queryName = queryName
	return q
}


func (q *KnnQuery) Source() (interface{}, error) {
	source := make(map[string]interface{})
	query := make(map[string]interface{})
	param := make(map[string]interface{})
	source["knn"] = query
	query[q.queryName] = param
	param["vector"] = q.vector
	param["k"] = q.K
	return source, nil
}
