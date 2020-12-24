package atinternet

type Properties []Property

type Property string

func (properties Properties) String() []string {

	propertiesString := []string{}

	for _, property := range properties {
		propertiesString = append(propertiesString, string(property))
	}

	return propertiesString
}

const (
	PAppCrash                  Property = "app_crash"
	PAppCrashClass             Property = "app_crash_class"
	PAppCrashScreen            Property = "app_crash_screen"
	PAppDsfs                   Property = "app_dsfs"
	PAppDsls                   Property = "app_dsls"
	PAppFsd                    Property = "app_fsd"
	PAppFsmn                   Property = "app_fsmn"
	PAppFsw                    Property = "app_fsw"
	PAppFswd                   Property = "app_fswd"
	PAppFswdn                  Property = "app_fswdn"
	PAppFswy                   Property = "app_fswy"
	PAppFsy                    Property = "app_fsy"
	PAppId                     Property = "app_id"
	PAppSc                     Property = "app_sc"
	PAppSessionStatus          Property = "app_session_status"
	PAppSessionid              Property = "app_sessionid"
	PAppVersion                Property = "app_version"
	PAppVisitorStatus          Property = "app_visitor_status"
	PArticleId                 Property = "article_id"
	PAudioId                   Property = "audio_id"
	PAvAdType                  Property = "av_ad_type"
	PAvBroadcastingType        Property = "av_broadcasting_type"
	PAvContent                 Property = "av_content"
	PAvContentDuration         Property = "av_content_duration"
	PAvContentGenre            Property = "av_content_genre"
	PAvContentId               Property = "av_content_id"
	PAvContentLinked           Property = "av_content_linked"
	PAvContentTheme1           Property = "av_content_theme1"
	PAvContentTheme2           Property = "av_content_theme2"
	PAvContentTheme3           Property = "av_content_theme3"
	PAvContentTimeConsumed     Property = "av_content_time_consumed"
	PAvContentType             Property = "av_content_type"
	PAvLanguage                Property = "av_language"
	PAvLaunchReason            Property = "av_launch_reason"
	PAvLocation                Property = "av_location"
	PAvPlaybackCompletionRate  Property = "av_playback_completion_rate"
	PAvPlaybackTotalDuration   Property = "av_playback_total_duration"
	PAvPlayer                  Property = "av_player"
	PAvSessionBackward         Property = "av_session_backward"
	PAvSessionBounce           Property = "av_session_bounce"
	PAvSessionContentDuration  Property = "av_session_content_duration"
	PAvSessionForward          Property = "av_session_forward"
	PAvSessionId               Property = "av_session_id"
	PAvSessionTime             Property = "av_session_time"
	PAvSubtitles               Property = "av_subtitles"
	PAvUrl                     Property = "av_url"
	PBrowser                   Property = "browser"
	PBrowserCookieAcceptance   Property = "browser_cookie_acceptance"
	PBrowserGroup              Property = "browser_group"
	PBrowserLanguage           Property = "browser_language"
	PBrowserVersion            Property = "browser_version"
	PCartCreationUtc           Property = "cart_creation_utc"
	PCartId                    Property = "cart_id"
	PCartLifetime              Property = "cart_lifetime"
	PCartNbdistinctproduct     Property = "cart_nbdistinctproduct"
	PCartQuantity              Property = "cart_quantity"
	PCartVersion               Property = "cart_version"
	PClick                     Property = "click"
	PClickChapter1             Property = "click_chapter1"
	PClickChapter2             Property = "click_chapter2"
	PClickChapter3             Property = "click_chapter3"
	PClickFullName             Property = "click_full_name"
	PConnectionIsp             Property = "connection_isp"
	PConnectionMonitor         Property = "connection_monitor"
	PConnectionOrganisation    Property = "connection_organisation"
	PConnectionType            Property = "connection_type"
	PDate                      Property = "date"
	PDateDay                   Property = "date_day"
	PDateDaynumber             Property = "date_daynumber"
	PDateMonth                 Property = "date_month"
	PDateMonthnumber           Property = "date_monthnumber"
	PDateWeek                  Property = "date_week"
	PDateYear                  Property = "date_year"
	PDateYearofweek            Property = "date_yearofweek"
	PDeviceBrand               Property = "device_brand"
	PDeviceDisplayWidth        Property = "device_display_width"
	PDeviceHour                Property = "device_hour"
	PDeviceName                Property = "device_name"
	PDeviceNameTech            Property = "device_name_tech"
	PDeviceScreenDiagonal      Property = "device_screen_diagonal"
	PDeviceScreenHeight        Property = "device_screen_height"
	PDeviceScreenWidth         Property = "device_screen_width"
	PDeviceType                Property = "device_type"
	PDomain                    Property = "domain"
	PEventCollectionPlatform   Property = "event_collection_platform"
	PEventCollectionVersion    Property = "event_collection_version"
	PEventHour                 Property = "event_hour"
	PEventId                   Property = "event_id"
	PEventMinute               Property = "event_minute"
	PEventName                 Property = "event_name"
	PEventPosition             Property = "event_position"
	PEventSecond               Property = "event_second"
	PEventTime                 Property = "event_time"
	PEventTimeUtc              Property = "event_time_utc"
	PEventUrl                  Property = "event_url"
	PExclusionCause            Property = "exclusion_cause"
	PExclusionType             Property = "exclusion_type"
	PGeencrypteerdeIdfa        Property = "geencrypteerde_idfa"
	PGeoCity                   Property = "geo_city"
	PGeoCityPc                 Property = "geo_city_pc"
	PGeoContinent              Property = "geo_continent"
	PGeoCountry                Property = "geo_country"
	PGeoLatitude               Property = "geo_latitude"
	PGeoLongitude              Property = "geo_longitude"
	PGeoMetro                  Property = "geo_metro"
	PGeoRegion                 Property = "geo_region"
	PGoalType                  Property = "goal_type"
	PIseClickRank              Property = "ise_click_rank"
	PIseKeyword                Property = "ise_keyword"
	PIsePage                   Property = "ise_page"
	PIseResult                 Property = "ise_result"
	PMvCreation                Property = "mv_creation"
	PMvTest                    Property = "mv_test"
	PMvWave                    Property = "mv_wave"
	PNoboColumn11              Property = "nobo_column_11"
	PNoboColumn12              Property = "nobo_column_12"
	PNoboColumn13              Property = "nobo_column_13"
	PNoboColumn14              Property = "nobo_column_14"
	PNoboColumn20              Property = "nobo_column_20"
	PNoboColumn21              Property = "nobo_column_21"
	PNoboColumn22              Property = "nobo_column_22"
	PNoboColumn23              Property = "nobo_column_23"
	PNoboColumn24              Property = "nobo_column_24"
	PNoboColumn25              Property = "nobo_column_25"
	PNoboColumn26              Property = "nobo_column_26"
	PNoboColumn27              Property = "nobo_column_27"
	PNoboColumn28              Property = "nobo_column_28"
	PNoboColumn29              Property = "nobo_column_29"
	PNoboColumn30              Property = "nobo_column_30"
	PNoboColumn4               Property = "nobo_column_4"
	PNoboColumn5               Property = "nobo_column_5"
	PNoboDevice                Property = "nobo_device"
	PNoboEventType             Property = "nobo_event_type"
	PNoboIdfv                  Property = "nobo_idfv"
	PNoboMediaType             Property = "nobo_media_type"
	PNoboPlatform              Property = "nobo_platform"
	PNoboSecundaryMediaBrand   Property = "nobo_secundary_media_brand"
	PNoboVisitorId             Property = "nobo_visitor_id"
	PNoboWebsiteName           Property = "nobo_website_name"
	PNosStatus                 Property = "nos_status"
	POffsiteConfirmation       Property = "offsite_confirmation"
	POnsiteadAdvertiser        Property = "onsitead_advertiser"
	POnsiteadCampaign          Property = "onsitead_campaign"
	POnsiteadCategory          Property = "onsitead_category"
	POnsiteadCreation          Property = "onsitead_creation"
	POnsiteadDetailedPlacement Property = "onsitead_detailed_placement"
	POnsiteadFormat            Property = "onsitead_format"
	POnsiteadGeneralPlacement  Property = "onsitead_general_placement"
	POnsiteadType              Property = "onsitead_type"
	POnsiteadUrl               Property = "onsitead_url"
	POnsiteadVariant           Property = "onsitead_variant"
	POs                        Property = "os"
	POsGroup                   Property = "os_group"
	POsVersion                 Property = "os_version"
	POsVersionName             Property = "os_version_name"
	PPage                      Property = "page"
	PPageAisle1                Property = "page_aisle1"
	PPageAisle2                Property = "page_aisle2"
	PPageAisle3                Property = "page_aisle3"
	PPageAisle4                Property = "page_aisle4"
	PPageAisle5                Property = "page_aisle5"
	PPageAisle6                Property = "page_aisle6"
	PPageChapter1              Property = "page_chapter1"
	PPageChapter2              Property = "page_chapter2"
	PPageChapter3              Property = "page_chapter3"
	PPageCustomcat1            Property = "page_customcat1"
	PPageCustomcat2            Property = "page_customcat2"
	PPageCustomcat3            Property = "page_customcat3"
	PPageDuration              Property = "page_duration"
	PPageFullName              Property = "page_full_name"
	PPagePosition              Property = "page_position"
	PPaymentMode               Property = "payment_mode"
	PPlatform                  Property = "platform"
	PPodcastEpisode            Property = "podcast_episode"
	PPodcastFeed               Property = "podcast_feed"
	PPodcastUrl                Property = "podcast_url"
	PPreviousDomain            Property = "previous_domain"
	PPrivacyStatus             Property = "privacy_status"
	PProduct                   Property = "product"
	PProductArticle            Property = "product_article"
	PProductBrand              Property = "product_brand"
	PProductCategory1          Property = "product_category1"
	PProductCategory2          Property = "product_category2"
	PProductCategory3          Property = "product_category3"
	PProductCategory4          Property = "product_category4"
	PProductCategory5          Property = "product_category5"
	PProductCategory6          Property = "product_category6"
	PProductDiscount           Property = "product_discount"
	PProductId                 Property = "product_id"
	PProductPlacement          Property = "product_placement"
	PProductPricetaxfree       Property = "product_pricetaxfree"
	PProductPricetaxincluded   Property = "product_pricetaxincluded"
	PProductPromocode          Property = "product_promocode"
	PProductQuantity           Property = "product_quantity"
	PProductStock              Property = "product_stock"
	PProductVariant            Property = "product_variant"
	PShippingCosttaxfree       Property = "shipping_costtaxfree"
	PShippingCosttaxincluded   Property = "shipping_costtaxincluded"
	PShippingDelivery          Property = "shipping_delivery"
	PSite                      Property = "site"
	PSiteEnv                   Property = "site_env"
	PSiteId                    Property = "site_id"
	PSiteLevel2                Property = "site_level2"
	PSourcesEvent              Property = "sources_event"
	PSrc                       Property = "src"
	PSrcAdChannel              Property = "src_ad_channel"
	PSrcAdDetailPlacement      Property = "src_ad_detail_placement"
	PSrcAdGeneralPlacement     Property = "src_ad_general_placement"
	PSrcAffIdentifier          Property = "src_aff_identifier"
	PSrcAffType                Property = "src_aff_type"
	PSrcCampaign               Property = "src_campaign"
	PSrcCampaignGroup          Property = "src_campaign_group"
	PSrcCreation               Property = "src_creation"
	PSrcDetail                 Property = "src_detail"
	PSrcEmailLink              Property = "src_email_link"
	PSrcEmailRecipient         Property = "src_email_recipient"
	PSrcEmailRecipientList     Property = "src_email_recipient_list"
	PSrcEmailSendDate          Property = "src_email_send_date"
	PSrcFormat                 Property = "src_format"
	PSrcMedium                 Property = "src_medium"
	PSrcOrganic                Property = "src_organic"
	PSrcOrganicDetail          Property = "src_organic_detail"
	PSrcPortalDomain           Property = "src_portal_domain"
	PSrcPortalSite             Property = "src_portal_site"
	PSrcPortalSiteId           Property = "src_portal_site_id"
	PSrcPortalUrl              Property = "src_portal_url"
	PSrcReferrerSiteDomain     Property = "src_referrer_site_domain"
	PSrcReferrerSiteUrl        Property = "src_referrer_site_url"
	PSrcReferrerUrl            Property = "src_referrer_url"
	PSrcSeCategory             Property = "src_se_category"
	PSrcSlNetwork              Property = "src_sl_network"
	PSrcSlTerm                 Property = "src_sl_term"
	PSrcType                   Property = "src_type"
	PSrcUrl                    Property = "src_url"
	PSrcUrlDomain              Property = "src_url_domain"
	PSrcVariant                Property = "src_variant"
	PTransactionDate           Property = "transaction_date"
	PTransactionFirstpurchase  Property = "transaction_firstpurchase"
	PTransactionId             Property = "transaction_id"
	PTransactionPromocode      Property = "transaction_promocode"
	PTransactionStatus         Property = "transaction_status"
	PUserCategory              Property = "user_category"
	PUserId                    Property = "user_id"
	PUserRecognition           Property = "user_recognition"
	PUtmCampaign               Property = "utm_campaign"
	PUtmContent                Property = "utm_content"
	PUtmMedium                 Property = "utm_medium"
	PUtmSource                 Property = "utm_source"
	PUtmTerm                   Property = "utm_term"
	PVideoId                   Property = "video_id"
	PVisitBounce               Property = "visit_bounce"
	PVisitConverted            Property = "visit_converted"
	PVisitDuration             Property = "visit_duration"
	PVisitEntryAisle1          Property = "visit_entry_aisle1"
	PVisitEntryAisle2          Property = "visit_entry_aisle2"
	PVisitEntryAisle3          Property = "visit_entry_aisle3"
	PVisitEntryAisle4          Property = "visit_entry_aisle4"
	PVisitEntryAisle5          Property = "visit_entry_aisle5"
	PVisitEntryAisle6          Property = "visit_entry_aisle6"
	PVisitEntrySiteLevel2      Property = "visit_entry_site_level2"
	PVisitEntrypage            Property = "visit_entrypage"
	PVisitEntrypageChapter1    Property = "visit_entrypage_chapter1"
	PVisitEntrypageChapter2    Property = "visit_entrypage_chapter2"
	PVisitEntrypageChapter3    Property = "visit_entrypage_chapter3"
	PVisitEntrypageFullName    Property = "visit_entrypage_full_name"
	PVisitExitSiteLevel2       Property = "visit_exit_site_level2"
	PVisitExitpage             Property = "visit_exitpage"
	PVisitExitpageChapter1     Property = "visit_exitpage_chapter1"
	PVisitExitpageChapter2     Property = "visit_exitpage_chapter2"
	PVisitExitpageChapter3     Property = "visit_exitpage_chapter3"
	PVisitExitpageFullName     Property = "visit_exitpage_full_name"
	PVisitHour                 Property = "visit_hour"
	PVisitId                   Property = "visit_id"
	PVisitImplicationDegree    Property = "visit_implication_degree"
	PVisitMinute               Property = "visit_minute"
	PVisitPageViews            Property = "visit_page_views"
	PVisitSales                Property = "visit_sales"
	PVisitSecond               Property = "visit_second"
	PVisitorId                 Property = "visitor_id"
	PVisitorPrivacyConsent     Property = "visitor_privacy_consent"
	PVisitorPrivacyMode        Property = "visitor_privacy_mode"
)
