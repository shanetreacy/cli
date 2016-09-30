package cloudcontrollerv2

type ServiceBinding struct {
	GUID string
}

func (client *CloudControllerClient) GetServiceBindings([]Query) ([]ServiceBinding, Warnings, error) {
	return nil, nil, nil
}

func (client *CloudControllerClient) DeleteServiceBinding(serviceBindingGUID string) (Warnings, error) {
	return nil, nil
}
