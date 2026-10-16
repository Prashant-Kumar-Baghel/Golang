package main

import "fmt"

//1st way (lets say if we want to add one more payment gateway then we have to change the code in makePayment method which is not a good practice. so we can use interface to achieve this.)
//In 1st way we are creating payment gateway as starting  we do add razaorpay payment gateway and then we do add stripe payment gateway then we have to change the code in makePayment method which is not a good practice. so we can use interface to achieve this. as if we do testing then we can  create mock payment but if gateway is hardcoded then we can't create mock payment gateway for testing. so we can use interface to achieve this.
// type razorpay struct {
// }

// func (r razorpay) pay(amount int) {
// 	//logic to make payment using razorpay payment gateway
// 	fmt.Println("make payment using razorpay", amount)
// }

// type payment struct {
// }

// func (p payment) makePayment(amount int) {
// 	//Payment using razorpay payment gateway
// 	// razorPaymentGw := razorpay{}
// 	// razorPaymentGw.pay(amount)

// 	//Payment using stripe payment gateway
// 	stripePaymentGw := stripe{}
// 	stripePaymentGw.pay(amount)
// }

// type stripe struct {
// }

// func (s stripe) pay(amount int) {
// 	//logic to make payment using stripe payment gateway .
// 	fmt.Println("make payment using stripe", amount)
// }

// func main() {

// 	newPayment := payment{}
// 	newPayment.makePayment(10)

// }

//2nd way ( lets say if we want to add one more payment gateway then we have to change the code in payment struct and makePayment method which is not a good practice. so we can use interface to achieve this. as if we do testing then we can  create mock payment but if gateway is hardcoded then we can't create mock payment gateway for testing. so we can use interface to achieve this.)
//In 2nd way we using reciever and fields in payment struct .
// type razorpay struct {
// }

// func (r razorpay) pay(amount int) {
// 	//logic to make payment using razorpay payment gateway
// 	fmt.Println("make payment using razorpay", amount)
// }

// type payment struct {

// 	// gateway razorpay
// 	gateway stripe
// }

// func (p payment) makePayment(amount int) {
// 	//Payment using razorpay payment gateway
// 	// p.gateway.pay(amount)

// 	//Payment using stripe payment gateway
// 	p.gateway.pay(amount)

// }

// type stripe struct {
// }

// func (s stripe) pay(amount int) {
// 	//logic to make payment using stripe payment gateway .
// 	fmt.Println("make payment using stripe", amount)
// }

// func main() {
// 	stripePaymentGw := stripe{}
// 	newPayment := payment{
// 		// gateway: razorPaymentGw,
// 		gateway: stripePaymentGw,
// 	}
// 	newPayment.makePayment(10)

// }

// 3rd way ( Interface)
//In interface we can use "er" suffix to name the interface. for example paymenter, payabler, reader, writer etc. but it is not mandatory to use "er" suffix to name the interface. we can use any name for the interface. but it is a good practice to use "er" suffix to name the interface because it gives us an idea that this is an interface and it has some methods which can be implemented by any struct.
//Below interface code is more flexible and we can add any payment gateway without changing the code in payment struct and makePayment method. we just need to implement the pay method of paymenter interface in the new payment gateway struct and then we can use it in payment struct. so this is a good practice to use interface to achieve this.
//By using interfaces we are not voilating open-close principle. open-close principle mean classes or method should open for extention not for modification.
type paymenter interface { //interface is a collection of methods which can be implemented by any struct. so we can use interface to achieve polymorphism in golang.
	//we can only integrate this interface with the struct which has pay method and refund method. so we can use any payment gateway which implements both methods of paymenter interface. if we try to integrate with that struct which doesnot has any of method then we get error
	pay(amount int)
	refund(amount int, account string)
}

type razorpay struct {
}

func (r razorpay) pay(amount int) { //In other languages we do explicitly to use interface on class by using "implement" keyword. but here if any struct has same methods as interface then it happens implecitly for that particular struct.
	//logic to make payment using razorpay payment gateway
	fmt.Println("make payment using razorpay", amount)
}

func (r razorpay) refund(amount int, account string) {
	//logic to refund payment using razorpay payment gateway
	fmt.Println("refund payment using razorpay", amount, account)
}

type payment struct {

	// gateway razorpay
	gateway paymenter //Way to integrate interface with struct .paymenter is an interface which can be implemented by any struct which has pay method and refund method. so we can use any payment gateway which implements both methods of paymenter interface.
}

func (p payment) makePayment(amount int) { // we don't need to integrate interface in makePayment method because we have already integrated interface in payment struct. so we can use any payment gateway which implements pay method and refund method of paymenter interface.
	//Payment using razorpay payment gateway
	// p.gateway.pay(amount)

	//Payment using razorpay payment gateway
	p.gateway.pay(amount)
	p.gateway.refund(14, "jdjndkd92009")

}

type stripe struct {
}

func (s stripe) pay(amount int) {
	//logic to make payment using stripe payment gateway .
	fmt.Println("make payment using stripe", amount)
}

func main() {
	// stripePaymentGw := stripe{}
	razorPaymentGw := razorpay{}
	newPayment := payment{
		gateway: razorPaymentGw,
		// gateway: stripePaymentGw,
	}
	newPayment.makePayment(10)

}
