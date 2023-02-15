package main

type Person struct {
	Name string
	Age  int
}

func (p *Person) SayHi1() {
	p.Name = "leon1"

}

func (p Person) SayHi2() {
	p.Name = "leon2"
}

func main() {

	//nums := []int{3, 2, 1, 5, 6, 4}
	//sort.FindLargestK(nums, 2)
	// arr := []int{4, 6, 2, 1, 9, 8, 3}
	// sort.HeapSort(arr)
	// fmt.Println(arr)
	//p1 := &Person{Name: "test", Age: 10}
	//fmt.Println("name1 : " + p1.Name)
	//p1.SayHi1()
	//fmt.Println("name2 : " + p1.Name)
	//
	//p2 := Person{Name: "test1", Age: 11}
	//fmt.Println("name3: " + p2.Name)
	//p2.SayHi2()
	//fmt.Println("name4 : " + p2.Name)
	//type resp struct {
	//	Errno int
	//}
	//
	//var r resp
	//b := "null"
	//fmt.Printf("%v\n", r)
	//if r != nil {
	//	fmt.Println("before bResp is not nil")
	//	fmt.Printf("before bResp address :%p\n", &r)
	//}
	//err := json.Unmarshal([]byte(b), &r)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//if r == nil {
	//	fmt.Println("after bResp is nil")
	//	fmt.Printf("after bResp address :%p\n", &r)
	//	if &r == nil {
	//		fmt.Println("after &bResp is nil")
	//	}
	//}
	// fmt.Println(r.Errno)
	//token := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiIsInppcCI6IkRFRiJ9.eNqcWFtv2zoM_iuDn4OD3m9PR5aZRKtt-dhS0mwohLTNNgNJEyQNsGHYfz-Sr5Isp-mebJCfqE8USZP-7e32T96dt8z3P_b5bvH670v-kn9frp_my3-e1ytv4D3td_nrYrfz7n57y8X3-fMvvltsyYt3d3p9eXt-eT7wJHAzf80XEvP1dHB-c3p98zjwNtv1t3y5kNbxertZb-dvC7VSmnydrwrxj_V6s9h-CvO9FOY79LLKX_Pdm0Sut97dt_lytxh4i9U8Xx6m-Lzfva1XeDnPV4qCNwUfozAUCaQRyTJCY-_xj4Q1gJSGICKIuYAoCekMIJNmWimmUYLimSXLGIaYQWqIk5R-BszM9TxDI0hoSDCxLKdSnFroIYlRjAkKa-kImAjpiMSWgHJWSxLJxsTQhMmDZiJBbGyghjQdUZagLJvSNGhUnIkMGKapJIQYIIwpjzswtbPljcKoJdNg7XEwjYdk5IKQwC0V5hLFsVdpraz4N54NIAQGOqJz1V1Ie8dZD4ejDteH0j2nMxJcPpHTJ9VTcUshIAwj7Qotcy17y2Dvvj5ieMyTQN6_67wQIz8EmUnNgsmZllNuN7aMq9QQEQ8ZKVFORu0Kn4f3VO6AVCgXTB8A83ZZ6Rk7DytjttgEazFXHtGhaE7i0FnW4EHlca9FgYKAUc1zmW6ncozB35JpMIt5V1rR7ip0IyZhpampGTRsoQ60iDjEFROHxrDTz0WW7Dgg8cipS2UQlVnQw86BsIi6EPqxiwhEmJEJYm7HuI3YRHoKiZl4WO4xoikxg0M3Q6dxN667Tn73VAoCmXTtFGRJigAiRMI-F9Wh6ziD4ld9way7LaUCj5F5sZWcoQfymfoOQ86Iq1QpYCCJy5472y2Ljjiz6PQjelPdjfF5IBXHYPrNWbTFBIVEr8z9_A9BOwdxghmJyn5l9jfSItjeAwt_ppqLnjWW1lravpol6B1IlSkfRpVcmgsxI73Hmuu7xD-CPkxWv0eMTUZVu2t41pJpMIFRzDNwa8LQadrV7tWqzmeE2wbVjhhC66AWJiCZ6isO7VAU1rJmzt7FuYtXBumEYKiSw0yXIY-DrCvqxZZVtahQZlR3O7WieWikda_abTxaSPtKrDZujOKRfEczOTuwCNiYNoCmmZcDSBtYjVum5IvGyRKLZExjEDGPfMtnult9iGFIrPml5GqPOwXbFP7jIJ8oIYLRezAmGjlmyGhOUpgQmIqMRxEyL9YEqI-cdoGm4xt4V9Jv-ZCxpoWvb6t2uTB9rsL9WKzb7iEWbf7pzaU9c8oLUJdeO1uaTdr2vC73IhsDaEpOxFi2AiJA2dinZWBIWZvtVl_S9e6xx_4retZVBRRzZToTPDU6lzL4BG4OUMRKaabe8tg0PHbw-mh8mBztMuJL99wHstVzm29GSGWpdIasSCnFID9RclYiZnxZiL4RsB4LrEHRHHaLknA40ItOMpO5LWLKyHDWm32iHDPNYhbI7Ykq6YLEEyqTQN_MPmskv8lkYnblRKVELPmogmcEp1sTkt5Faku3pqrOPctUJ9-jq_4F2WmDJrLpb_sv2W93ZClEdAIdcXFmijgbiwwlKjQwT73Hgff2a6P-qFVlP5-_eXenV1fX55c3Fxcn6t_aTmpvbz_5Z77UL35uKv3VydnFyZ__AQAA__8.kFrgbwt3PwodyUInUDhPqi6JxcvpVwZ045L_sE_xVzuT9e_5PkuaTFH7thYDM_UGrBmHBDOKIu20ow2JmapN4Q"
	//
	//type Business struct {
	//	LegacyUserID    int64    `json:"legacyUserId"`
	//	Email           string   `json:"email"`
	//	Name            string   `json:"name"`
	//	Profile         string   `json:"profile"`
	//	Companies       []int64  `json:"companies"`
	//	IsAdministrator bool     `json:"is_admin"`
	//	CustomClaims    []string `json:"custom_claims"`
	//}
	//
	//type MyClaims struct {
	//	Type     string   `json:"type"`
	//	Business Business `json:"business"`
	//	Claims   []string `json:"claims"`
	//	jwt.StandardClaims
	//}
	//
	//// 解压 拼装
	//tokens := strings.Split(token, ".")
	//// 1. 验证token是否合法
	////key := []byte("xabaduba")
	//key, _ := base64.URLEncoding.DecodeString("xabaduba")
	//mac := hmac.New(crypto.SHA512.New, []byte(key))
	//mac.Write([]byte(tokens[0] + "." + tokens[1]))
	//if strings.TrimRight(base64.URLEncoding.EncodeToString(mac.Sum(nil)), "=") == tokens[2] {
	//	fmt.Println("verified token")
	//}
	//
	//b, _ := base64.URLEncoding.DecodeString(tokens[1])
	//buf := bytes.NewBuffer(b)
	//r, _ := zlib.NewReader(buf)
	//body, _ := ioutil.ReadAll(r)
	//var bus MyClaims
	//if err := json.Unmarshal(body, &bus); err != nil {
	//	panic(err)
	//}
	//pp.Println(bus)
	////进行base64编码
	//res := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	//fmt.Println(res)
	//s, err := ioutil.ReadAll(i)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(s)
	//t, err := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
	//	return []byte("abgajbkga-12slidbvhsjaf-127235987123-asdkvbjawjkhg-1974nc1uht-!@#$%"), nil
	//})
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(t)
	//tem, _ := t.Claims.(MyClaims)
	//fmt.Println(tem)
}
