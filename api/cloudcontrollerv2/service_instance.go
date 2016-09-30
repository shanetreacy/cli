package cloudcontrollerv2

type ServiceInstance struct {
	GUID string
	Name string
}

func (client *CloudControllerClient) GetServiceInstances([]Query) ([]ServiceInstance, Warnings, error) {
	return nil, nil, nil
}
