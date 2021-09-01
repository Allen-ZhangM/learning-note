package main

import "fmt"

func main() {
	smsOTP := &sms{}
	o := otp{
		iOtp: smsOTP,
	}
	_ = o.genAndSendOTP(4)

	fmt.Println("")
	emailOTP := &email{}
	o = otp{
		iOtp: emailOTP,
	}
	_ = o.genAndSendOTP(4)

}
