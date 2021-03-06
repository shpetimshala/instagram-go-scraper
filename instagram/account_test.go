package instagram

import "testing"

func Test_getFromAccountPage(t *testing.T) {
	username := "vorkytaka"
	fullname := "Konstantin"
	id := "248188406"
	biography := ""
	verified := false

	jsonBody := []byte("{\"user\":{\"biography\":null,\"blocked_by_viewer\":false,\"country_block\":false,\"external_url\":null,\"external_url_linkshimmed\":null,\"followed_by\":{\"count\":116},\"followed_by_viewer\":false,\"follows\":{\"count\":67},\"follows_viewer\":false,\"full_name\":\"Konstantin\",\"has_blocked_viewer\":false,\"has_requested_viewer\":false,\"id\":\"248188406\",\"is_private\":false,\"is_verified\":false,\"profile_pic_url\":\"https://scontent-arn2-1.cdninstagram.com/t51.2885-19/s150x150/17125816_320904131658190_1521093063361953792_a.jpg\",\"profile_pic_url_hd\":\"https://scontent-arn2-1.cdninstagram.com/t51.2885-19/s320x320/17125816_320904131658190_1521093063361953792_a.jpg\",\"requested_by_viewer\":false,\"username\":\"vorkytaka\",\"connected_fb_page\":null,\"media\":{\"nodes\":[{\"__typename\":\"GraphImage\",\"id\":\"1396549367405712353\",\"comments_disabled\":false,\"dimensions\":{\"height\":1180,\"width\":1080},\"gating_info\":null,\"media_preview\":\"ACcqjNj6Gk+wE960FI7kCpkboRzmmIzorcRem7ucZx9KckaMGjuEAyeGHUe+f6dK03hVs5GM45FZl8zlSqg4xj9fz9qT8hrzMFk2kjrzRQ8bL97iigDauD8uB1Jx+tXIZMD6HH6VWnjyNw4xzis3zWU4zTJ2OqVwaw5Jj85/3v60tvcMxx680r246knHcetOxPNr6GMzs/3jmirDxDmikaXLXmsqnzTg9Ao5PuT9PTOapMdxyPpQ3U/U00UgL9rOsZCnoePz/rmpWuVJxkVmUtO5LSLE7AjAoqsaKW41of/Z\",\"owner\":{\"id\":\"248188406\"},\"thumbnail_src\":\"https://scontent-arn2-1.cdninstagram.com/t51.2885-15/s640x640/sh0.08/e35/c0.36.770.770/15258945_207881816333197_1897681735815004160_n.jpg\",\"is_video\":false,\"code\":\"BNhikAbg5Ph\",\"date\":1480701679,\"display_src\":\"https://scontent-arn2-1.cdninstagram.com/t51.2885-15/e35/15258945_207881816333197_1897681735815004160_n.jpg\",\"caption\":\"Тот момент когда бычишь\",\"comments\":{\"count\":0},\"likes\":{\"count\":91}},{\"__typename\":\"GraphImage\",\"id\":\"1392739918043798279\",\"comments_disabled\":false,\"dimensions\":{\"height\":1080,\"width\":1080},\"gating_info\":null,\"media_preview\":\"ACoqyJppfNdQ7DDMPvH1PvSPJcR4y78+jn/GrDxfvnP+238zWv8AY4xFlfnyMtnHpn/IHNDfLbQFZmIs8wjY73/hx8x9T70ts1zKcqZHVeuGP+NWUhHlsCP7v86tW8v2dCmDg56Yzk/WradtFdkKavZmRPdSM/yM4A/2j/jXQ2+TEhJJJVc/kKy5U3ksQATjp7DFbkCYjT/dX+QpSVkgUlJuxiSSxCRwWOQzZ+U+p96T7TGON5x6bT369+/eqkigzSf77f8AoRq4mnMTkjA9KbfKlzDUU3oOjni2sdxx8v8ACfX60hnibox/75P+NE8UaoyKMP8ALkDnjP8An3qG3szJkHO7sBj8z7Uud25tkHs43s9/UkNxEB94/wDfJ/xrftyDEhHIKr29hXKPbkZU9R1rqoHPlrgADaOPwFKd2kKKim7HJyybJn/32/ma2rdLq7Gf9TEec45PuO/48CrS28RfOxc5znaM59elaCmplK9tC0Zl5bxWsHyDksuWPU9ep/pWWJtvKnBrpZEVxhgGHuM/zpGgj/ur+Qqoy0s1czlG7vc5dmZ8kc+prcgOI1H+yP5CpHjQcbRj6CgAAYFEpcxUY2P/2Q==\",\"owner\":{\"id\":\"248188406\"},\"thumbnail_src\":\"https://scontent-arn2-1.cdninstagram.com/t51.2885-15/s640x640/sh0.08/e35/15275607_252605521823760_3549615278427996160_n.jpg\",\"is_video\":false,\"code\":\"BNUAZOPgdsH\",\"date\":1480247557,\"display_src\":\"https://scontent-arn2-1.cdninstagram.com/t51.2885-15/e35/15275607_252605521823760_3549615278427996160_n.jpg\",\"caption\":\"Тот момент, когда «да отъ*бись ты со своими фотками»\",\"comments\":{\"count\":0},\"likes\":{\"count\":43}},{\"__typename\":\"GraphImage\",\"id\":\"1362935275164261893\",\"comments_disabled\":false,\"dimensions\":{\"height\":1068,\"width\":1080},\"gating_info\":null,\"media_preview\":\"ACop6Q9KzMVdmbAqiKpEjrgbUHzBSM9fftxWQpJPrWrLGZIipGRkEVUto8HIUnaD9aaqOOhLpRmuZmzbk+WueuKmyKzjKoHzZHseDTPtEf8AeFTdMe2hXWCUHluPTrz+NT21oUbzZDk9hng+5/wqeNQeW4H609n3D0FYxv8AFLRDsh2/dUPmxx9WUfiKqXNz9nbZgnI4I9e3FYbSFutbpXB6HQXSRzEDPzdOD2P+eKof2e/ZhT7VGYBmbgAcc5/wq95yf3hXM5tOyG0jM+25brgd/wDP+TTDctIcBsD6Dgd8D2qi33j/AJ7Cnnt9P61q0S2auqYcKQDvGACO4x/k/pWbHC7DHAPr/n09q1G/1Q/65rTU6f59Kpu1kityk0pUqozkLhgOm7mqxJqWL7g/36R/vH6miyUmjM//2Q==\",\"owner\":{\"id\":\"248188406\"},\"thumbnail_src\":\"https://scontent-arn2-1.cdninstagram.com/t51.2885-15/s640x640/sh0.08/e35/c4.0.773.773/14736202_1793111674277119_6228540617513762816_n.jpg\",\"is_video\":false,\"code\":\"BLqHmlgA94F\",\"date\":1476694567,\"display_src\":\"https://scontent-arn2-1.cdninstagram.com/t51.2885-15/e35/14736202_1793111674277119_6228540617513762816_n.jpg\",\"comments\":{\"count\":11},\"likes\":{\"count\":56}},{\"__typename\":\"GraphImage\",\"id\":\"1326001438517766112\",\"comments_disabled\":false,\"dimensions\":{\"height\":955,\"width\":1080},\"gating_info\":null,\"media_preview\":\"AColvioWuArhecZwT2FSk4H6VGsG19wz3zk8fy5pAWAQeRyKWmIuxQD2pwYGpulvoVbsLSUtJinvsIy72VkdQDjjP15oGpKOG6+3NU7ycTFQAflP+elOgVIXLSKWZTx6D368mq6aku99C81wSRuG3IyM+9TJID3rKvLtJipAIZc9ccg0RXCng4rmnF3ubR2sbgen1lLcBTjPGM+9T/b09DVU09bkyMXrSlyetFFdBAxuevNM2CiigBw46U7JoooA/9k=\",\"owner\":{\"id\":\"248188406\"},\"thumbnail_src\":\"https://scontent-arn2-1.cdninstagram.com/t51.2885-15/s640x640/sh0.08/e35/c62.0.955.955/13687113_1035654386530481_70963703_n.jpg\",\"is_video\":false,\"code\":\"BJm50hqgifg\",\"date\":1472291711,\"display_src\":\"https://scontent-arn2-1.cdninstagram.com/t51.2885-15/e35/13687113_1035654386530481_70963703_n.jpg\",\"caption\":\"Последняя надежда нашей сборной\",\"comments\":{\"count\":3},\"likes\":{\"count\":61}},{\"__typename\":\"GraphImage\",\"id\":\"1323175264271763418\",\"comments_disabled\":false,\"dimensions\":{\"height\":1052,\"width\":1080},\"gating_info\":null,\"media_preview\":\"ACooy4FBf8KsSRblIHUciq9qMtn2rRjGDk9KzNuhVg+UZxkVZZA5zjAHap0gXbinMnBp3JsUJUGKg5qy52gk1SzQgZHCzK2R1rQFwgIDHn9PxqnEPmz2qYxjOTyRSKNMfNyvSmvMiYDHBP8AnmqouTDhQPxp5jWXDke9AieSEyJuA4z/ACqm1qQTXRRqBAM4GKz/ADY/7y/nVIzbMSMbT1Bz6U+Q8Y9OhooqTQhPz960I32pz2oooAqSzoc9WJ9Sao5+lFFWZH//2Q==\",\"owner\":{\"id\":\"248188406\"},\"thumbnail_src\":\"https://scontent-arn2-1.cdninstagram.com/t51.2885-15/s640x640/sh0.08/e35/c14.0.1052.1052/14072749_252283598503899_859117132_n.jpg\",\"is_video\":false,\"code\":\"BJc3ORygIfa\",\"date\":1471954804,\"display_src\":\"https://scontent-arn2-1.cdninstagram.com/t51.2885-15/e35/14072749_252283598503899_859117132_n.jpg\",\"comments\":{\"count\":1},\"likes\":{\"count\":36}},{\"__typename\":\"GraphImage\",\"id\":\"1283260814970678581\",\"comments_disabled\":false,\"dimensions\":{\"height\":938,\"width\":750},\"gating_info\":null,\"media_preview\":\"ACEqpyQMGDSkfMR06c9avsMYHp/Smtb7wfM+Z279h6Bfb370sZ3qGPJI/I9x+dZt3+RvFWv56g9o5+YNx2wcUgzjnqOD9RxWiuSm09cdKoBSy7gMdcgdj3FSnca0I9tFLn6/kaKYy7JEw5Qg47H+h/xH41DBGYjzk7m6AcAnk5PoOvYe1PguPP4HDA8/408opYkls9CASAcetJ6aMzTuronUg5PXB/OokGAfdiao3d+bdwgUEYBHOMZ/Cp42OD6bjSs1r3BtE1FQbqKdiLjUKWcZZvvH+fpUcE4n5T5ZB1U9G9x7/wCTWXcklhn0qHJBBHBFa8t1fqY81rJbF28QzyZXhgACpHIxn8Knt5ThlJztP86fqPCo3cnGe+MdKrpwlLdFvRk/nUVUzRTsZ8zP/9k=\",\"owner\":{\"id\":\"248188406\"},\"thumbnail_src\":\"https://scontent-arn2-1.cdninstagram.com/t51.2885-15/s640x640/sh0.08/e35/c0.85.682.682/13573621_1677236349025866_1510796872_n.jpg\",\"is_video\":false,\"code\":\"BHPDumGgW01\",\"date\":1467196631,\"display_src\":\"https://scontent-arn2-1.cdninstagram.com/t51.2885-15/e35/13573621_1677236349025866_1510796872_n.jpg\",\"caption\":\"Немножко эротики вам в ленту\",\"comments\":{\"count\":4},\"likes\":{\"count\":97}},{\"__typename\":\"GraphVideo\",\"id\":\"970208779275943528\",\"comments_disabled\":false,\"dimensions\":{\"height\":640,\"width\":640},\"gating_info\":null,\"media_preview\":null,\"owner\":{\"id\":\"248188406\"},\"thumbnail_src\":\"https://scontent-arn2-1.cdninstagram.com/t51.2885-15/e15/11176395_926076444103554_1888764605_n.jpg\",\"is_video\":true,\"code\":\"12376OtT5o\",\"date\":1429877921,\"display_src\":\"https://scontent-arn2-1.cdninstagram.com/t51.2885-15/e15/11176395_926076444103554_1888764605_n.jpg\",\"video_views\":0,\"caption\":\"Зов джунглей\",\"comments\":{\"count\":4},\"likes\":{\"count\":341}},{\"__typename\":\"GraphImage\",\"id\":\"874542900232666913\",\"comments_disabled\":false,\"dimensions\":{\"height\":640,\"width\":640},\"gating_info\":null,\"media_preview\":null,\"owner\":{\"id\":\"248188406\"},\"thumbnail_src\":\"https://scontent-arn2-1.cdninstagram.com/t51.2885-15/e15/10838337_330322153806719_1814034824_n.jpg\",\"is_video\":false,\"code\":\"wjACJdNT8h\",\"date\":1418473659,\"display_src\":\"https://scontent-arn2-1.cdninstagram.com/t51.2885-15/e15/10838337_330322153806719_1814034824_n.jpg\",\"caption\":\"Дружище, я спросил: есть че?\",\"comments\":{\"count\":1},\"likes\":{\"count\":47}},{\"__typename\":\"GraphImage\",\"id\":\"680496063176850510\",\"comments_disabled\":false,\"dimensions\":{\"height\":640,\"width\":640},\"gating_info\":null,\"media_preview\":null,\"owner\":{\"id\":\"248188406\"},\"thumbnail_src\":\"https://scontent-arn2-1.cdninstagram.com/t51.2885-15/e15/11417313_477504195757488_1603614850_n.jpg\",\"is_video\":false,\"code\":\"lxm5BktTxO\",\"date\":1395341473,\"display_src\":\"https://scontent-arn2-1.cdninstagram.com/t51.2885-15/e15/11417313_477504195757488_1603614850_n.jpg\",\"caption\":\"Принял эстафету от @marialexsandrovna и @barannikl. Передавать некому. Прочитал - значит принял. :)\",\"comments\":{\"count\":1},\"likes\":{\"count\":41}},{\"__typename\":\"GraphImage\",\"id\":\"540482521087164383\",\"comments_disabled\":false,\"dimensions\":{\"height\":612,\"width\":612},\"gating_info\":null,\"media_preview\":null,\"owner\":{\"id\":\"248188406\"},\"thumbnail_src\":\"https://scontent-arn2-1.cdninstagram.com/t51.2885-15/e15/11382920_486955271468994_1738264589_n.jpg\",\"is_video\":false,\"code\":\"eALf1EtT_f\",\"date\":1378650559,\"display_src\":\"https://scontent-arn2-1.cdninstagram.com/t51.2885-15/e15/11382920_486955271468994_1738264589_n.jpg\",\"caption\":\"Не знаю, как бы я высыпался, если бы не физика.\",\"comments\":{\"count\":0},\"likes\":{\"count\":40}},{\"__typename\":\"GraphImage\",\"id\":\"512999832411258539\",\"comments_disabled\":false,\"dimensions\":{\"height\":612,\"width\":612},\"gating_info\":null,\"media_preview\":null,\"owner\":{\"id\":\"248188406\"},\"thumbnail_src\":\"https://scontent-arn2-1.cdninstagram.com/t51.2885-15/e15/11325321_112286479105192_62945882_n.jpg\",\"is_video\":false,\"code\":\"ceiqEstT6r\",\"date\":1375374367,\"display_src\":\"https://scontent-arn2-1.cdninstagram.com/t51.2885-15/e15/11325321_112286479105192_62945882_n.jpg\",\"caption\":\"Дружище, есть че?\",\"comments\":{\"count\":0},\"likes\":{\"count\":33}},{\"__typename\":\"GraphImage\",\"id\":\"441358231205657855\",\"comments_disabled\":false,\"dimensions\":{\"height\":612,\"width\":612},\"gating_info\":null,\"media_preview\":null,\"owner\":{\"id\":\"248188406\"},\"thumbnail_src\":\"https://scontent-arn2-1.cdninstagram.com/t51.2885-15/e15/11330527_1587491721510550_960812045_n.jpg\",\"is_video\":false,\"code\":\"YgBPkNtTz_\",\"date\":1366834022,\"display_src\":\"https://scontent-arn2-1.cdninstagram.com/t51.2885-15/e15/11330527_1587491721510550_960812045_n.jpg\",\"caption\":\"Всем добра и расслабона.\",\"comments\":{\"count\":8},\"likes\":{\"count\":26}}],\"count\":12,\"page_info\":{\"has_next_page\":false,\"end_cursor\":\"AQBjZIMQGqq5EE74hgoVt6WwSHezmmG8OF2ytL6fyVaeVW9Wv3-1BqtmKEUIf9GokDo\"}}},\"logging_page_id\":\"profilePage_248188406\"}")

	account, err := getFromAccountPage(jsonBody)

	if err != nil {
		t.Error(err)
	}
	if account.Username != username {
		t.Errorf("Account username is incorrect.\nExpect %s, get %s.", account.Username, username)
	}
	if account.FullName != fullname {
		t.Errorf("Account fullname is incorrect.\nExpect %s, get %s.", account.FullName, fullname)
	}
	if account.ID != id {
		t.Errorf("Account id is incorrect.\nExpect %s, get %s.", account.ID, id)
	}
	if account.Biography != biography {
		t.Errorf("Account biography is incorrect.\nExpect %s, get %s.", account.Biography, biography)
	}
	if account.Verified != verified {
		t.Errorf("Account verified field is incorrect.\nExpect %b, get %b.", account.Verified, verified)
	}
	if account.MediaCount == 0 {
		t.Error("Account has empty media count.")
	}
	if account.Followers == 0 {
		t.Error("Account has empty followers count.")
	}
	if account.Follows == 0 {
		t.Error("Account has empty following count.")
	}
}

func Test_getFromSearchPage(t *testing.T) {
	fullname := "Konstantin"
	id := "248188406"
	jsonBody := []byte("{\"users\":[{\"position\":0,\"user\":{\"pk\":\"248188406\",\"username\":\"vorkytaka\",\"full_name\":\"Konstantin\",\"is_private\":false,\"profile_pic_url\":\"https://scontent-arn2-1.cdninstagram.com/t51.2885-19/s150x150/17125816_320904131658190_1521093063361953792_a.jpg\",\"profile_pic_id\":\"1467642282546943639_248188406\",\"is_verified\":false,\"has_anonymous_profile_picture\":false,\"follower_count\":116,\"byline\":\"Взаимные подписчики: 46\",\"social_context\":\"Подписчики: nahuyaggre+56\",\"search_social_context\":\"Подписчики: nahuyaggre+56\",\"mutual_followers_count\":46,\"following\":false,\"outgoing_request\":false}}],\"places\":[],\"hashtags\":[],\"has_more\":false,\"rank_token\":\"3d0af3b8-e35f-48b9-8b12-e01f45145e41\",\"status\":\"ok\"}")

	accounts, err := getFromSearchPage(jsonBody)
	if err != nil {
		t.Error(err)
	}
	if accounts[0].ID != id {
		t.Error("Incorrect account ID.")
	}
	if accounts[0].FullName != fullname {
		t.Error("Incorrect account fullname.")
	}
	if accounts[0].Private {
		t.Error("Incorrect account private field.")
	}
	if accounts[0].Verified {
		t.Error("Incorrect account verified field.")
	}
	if accounts[0].Followers == 0 {
		t.Error("Incorrect account followers count.")
	}
}
