package cli

import "github.com/yourusername/yourproject/application"

func Run(service application.ProductService, action string, productId string, prdocutName string, price float64) (string, error) {
	var result = ""
	switch action {
	case "create":
		product, err := service.Create(prdocutName, price)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product ID %s with the name %s has been created with the price %f and status $s", 
			product.GetId(), product.GetName(), product.GetPrice(), product.GetStatus())
	case "enable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		res, err := service.Enable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product %s has been enabled", res.GetName())
	case "disable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		res, err := service.Disable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product %s has been disabled", res.GetName())
	default:
		res, err := service.Get(productId)
		if err != nil {	
			return result, err	
		}
		result = fmt.Sprintf("Product ID: %s \nName: %s \nPrice: %f \nStatus: $s",
			res.GetId(), res.GetName(), res.GetPrice(), res.GetStatus())
	}
	return result, nil
}