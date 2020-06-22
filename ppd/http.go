package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"strings"
)

func HttpGet(url string) (ret string, err error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("http error %s\n", err)
	}
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.131 Safari/537.36")
	req.Header.Add("referer", "https://www.zhihu.com/")
	// req.Header.Add("x-ab-param", `zr_training_boost=false;zr_test_aa1=0;se_relation_1=2;se_college=default;se_ffzx_jushen1=0;se_oneboxtopic=1;se_sug_term=0;tp_club_new=1;tsp_hotlist_ui=3;li_yxzl_new_style_a=1;li_se_section=1;se_bert_eng=0;se_cardrank_2=1;se_new_bert=0;se_v045=0;se_cardrank_3=0;tp_club_top=0;tp_header_style=1;tp_club_flow_ai=1;se_billboardsearch=0;se_video_tab=0;se_bsi=0;tp_m_intro_re_topic=1;ls_videoad=2;zr_search_sim2=1;se_video_dnn=1;li_svip_cardshow=1;li_catalog_card=1;zr_intervene=0;soc_feed_intelligent=0;tp_club_reactionv2=0;tp_dingyue_video=0;tsp_ad_cardredesign=0;zr_art_rec=base;zr_ans_rec=gbrank;tp_club_fdv4=0;pf_adjust=1;li_paid_answer_exp=0;se_clubrank=1;se_clarify=1;se_v043=0;pf_creator_card=1;pf_foltopic_usernum=50;ls_recommend_test=5;ls_video_commercial=0;li_svip_tab_search=1;li_panswer_topic=1;zr_training_first=true;se_topicfeed=0;se_v048=0;tp_club_qa_entrance=1;tp_topic_tab=0;tp_club_feedv3=0;qap_question_visitor= 0;zw_sameq_sorce=999;se_aa_base=1;se_cardrank_4=1;se_multianswer=2;se_v046=0;tp_movie_ux=2;zr_km_answer=open_cvr;se_whitelist=1;se_backsearch=0;se_relationship=1;se_new_cbert=0;se_searchvideo=3;se_topic_wei=0;tp_sft=a;tp_move_scorecard=0;soc_iosweeklynew=2;ug_newtag=1;zr_search_topic=1;se_club_boost=1;ug_follow_topic_1=2;se_hotsearch=1;top_test_4_liguangyi=1;pf_noti_entry_num=0;zr_search_paid=1;zr_rel_search=base;se_content0=1;se_videobox=0;tp_meta_card=0;tp_topic_entry=0;top_hotcommerce=1;li_topics_search=0;li_salt_hot=1;se_wil_act=0;tp_discover=1;top_quality=0;ug_goodcomment_0=1;tp_club_feed=0;top_root=0;li_vip_verti_search=0;li_answer_card=0;li_video_section=1;qap_question_author=0;zr_slot_training=2;top_v_album=1;top_universalebook=1;ls_fmp4=0;li_car_meta=0;se_hotmore=2;li_viptab_name=0;se_expired_ob=0;se_web0answer=0;se_adsrank=4;soc_adweeklynew=2;pf_fuceng=1;pf_profile2_tab=0;qap_labeltype=1;zr_expslotpaid=1;tp_club_entrance=1;tsp_ios_cardredesign=0;se_v044=0;se_v049=0;tp_club_bt=0;tp_topic_style=0;li_ebook_gen_search=2;zr_search_sims=0;se_click_v_v=1;se_col_boost=1;se_v040=0;tp_score_1=a;se_colorfultab=1;se_searchwiki=1;se_v040_2=2;pf_newguide_vertical=0;zr_zr_search_sims=0;zr_rec_answer_cp=open;zr_slotpaidexp=8;top_ebook=0;li_training_chapter=1;se_entity22=1;se_mobilecard=0;se_hotsearch_2=1;se_multi_images=0;tp_club__entrance2=1;tp_topic_tab_new=0-0-0;soc_notification=1`)
	setCookie(`_zap=77077964-ee0b-4611-8d86-d3fc0b2ec4b5; d_c0="ABAgH6aKdRCPTlzOmr_bbHzrLsiuxYkjsT8=|1575545990"; _ga=GA1.2.679528324.1583295687; capsion_ticket="2|1:0|10:1591171692|14:capsion_ticket|44:MDM2ODJkNDE0OGEwNDhlMDlkMWIwMTQ0M2Y0MWJmMzA=|93b246435201348e70258784fbb55941789745f302236ae8d7fdc8c1594402be"; z_c0="2|1:0|10:1591171695|4:z_c0|92:Mi4xWUdzSUFnQUFBQUFBRUNBZnBvcDFFQ1lBQUFCZ0FsVk5iNmpFWHdEb2NnNmNOVXJsM2Rwd0dPdkJEZDk3eERSTHhR|b96f978c8398b4d249752090c2916b5546389dfe07ec570f43fdab0c703cc85d"; q_c1=82c5f4f606e54a0f810f1f3e2aacadfb|1591949949000|1575888634000; _gid=GA1.2.606153531.1592194675; _xsrf=ed941e23-2366-45bc-8fdf-fc8ae4a3222e; SESSIONID=ta4NpZCzdny5A5RQHo7rJq5OyLsj3jWpqWfyFBzsZTA; JOID=UV4WCk6KkcGO-S7SB4CCGLqyVaEW-_X_zI9xlHfo-r2xtB2FUzxGLtr6Lt0BHvy_P50kYuwO2kESmVgk3act02E=; osd=UFkQBkuLlseC_C_VAYyHGb20WaQX_PPzyY52knvt-7q3uBiEVDpKK9v9KNEEH_u5M5glZeoC30AVn1Qh3KAr32Q=; tshl=; Hm_lvt_98beee57fd2ef70ccdd5ca52b9740c49=1592470764,1592470776,1592471508,1592472184; tst=r; Hm_lpvt_98beee57fd2ef70ccdd5ca52b9740c49=1592533296; KLBRSID=81978cf28cf03c58e07f705c156aa833|1592533721|1592532175`, req)
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("http error %s\n", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("http error %s\n", err)
	}
	return string(body), nil
}

func urlToDocument(url string) *goquery.Document {
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Errorf("urlToDocument error %s\n", err)
	}
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.88 Safari/537.36")
	request.Header.Add("Referer", "https://www.zhihu.com/")
	response, _ := client.Do(request)
	if err != nil {
		fmt.Errorf("urlToDocument error %s\n", err)
	}
	dom, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		fmt.Errorf("urlToDocument error %s\n", err)

	}
	return dom

}

func setCookie(cookies string, req *http.Request) {
	for _, i := range strings.Split(cookies, ";") {
		j := strings.Split(i, "=")
		k := j[0]
		v := strings.Trim(strings.Join(j[1:], "="), `"`)
		cookie := http.Cookie{Name: k, Value: v, Path: "/", MaxAge: 86400}
		req.AddCookie(&cookie)
	}
}

func GetRecommendsQuestions() {
	ret, err := HttpGet("https://www.zhihu.com/api/v3/feed/topstory/recommend?session_token=8ddff5c79509674c2291dd49176971ac&desktop=true&page_number=3&limit=6&action=down&after_id=11&ad_interval=-1")
	if err != nil {
		fmt.Println(err)
	}
	retDatas := gjson.Get(ret, "data.#.target.question")
	for _, questionInfo := range retDatas.Array() {
		title := gjson.Get(questionInfo.Raw, "title")
		// commentCount := gjson.Get(questionInfo.Raw, "comment_count")
		// followersCount := gjson.Get(questionInfo.Raw, "followers_count")
		url := gjson.Get(questionInfo.Raw, "url")
		fmt.Println(url, title)
	}
}

func main() {
	ret, err := HttpGet("https://www.zhihu.com/api/v4/questions/39957226/answers?include=data%5B*%5D.is_normal%2Cadmin_closed_comment%2Creward_info%2Cis_collapsed%2Cannotation_action%2Cannotation_detail%2Ccollapse_reason%2Cis_sticky%2Ccollapsed_by%2Csuggest_edit%2Ccomment_count%2Ccan_comment%2Ccontent%2Ceditable_content%2Cvoteup_count%2Creshipment_settings%2Ccomment_permission%2Ccreated_time%2Cupdated_time%2Creview_info%2Crelevant_info%2Cquestion%2Cexcerpt%2Crelationship.is_authorized%2Cis_author%2Cvoting%2Cis_thanked%2Cis_nothelp%2Cis_labeled%2Cis_recognized%2Cpaid_info%2Cpaid_info_content%3Bdata%5B*%5D.mark_infos%5B*%5D.url%3Bdata%5B*%5D.author.follower_count%2Cbadge%5B*%5D.topics&offset=3&limit=5&sort_by=default&platform=desktop")
	if err != nil {
		fmt.Println(err)
	}
	retDatas := gjson.Get(ret, "data")
	for _, questionInfo := range retDatas.Array() {
		title := gjson.Get(questionInfo.Raw, "title")
		// commentCount := gjson.Get(questionInfo.Raw, "comment_count")
		// followersCount := gjson.Get(questionInfo.Raw, "followers_count")
		content := gjson.Get(questionInfo.Raw, "content")
		fmt.Println(title, content)
	}
}
