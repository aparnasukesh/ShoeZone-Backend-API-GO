package util

import (
	"crypto/rand"
	"crypto/tls"
	"fmt"
	"os"

	gomail "gopkg.in/mail.v2"
)

func Otpgeneration(emails string) string {

	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", os.Getenv("EMAIL"))

	// Set E-Mail receivers
	m.SetHeader("To", emails)

	// Set E-Mail subject
	m.SetHeader("Subject", "OTP to verify your Gmail")

	//otp generation
	onetimepassword, err := GenCaptchaCode()
	if err != nil {
		panic(err)
	}

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", onetimepassword+" is your OTP to register to ShoeZone. Thank you registering to our site. Dont't give this code to anyone")

	// Settings for SMTP server
	d := gomail.NewDialer("smtp.gmail.com", 587, os.Getenv("EMAIL"), os.Getenv("PASSWORD"))

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	} else {
		fmt.Println("OTP has been sent successfully")
	}
	return onetimepassword
}

func GenCaptchaCode() (string, error) {

	codes := make([]byte, 6)
	if _, err := rand.Read(codes); err != nil {
		return "", err
	}

	for i := 0; i < 6; i++ {
		codes[i] = uint8(48 + (codes[i] % 10))
	}

	return string(codes), nil
}
