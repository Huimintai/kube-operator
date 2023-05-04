package util

type DataRef struct {
	Name      string `json:"name"`
	NameSpace string `json:"namespace"`
}

type ConfigMapRef = DataRef

type SecretRef = DataRef

type PodRef = DataRef

type JobRef = DataRef

func (dr *DataRef) IsEmpty() bool {
	if dr == nil || len(dr.Name) == 0 {
		return true
	}
	return false
}
