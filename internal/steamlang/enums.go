// Generated code
// DO NOT EDIT

package steamlang

type EMsg int32

const (
	EMsg_Invalid                                                  EMsg = 0
	EMsg_Multi                                                         = 1
	EMsg_BaseGeneral                                                   = 100
	EMsg_GenericReply                                                  = 100
	EMsg_DestJobFailed                                                 = 113
	EMsg_Alert                                                         = 115
	EMsg_SCIDRequest                                                   = 120
	EMsg_SCIDResponse                                                  = 121
	EMsg_JobHeartbeat                                                  = 123
	EMsg_HubConnect                                                    = 124
	EMsg_Subscribe                                                     = 126
	EMsg_RouteMessage                                                  = 127
	EMsg_RemoteSysID                                                   = 128
	EMsg_AMCreateAccountResponse                                       = 129
	EMsg_WGRequest                                                     = 130
	EMsg_WGResponse                                                    = 131
	EMsg_KeepAlive                                                     = 132
	EMsg_WebAPIJobRequest                                              = 133
	EMsg_WebAPIJobResponse                                             = 134
	EMsg_ClientSessionStart                                            = 135
	EMsg_ClientSessionEnd                                              = 136
	EMsg_ClientSessionUpdateAuthTicket                                 = 137
	EMsg_StatsDeprecated                                               = 138 // Deprecated
	EMsg_Ping                                                          = 139
	EMsg_PingResponse                                                  = 140
	EMsg_Stats                                                         = 141
	EMsg_RequestFullStatsBlock                                         = 142
	EMsg_LoadDBOCacheItem                                              = 143
	EMsg_LoadDBOCacheItemResponse                                      = 144
	EMsg_InvalidateDBOCacheItems                                       = 145
	EMsg_ServiceMethod                                                 = 146
	EMsg_ServiceMethodResponse                                         = 147
	EMsg_BaseShell                                                     = 200
	EMsg_AssignSysID                                                   = 200
	EMsg_Exit                                                          = 201
	EMsg_DirRequest                                                    = 202
	EMsg_DirResponse                                                   = 203
	EMsg_ZipRequest                                                    = 204
	EMsg_ZipResponse                                                   = 205
	EMsg_UpdateRecordResponse                                          = 215
	EMsg_UpdateCreditCardRequest                                       = 221
	EMsg_UpdateUserBanResponse                                         = 225
	EMsg_PrepareToExit                                                 = 226
	EMsg_ContentDescriptionUpdate                                      = 227
	EMsg_TestResetServer                                               = 228
	EMsg_UniverseChanged                                               = 229
	EMsg_ShellConfigInfoUpdate                                         = 230
	EMsg_RequestWindowsEventLogEntries                                 = 233
	EMsg_ProvideWindowsEventLogEntries                                 = 234
	EMsg_ShellSearchLogs                                               = 235
	EMsg_ShellSearchLogsResponse                                       = 236
	EMsg_ShellCheckWindowsUpdates                                      = 237
	EMsg_ShellCheckWindowsUpdatesResponse                              = 238
	EMsg_ShellFlushUserLicenseCache                                    = 239
	EMsg_BaseGM                                                        = 300
	EMsg_Heartbeat                                                     = 300
	EMsg_ShellFailed                                                   = 301
	EMsg_ExitShells                                                    = 307
	EMsg_ExitShell                                                     = 308
	EMsg_GracefulExitShell                                             = 309
	EMsg_NotifyWatchdog                                                = 314
	EMsg_LicenseProcessingComplete                                     = 316
	EMsg_SetTestFlag                                                   = 317
	EMsg_QueuedEmailsComplete                                          = 318
	EMsg_GMReportPHPError                                              = 319
	EMsg_GMDRMSync                                                     = 320
	EMsg_PhysicalBoxInventory                                          = 321
	EMsg_UpdateConfigFile                                              = 322
	EMsg_TestInitDB                                                    = 323
	EMsg_GMWriteConfigToSQL                                            = 324
	EMsg_GMLoadActivationCodes                                         = 325
	EMsg_GMQueueForFBS                                                 = 326
	EMsg_GMSchemaConversionResults                                     = 327
	EMsg_GMSchemaConversionResultsResponse                             = 328
	EMsg_GMWriteShellFailureToSQL                                      = 329
	EMsg_BaseAIS                                                       = 400
	EMsg_AISRefreshContentDescription                                  = 401
	EMsg_AISRequestContentDescription                                  = 402
	EMsg_AISUpdateAppInfo                                              = 403
	EMsg_AISUpdatePackageInfo                                          = 404
	EMsg_AISGetPackageChangeNumber                                     = 405
	EMsg_AISGetPackageChangeNumberResponse                             = 406
	EMsg_AISAppInfoTableChanged                                        = 407
	EMsg_AISUpdatePackageInfoResponse                                  = 408
	EMsg_AISCreateMarketingMessage                                     = 409
	EMsg_AISCreateMarketingMessageResponse                             = 410
	EMsg_AISGetMarketingMessage                                        = 411
	EMsg_AISGetMarketingMessageResponse                                = 412
	EMsg_AISUpdateMarketingMessage                                     = 413
	EMsg_AISUpdateMarketingMessageResponse                             = 414
	EMsg_AISRequestMarketingMessageUpdate                              = 415
	EMsg_AISDeleteMarketingMessage                                     = 416
	EMsg_AISGetMarketingTreatments                                     = 419
	EMsg_AISGetMarketingTreatmentsResponse                             = 420
	EMsg_AISRequestMarketingTreatmentUpdate                            = 421
	EMsg_AISTestAddPackage                                             = 422
	EMsg_AIGetAppGCFlags                                               = 423
	EMsg_AIGetAppGCFlagsResponse                                       = 424
	EMsg_AIGetAppList                                                  = 425
	EMsg_AIGetAppListResponse                                          = 426
	EMsg_AIGetAppInfo                                                  = 427
	EMsg_AIGetAppInfoResponse                                          = 428
	EMsg_AISGetCouponDefinition                                        = 429
	EMsg_AISGetCouponDefinitionResponse                                = 430
	EMsg_BaseAM                                                        = 500
	EMsg_AMUpdateUserBanRequest                                        = 504
	EMsg_AMAddLicense                                                  = 505
	EMsg_AMBeginProcessingLicenses                                     = 507
	EMsg_AMSendSystemIMToUser                                          = 508
	EMsg_AMExtendLicense                                               = 509
	EMsg_AMAddMinutesToLicense                                         = 510
	EMsg_AMCancelLicense                                               = 511
	EMsg_AMInitPurchase                                                = 512
	EMsg_AMPurchaseResponse                                            = 513
	EMsg_AMGetFinalPrice                                               = 514
	EMsg_AMGetFinalPriceResponse                                       = 515
	EMsg_AMGetLegacyGameKey                                            = 516
	EMsg_AMGetLegacyGameKeyResponse                                    = 517
	EMsg_AMFindHungTransactions                                        = 518
	EMsg_AMSetAccountTrustedRequest                                    = 519
	EMsg_AMCompletePurchase                                            = 521
	EMsg_AMCancelPurchase                                              = 522
	EMsg_AMNewChallenge                                                = 523
	EMsg_AMFixPendingPurchaseResponse                                  = 526
	EMsg_AMIsUserBanned                                                = 527
	EMsg_AMRegisterKey                                                 = 528
	EMsg_AMLoadActivationCodes                                         = 529
	EMsg_AMLoadActivationCodesResponse                                 = 530
	EMsg_AMLookupKeyResponse                                           = 531
	EMsg_AMLookupKey                                                   = 532
	EMsg_AMChatCleanup                                                 = 533
	EMsg_AMClanCleanup                                                 = 534
	EMsg_AMFixPendingRefund                                            = 535
	EMsg_AMReverseChargeback                                           = 536
	EMsg_AMReverseChargebackResponse                                   = 537
	EMsg_AMClanCleanupList                                             = 538
	EMsg_AMGetLicenses                                                 = 539
	EMsg_AMGetLicensesResponse                                         = 540
	EMsg_AllowUserToPlayQuery                                          = 550
	EMsg_AllowUserToPlayResponse                                       = 551
	EMsg_AMVerfiyUser                                                  = 552
	EMsg_AMClientNotPlaying                                            = 553
	EMsg_ClientRequestFriendship                                       = 554
	EMsg_AMRelayPublishStatus                                          = 555
	EMsg_AMResetCommunityContent                                       = 556
	EMsg_AMPrimePersonaStateCache                                      = 557
	EMsg_AMAllowUserContentQuery                                       = 558
	EMsg_AMAllowUserContentResponse                                    = 559
	EMsg_AMInitPurchaseResponse                                        = 560
	EMsg_AMRevokePurchaseResponse                                      = 561
	EMsg_AMLockProfile                                                 = 562
	EMsg_AMRefreshGuestPasses                                          = 563
	EMsg_AMInviteUserToClan                                            = 564
	EMsg_AMAcknowledgeClanInvite                                       = 565
	EMsg_AMGrantGuestPasses                                            = 566
	EMsg_AMClanDataUpdated                                             = 567
	EMsg_AMReloadAccount                                               = 568
	EMsg_AMClientChatMsgRelay                                          = 569
	EMsg_AMChatMulti                                                   = 570
	EMsg_AMClientChatInviteRelay                                       = 571
	EMsg_AMChatInvite                                                  = 572
	EMsg_AMClientJoinChatRelay                                         = 573
	EMsg_AMClientChatMemberInfoRelay                                   = 574
	EMsg_AMPublishChatMemberInfo                                       = 575
	EMsg_AMClientAcceptFriendInvite                                    = 576
	EMsg_AMChatEnter                                                   = 577
	EMsg_AMClientPublishRemovalFromSource                              = 578
	EMsg_AMChatActionResult                                            = 579
	EMsg_AMFindAccounts                                                = 580
	EMsg_AMFindAccountsResponse                                        = 581
	EMsg_AMSetAccountFlags                                             = 584
	EMsg_AMCreateClan                                                  = 586
	EMsg_AMCreateClanResponse                                          = 587
	EMsg_AMGetClanDetails                                              = 588
	EMsg_AMGetClanDetailsResponse                                      = 589
	EMsg_AMSetPersonaName                                              = 590
	EMsg_AMSetAvatar                                                   = 591
	EMsg_AMAuthenticateUser                                            = 592
	EMsg_AMAuthenticateUserResponse                                    = 593
	EMsg_AMGetAccountFriendsCount                                      = 594
	EMsg_AMGetAccountFriendsCountResponse                              = 595
	EMsg_AMP2PIntroducerMessage                                        = 596
	EMsg_ClientChatAction                                              = 597
	EMsg_AMClientChatActionRelay                                       = 598
	EMsg_BaseVS                                                        = 600
	EMsg_ReqChallenge                                                  = 600
	EMsg_VACResponse                                                   = 601
	EMsg_ReqChallengeTest                                              = 602
	EMsg_VSMarkCheat                                                   = 604
	EMsg_VSAddCheat                                                    = 605
	EMsg_VSPurgeCodeModDB                                              = 606
	EMsg_VSGetChallengeResults                                         = 607
	EMsg_VSChallengeResultText                                         = 608
	EMsg_VSReportLingerer                                              = 609
	EMsg_VSRequestManagedChallenge                                     = 610
	EMsg_VSLoadDBFinished                                              = 611
	EMsg_BaseDRMS                                                      = 625
	EMsg_DRMBuildBlobRequest                                           = 628
	EMsg_DRMBuildBlobResponse                                          = 629
	EMsg_DRMResolveGuidRequest                                         = 630
	EMsg_DRMResolveGuidResponse                                        = 631
	EMsg_DRMVariabilityReport                                          = 633
	EMsg_DRMVariabilityReportResponse                                  = 634
	EMsg_DRMStabilityReport                                            = 635
	EMsg_DRMStabilityReportResponse                                    = 636
	EMsg_DRMDetailsReportRequest                                       = 637
	EMsg_DRMDetailsReportResponse                                      = 638
	EMsg_DRMProcessFile                                                = 639
	EMsg_DRMAdminUpdate                                                = 640
	EMsg_DRMAdminUpdateResponse                                        = 641
	EMsg_DRMSync                                                       = 642
	EMsg_DRMSyncResponse                                               = 643
	EMsg_DRMProcessFileResponse                                        = 644
	EMsg_DRMEmptyGuidCache                                             = 645
	EMsg_DRMEmptyGuidCacheResponse                                     = 646
	EMsg_BaseCS                                                        = 650
	EMsg_CSUserContentRequest                                          = 652 // Deprecated
	EMsg_BaseClient                                                    = 700
	EMsg_ClientLogOn_Deprecated                                        = 701 // Deprecated
	EMsg_ClientAnonLogOn_Deprecated                                    = 702 // Deprecated
	EMsg_ClientHeartBeat                                               = 703
	EMsg_ClientVACResponse                                             = 704
	EMsg_ClientGamesPlayed_obsolete                                    = 705 // Deprecated
	EMsg_ClientLogOff                                                  = 706
	EMsg_ClientNoUDPConnectivity                                       = 707
	EMsg_ClientInformOfCreateAccount                                   = 708
	EMsg_ClientAckVACBan                                               = 709
	EMsg_ClientConnectionStats                                         = 710
	EMsg_ClientInitPurchase                                            = 711
	EMsg_ClientPingResponse                                            = 712
	EMsg_ClientRemoveFriend                                            = 714
	EMsg_ClientGamesPlayedNoDataBlob                                   = 715
	EMsg_ClientChangeStatus                                            = 716
	EMsg_ClientVacStatusResponse                                       = 717
	EMsg_ClientFriendMsg                                               = 718
	EMsg_ClientGameConnect_obsolete                                    = 719 // Deprecated
	EMsg_ClientGamesPlayed2_obsolete                                   = 720 // Deprecated
	EMsg_ClientGameEnded_obsolete                                      = 721 // Deprecated
	EMsg_ClientGetFinalPrice                                           = 722
	EMsg_ClientSystemIM                                                = 726
	EMsg_ClientSystemIMAck                                             = 727
	EMsg_ClientGetLicenses                                             = 728
	EMsg_ClientCancelLicense                                           = 729 // Deprecated
	EMsg_ClientGetLegacyGameKey                                        = 730
	EMsg_ClientContentServerLogOn_Deprecated                           = 731 // Deprecated
	EMsg_ClientAckVACBan2                                              = 732
	EMsg_ClientAckMessageByGID                                         = 735
	EMsg_ClientGetPurchaseReceipts                                     = 736
	EMsg_ClientAckPurchaseReceipt                                      = 737
	EMsg_ClientGamesPlayed3_obsolete                                   = 738 // Deprecated
	EMsg_ClientSendGuestPass                                           = 739
	EMsg_ClientAckGuestPass                                            = 740
	EMsg_ClientRedeemGuestPass                                         = 741
	EMsg_ClientGamesPlayed                                             = 742
	EMsg_ClientRegisterKey                                             = 743
	EMsg_ClientInviteUserToClan                                        = 744
	EMsg_ClientAcknowledgeClanInvite                                   = 745
	EMsg_ClientPurchaseWithMachineID                                   = 746
	EMsg_ClientAppUsageEvent                                           = 747
	EMsg_ClientGetGiftTargetList                                       = 748 // Deprecated
	EMsg_ClientGetGiftTargetListResponse                               = 749 // Deprecated
	EMsg_ClientLogOnResponse                                           = 751
	EMsg_ClientVACChallenge                                            = 753
	EMsg_ClientSetHeartbeatRate                                        = 755
	EMsg_ClientNotLoggedOnDeprecated                                   = 756 // Deprecated
	EMsg_ClientLoggedOff                                               = 757
	EMsg_GSApprove                                                     = 758
	EMsg_GSDeny                                                        = 759
	EMsg_GSKick                                                        = 760
	EMsg_ClientCreateAcctResponse                                      = 761
	EMsg_ClientPurchaseResponse                                        = 763
	EMsg_ClientPing                                                    = 764
	EMsg_ClientNOP                                                     = 765
	EMsg_ClientPersonaState                                            = 766
	EMsg_ClientFriendsList                                             = 767
	EMsg_ClientAccountInfo                                             = 768
	EMsg_ClientVacStatusQuery                                          = 770
	EMsg_ClientNewsUpdate                                              = 771
	EMsg_ClientGameConnectDeny                                         = 773
	EMsg_GSStatusReply                                                 = 774
	EMsg_ClientGetFinalPriceResponse                                   = 775
	EMsg_ClientGameConnectTokens                                       = 779
	EMsg_ClientLicenseList                                             = 780
	EMsg_ClientCancelLicenseResponse                                   = 781 // Deprecated
	EMsg_ClientVACBanStatus                                            = 782
	EMsg_ClientCMList                                                  = 783
	EMsg_ClientEncryptPct                                              = 784
	EMsg_ClientGetLegacyGameKeyResponse                                = 785
	EMsg_ClientFavoritesList                                           = 786
	EMsg_CSUserContentApprove                                          = 787 // Deprecated
	EMsg_CSUserContentDeny                                             = 788 // Deprecated
	EMsg_ClientInitPurchaseResponse                                    = 789
	EMsg_ClientAddFriend                                               = 791
	EMsg_ClientAddFriendResponse                                       = 792
	EMsg_ClientInviteFriend                                            = 793
	EMsg_ClientInviteFriendResponse                                    = 794
	EMsg_ClientSendGuestPassResponse                                   = 795 // Deprecated
	EMsg_ClientAckGuestPassResponse                                    = 796
	EMsg_ClientRedeemGuestPassResponse                                 = 797
	EMsg_ClientUpdateGuestPassesList                                   = 798
	EMsg_ClientChatMsg                                                 = 799
	EMsg_ClientChatInvite                                              = 800
	EMsg_ClientJoinChat                                                = 801
	EMsg_ClientChatMemberInfo                                          = 802
	EMsg_ClientLogOnWithCredentials_Deprecated                         = 803 // Deprecated
	EMsg_ClientPasswordChangeResponse                                  = 805
	EMsg_ClientChatEnter                                               = 807
	EMsg_ClientFriendRemovedFromSource                                 = 808
	EMsg_ClientCreateChat                                              = 809
	EMsg_ClientCreateChatResponse                                      = 810
	EMsg_ClientUpdateChatMetadata                                      = 811
	EMsg_ClientP2PIntroducerMessage                                    = 813
	EMsg_ClientChatActionResult                                        = 814
	EMsg_ClientRequestFriendData                                       = 815
	EMsg_ClientGetUserStats                                            = 818
	EMsg_ClientGetUserStatsResponse                                    = 819
	EMsg_ClientStoreUserStats                                          = 820
	EMsg_ClientStoreUserStatsResponse                                  = 821
	EMsg_ClientClanState                                               = 822
	EMsg_ClientServiceModule                                           = 830
	EMsg_ClientServiceCall                                             = 831
	EMsg_ClientServiceCallResponse                                     = 832
	EMsg_ClientPackageInfoRequest                                      = 833
	EMsg_ClientPackageInfoResponse                                     = 834
	EMsg_ClientNatTraversalStatEvent                                   = 839
	EMsg_ClientAppInfoRequest                                          = 840
	EMsg_ClientAppInfoResponse                                         = 841
	EMsg_ClientSteamUsageEvent                                         = 842
	EMsg_ClientCheckPassword                                           = 845
	EMsg_ClientResetPassword                                           = 846
	EMsg_ClientCheckPasswordResponse                                   = 848
	EMsg_ClientResetPasswordResponse                                   = 849
	EMsg_ClientSessionToken                                            = 850
	EMsg_ClientDRMProblemReport                                        = 851
	EMsg_ClientSetIgnoreFriend                                         = 855
	EMsg_ClientSetIgnoreFriendResponse                                 = 856
	EMsg_ClientGetAppOwnershipTicket                                   = 857
	EMsg_ClientGetAppOwnershipTicketResponse                           = 858
	EMsg_ClientGetLobbyListResponse                                    = 860
	EMsg_ClientGetLobbyMetadata                                        = 861
	EMsg_ClientGetLobbyMetadataResponse                                = 862
	EMsg_ClientVTTCert                                                 = 863
	EMsg_ClientAppInfoUpdate                                           = 866
	EMsg_ClientAppInfoChanges                                          = 867
	EMsg_ClientServerList                                              = 880
	EMsg_ClientEmailChangeResponse                                     = 891
	EMsg_ClientSecretQAChangeResponse                                  = 892
	EMsg_ClientDRMBlobRequest                                          = 896
	EMsg_ClientDRMBlobResponse                                         = 897
	EMsg_ClientLookupKey                                               = 898
	EMsg_ClientLookupKeyResponse                                       = 899
	EMsg_BaseGameServer                                                = 900
	EMsg_GSDisconnectNotice                                            = 901
	EMsg_GSStatus                                                      = 903
	EMsg_GSUserPlaying                                                 = 905
	EMsg_GSStatus2                                                     = 906
	EMsg_GSStatusUpdate_Unused                                         = 907
	EMsg_GSServerType                                                  = 908
	EMsg_GSPlayerList                                                  = 909
	EMsg_GSGetUserAchievementStatus                                    = 910
	EMsg_GSGetUserAchievementStatusResponse                            = 911
	EMsg_GSGetPlayStats                                                = 918
	EMsg_GSGetPlayStatsResponse                                        = 919
	EMsg_GSGetUserGroupStatus                                          = 920
	EMsg_AMGetUserGroupStatus                                          = 921
	EMsg_AMGetUserGroupStatusResponse                                  = 922
	EMsg_GSGetUserGroupStatusResponse                                  = 923
	EMsg_GSGetReputation                                               = 936
	EMsg_GSGetReputationResponse                                       = 937
	EMsg_GSAssociateWithClan                                           = 938
	EMsg_GSAssociateWithClanResponse                                   = 939
	EMsg_GSComputeNewPlayerCompatibility                               = 940
	EMsg_GSComputeNewPlayerCompatibilityResponse                       = 941
	EMsg_BaseAdmin                                                     = 1000
	EMsg_AdminCmd                                                      = 1000
	EMsg_AdminCmdResponse                                              = 1004
	EMsg_AdminLogListenRequest                                         = 1005
	EMsg_AdminLogEvent                                                 = 1006
	EMsg_LogSearchRequest                                              = 1007
	EMsg_LogSearchResponse                                             = 1008
	EMsg_LogSearchCancel                                               = 1009
	EMsg_UniverseData                                                  = 1010
	EMsg_RequestStatHistory                                            = 1014
	EMsg_StatHistory                                                   = 1015
	EMsg_AdminPwLogon                                                  = 1017
	EMsg_AdminPwLogonResponse                                          = 1018
	EMsg_AdminSpew                                                     = 1019
	EMsg_AdminConsoleTitle                                             = 1020
	EMsg_AdminGCSpew                                                   = 1023
	EMsg_AdminGCCommand                                                = 1024
	EMsg_AdminGCGetCommandList                                         = 1025
	EMsg_AdminGCGetCommandListResponse                                 = 1026
	EMsg_FBSConnectionData                                             = 1027
	EMsg_AdminMsgSpew                                                  = 1028
	EMsg_BaseFBS                                                       = 1100
	EMsg_FBSReqVersion                                                 = 1100
	EMsg_FBSVersionInfo                                                = 1101
	EMsg_FBSForceRefresh                                               = 1102
	EMsg_FBSForceBounce                                                = 1103
	EMsg_FBSDeployPackage                                              = 1104
	EMsg_FBSDeployResponse                                             = 1105
	EMsg_FBSUpdateBootstrapper                                         = 1106
	EMsg_FBSSetState                                                   = 1107
	EMsg_FBSApplyOSUpdates                                             = 1108
	EMsg_FBSRunCMDScript                                               = 1109
	EMsg_FBSRebootBox                                                  = 1110
	EMsg_FBSSetBigBrotherMode                                          = 1111
	EMsg_FBSMinidumpServer                                             = 1112
	EMsg_FBSSetShellCount_obsolete                                     = 1113 // Deprecated
	EMsg_FBSDeployHotFixPackage                                        = 1114
	EMsg_FBSDeployHotFixResponse                                       = 1115
	EMsg_FBSDownloadHotFix                                             = 1116
	EMsg_FBSDownloadHotFixResponse                                     = 1117
	EMsg_FBSUpdateTargetConfigFile                                     = 1118
	EMsg_FBSApplyAccountCred                                           = 1119
	EMsg_FBSApplyAccountCredResponse                                   = 1120
	EMsg_FBSSetShellCount                                              = 1121
	EMsg_FBSTerminateShell                                             = 1122
	EMsg_FBSQueryGMForRequest                                          = 1123
	EMsg_FBSQueryGMResponse                                            = 1124
	EMsg_FBSTerminateZombies                                           = 1125
	EMsg_FBSInfoFromBootstrapper                                       = 1126
	EMsg_FBSRebootBoxResponse                                          = 1127
	EMsg_FBSBootstrapperPackageRequest                                 = 1128
	EMsg_FBSBootstrapperPackageResponse                                = 1129
	EMsg_FBSBootstrapperGetPackageChunk                                = 1130
	EMsg_FBSBootstrapperGetPackageChunkResponse                        = 1131
	EMsg_FBSBootstrapperPackageTransferProgress                        = 1132
	EMsg_FBSRestartBootstrapper                                        = 1133
	EMsg_BaseFileXfer                                                  = 1200
	EMsg_FileXferRequest                                               = 1200
	EMsg_FileXferResponse                                              = 1201
	EMsg_FileXferData                                                  = 1202
	EMsg_FileXferEnd                                                   = 1203
	EMsg_FileXferDataAck                                               = 1204
	EMsg_BaseChannelAuth                                               = 1300
	EMsg_ChannelAuthChallenge                                          = 1300
	EMsg_ChannelAuthResponse                                           = 1301
	EMsg_ChannelAuthResult                                             = 1302
	EMsg_ChannelEncryptRequest                                         = 1303
	EMsg_ChannelEncryptResponse                                        = 1304
	EMsg_ChannelEncryptResult                                          = 1305
	EMsg_BaseBS                                                        = 1400
	EMsg_BSPurchaseStart                                               = 1401
	EMsg_BSPurchaseResponse                                            = 1402
	EMsg_BSSettleNOVA                                                  = 1404
	EMsg_BSSettleComplete                                              = 1406
	EMsg_BSBannedRequest                                               = 1407
	EMsg_BSInitPayPalTxn                                               = 1408
	EMsg_BSInitPayPalTxnResponse                                       = 1409
	EMsg_BSGetPayPalUserInfo                                           = 1410
	EMsg_BSGetPayPalUserInfoResponse                                   = 1411
	EMsg_BSRefundTxn                                                   = 1413
	EMsg_BSRefundTxnResponse                                           = 1414
	EMsg_BSGetEvents                                                   = 1415
	EMsg_BSChaseRFRRequest                                             = 1416
	EMsg_BSPaymentInstrBan                                             = 1417
	EMsg_BSPaymentInstrBanResponse                                     = 1418
	EMsg_BSProcessGCReports                                            = 1419
	EMsg_BSProcessPPReports                                            = 1420
	EMsg_BSInitGCBankXferTxn                                           = 1421
	EMsg_BSInitGCBankXferTxnResponse                                   = 1422
	EMsg_BSQueryGCBankXferTxn                                          = 1423
	EMsg_BSQueryGCBankXferTxnResponse                                  = 1424
	EMsg_BSCommitGCTxn                                                 = 1425
	EMsg_BSQueryTransactionStatus                                      = 1426
	EMsg_BSQueryTransactionStatusResponse                              = 1427
	EMsg_BSQueryCBOrderStatus                                          = 1428
	EMsg_BSQueryCBOrderStatusResponse                                  = 1429
	EMsg_BSRunRedFlagReport                                            = 1430
	EMsg_BSQueryPaymentInstUsage                                       = 1431
	EMsg_BSQueryPaymentInstResponse                                    = 1432
	EMsg_BSQueryTxnExtendedInfo                                        = 1433
	EMsg_BSQueryTxnExtendedInfoResponse                                = 1434
	EMsg_BSUpdateConversionRates                                       = 1435
	EMsg_BSProcessUSBankReports                                        = 1436
	EMsg_BSPurchaseRunFraudChecks                                      = 1437
	EMsg_BSPurchaseRunFraudChecksResponse                              = 1438
	EMsg_BSStartShippingJobs                                           = 1439
	EMsg_BSQueryBankInformation                                        = 1440
	EMsg_BSQueryBankInformationResponse                                = 1441
	EMsg_BSValidateXsollaSignature                                     = 1445
	EMsg_BSValidateXsollaSignatureResponse                             = 1446
	EMsg_BSQiwiWalletInvoice                                           = 1448
	EMsg_BSQiwiWalletInvoiceResponse                                   = 1449
	EMsg_BSUpdateInventoryFromProPack                                  = 1450
	EMsg_BSUpdateInventoryFromProPackResponse                          = 1451
	EMsg_BSSendShippingRequest                                         = 1452
	EMsg_BSSendShippingRequestResponse                                 = 1453
	EMsg_BSGetProPackOrderStatus                                       = 1454
	EMsg_BSGetProPackOrderStatusResponse                               = 1455
	EMsg_BSCheckJobRunning                                             = 1456
	EMsg_BSCheckJobRunningResponse                                     = 1457
	EMsg_BSResetPackagePurchaseRateLimit                               = 1458
	EMsg_BSResetPackagePurchaseRateLimitResponse                       = 1459
	EMsg_BSUpdatePaymentData                                           = 1460
	EMsg_BSUpdatePaymentDataResponse                                   = 1461
	EMsg_BSGetBillingAddress                                           = 1462
	EMsg_BSGetBillingAddressResponse                                   = 1463
	EMsg_BSGetCreditCardInfo                                           = 1464
	EMsg_BSGetCreditCardInfoResponse                                   = 1465
	EMsg_BSRemoveExpiredPaymentData                                    = 1468
	EMsg_BSRemoveExpiredPaymentDataResponse                            = 1469
	EMsg_BSConvertToCurrentKeys                                        = 1470
	EMsg_BSConvertToCurrentKeysResponse                                = 1471
	EMsg_BSInitPurchase                                                = 1472
	EMsg_BSInitPurchaseResponse                                        = 1473
	EMsg_BSCompletePurchase                                            = 1474
	EMsg_BSCompletePurchaseResponse                                    = 1475
	EMsg_BSPruneCardUsageStats                                         = 1476
	EMsg_BSPruneCardUsageStatsResponse                                 = 1477
	EMsg_BSStoreBankInformation                                        = 1478
	EMsg_BSStoreBankInformationResponse                                = 1479
	EMsg_BSVerifyPOSAKey                                               = 1480
	EMsg_BSVerifyPOSAKeyResponse                                       = 1481
	EMsg_BSReverseRedeemPOSAKey                                        = 1482
	EMsg_BSReverseRedeemPOSAKeyResponse                                = 1483
	EMsg_BSQueryFindCreditCard                                         = 1484
	EMsg_BSQueryFindCreditCardResponse                                 = 1485
	EMsg_BSStatusInquiryPOSAKey                                        = 1486
	EMsg_BSStatusInquiryPOSAKeyResponse                                = 1487
	EMsg_BSValidateMoPaySignature                                      = 1488
	EMsg_BSValidateMoPaySignatureResponse                              = 1489
	EMsg_BSMoPayConfirmProductDelivery                                 = 1490
	EMsg_BSMoPayConfirmProductDeliveryResponse                         = 1491
	EMsg_BSGenerateMoPayMD5                                            = 1492
	EMsg_BSGenerateMoPayMD5Response                                    = 1493
	EMsg_BSBoaCompraConfirmProductDelivery                             = 1494
	EMsg_BSBoaCompraConfirmProductDeliveryResponse                     = 1495
	EMsg_BSGenerateBoaCompraMD5                                        = 1496
	EMsg_BSGenerateBoaCompraMD5Response                                = 1497
	EMsg_BaseATS                                                       = 1500
	EMsg_ATSStartStressTest                                            = 1501
	EMsg_ATSStopStressTest                                             = 1502
	EMsg_ATSRunFailServerTest                                          = 1503
	EMsg_ATSUFSPerfTestTask                                            = 1504
	EMsg_ATSUFSPerfTestResponse                                        = 1505
	EMsg_ATSCycleTCM                                                   = 1506
	EMsg_ATSInitDRMSStressTest                                         = 1507
	EMsg_ATSCallTest                                                   = 1508
	EMsg_ATSCallTestReply                                              = 1509
	EMsg_ATSStartExternalStress                                        = 1510
	EMsg_ATSExternalStressJobStart                                     = 1511
	EMsg_ATSExternalStressJobQueued                                    = 1512
	EMsg_ATSExternalStressJobRunning                                   = 1513
	EMsg_ATSExternalStressJobStopped                                   = 1514
	EMsg_ATSExternalStressJobStopAll                                   = 1515
	EMsg_ATSExternalStressActionResult                                 = 1516
	EMsg_ATSStarted                                                    = 1517
	EMsg_ATSCSPerfTestTask                                             = 1518
	EMsg_ATSCSPerfTestResponse                                         = 1519
	EMsg_BaseDP                                                        = 1600
	EMsg_DPSetPublishingState                                          = 1601
	EMsg_DPGamePlayedStats                                             = 1602
	EMsg_DPUniquePlayersStat                                           = 1603
	EMsg_DPVacInfractionStats                                          = 1605
	EMsg_DPVacBanStats                                                 = 1606
	EMsg_DPBlockingStats                                               = 1607
	EMsg_DPNatTraversalStats                                           = 1608
	EMsg_DPSteamUsageEvent                                             = 1609
	EMsg_DPVacCertBanStats                                             = 1610
	EMsg_DPVacCafeBanStats                                             = 1611
	EMsg_DPCloudStats                                                  = 1612
	EMsg_DPAchievementStats                                            = 1613
	EMsg_DPAccountCreationStats                                        = 1614
	EMsg_DPGetPlayerCount                                              = 1615
	EMsg_DPGetPlayerCountResponse                                      = 1616
	EMsg_DPGameServersPlayersStats                                     = 1617
	EMsg_DPDownloadRateStatistics                                      = 1618
	EMsg_DPFacebookStatistics                                          = 1619
	EMsg_ClientDPCheckSpecialSurvey                                    = 1620
	EMsg_ClientDPCheckSpecialSurveyResponse                            = 1621
	EMsg_ClientDPSendSpecialSurveyResponse                             = 1622
	EMsg_ClientDPSendSpecialSurveyResponseReply                        = 1623
	EMsg_DPStoreSaleStatistics                                         = 1624
	EMsg_ClientDPUpdateAppJobReport                                    = 1625
	EMsg_ClientDPSteam2AppStarted                                      = 1627
	EMsg_DPUpdateContentEvent                                          = 1626
	EMsg_BaseCM                                                        = 1700
	EMsg_CMSetAllowState                                               = 1701
	EMsg_CMSpewAllowState                                              = 1702
	EMsg_CMAppInfoResponseDeprecated                                   = 1703 // Deprecated
	EMsg_BaseDSS                                                       = 1800
	EMsg_DSSNewFile                                                    = 1801
	EMsg_DSSCurrentFileList                                            = 1802
	EMsg_DSSSynchList                                                  = 1803
	EMsg_DSSSynchListResponse                                          = 1804
	EMsg_DSSSynchSubscribe                                             = 1805
	EMsg_DSSSynchUnsubscribe                                           = 1806
	EMsg_BaseEPM                                                       = 1900
	EMsg_EPMStartProcess                                               = 1901
	EMsg_EPMStopProcess                                                = 1902
	EMsg_EPMRestartProcess                                             = 1903
	EMsg_BaseGC                                                        = 2200
	EMsg_GCSendClient                                                  = 2200
	EMsg_AMRelayToGC                                                   = 2201
	EMsg_GCUpdatePlayedState                                           = 2202
	EMsg_GCCmdRevive                                                   = 2203
	EMsg_GCCmdBounce                                                   = 2204
	EMsg_GCCmdForceBounce                                              = 2205
	EMsg_GCCmdDown                                                     = 2206
	EMsg_GCCmdDeploy                                                   = 2207
	EMsg_GCCmdDeployResponse                                           = 2208
	EMsg_GCCmdSwitch                                                   = 2209
	EMsg_AMRefreshSessions                                             = 2210
	EMsg_GCUpdateGSState                                               = 2211
	EMsg_GCAchievementAwarded                                          = 2212
	EMsg_GCSystemMessage                                               = 2213
	EMsg_GCValidateSession                                             = 2214
	EMsg_GCValidateSessionResponse                                     = 2215
	EMsg_GCCmdStatus                                                   = 2216
	EMsg_GCRegisterWebInterfaces                                       = 2217 // Deprecated
	EMsg_GCRegisterWebInterfaces_Deprecated                            = 2217 // Deprecated
	EMsg_GCGetAccountDetails                                           = 2218 // Deprecated
	EMsg_GCGetAccountDetails_DEPRECATED                                = 2218 // Deprecated
	EMsg_GCInterAppMessage                                             = 2219
	EMsg_GCGetEmailTemplate                                            = 2220
	EMsg_GCGetEmailTemplateResponse                                    = 2221
	EMsg_ISRelayToGCH                                                  = 2222
	EMsg_GCHRelayClientToIS                                            = 2223
	EMsg_GCHUpdateSession                                              = 2224
	EMsg_GCHRequestUpdateSession                                       = 2225
	EMsg_GCHRequestStatus                                              = 2226
	EMsg_GCHRequestStatusResponse                                      = 2227
	EMsg_BaseP2P                                                       = 2500
	EMsg_P2PIntroducerMessage                                          = 2502
	EMsg_BaseSM                                                        = 2900
	EMsg_SMExpensiveReport                                             = 2902
	EMsg_SMHourlyReport                                                = 2903
	EMsg_SMFishingReport                                               = 2904
	EMsg_SMPartitionRenames                                            = 2905
	EMsg_SMMonitorSpace                                                = 2906
	EMsg_SMGetSchemaConversionResults                                  = 2907
	EMsg_SMGetSchemaConversionResultsResponse                          = 2908
	EMsg_BaseTest                                                      = 3000
	EMsg_FailServer                                                    = 3000
	EMsg_JobHeartbeatTest                                              = 3001
	EMsg_JobHeartbeatTestResponse                                      = 3002
	EMsg_BaseFTSRange                                                  = 3100
	EMsg_FTSGetBrowseCounts                                            = 3101
	EMsg_FTSGetBrowseCountsResponse                                    = 3102
	EMsg_FTSBrowseClans                                                = 3103
	EMsg_FTSBrowseClansResponse                                        = 3104
	EMsg_FTSSearchClansByLocation                                      = 3105
	EMsg_FTSSearchClansByLocationResponse                              = 3106
	EMsg_FTSSearchPlayersByLocation                                    = 3107
	EMsg_FTSSearchPlayersByLocationResponse                            = 3108
	EMsg_FTSClanDeleted                                                = 3109
	EMsg_FTSSearch                                                     = 3110
	EMsg_FTSSearchResponse                                             = 3111
	EMsg_FTSSearchStatus                                               = 3112
	EMsg_FTSSearchStatusResponse                                       = 3113
	EMsg_FTSGetGSPlayStats                                             = 3114
	EMsg_FTSGetGSPlayStatsResponse                                     = 3115
	EMsg_FTSGetGSPlayStatsForServer                                    = 3116
	EMsg_FTSGetGSPlayStatsForServerResponse                            = 3117
	EMsg_FTSReportIPUpdates                                            = 3118
	EMsg_BaseCCSRange                                                  = 3150
	EMsg_CCSGetComments                                                = 3151
	EMsg_CCSGetCommentsResponse                                        = 3152
	EMsg_CCSAddComment                                                 = 3153
	EMsg_CCSAddCommentResponse                                         = 3154
	EMsg_CCSDeleteComment                                              = 3155
	EMsg_CCSDeleteCommentResponse                                      = 3156
	EMsg_CCSPreloadComments                                            = 3157
	EMsg_CCSNotifyCommentCount                                         = 3158
	EMsg_CCSGetCommentsForNews                                         = 3159
	EMsg_CCSGetCommentsForNewsResponse                                 = 3160
	EMsg_CCSDeleteAllCommentsByAuthor                                  = 3161
	EMsg_CCSDeleteAllCommentsByAuthorResponse                          = 3162
	EMsg_BaseLBSRange                                                  = 3200
	EMsg_LBSSetScore                                                   = 3201
	EMsg_LBSSetScoreResponse                                           = 3202
	EMsg_LBSFindOrCreateLB                                             = 3203
	EMsg_LBSFindOrCreateLBResponse                                     = 3204
	EMsg_LBSGetLBEntries                                               = 3205
	EMsg_LBSGetLBEntriesResponse                                       = 3206
	EMsg_LBSGetLBList                                                  = 3207
	EMsg_LBSGetLBListResponse                                          = 3208
	EMsg_LBSSetLBDetails                                               = 3209
	EMsg_LBSDeleteLB                                                   = 3210
	EMsg_LBSDeleteLBEntry                                              = 3211
	EMsg_LBSResetLB                                                    = 3212
	EMsg_BaseOGS                                                       = 3400
	EMsg_OGSBeginSession                                               = 3401
	EMsg_OGSBeginSessionResponse                                       = 3402
	EMsg_OGSEndSession                                                 = 3403
	EMsg_OGSEndSessionResponse                                         = 3404
	EMsg_OGSWriteAppSessionRow                                         = 3406
	EMsg_BaseBRP                                                       = 3600
	EMsg_BRPStartShippingJobs                                          = 3601
	EMsg_BRPProcessUSBankReports                                       = 3602
	EMsg_BRPProcessGCReports                                           = 3603
	EMsg_BRPProcessPPReports                                           = 3604
	EMsg_BRPSettleNOVA                                                 = 3605
	EMsg_BRPSettleCB                                                   = 3606
	EMsg_BRPCommitGC                                                   = 3607
	EMsg_BRPCommitGCResponse                                           = 3608
	EMsg_BRPFindHungTransactions                                       = 3609
	EMsg_BRPCheckFinanceCloseOutDate                                   = 3610
	EMsg_BRPProcessLicenses                                            = 3611
	EMsg_BRPProcessLicensesResponse                                    = 3612
	EMsg_BRPRemoveExpiredPaymentData                                   = 3613
	EMsg_BRPRemoveExpiredPaymentDataResponse                           = 3614
	EMsg_BRPConvertToCurrentKeys                                       = 3615
	EMsg_BRPConvertToCurrentKeysResponse                               = 3616
	EMsg_BRPPruneCardUsageStats                                        = 3617
	EMsg_BRPPruneCardUsageStatsResponse                                = 3618
	EMsg_BRPCheckActivationCodes                                       = 3619
	EMsg_BRPCheckActivationCodesResponse                               = 3620
	EMsg_BaseAMRange2                                                  = 4000
	EMsg_AMCreateChat                                                  = 4001
	EMsg_AMCreateChatResponse                                          = 4002
	EMsg_AMUpdateChatMetadata                                          = 4003
	EMsg_AMPublishChatMetadata                                         = 4004
	EMsg_AMSetProfileURL                                               = 4005
	EMsg_AMGetAccountEmailAddress                                      = 4006
	EMsg_AMGetAccountEmailAddressResponse                              = 4007
	EMsg_AMRequestFriendData                                           = 4008
	EMsg_AMRouteToClients                                              = 4009
	EMsg_AMLeaveClan                                                   = 4010
	EMsg_AMClanPermissions                                             = 4011
	EMsg_AMClanPermissionsResponse                                     = 4012
	EMsg_AMCreateClanEvent                                             = 4013
	EMsg_AMCreateClanEventResponse                                     = 4014
	EMsg_AMUpdateClanEvent                                             = 4015
	EMsg_AMUpdateClanEventResponse                                     = 4016
	EMsg_AMGetClanEvents                                               = 4017
	EMsg_AMGetClanEventsResponse                                       = 4018
	EMsg_AMDeleteClanEvent                                             = 4019
	EMsg_AMDeleteClanEventResponse                                     = 4020
	EMsg_AMSetClanPermissionSettings                                   = 4021
	EMsg_AMSetClanPermissionSettingsResponse                           = 4022
	EMsg_AMGetClanPermissionSettings                                   = 4023
	EMsg_AMGetClanPermissionSettingsResponse                           = 4024
	EMsg_AMPublishChatRoomInfo                                         = 4025
	EMsg_ClientChatRoomInfo                                            = 4026
	EMsg_AMCreateClanAnnouncement                                      = 4027
	EMsg_AMCreateClanAnnouncementResponse                              = 4028
	EMsg_AMUpdateClanAnnouncement                                      = 4029
	EMsg_AMUpdateClanAnnouncementResponse                              = 4030
	EMsg_AMGetClanAnnouncementsCount                                   = 4031
	EMsg_AMGetClanAnnouncementsCountResponse                           = 4032
	EMsg_AMGetClanAnnouncements                                        = 4033
	EMsg_AMGetClanAnnouncementsResponse                                = 4034
	EMsg_AMDeleteClanAnnouncement                                      = 4035
	EMsg_AMDeleteClanAnnouncementResponse                              = 4036
	EMsg_AMGetSingleClanAnnouncement                                   = 4037
	EMsg_AMGetSingleClanAnnouncementResponse                           = 4038
	EMsg_AMGetClanHistory                                              = 4039
	EMsg_AMGetClanHistoryResponse                                      = 4040
	EMsg_AMGetClanPermissionBits                                       = 4041
	EMsg_AMGetClanPermissionBitsResponse                               = 4042
	EMsg_AMSetClanPermissionBits                                       = 4043
	EMsg_AMSetClanPermissionBitsResponse                               = 4044
	EMsg_AMSessionInfoRequest                                          = 4045
	EMsg_AMSessionInfoResponse                                         = 4046
	EMsg_AMValidateWGToken                                             = 4047
	EMsg_AMGetSingleClanEvent                                          = 4048
	EMsg_AMGetSingleClanEventResponse                                  = 4049
	EMsg_AMGetClanRank                                                 = 4050
	EMsg_AMGetClanRankResponse                                         = 4051
	EMsg_AMSetClanRank                                                 = 4052
	EMsg_AMSetClanRankResponse                                         = 4053
	EMsg_AMGetClanPOTW                                                 = 4054
	EMsg_AMGetClanPOTWResponse                                         = 4055
	EMsg_AMSetClanPOTW                                                 = 4056
	EMsg_AMSetClanPOTWResponse                                         = 4057
	EMsg_AMRequestChatMetadata                                         = 4058
	EMsg_AMDumpUser                                                    = 4059
	EMsg_AMKickUserFromClan                                            = 4060
	EMsg_AMAddFounderToClan                                            = 4061
	EMsg_AMValidateWGTokenResponse                                     = 4062
	EMsg_AMSetCommunityState                                           = 4063
	EMsg_AMSetAccountDetails                                           = 4064
	EMsg_AMGetChatBanList                                              = 4065
	EMsg_AMGetChatBanListResponse                                      = 4066
	EMsg_AMUnBanFromChat                                               = 4067
	EMsg_AMSetClanDetails                                              = 4068
	EMsg_AMGetAccountLinks                                             = 4069
	EMsg_AMGetAccountLinksResponse                                     = 4070
	EMsg_AMSetAccountLinks                                             = 4071
	EMsg_AMSetAccountLinksResponse                                     = 4072
	EMsg_AMGetUserGameStats                                            = 4073
	EMsg_AMGetUserGameStatsResponse                                    = 4074
	EMsg_AMCheckClanMembership                                         = 4075
	EMsg_AMGetClanMembers                                              = 4076
	EMsg_AMGetClanMembersResponse                                      = 4077
	EMsg_AMJoinPublicClan                                              = 4078
	EMsg_AMNotifyChatOfClanChange                                      = 4079
	EMsg_AMResubmitPurchase                                            = 4080
	EMsg_AMAddFriend                                                   = 4081
	EMsg_AMAddFriendResponse                                           = 4082
	EMsg_AMRemoveFriend                                                = 4083
	EMsg_AMDumpClan                                                    = 4084
	EMsg_AMChangeClanOwner                                             = 4085
	EMsg_AMCancelEasyCollect                                           = 4086
	EMsg_AMCancelEasyCollectResponse                                   = 4087
	EMsg_AMGetClanMembershipList                                       = 4088
	EMsg_AMGetClanMembershipListResponse                               = 4089
	EMsg_AMClansInCommon                                               = 4090
	EMsg_AMClansInCommonResponse                                       = 4091
	EMsg_AMIsValidAccountID                                            = 4092
	EMsg_AMConvertClan                                                 = 4093
	EMsg_AMGetGiftTargetListRelay                                      = 4094
	EMsg_AMWipeFriendsList                                             = 4095
	EMsg_AMSetIgnored                                                  = 4096
	EMsg_AMClansInCommonCountResponse                                  = 4097
	EMsg_AMFriendsList                                                 = 4098
	EMsg_AMFriendsListResponse                                         = 4099
	EMsg_AMFriendsInCommon                                             = 4100
	EMsg_AMFriendsInCommonResponse                                     = 4101
	EMsg_AMFriendsInCommonCountResponse                                = 4102
	EMsg_AMClansInCommonCount                                          = 4103
	EMsg_AMChallengeVerdict                                            = 4104
	EMsg_AMChallengeNotification                                       = 4105
	EMsg_AMFindGSByIP                                                  = 4106
	EMsg_AMFoundGSByIP                                                 = 4107
	EMsg_AMGiftRevoked                                                 = 4108
	EMsg_AMCreateAccountRecord                                         = 4109
	EMsg_AMUserClanList                                                = 4110
	EMsg_AMUserClanListResponse                                        = 4111
	EMsg_AMGetAccountDetails2                                          = 4112
	EMsg_AMGetAccountDetailsResponse2                                  = 4113
	EMsg_AMSetCommunityProfileSettings                                 = 4114
	EMsg_AMSetCommunityProfileSettingsResponse                         = 4115
	EMsg_AMGetCommunityPrivacyState                                    = 4116
	EMsg_AMGetCommunityPrivacyStateResponse                            = 4117
	EMsg_AMCheckClanInviteRateLimiting                                 = 4118
	EMsg_AMGetUserAchievementStatus                                    = 4119
	EMsg_AMGetIgnored                                                  = 4120
	EMsg_AMGetIgnoredResponse                                          = 4121
	EMsg_AMSetIgnoredResponse                                          = 4122
	EMsg_AMSetFriendRelationshipNone                                   = 4123
	EMsg_AMGetFriendRelationship                                       = 4124
	EMsg_AMGetFriendRelationshipResponse                               = 4125
	EMsg_AMServiceModulesCache                                         = 4126
	EMsg_AMServiceModulesCall                                          = 4127
	EMsg_AMServiceModulesCallResponse                                  = 4128
	EMsg_AMGetCaptchaDataForIP                                         = 4129
	EMsg_AMGetCaptchaDataForIPResponse                                 = 4130
	EMsg_AMValidateCaptchaDataForIP                                    = 4131
	EMsg_AMValidateCaptchaDataForIPResponse                            = 4132
	EMsg_AMTrackFailedAuthByIP                                         = 4133
	EMsg_AMGetCaptchaDataByGID                                         = 4134
	EMsg_AMGetCaptchaDataByGIDResponse                                 = 4135
	EMsg_AMGetLobbyList                                                = 4136
	EMsg_AMGetLobbyListResponse                                        = 4137
	EMsg_AMGetLobbyMetadata                                            = 4138
	EMsg_AMGetLobbyMetadataResponse                                    = 4139
	EMsg_CommunityAddFriendNews                                        = 4140
	EMsg_AMAddClanNews                                                 = 4141
	EMsg_AMWriteNews                                                   = 4142
	EMsg_AMFindClanUser                                                = 4143
	EMsg_AMFindClanUserResponse                                        = 4144
	EMsg_AMBanFromChat                                                 = 4145
	EMsg_AMGetUserHistoryResponse                                      = 4146
	EMsg_AMGetUserNewsSubscriptions                                    = 4147
	EMsg_AMGetUserNewsSubscriptionsResponse                            = 4148
	EMsg_AMSetUserNewsSubscriptions                                    = 4149
	EMsg_AMGetUserNews                                                 = 4150
	EMsg_AMGetUserNewsResponse                                         = 4151
	EMsg_AMSendQueuedEmails                                            = 4152
	EMsg_AMSetLicenseFlags                                             = 4153
	EMsg_AMGetUserHistory                                              = 4154
	EMsg_CommunityDeleteUserNews                                       = 4155
	EMsg_AMAllowUserFilesRequest                                       = 4156
	EMsg_AMAllowUserFilesResponse                                      = 4157
	EMsg_AMGetAccountStatus                                            = 4158
	EMsg_AMGetAccountStatusResponse                                    = 4159
	EMsg_AMEditBanReason                                               = 4160
	EMsg_AMCheckClanMembershipResponse                                 = 4161
	EMsg_AMProbeClanMembershipList                                     = 4162
	EMsg_AMProbeClanMembershipListResponse                             = 4163
	EMsg_AMGetFriendsLobbies                                           = 4165
	EMsg_AMGetFriendsLobbiesResponse                                   = 4166
	EMsg_AMGetUserFriendNewsResponse                                   = 4172
	EMsg_CommunityGetUserFriendNews                                    = 4173
	EMsg_AMGetUserClansNewsResponse                                    = 4174
	EMsg_AMGetUserClansNews                                            = 4175
	EMsg_AMStoreInitPurchase                                           = 4176
	EMsg_AMStoreInitPurchaseResponse                                   = 4177
	EMsg_AMStoreGetFinalPrice                                          = 4178
	EMsg_AMStoreGetFinalPriceResponse                                  = 4179
	EMsg_AMStoreCompletePurchase                                       = 4180
	EMsg_AMStoreCancelPurchase                                         = 4181
	EMsg_AMStorePurchaseResponse                                       = 4182
	EMsg_AMCreateAccountRecordInSteam3                                 = 4183
	EMsg_AMGetPreviousCBAccount                                        = 4184
	EMsg_AMGetPreviousCBAccountResponse                                = 4185
	EMsg_AMUpdateBillingAddress                                        = 4186
	EMsg_AMUpdateBillingAddressResponse                                = 4187
	EMsg_AMGetBillingAddress                                           = 4188
	EMsg_AMGetBillingAddressResponse                                   = 4189
	EMsg_AMGetUserLicenseHistory                                       = 4190
	EMsg_AMGetUserLicenseHistoryResponse                               = 4191
	EMsg_AMSupportChangePassword                                       = 4194
	EMsg_AMSupportChangeEmail                                          = 4195
	EMsg_AMSupportChangeSecretQA                                       = 4196
	EMsg_AMResetUserVerificationGSByIP                                 = 4197
	EMsg_AMUpdateGSPlayStats                                           = 4198
	EMsg_AMSupportEnableOrDisable                                      = 4199
	EMsg_AMGetComments                                                 = 4200
	EMsg_AMGetCommentsResponse                                         = 4201
	EMsg_AMAddComment                                                  = 4202
	EMsg_AMAddCommentResponse                                          = 4203
	EMsg_AMDeleteComment                                               = 4204
	EMsg_AMDeleteCommentResponse                                       = 4205
	EMsg_AMGetPurchaseStatus                                           = 4206
	EMsg_AMSupportIsAccountEnabled                                     = 4209
	EMsg_AMSupportIsAccountEnabledResponse                             = 4210
	EMsg_AMGetUserStats                                                = 4211
	EMsg_AMSupportKickSession                                          = 4212
	EMsg_AMGSSearch                                                    = 4213
	EMsg_MarketingMessageUpdate                                        = 4216
	EMsg_AMRouteFriendMsg                                              = 4219
	EMsg_AMTicketAuthRequestOrResponse                                 = 4220
	EMsg_AMVerifyDepotManagementRights                                 = 4222
	EMsg_AMVerifyDepotManagementRightsResponse                         = 4223
	EMsg_AMAddFreeLicense                                              = 4224
	EMsg_AMGetUserFriendsMinutesPlayed                                 = 4225
	EMsg_AMGetUserFriendsMinutesPlayedResponse                         = 4226
	EMsg_AMGetUserMinutesPlayed                                        = 4227
	EMsg_AMGetUserMinutesPlayedResponse                                = 4228
	EMsg_AMValidateEmailLink                                           = 4231
	EMsg_AMValidateEmailLinkResponse                                   = 4232
	EMsg_AMAddUsersToMarketingTreatment                                = 4234
	EMsg_AMStoreUserStats                                              = 4236
	EMsg_AMGetUserGameplayInfo                                         = 4237
	EMsg_AMGetUserGameplayInfoResponse                                 = 4238
	EMsg_AMGetCardList                                                 = 4239
	EMsg_AMGetCardListResponse                                         = 4240
	EMsg_AMDeleteStoredCard                                            = 4241
	EMsg_AMRevokeLegacyGameKeys                                        = 4242
	EMsg_AMGetWalletDetails                                            = 4244
	EMsg_AMGetWalletDetailsResponse                                    = 4245
	EMsg_AMDeleteStoredPaymentInfo                                     = 4246
	EMsg_AMGetStoredPaymentSummary                                     = 4247
	EMsg_AMGetStoredPaymentSummaryResponse                             = 4248
	EMsg_AMGetWalletConversionRate                                     = 4249
	EMsg_AMGetWalletConversionRateResponse                             = 4250
	EMsg_AMConvertWallet                                               = 4251
	EMsg_AMConvertWalletResponse                                       = 4252
	EMsg_AMRelayGetFriendsWhoPlayGame                                  = 4253
	EMsg_AMRelayGetFriendsWhoPlayGameResponse                          = 4254
	EMsg_AMSetPreApproval                                              = 4255
	EMsg_AMSetPreApprovalResponse                                      = 4256
	EMsg_AMMarketingTreatmentUpdate                                    = 4257
	EMsg_AMCreateRefund                                                = 4258
	EMsg_AMCreateRefundResponse                                        = 4259
	EMsg_AMCreateChargeback                                            = 4260
	EMsg_AMCreateChargebackResponse                                    = 4261
	EMsg_AMCreateDispute                                               = 4262
	EMsg_AMCreateDisputeResponse                                       = 4263
	EMsg_AMClearDispute                                                = 4264
	EMsg_AMClearDisputeResponse                                        = 4265
	EMsg_AMPlayerNicknameList                                          = 4266
	EMsg_AMPlayerNicknameListResponse                                  = 4267
	EMsg_AMSetDRMTestConfig                                            = 4268
	EMsg_AMGetUserCurrentGameInfo                                      = 4269
	EMsg_AMGetUserCurrentGameInfoResponse                              = 4270
	EMsg_AMGetGSPlayerList                                             = 4271
	EMsg_AMGetGSPlayerListResponse                                     = 4272
	EMsg_AMUpdatePersonaStateCache                                     = 4275
	EMsg_AMGetGameMembers                                              = 4276
	EMsg_AMGetGameMembersResponse                                      = 4277
	EMsg_AMGetSteamIDForMicroTxn                                       = 4278
	EMsg_AMGetSteamIDForMicroTxnResponse                               = 4279
	EMsg_AMAddPublisherUser                                            = 4280
	EMsg_AMRemovePublisherUser                                         = 4281
	EMsg_AMGetUserLicenseList                                          = 4282
	EMsg_AMGetUserLicenseListResponse                                  = 4283
	EMsg_AMReloadGameGroupPolicy                                       = 4284
	EMsg_AMAddFreeLicenseResponse                                      = 4285
	EMsg_AMVACStatusUpdate                                             = 4286
	EMsg_AMGetAccountDetails                                           = 4287
	EMsg_AMGetAccountDetailsResponse                                   = 4288
	EMsg_AMGetPlayerLinkDetails                                        = 4289
	EMsg_AMGetPlayerLinkDetailsResponse                                = 4290
	EMsg_AMSubscribeToPersonaFeed                                      = 4291
	EMsg_AMGetUserVacBanList                                           = 4292
	EMsg_AMGetUserVacBanListResponse                                   = 4293
	EMsg_AMGetAccountFlagsForWGSpoofing                                = 4294
	EMsg_AMGetAccountFlagsForWGSpoofingResponse                        = 4295
	EMsg_AMGetFriendsWishlistInfo                                      = 4296
	EMsg_AMGetFriendsWishlistInfoResponse                              = 4297
	EMsg_AMGetClanOfficers                                             = 4298
	EMsg_AMGetClanOfficersResponse                                     = 4299
	EMsg_AMNameChange                                                  = 4300
	EMsg_AMGetNameHistory                                              = 4301
	EMsg_AMGetNameHistoryResponse                                      = 4302
	EMsg_AMUpdateProviderStatus                                        = 4305
	EMsg_AMClearPersonaMetadataBlob                                    = 4306
	EMsg_AMSupportRemoveAccountSecurity                                = 4307
	EMsg_AMIsAccountInCaptchaGracePeriod                               = 4308
	EMsg_AMIsAccountInCaptchaGracePeriodResponse                       = 4309
	EMsg_AMAccountPS3Unlink                                            = 4310
	EMsg_AMAccountPS3UnlinkResponse                                    = 4311
	EMsg_AMStoreUserStatsResponse                                      = 4312
	EMsg_AMGetAccountPSNInfo                                           = 4313
	EMsg_AMGetAccountPSNInfoResponse                                   = 4314
	EMsg_AMAuthenticatedPlayerList                                     = 4315
	EMsg_AMGetUserGifts                                                = 4316
	EMsg_AMGetUserGiftsResponse                                        = 4317
	EMsg_AMTransferLockedGifts                                         = 4320
	EMsg_AMTransferLockedGiftsResponse                                 = 4321
	EMsg_AMPlayerHostedOnGameServer                                    = 4322
	EMsg_AMGetAccountBanInfo                                           = 4323
	EMsg_AMGetAccountBanInfoResponse                                   = 4324
	EMsg_AMRecordBanEnforcement                                        = 4325
	EMsg_AMRollbackGiftTransfer                                        = 4326
	EMsg_AMRollbackGiftTransferResponse                                = 4327
	EMsg_AMHandlePendingTransaction                                    = 4328
	EMsg_AMRequestClanDetails                                          = 4329
	EMsg_AMDeleteStoredPaypalAgreement                                 = 4330
	EMsg_AMGameServerUpdate                                            = 4331
	EMsg_AMGameServerRemove                                            = 4332
	EMsg_AMGetPaypalAgreements                                         = 4333
	EMsg_AMGetPaypalAgreementsResponse                                 = 4334
	EMsg_AMGameServerPlayerCompatibilityCheck                          = 4335
	EMsg_AMGameServerPlayerCompatibilityCheckResponse                  = 4336
	EMsg_AMRenewLicense                                                = 4337
	EMsg_AMGetAccountCommunityBanInfo                                  = 4338
	EMsg_AMGetAccountCommunityBanInfoResponse                          = 4339
	EMsg_AMGameServerAccountChangePassword                             = 4340
	EMsg_AMGameServerAccountDeleteAccount                              = 4341
	EMsg_AMRenewAgreement                                              = 4342
	EMsg_AMSendEmail                                                   = 4343
	EMsg_AMXsollaPayment                                               = 4344
	EMsg_AMXsollaPaymentResponse                                       = 4345
	EMsg_AMAcctAllowedToPurchase                                       = 4346
	EMsg_AMAcctAllowedToPurchaseResponse                               = 4347
	EMsg_AMSwapKioskDeposit                                            = 4348
	EMsg_AMSwapKioskDepositResponse                                    = 4349
	EMsg_AMSetUserGiftUnowned                                          = 4350
	EMsg_AMSetUserGiftUnownedResponse                                  = 4351
	EMsg_AMClaimUnownedUserGift                                        = 4352
	EMsg_AMClaimUnownedUserGiftResponse                                = 4353
	EMsg_AMSetClanName                                                 = 4354
	EMsg_AMSetClanNameResponse                                         = 4355
	EMsg_AMGrantCoupon                                                 = 4356
	EMsg_AMGrantCouponResponse                                         = 4357
	EMsg_AMIsPackageRestrictedInUserCountry                            = 4358
	EMsg_AMIsPackageRestrictedInUserCountryResponse                    = 4359
	EMsg_AMHandlePendingTransactionResponse                            = 4360
	EMsg_AMGrantGuestPasses2                                           = 4361
	EMsg_AMGrantGuestPasses2Response                                   = 4362
	EMsg_AMSessionQuery                                                = 4363
	EMsg_AMSessionQueryResponse                                        = 4364
	EMsg_AMGetPlayerBanDetails                                         = 4365
	EMsg_AMGetPlayerBanDetailsResponse                                 = 4366
	EMsg_AMFinalizePurchase                                            = 4367
	EMsg_AMFinalizePurchaseResponse                                    = 4368
	EMsg_AMPersonaChangeResponse                                       = 4372
	EMsg_AMGetClanDetailsForForumCreation                              = 4373
	EMsg_AMGetClanDetailsForForumCreationResponse                      = 4374
	EMsg_AMGetPendingNotificationCount                                 = 4375
	EMsg_AMGetPendingNotificationCountResponse                         = 4376
	EMsg_AMPasswordHashUpgrade                                         = 4377
	EMsg_AMMoPayPayment                                                = 4378
	EMsg_AMMoPayPaymentResponse                                        = 4379
	EMsg_AMBoaCompraPayment                                            = 4380
	EMsg_AMBoaCompraPaymentResponse                                    = 4381
	EMsg_AMExpireCaptchaByGID                                          = 4382
	EMsg_AMCompleteExternalPurchase                                    = 4383
	EMsg_AMCompleteExternalPurchaseResponse                            = 4384
	EMsg_AMResolveNegativeWalletCredits                                = 4385
	EMsg_AMResolveNegativeWalletCreditsResponse                        = 4386
	EMsg_AMPayelpPayment                                               = 4387
	EMsg_AMPayelpPaymentResponse                                       = 4388
	EMsg_AMPlayerGetClanBasicDetails                                   = 4389
	EMsg_AMPlayerGetClanBasicDetailsResponse                           = 4390
	EMsg_BasePSRange                                                   = 5000
	EMsg_PSCreateShoppingCart                                          = 5001
	EMsg_PSCreateShoppingCartResponse                                  = 5002
	EMsg_PSIsValidShoppingCart                                         = 5003
	EMsg_PSIsValidShoppingCartResponse                                 = 5004
	EMsg_PSAddPackageToShoppingCart                                    = 5005
	EMsg_PSAddPackageToShoppingCartResponse                            = 5006
	EMsg_PSRemoveLineItemFromShoppingCart                              = 5007
	EMsg_PSRemoveLineItemFromShoppingCartResponse                      = 5008
	EMsg_PSGetShoppingCartContents                                     = 5009
	EMsg_PSGetShoppingCartContentsResponse                             = 5010
	EMsg_PSAddWalletCreditToShoppingCart                               = 5011
	EMsg_PSAddWalletCreditToShoppingCartResponse                       = 5012
	EMsg_BaseUFSRange                                                  = 5200
	EMsg_ClientUFSUploadFileRequest                                    = 5202
	EMsg_ClientUFSUploadFileResponse                                   = 5203
	EMsg_ClientUFSUploadFileChunk                                      = 5204
	EMsg_ClientUFSUploadFileFinished                                   = 5205
	EMsg_ClientUFSGetFileListForApp                                    = 5206
	EMsg_ClientUFSGetFileListForAppResponse                            = 5207
	EMsg_ClientUFSDownloadRequest                                      = 5210
	EMsg_ClientUFSDownloadResponse                                     = 5211
	EMsg_ClientUFSDownloadChunk                                        = 5212
	EMsg_ClientUFSLoginRequest                                         = 5213
	EMsg_ClientUFSLoginResponse                                        = 5214
	EMsg_UFSReloadPartitionInfo                                        = 5215
	EMsg_ClientUFSTransferHeartbeat                                    = 5216
	EMsg_UFSSynchronizeFile                                            = 5217
	EMsg_UFSSynchronizeFileResponse                                    = 5218
	EMsg_ClientUFSDeleteFileRequest                                    = 5219
	EMsg_ClientUFSDeleteFileResponse                                   = 5220
	EMsg_UFSDownloadRequest                                            = 5221
	EMsg_UFSDownloadResponse                                           = 5222
	EMsg_UFSDownloadChunk                                              = 5223
	EMsg_ClientUFSGetUGCDetails                                        = 5226
	EMsg_ClientUFSGetUGCDetailsResponse                                = 5227
	EMsg_UFSUpdateFileFlags                                            = 5228
	EMsg_UFSUpdateFileFlagsResponse                                    = 5229
	EMsg_ClientUFSGetSingleFileInfo                                    = 5230
	EMsg_ClientUFSGetSingleFileInfoResponse                            = 5231
	EMsg_ClientUFSShareFile                                            = 5232
	EMsg_ClientUFSShareFileResponse                                    = 5233
	EMsg_UFSReloadAccount                                              = 5234
	EMsg_UFSReloadAccountResponse                                      = 5235
	EMsg_UFSUpdateRecordBatched                                        = 5236
	EMsg_UFSUpdateRecordBatchedResponse                                = 5237
	EMsg_UFSMigrateFile                                                = 5238
	EMsg_UFSMigrateFileResponse                                        = 5239
	EMsg_UFSGetUGCURLs                                                 = 5240
	EMsg_UFSGetUGCURLsResponse                                         = 5241
	EMsg_UFSHttpUploadFileFinishRequest                                = 5242
	EMsg_UFSHttpUploadFileFinishResponse                               = 5243
	EMsg_UFSDownloadStartRequest                                       = 5244
	EMsg_UFSDownloadStartResponse                                      = 5245
	EMsg_UFSDownloadChunkRequest                                       = 5246
	EMsg_UFSDownloadChunkResponse                                      = 5247
	EMsg_UFSDownloadFinishRequest                                      = 5248
	EMsg_UFSDownloadFinishResponse                                     = 5249
	EMsg_UFSFlushURLCache                                              = 5250
	EMsg_UFSUploadCommit                                               = 5251
	EMsg_UFSUploadCommitResponse                                       = 5252
	EMsg_BaseClient2                                                   = 5400
	EMsg_ClientRequestForgottenPasswordEmail                           = 5401
	EMsg_ClientRequestForgottenPasswordEmailResponse                   = 5402
	EMsg_ClientCreateAccountResponse                                   = 5403
	EMsg_ClientResetForgottenPassword                                  = 5404
	EMsg_ClientResetForgottenPasswordResponse                          = 5405
	EMsg_ClientCreateAccount2                                          = 5406
	EMsg_ClientInformOfResetForgottenPassword                          = 5407
	EMsg_ClientInformOfResetForgottenPasswordResponse                  = 5408
	EMsg_ClientAnonUserLogOn_Deprecated                                = 5409 // Deprecated
	EMsg_ClientGamesPlayedWithDataBlob                                 = 5410
	EMsg_ClientUpdateUserGameInfo                                      = 5411
	EMsg_ClientFileToDownload                                          = 5412
	EMsg_ClientFileToDownloadResponse                                  = 5413
	EMsg_ClientLBSSetScore                                             = 5414
	EMsg_ClientLBSSetScoreResponse                                     = 5415
	EMsg_ClientLBSFindOrCreateLB                                       = 5416
	EMsg_ClientLBSFindOrCreateLBResponse                               = 5417
	EMsg_ClientLBSGetLBEntries                                         = 5418
	EMsg_ClientLBSGetLBEntriesResponse                                 = 5419
	EMsg_ClientMarketingMessageUpdate                                  = 5420 // Deprecated
	EMsg_ClientChatDeclined                                            = 5426
	EMsg_ClientFriendMsgIncoming                                       = 5427
	EMsg_ClientAuthList_Deprecated                                     = 5428 // Deprecated
	EMsg_ClientTicketAuthComplete                                      = 5429
	EMsg_ClientIsLimitedAccount                                        = 5430
	EMsg_ClientRequestAuthList                                         = 5431
	EMsg_ClientAuthList                                                = 5432
	EMsg_ClientStat                                                    = 5433
	EMsg_ClientP2PConnectionInfo                                       = 5434
	EMsg_ClientP2PConnectionFailInfo                                   = 5435
	EMsg_ClientGetNumberOfCurrentPlayers                               = 5436
	EMsg_ClientGetNumberOfCurrentPlayersResponse                       = 5437
	EMsg_ClientGetDepotDecryptionKey                                   = 5438
	EMsg_ClientGetDepotDecryptionKeyResponse                           = 5439
	EMsg_GSPerformHardwareSurvey                                       = 5440
	EMsg_ClientGetAppBetaPasswords                                     = 5441
	EMsg_ClientGetAppBetaPasswordsResponse                             = 5442
	EMsg_ClientEnableTestLicense                                       = 5443
	EMsg_ClientEnableTestLicenseResponse                               = 5444
	EMsg_ClientDisableTestLicense                                      = 5445
	EMsg_ClientDisableTestLicenseResponse                              = 5446
	EMsg_ClientRequestValidationMail                                   = 5448
	EMsg_ClientRequestValidationMailResponse                           = 5449
	EMsg_ClientToGC                                                    = 5452
	EMsg_ClientFromGC                                                  = 5453
	EMsg_ClientRequestChangeMail                                       = 5454
	EMsg_ClientRequestChangeMailResponse                               = 5455
	EMsg_ClientEmailAddrInfo                                           = 5456
	EMsg_ClientPasswordChange3                                         = 5457
	EMsg_ClientEmailChange3                                            = 5458
	EMsg_ClientPersonalQAChange3                                       = 5459
	EMsg_ClientResetForgottenPassword3                                 = 5460
	EMsg_ClientRequestForgottenPasswordEmail3                          = 5461
	EMsg_ClientCreateAccount3                                          = 5462
	EMsg_ClientNewLoginKey                                             = 5463
	EMsg_ClientNewLoginKeyAccepted                                     = 5464
	EMsg_ClientLogOnWithHash_Deprecated                                = 5465 // Deprecated
	EMsg_ClientStoreUserStats2                                         = 5466
	EMsg_ClientStatsUpdated                                            = 5467
	EMsg_ClientActivateOEMLicense                                      = 5468
	EMsg_ClientRegisterOEMMachine                                      = 5469
	EMsg_ClientRegisterOEMMachineResponse                              = 5470
	EMsg_ClientRequestedClientStats                                    = 5480
	EMsg_ClientStat2Int32                                              = 5481
	EMsg_ClientStat2                                                   = 5482
	EMsg_ClientVerifyPassword                                          = 5483
	EMsg_ClientVerifyPasswordResponse                                  = 5484
	EMsg_ClientDRMDownloadRequest                                      = 5485
	EMsg_ClientDRMDownloadResponse                                     = 5486
	EMsg_ClientDRMFinalResult                                          = 5487
	EMsg_ClientGetFriendsWhoPlayGame                                   = 5488
	EMsg_ClientGetFriendsWhoPlayGameResponse                           = 5489
	EMsg_ClientOGSBeginSession                                         = 5490
	EMsg_ClientOGSBeginSessionResponse                                 = 5491
	EMsg_ClientOGSEndSession                                           = 5492
	EMsg_ClientOGSEndSessionResponse                                   = 5493
	EMsg_ClientOGSWriteRow                                             = 5494
	EMsg_ClientDRMTest                                                 = 5495
	EMsg_ClientDRMTestResult                                           = 5496
	EMsg_ClientServerUnavailable                                       = 5500
	EMsg_ClientServersAvailable                                        = 5501
	EMsg_ClientRegisterAuthTicketWithCM                                = 5502
	EMsg_ClientGCMsgFailed                                             = 5503
	EMsg_ClientMicroTxnAuthRequest                                     = 5504
	EMsg_ClientMicroTxnAuthorize                                       = 5505
	EMsg_ClientMicroTxnAuthorizeResponse                               = 5506
	EMsg_ClientAppMinutesPlayedData                                    = 5507
	EMsg_ClientGetMicroTxnInfo                                         = 5508
	EMsg_ClientGetMicroTxnInfoResponse                                 = 5509
	EMsg_ClientMarketingMessageUpdate2                                 = 5510
	EMsg_ClientDeregisterWithServer                                    = 5511
	EMsg_ClientSubscribeToPersonaFeed                                  = 5512
	EMsg_ClientLogon                                                   = 5514
	EMsg_ClientGetClientDetails                                        = 5515
	EMsg_ClientGetClientDetailsResponse                                = 5516
	EMsg_ClientReportOverlayDetourFailure                              = 5517
	EMsg_ClientGetClientAppList                                        = 5518
	EMsg_ClientGetClientAppListResponse                                = 5519
	EMsg_ClientInstallClientApp                                        = 5520
	EMsg_ClientInstallClientAppResponse                                = 5521
	EMsg_ClientUninstallClientApp                                      = 5522
	EMsg_ClientUninstallClientAppResponse                              = 5523
	EMsg_ClientSetClientAppUpdateState                                 = 5524
	EMsg_ClientSetClientAppUpdateStateResponse                         = 5525
	EMsg_ClientRequestEncryptedAppTicket                               = 5526
	EMsg_ClientRequestEncryptedAppTicketResponse                       = 5527
	EMsg_ClientWalletInfoUpdate                                        = 5528
	EMsg_ClientLBSSetUGC                                               = 5529
	EMsg_ClientLBSSetUGCResponse                                       = 5530
	EMsg_ClientAMGetClanOfficers                                       = 5531
	EMsg_ClientAMGetClanOfficersResponse                               = 5532
	EMsg_ClientCheckFileSignature                                      = 5533
	EMsg_ClientCheckFileSignatureResponse                              = 5534
	EMsg_ClientFriendProfileInfo                                       = 5535
	EMsg_ClientFriendProfileInfoResponse                               = 5536
	EMsg_ClientUpdateMachineAuth                                       = 5537
	EMsg_ClientUpdateMachineAuthResponse                               = 5538
	EMsg_ClientReadMachineAuth                                         = 5539
	EMsg_ClientReadMachineAuthResponse                                 = 5540
	EMsg_ClientRequestMachineAuth                                      = 5541
	EMsg_ClientRequestMachineAuthResponse                              = 5542
	EMsg_ClientScreenshotsChanged                                      = 5543
	EMsg_ClientEmailChange4                                            = 5544
	EMsg_ClientEmailChangeResponse4                                    = 5545
	EMsg_ClientGetCDNAuthToken                                         = 5546
	EMsg_ClientGetCDNAuthTokenResponse                                 = 5547
	EMsg_ClientDownloadRateStatistics                                  = 5548
	EMsg_ClientRequestAccountData                                      = 5549
	EMsg_ClientRequestAccountDataResponse                              = 5550
	EMsg_ClientResetForgottenPassword4                                 = 5551
	EMsg_ClientHideFriend                                              = 5552
	EMsg_ClientFriendsGroupsList                                       = 5553
	EMsg_ClientGetClanActivityCounts                                   = 5554
	EMsg_ClientGetClanActivityCountsResponse                           = 5555
	EMsg_ClientOGSReportString                                         = 5556
	EMsg_ClientOGSReportBug                                            = 5557
	EMsg_ClientSentLogs                                                = 5558
	EMsg_ClientLogonGameServer                                         = 5559
	EMsg_AMClientCreateFriendsGroup                                    = 5560
	EMsg_AMClientCreateFriendsGroupResponse                            = 5561
	EMsg_AMClientDeleteFriendsGroup                                    = 5562
	EMsg_AMClientDeleteFriendsGroupResponse                            = 5563
	EMsg_AMClientRenameFriendsGroup                                    = 5564
	EMsg_AMClientRenameFriendsGroupResponse                            = 5565
	EMsg_AMClientAddFriendToGroup                                      = 5566
	EMsg_AMClientAddFriendToGroupResponse                              = 5567
	EMsg_AMClientRemoveFriendFromGroup                                 = 5568
	EMsg_AMClientRemoveFriendFromGroupResponse                         = 5569
	EMsg_ClientAMGetPersonaNameHistory                                 = 5570
	EMsg_ClientAMGetPersonaNameHistoryResponse                         = 5571
	EMsg_ClientRequestFreeLicense                                      = 5572
	EMsg_ClientRequestFreeLicenseResponse                              = 5573
	EMsg_ClientDRMDownloadRequestWithCrashData                         = 5574
	EMsg_ClientAuthListAck                                             = 5575
	EMsg_ClientItemAnnouncements                                       = 5576
	EMsg_ClientRequestItemAnnouncements                                = 5577
	EMsg_ClientFriendMsgEchoToSender                                   = 5578
	EMsg_ClientChangeSteamGuardOptions                                 = 5579
	EMsg_ClientChangeSteamGuardOptionsResponse                         = 5580
	EMsg_ClientOGSGameServerPingSample                                 = 5581
	EMsg_ClientCommentNotifications                                    = 5582
	EMsg_ClientRequestCommentNotifications                             = 5583
	EMsg_ClientPersonaChangeResponse                                   = 5584
	EMsg_ClientRequestWebAPIAuthenticateUserNonce                      = 5585
	EMsg_ClientRequestWebAPIAuthenticateUserNonceResponse              = 5586
	EMsg_ClientPlayerNicknameList                                      = 5587
	EMsg_AMClientSetPlayerNickname                                     = 5588
	EMsg_AMClientSetPlayerNicknameResponse                             = 5589
	EMsg_ClientRequestOAuthTokenForApp                                 = 5590 // Deprecated
	EMsg_ClientRequestOAuthTokenForAppResponse                         = 5591 // Deprecated
	EMsg_ClientCreateAccountProto                                      = 5590
	EMsg_ClientCreateAccountProtoResponse                              = 5591
	EMsg_ClientGetNumberOfCurrentPlayersDP                             = 5592
	EMsg_ClientGetNumberOfCurrentPlayersDPResponse                     = 5593
	EMsg_ClientServiceMethod                                           = 5594
	EMsg_ClientServiceMethodResponse                                   = 5595
	EMsg_ClientFriendUserStatusPublished                               = 5596
	EMsg_ClientCurrentUIMode                                           = 5597
	EMsg_ClientVanityURLChangedNotification                            = 5598
	EMsg_ClientUserNotifications                                       = 5599
	EMsg_BaseDFS                                                       = 5600
	EMsg_DFSGetFile                                                    = 5601
	EMsg_DFSInstallLocalFile                                           = 5602
	EMsg_DFSConnection                                                 = 5603
	EMsg_DFSConnectionReply                                            = 5604
	EMsg_ClientDFSAuthenticateRequest                                  = 5605
	EMsg_ClientDFSAuthenticateResponse                                 = 5606
	EMsg_ClientDFSEndSession                                           = 5607
	EMsg_DFSPurgeFile                                                  = 5608
	EMsg_DFSRouteFile                                                  = 5609
	EMsg_DFSGetFileFromServer                                          = 5610
	EMsg_DFSAcceptedResponse                                           = 5611
	EMsg_DFSRequestPingback                                            = 5612
	EMsg_DFSRecvTransmitFile                                           = 5613
	EMsg_DFSSendTransmitFile                                           = 5614
	EMsg_DFSRequestPingback2                                           = 5615
	EMsg_DFSResponsePingback2                                          = 5616
	EMsg_ClientDFSDownloadStatus                                       = 5617
	EMsg_DFSStartTransfer                                              = 5618
	EMsg_DFSTransferComplete                                           = 5619
	EMsg_BaseMDS                                                       = 5800
	EMsg_ClientMDSLoginRequest                                         = 5801
	EMsg_ClientMDSLoginResponse                                        = 5802
	EMsg_ClientMDSUploadManifestRequest                                = 5803
	EMsg_ClientMDSUploadManifestResponse                               = 5804
	EMsg_ClientMDSTransmitManifestDataChunk                            = 5805
	EMsg_ClientMDSHeartbeat                                            = 5806
	EMsg_ClientMDSUploadDepotChunks                                    = 5807
	EMsg_ClientMDSUploadDepotChunksResponse                            = 5808
	EMsg_ClientMDSInitDepotBuildRequest                                = 5809
	EMsg_ClientMDSInitDepotBuildResponse                               = 5810
	EMsg_AMToMDSGetDepotDecryptionKey                                  = 5812
	EMsg_MDSToAMGetDepotDecryptionKeyResponse                          = 5813
	EMsg_MDSGetVersionsForDepot                                        = 5814
	EMsg_MDSGetVersionsForDepotResponse                                = 5815
	EMsg_MDSSetPublicVersionForDepot                                   = 5816 // Deprecated
	EMsg_MDSSetPublicVersionForDepotResponse                           = 5817 // Deprecated
	EMsg_ClientMDSInitWorkshopBuildRequest                             = 5816
	EMsg_ClientMDSInitWorkshopBuildResponse                            = 5817
	EMsg_ClientMDSGetDepotManifest                                     = 5818
	EMsg_ClientMDSGetDepotManifestResponse                             = 5819
	EMsg_ClientMDSGetDepotManifestChunk                                = 5820
	EMsg_ClientMDSUploadRateTest                                       = 5823
	EMsg_ClientMDSUploadRateTestResponse                               = 5824
	EMsg_MDSDownloadDepotChunksAck                                     = 5825
	EMsg_MDSContentServerStatsBroadcast                                = 5826
	EMsg_MDSContentServerConfigRequest                                 = 5827
	EMsg_MDSContentServerConfig                                        = 5828
	EMsg_MDSGetDepotManifest                                           = 5829
	EMsg_MDSGetDepotManifestResponse                                   = 5830
	EMsg_MDSGetDepotManifestChunk                                      = 5831
	EMsg_MDSGetDepotChunk                                              = 5832
	EMsg_MDSGetDepotChunkResponse                                      = 5833
	EMsg_MDSGetDepotChunkChunk                                         = 5834
	EMsg_MDSUpdateContentServerConfig                                  = 5835
	EMsg_MDSGetServerListForUser                                       = 5836
	EMsg_MDSGetServerListForUserResponse                               = 5837
	EMsg_ClientMDSRegisterAppBuild                                     = 5838
	EMsg_ClientMDSRegisterAppBuildResponse                             = 5839
	EMsg_ClientMDSSetAppBuildLive                                      = 5840
	EMsg_ClientMDSSetAppBuildLiveResponse                              = 5841
	EMsg_ClientMDSGetPrevDepotBuild                                    = 5842
	EMsg_ClientMDSGetPrevDepotBuildResponse                            = 5843
	EMsg_MDSToCSFlushChunk                                             = 5844
	EMsg_ClientMDSSignInstallScript                                    = 5845
	EMsg_ClientMDSSignInstallScriptResponse                            = 5846
	EMsg_CSBase                                                        = 6200
	EMsg_CSPing                                                        = 6201
	EMsg_CSPingResponse                                                = 6202
	EMsg_GMSBase                                                       = 6400
	EMsg_GMSGameServerReplicate                                        = 6401
	EMsg_ClientGMSServerQuery                                          = 6403
	EMsg_GMSClientServerQueryResponse                                  = 6404
	EMsg_AMGMSGameServerUpdate                                         = 6405
	EMsg_AMGMSGameServerRemove                                         = 6406
	EMsg_GameServerOutOfDate                                           = 6407
	EMsg_ClientAuthorizeLocalDeviceRequest                             = 6501
	EMsg_ClientAuthorizeLocalDevice                                    = 6502
	EMsg_ClientDeauthorizeDeviceRequest                                = 6503
	EMsg_ClientDeauthorizeDevice                                       = 6504
	EMsg_ClientUseLocalDeviceAuthorizations                            = 6505
	EMsg_ClientGetAuthorizedDevices                                    = 6506
	EMsg_ClientGetAuthorizedDevicesResponse                            = 6507
	EMsg_MMSBase                                                       = 6600
	EMsg_ClientMMSCreateLobby                                          = 6601
	EMsg_ClientMMSCreateLobbyResponse                                  = 6602
	EMsg_ClientMMSJoinLobby                                            = 6603
	EMsg_ClientMMSJoinLobbyResponse                                    = 6604
	EMsg_ClientMMSLeaveLobby                                           = 6605
	EMsg_ClientMMSLeaveLobbyResponse                                   = 6606
	EMsg_ClientMMSGetLobbyList                                         = 6607
	EMsg_ClientMMSGetLobbyListResponse                                 = 6608
	EMsg_ClientMMSSetLobbyData                                         = 6609
	EMsg_ClientMMSSetLobbyDataResponse                                 = 6610
	EMsg_ClientMMSGetLobbyData                                         = 6611
	EMsg_ClientMMSLobbyData                                            = 6612
	EMsg_ClientMMSSendLobbyChatMsg                                     = 6613
	EMsg_ClientMMSLobbyChatMsg                                         = 6614
	EMsg_ClientMMSSetLobbyOwner                                        = 6615
	EMsg_ClientMMSSetLobbyOwnerResponse                                = 6616
	EMsg_ClientMMSSetLobbyGameServer                                   = 6617
	EMsg_ClientMMSLobbyGameServerSet                                   = 6618
	EMsg_ClientMMSUserJoinedLobby                                      = 6619
	EMsg_ClientMMSUserLeftLobby                                        = 6620
	EMsg_ClientMMSInviteToLobby                                        = 6621
	EMsg_ClientMMSFlushFrenemyListCache                                = 6622
	EMsg_ClientMMSFlushFrenemyListCacheResponse                        = 6623
	EMsg_ClientMMSSetLobbyLinked                                       = 6624
	EMsg_NonStdMsgBase                                                 = 6800
	EMsg_NonStdMsgMemcached                                            = 6801
	EMsg_NonStdMsgHTTPServer                                           = 6802
	EMsg_NonStdMsgHTTPClient                                           = 6803
	EMsg_NonStdMsgWGResponse                                           = 6804
	EMsg_NonStdMsgPHPSimulator                                         = 6805
	EMsg_NonStdMsgChase                                                = 6806
	EMsg_NonStdMsgDFSTransfer                                          = 6807
	EMsg_NonStdMsgTests                                                = 6808
	EMsg_NonStdMsgUMQpipeAAPL                                          = 6809
	EMsg_NonStdMsgSyslog                                               = 6810
	EMsg_NonStdMsgLogsink                                              = 6811
	EMsg_UDSBase                                                       = 7000
	EMsg_ClientUDSP2PSessionStarted                                    = 7001
	EMsg_ClientUDSP2PSessionEnded                                      = 7002
	EMsg_UDSRenderUserAuth                                             = 7003
	EMsg_UDSRenderUserAuthResponse                                     = 7004
	EMsg_ClientUDSInviteToGame                                         = 7005
	EMsg_UDSFindSession                                                = 7006
	EMsg_UDSFindSessionResponse                                        = 7007
	EMsg_MPASBase                                                      = 7100
	EMsg_MPASVacBanReset                                               = 7101
	EMsg_KGSBase                                                       = 7200
	EMsg_KGSAllocateKeyRange                                           = 7201
	EMsg_KGSAllocateKeyRangeResponse                                   = 7202
	EMsg_KGSGenerateKeys                                               = 7203
	EMsg_KGSGenerateKeysResponse                                       = 7204
	EMsg_KGSRemapKeys                                                  = 7205
	EMsg_KGSRemapKeysResponse                                          = 7206
	EMsg_KGSGenerateGameStopWCKeys                                     = 7207
	EMsg_KGSGenerateGameStopWCKeysResponse                             = 7208
	EMsg_UCMBase                                                       = 7300
	EMsg_ClientUCMAddScreenshot                                        = 7301
	EMsg_ClientUCMAddScreenshotResponse                                = 7302
	EMsg_UCMValidateObjectExists                                       = 7303
	EMsg_UCMValidateObjectExistsResponse                               = 7304
	EMsg_UCMResetCommunityContent                                      = 7307
	EMsg_UCMResetCommunityContentResponse                              = 7308
	EMsg_ClientUCMDeleteScreenshot                                     = 7309
	EMsg_ClientUCMDeleteScreenshotResponse                             = 7310
	EMsg_ClientUCMPublishFile                                          = 7311
	EMsg_ClientUCMPublishFileResponse                                  = 7312
	EMsg_ClientUCMGetPublishedFileDetails                              = 7313 // Deprecated
	EMsg_ClientUCMGetPublishedFileDetailsResponse                      = 7314 // Deprecated
	EMsg_ClientUCMDeletePublishedFile                                  = 7315
	EMsg_ClientUCMDeletePublishedFileResponse                          = 7316
	EMsg_ClientUCMEnumerateUserPublishedFiles                          = 7317
	EMsg_ClientUCMEnumerateUserPublishedFilesResponse                  = 7318
	EMsg_ClientUCMSubscribePublishedFile                               = 7319
	EMsg_ClientUCMSubscribePublishedFileResponse                       = 7320
	EMsg_ClientUCMEnumerateUserSubscribedFiles                         = 7321
	EMsg_ClientUCMEnumerateUserSubscribedFilesResponse                 = 7322
	EMsg_ClientUCMUnsubscribePublishedFile                             = 7323
	EMsg_ClientUCMUnsubscribePublishedFileResponse                     = 7324
	EMsg_ClientUCMUpdatePublishedFile                                  = 7325
	EMsg_ClientUCMUpdatePublishedFileResponse                          = 7326
	EMsg_UCMUpdatePublishedFile                                        = 7327
	EMsg_UCMUpdatePublishedFileResponse                                = 7328
	EMsg_UCMDeletePublishedFile                                        = 7329
	EMsg_UCMDeletePublishedFileResponse                                = 7330
	EMsg_UCMUpdatePublishedFileStat                                    = 7331
	EMsg_UCMUpdatePublishedFileBan                                     = 7332
	EMsg_UCMUpdatePublishedFileBanResponse                             = 7333
	EMsg_UCMUpdateTaggedScreenshot                                     = 7334
	EMsg_UCMAddTaggedScreenshot                                        = 7335
	EMsg_UCMRemoveTaggedScreenshot                                     = 7336
	EMsg_UCMReloadPublishedFile                                        = 7337
	EMsg_UCMReloadUserFileListCaches                                   = 7338
	EMsg_UCMPublishedFileReported                                      = 7339
	EMsg_UCMUpdatePublishedFileIncompatibleStatus                      = 7340
	EMsg_UCMPublishedFilePreviewAdd                                    = 7341
	EMsg_UCMPublishedFilePreviewAddResponse                            = 7342
	EMsg_UCMPublishedFilePreviewRemove                                 = 7343
	EMsg_UCMPublishedFilePreviewRemoveResponse                         = 7344
	EMsg_UCMPublishedFilePreviewChangeSortOrder                        = 7345
	EMsg_UCMPublishedFilePreviewChangeSortOrderResponse                = 7346
	EMsg_ClientUCMPublishedFileSubscribed                              = 7347
	EMsg_ClientUCMPublishedFileUnsubscribed                            = 7348
	EMsg_UCMPublishedFileSubscribed                                    = 7349
	EMsg_UCMPublishedFileUnsubscribed                                  = 7350
	EMsg_UCMPublishFile                                                = 7351
	EMsg_UCMPublishFileResponse                                        = 7352
	EMsg_UCMPublishedFileChildAdd                                      = 7353
	EMsg_UCMPublishedFileChildAddResponse                              = 7354
	EMsg_UCMPublishedFileChildRemove                                   = 7355
	EMsg_UCMPublishedFileChildRemoveResponse                           = 7356
	EMsg_UCMPublishedFileChildChangeSortOrder                          = 7357
	EMsg_UCMPublishedFileChildChangeSortOrderResponse                  = 7358
	EMsg_UCMPublishedFileParentChanged                                 = 7359
	EMsg_ClientUCMGetPublishedFilesForUser                             = 7360
	EMsg_ClientUCMGetPublishedFilesForUserResponse                     = 7361
	EMsg_UCMGetPublishedFilesForUser                                   = 7362
	EMsg_UCMGetPublishedFilesForUserResponse                           = 7363
	EMsg_ClientUCMSetUserPublishedFileAction                           = 7364
	EMsg_ClientUCMSetUserPublishedFileActionResponse                   = 7365
	EMsg_ClientUCMEnumeratePublishedFilesByUserAction                  = 7366
	EMsg_ClientUCMEnumeratePublishedFilesByUserActionResponse          = 7367
	EMsg_ClientUCMPublishedFileDeleted                                 = 7368
	EMsg_UCMGetUserSubscribedFiles                                     = 7369
	EMsg_UCMGetUserSubscribedFilesResponse                             = 7370
	EMsg_UCMFixStatsPublishedFile                                      = 7371
	EMsg_UCMDeleteOldScreenshot                                        = 7372
	EMsg_UCMDeleteOldScreenshotResponse                                = 7373
	EMsg_UCMDeleteOldVideo                                             = 7374
	EMsg_UCMDeleteOldVideoResponse                                     = 7375
	EMsg_UCMUpdateOldScreenshotPrivacy                                 = 7376
	EMsg_UCMUpdateOldScreenshotPrivacyResponse                         = 7377
	EMsg_ClientUCMEnumerateUserSubscribedFilesWithUpdates              = 7378
	EMsg_ClientUCMEnumerateUserSubscribedFilesWithUpdatesResponse      = 7379
	EMsg_UCMPublishedFileContentUpdated                                = 7380
	EMsg_UCMPublishedFileUpdated                                       = 7381
	EMsg_FSBase                                                        = 7500
	EMsg_ClientRichPresenceUpload                                      = 7501
	EMsg_ClientRichPresenceRequest                                     = 7502
	EMsg_ClientRichPresenceInfo                                        = 7503
	EMsg_FSRichPresenceRequest                                         = 7504
	EMsg_FSRichPresenceResponse                                        = 7505
	EMsg_FSComputeFrenematrix                                          = 7506
	EMsg_FSComputeFrenematrixResponse                                  = 7507
	EMsg_FSPlayStatusNotification                                      = 7508
	EMsg_FSPublishPersonaStatus                                        = 7509
	EMsg_FSAddOrRemoveFollower                                         = 7510
	EMsg_FSAddOrRemoveFollowerResponse                                 = 7511
	EMsg_FSUpdateFollowingList                                         = 7512
	EMsg_FSCommentNotification                                         = 7513
	EMsg_FSCommentNotificationViewed                                   = 7514
	EMsg_ClientFSGetFollowerCount                                      = 7515
	EMsg_ClientFSGetFollowerCountResponse                              = 7516
	EMsg_ClientFSGetIsFollowing                                        = 7517
	EMsg_ClientFSGetIsFollowingResponse                                = 7518
	EMsg_ClientFSEnumerateFollowingList                                = 7519
	EMsg_ClientFSEnumerateFollowingListResponse                        = 7520
	EMsg_FSGetPendingNotificationCount                                 = 7521
	EMsg_FSGetPendingNotificationCountResponse                         = 7522
	EMsg_ClientFSOfflineMessageNotification                            = 7523
	EMsg_ClientFSRequestOfflineMessageCount                            = 7524
	EMsg_ClientFSGetFriendMessageHistory                               = 7525
	EMsg_ClientFSGetFriendMessageHistoryResponse                       = 7526
	EMsg_ClientFSGetFriendMessageHistoryForOfflineMessages             = 7527
	EMsg_ClientFSGetFriendsSteamLevels                                 = 7528
	EMsg_ClientFSGetFriendsSteamLevelsResponse                         = 7529
	EMsg_DRMRange2                                                     = 7600
	EMsg_CEGVersionSetEnableDisableRequest                             = 7600
	EMsg_CEGVersionSetEnableDisableResponse                            = 7601
	EMsg_CEGPropStatusDRMSRequest                                      = 7602
	EMsg_CEGPropStatusDRMSResponse                                     = 7603
	EMsg_CEGWhackFailureReportRequest                                  = 7604
	EMsg_CEGWhackFailureReportResponse                                 = 7605
	EMsg_DRMSFetchVersionSet                                           = 7606
	EMsg_DRMSFetchVersionSetResponse                                   = 7607
	EMsg_EconBase                                                      = 7700
	EMsg_EconTrading_InitiateTradeRequest                              = 7701
	EMsg_EconTrading_InitiateTradeProposed                             = 7702
	EMsg_EconTrading_InitiateTradeResponse                             = 7703
	EMsg_EconTrading_InitiateTradeResult                               = 7704
	EMsg_EconTrading_StartSession                                      = 7705
	EMsg_EconTrading_CancelTradeRequest                                = 7706
	EMsg_EconFlushInventoryCache                                       = 7707
	EMsg_EconFlushInventoryCacheResponse                               = 7708
	EMsg_EconCDKeyProcessTransaction                                   = 7711
	EMsg_EconCDKeyProcessTransactionResponse                           = 7712
	EMsg_EconGetErrorLogs                                              = 7713
	EMsg_EconGetErrorLogsResponse                                      = 7714
	EMsg_RMRange                                                       = 7800
	EMsg_RMTestVerisignOTP                                             = 7800
	EMsg_RMTestVerisignOTPResponse                                     = 7801
	EMsg_RMDeleteMemcachedKeys                                         = 7803
	EMsg_RMRemoteInvoke                                                = 7804
	EMsg_BadLoginIPList                                                = 7805
	EMsg_UGSBase                                                       = 7900
	EMsg_UGSUpdateGlobalStats                                          = 7900
	EMsg_ClientUGSGetGlobalStats                                       = 7901
	EMsg_ClientUGSGetGlobalStatsResponse                               = 7902
	EMsg_StoreBase                                                     = 8000
	EMsg_StoreUpdateRecommendationCount                                = 8000
	EMsg_UMQBase                                                       = 8100
	EMsg_UMQLogonRequest                                               = 8100
	EMsg_UMQLogonResponse                                              = 8101
	EMsg_UMQLogoffRequest                                              = 8102
	EMsg_UMQLogoffResponse                                             = 8103
	EMsg_UMQSendChatMessage                                            = 8104
	EMsg_UMQIncomingChatMessage                                        = 8105
	EMsg_UMQPoll                                                       = 8106
	EMsg_UMQPollResults                                                = 8107
	EMsg_UMQ2AM_ClientMsgBatch                                         = 8108
	EMsg_UMQEnqueueMobileSalePromotions                                = 8109
	EMsg_UMQEnqueueMobileAnnouncements                                 = 8110
	EMsg_WorkshopBase                                                  = 8200
	EMsg_WorkshopAcceptTOSRequest                                      = 8200
	EMsg_WorkshopAcceptTOSResponse                                     = 8201
	EMsg_WebAPIBase                                                    = 8300
	EMsg_WebAPIValidateOAuth2Token                                     = 8300
	EMsg_WebAPIValidateOAuth2TokenResponse                             = 8301
	EMsg_WebAPIInvalidateTokensForAccount                              = 8302
	EMsg_WebAPIRegisterGCInterfaces                                    = 8303
	EMsg_WebAPIInvalidateOAuthClientCache                              = 8304
	EMsg_WebAPIInvalidateOAuthTokenCache                               = 8305
	EMsg_BackpackBase                                                  = 8400
	EMsg_BackpackAddToCurrency                                         = 8401
	EMsg_BackpackAddToCurrencyResponse                                 = 8402
	EMsg_CREBase                                                       = 8500
	EMsg_CRERankByTrend                                                = 8501
	EMsg_CRERankByTrendResponse                                        = 8502
	EMsg_CREItemVoteSummary                                            = 8503
	EMsg_CREItemVoteSummaryResponse                                    = 8504
	EMsg_CRERankByVote                                                 = 8505
	EMsg_CRERankByVoteResponse                                         = 8506
	EMsg_CREUpdateUserPublishedItemVote                                = 8507
	EMsg_CREUpdateUserPublishedItemVoteResponse                        = 8508
	EMsg_CREGetUserPublishedItemVoteDetails                            = 8509
	EMsg_CREGetUserPublishedItemVoteDetailsResponse                    = 8510
	EMsg_CREEnumeratePublishedFiles                                    = 8511
	EMsg_CREEnumeratePublishedFilesResponse                            = 8512
	EMsg_CREPublishedFileVoteAdded                                     = 8513
	EMsg_SecretsBase                                                   = 8600
	EMsg_SecretsRequestCredentialPair                                  = 8600
	EMsg_SecretsCredentialPairResponse                                 = 8601
	EMsg_SecretsRequestServerIdentity                                  = 8602
	EMsg_SecretsServerIdentityResponse                                 = 8603
	EMsg_SecretsUpdateServerIdentities                                 = 8604
	EMsg_BoxMonitorBase                                                = 8700
	EMsg_BoxMonitorReportRequest                                       = 8700
	EMsg_BoxMonitorReportResponse                                      = 8701
	EMsg_LogsinkBase                                                   = 8800
	EMsg_LogsinkWriteReport                                            = 8800
	EMsg_PICSBase                                                      = 8900
	EMsg_ClientPICSChangesSinceRequest                                 = 8901
	EMsg_ClientPICSChangesSinceResponse                                = 8902
	EMsg_ClientPICSProductInfoRequest                                  = 8903
	EMsg_ClientPICSProductInfoResponse                                 = 8904
	EMsg_ClientPICSAccessTokenRequest                                  = 8905
	EMsg_ClientPICSAccessTokenResponse                                 = 8906
	EMsg_WorkerProcess                                                 = 9000
	EMsg_WorkerProcessPingRequest                                      = 9000
	EMsg_WorkerProcessPingResponse                                     = 9001
	EMsg_WorkerProcessShutdown                                         = 9002
	EMsg_DRMWorkerProcess                                              = 9100
	EMsg_DRMWorkerProcessDRMAndSign                                    = 9100
	EMsg_DRMWorkerProcessDRMAndSignResponse                            = 9101
	EMsg_DRMWorkerProcessSteamworksInfoRequest                         = 9102
	EMsg_DRMWorkerProcessSteamworksInfoResponse                        = 9103
	EMsg_DRMWorkerProcessInstallDRMDLLRequest                          = 9104
	EMsg_DRMWorkerProcessInstallDRMDLLResponse                         = 9105
	EMsg_DRMWorkerProcessSecretIdStringRequest                         = 9106
	EMsg_DRMWorkerProcessSecretIdStringResponse                        = 9107
	EMsg_DRMWorkerProcessGetDRMGuidsFromFileRequest                    = 9108
	EMsg_DRMWorkerProcessGetDRMGuidsFromFileResponse                   = 9109
	EMsg_DRMWorkerProcessInstallProcessedFilesRequest                  = 9110
	EMsg_DRMWorkerProcessInstallProcessedFilesResponse                 = 9111
	EMsg_DRMWorkerProcessExamineBlobRequest                            = 9112
	EMsg_DRMWorkerProcessExamineBlobResponse                           = 9113
	EMsg_DRMWorkerProcessDescribeSecretRequest                         = 9114
	EMsg_DRMWorkerProcessDescribeSecretResponse                        = 9115
	EMsg_DRMWorkerProcessBackfillOriginalRequest                       = 9116
	EMsg_DRMWorkerProcessBackfillOriginalResponse                      = 9117
	EMsg_DRMWorkerProcessValidateDRMDLLRequest                         = 9118
	EMsg_DRMWorkerProcessValidateDRMDLLResponse                        = 9119
	EMsg_DRMWorkerProcessValidateFileRequest                           = 9120
	EMsg_DRMWorkerProcessValidateFileResponse                          = 9121
	EMsg_DRMWorkerProcessSplitAndInstallRequest                        = 9122
	EMsg_DRMWorkerProcessSplitAndInstallResponse                       = 9123
	EMsg_DRMWorkerProcessGetBlobRequest                                = 9124
	EMsg_DRMWorkerProcessGetBlobResponse                               = 9125
	EMsg_DRMWorkerProcessEvaluateCrashRequest                          = 9126
	EMsg_DRMWorkerProcessEvaluateCrashResponse                         = 9127
	EMsg_DRMWorkerProcessAnalyzeFileRequest                            = 9128
	EMsg_DRMWorkerProcessAnalyzeFileResponse                           = 9129
	EMsg_DRMWorkerProcessUnpackBlobRequest                             = 9130
	EMsg_DRMWorkerProcessUnpackBlobResponse                            = 9131
	EMsg_DRMWorkerProcessInstallAllRequest                             = 9132
	EMsg_DRMWorkerProcessInstallAllResponse                            = 9133
	EMsg_TestWorkerProcess                                             = 9200
	EMsg_TestWorkerProcessLoadUnloadModuleRequest                      = 9200
	EMsg_TestWorkerProcessLoadUnloadModuleResponse                     = 9201
	EMsg_TestWorkerProcessServiceModuleCallRequest                     = 9202
	EMsg_TestWorkerProcessServiceModuleCallResponse                    = 9203
	EMsg_ClientGetEmoticonList                                         = 9330
	EMsg_ClientEmoticonList                                            = 9331
	EMsg_ClientSharedLibraryBase                                       = 9400
	EMsg_ClientSharedLicensesLockStatus                                = 9403 // Deprecated
	EMsg_ClientSharedLicensesStopPlaying                               = 9404 // Deprecated
	EMsg_ClientSharedLibraryLockStatus                                 = 9405
	EMsg_ClientSharedLibraryStopPlaying                                = 9406
	EMsg_ClientUnlockStreaming                                         = 9507
	EMsg_ClientUnlockStreamingResponse                                 = 9508
	EMsg_ClientPlayingSessionState                                     = 9600
	EMsg_ClientKickPlayingSession                                      = 9601
)

func (e EMsg) String() string {
	switch e {
	case EMsg_Invalid:
		return "EMsg_Invalid"
	case EMsg_Multi:
		return "EMsg_Multi"
	case EMsg_BaseGeneral:
		return "EMsg_BaseGeneral"
	case EMsg_DestJobFailed:
		return "EMsg_DestJobFailed"
	case EMsg_Alert:
		return "EMsg_Alert"
	case EMsg_SCIDRequest:
		return "EMsg_SCIDRequest"
	case EMsg_SCIDResponse:
		return "EMsg_SCIDResponse"
	case EMsg_JobHeartbeat:
		return "EMsg_JobHeartbeat"
	case EMsg_HubConnect:
		return "EMsg_HubConnect"
	case EMsg_Subscribe:
		return "EMsg_Subscribe"
	case EMsg_RouteMessage:
		return "EMsg_RouteMessage"
	case EMsg_RemoteSysID:
		return "EMsg_RemoteSysID"
	case EMsg_AMCreateAccountResponse:
		return "EMsg_AMCreateAccountResponse"
	case EMsg_WGRequest:
		return "EMsg_WGRequest"
	case EMsg_WGResponse:
		return "EMsg_WGResponse"
	case EMsg_KeepAlive:
		return "EMsg_KeepAlive"
	case EMsg_WebAPIJobRequest:
		return "EMsg_WebAPIJobRequest"
	case EMsg_WebAPIJobResponse:
		return "EMsg_WebAPIJobResponse"
	case EMsg_ClientSessionStart:
		return "EMsg_ClientSessionStart"
	case EMsg_ClientSessionEnd:
		return "EMsg_ClientSessionEnd"
	case EMsg_ClientSessionUpdateAuthTicket:
		return "EMsg_ClientSessionUpdateAuthTicket"
	case EMsg_StatsDeprecated:
		return "EMsg_StatsDeprecated"
	case EMsg_Ping:
		return "EMsg_Ping"
	case EMsg_PingResponse:
		return "EMsg_PingResponse"
	case EMsg_Stats:
		return "EMsg_Stats"
	case EMsg_RequestFullStatsBlock:
		return "EMsg_RequestFullStatsBlock"
	case EMsg_LoadDBOCacheItem:
		return "EMsg_LoadDBOCacheItem"
	case EMsg_LoadDBOCacheItemResponse:
		return "EMsg_LoadDBOCacheItemResponse"
	case EMsg_InvalidateDBOCacheItems:
		return "EMsg_InvalidateDBOCacheItems"
	case EMsg_ServiceMethod:
		return "EMsg_ServiceMethod"
	case EMsg_ServiceMethodResponse:
		return "EMsg_ServiceMethodResponse"
	case EMsg_BaseShell:
		return "EMsg_BaseShell"
	case EMsg_Exit:
		return "EMsg_Exit"
	case EMsg_DirRequest:
		return "EMsg_DirRequest"
	case EMsg_DirResponse:
		return "EMsg_DirResponse"
	case EMsg_ZipRequest:
		return "EMsg_ZipRequest"
	case EMsg_ZipResponse:
		return "EMsg_ZipResponse"
	case EMsg_UpdateRecordResponse:
		return "EMsg_UpdateRecordResponse"
	case EMsg_UpdateCreditCardRequest:
		return "EMsg_UpdateCreditCardRequest"
	case EMsg_UpdateUserBanResponse:
		return "EMsg_UpdateUserBanResponse"
	case EMsg_PrepareToExit:
		return "EMsg_PrepareToExit"
	case EMsg_ContentDescriptionUpdate:
		return "EMsg_ContentDescriptionUpdate"
	case EMsg_TestResetServer:
		return "EMsg_TestResetServer"
	case EMsg_UniverseChanged:
		return "EMsg_UniverseChanged"
	case EMsg_ShellConfigInfoUpdate:
		return "EMsg_ShellConfigInfoUpdate"
	case EMsg_RequestWindowsEventLogEntries:
		return "EMsg_RequestWindowsEventLogEntries"
	case EMsg_ProvideWindowsEventLogEntries:
		return "EMsg_ProvideWindowsEventLogEntries"
	case EMsg_ShellSearchLogs:
		return "EMsg_ShellSearchLogs"
	case EMsg_ShellSearchLogsResponse:
		return "EMsg_ShellSearchLogsResponse"
	case EMsg_ShellCheckWindowsUpdates:
		return "EMsg_ShellCheckWindowsUpdates"
	case EMsg_ShellCheckWindowsUpdatesResponse:
		return "EMsg_ShellCheckWindowsUpdatesResponse"
	case EMsg_ShellFlushUserLicenseCache:
		return "EMsg_ShellFlushUserLicenseCache"
	case EMsg_BaseGM:
		return "EMsg_BaseGM"
	case EMsg_ShellFailed:
		return "EMsg_ShellFailed"
	case EMsg_ExitShells:
		return "EMsg_ExitShells"
	case EMsg_ExitShell:
		return "EMsg_ExitShell"
	case EMsg_GracefulExitShell:
		return "EMsg_GracefulExitShell"
	case EMsg_NotifyWatchdog:
		return "EMsg_NotifyWatchdog"
	case EMsg_LicenseProcessingComplete:
		return "EMsg_LicenseProcessingComplete"
	case EMsg_SetTestFlag:
		return "EMsg_SetTestFlag"
	case EMsg_QueuedEmailsComplete:
		return "EMsg_QueuedEmailsComplete"
	case EMsg_GMReportPHPError:
		return "EMsg_GMReportPHPError"
	case EMsg_GMDRMSync:
		return "EMsg_GMDRMSync"
	case EMsg_PhysicalBoxInventory:
		return "EMsg_PhysicalBoxInventory"
	case EMsg_UpdateConfigFile:
		return "EMsg_UpdateConfigFile"
	case EMsg_TestInitDB:
		return "EMsg_TestInitDB"
	case EMsg_GMWriteConfigToSQL:
		return "EMsg_GMWriteConfigToSQL"
	case EMsg_GMLoadActivationCodes:
		return "EMsg_GMLoadActivationCodes"
	case EMsg_GMQueueForFBS:
		return "EMsg_GMQueueForFBS"
	case EMsg_GMSchemaConversionResults:
		return "EMsg_GMSchemaConversionResults"
	case EMsg_GMSchemaConversionResultsResponse:
		return "EMsg_GMSchemaConversionResultsResponse"
	case EMsg_GMWriteShellFailureToSQL:
		return "EMsg_GMWriteShellFailureToSQL"
	case EMsg_BaseAIS:
		return "EMsg_BaseAIS"
	case EMsg_AISRefreshContentDescription:
		return "EMsg_AISRefreshContentDescription"
	case EMsg_AISRequestContentDescription:
		return "EMsg_AISRequestContentDescription"
	case EMsg_AISUpdateAppInfo:
		return "EMsg_AISUpdateAppInfo"
	case EMsg_AISUpdatePackageInfo:
		return "EMsg_AISUpdatePackageInfo"
	case EMsg_AISGetPackageChangeNumber:
		return "EMsg_AISGetPackageChangeNumber"
	case EMsg_AISGetPackageChangeNumberResponse:
		return "EMsg_AISGetPackageChangeNumberResponse"
	case EMsg_AISAppInfoTableChanged:
		return "EMsg_AISAppInfoTableChanged"
	case EMsg_AISUpdatePackageInfoResponse:
		return "EMsg_AISUpdatePackageInfoResponse"
	case EMsg_AISCreateMarketingMessage:
		return "EMsg_AISCreateMarketingMessage"
	case EMsg_AISCreateMarketingMessageResponse:
		return "EMsg_AISCreateMarketingMessageResponse"
	case EMsg_AISGetMarketingMessage:
		return "EMsg_AISGetMarketingMessage"
	case EMsg_AISGetMarketingMessageResponse:
		return "EMsg_AISGetMarketingMessageResponse"
	case EMsg_AISUpdateMarketingMessage:
		return "EMsg_AISUpdateMarketingMessage"
	case EMsg_AISUpdateMarketingMessageResponse:
		return "EMsg_AISUpdateMarketingMessageResponse"
	case EMsg_AISRequestMarketingMessageUpdate:
		return "EMsg_AISRequestMarketingMessageUpdate"
	case EMsg_AISDeleteMarketingMessage:
		return "EMsg_AISDeleteMarketingMessage"
	case EMsg_AISGetMarketingTreatments:
		return "EMsg_AISGetMarketingTreatments"
	case EMsg_AISGetMarketingTreatmentsResponse:
		return "EMsg_AISGetMarketingTreatmentsResponse"
	case EMsg_AISRequestMarketingTreatmentUpdate:
		return "EMsg_AISRequestMarketingTreatmentUpdate"
	case EMsg_AISTestAddPackage:
		return "EMsg_AISTestAddPackage"
	case EMsg_AIGetAppGCFlags:
		return "EMsg_AIGetAppGCFlags"
	case EMsg_AIGetAppGCFlagsResponse:
		return "EMsg_AIGetAppGCFlagsResponse"
	case EMsg_AIGetAppList:
		return "EMsg_AIGetAppList"
	case EMsg_AIGetAppListResponse:
		return "EMsg_AIGetAppListResponse"
	case EMsg_AIGetAppInfo:
		return "EMsg_AIGetAppInfo"
	case EMsg_AIGetAppInfoResponse:
		return "EMsg_AIGetAppInfoResponse"
	case EMsg_AISGetCouponDefinition:
		return "EMsg_AISGetCouponDefinition"
	case EMsg_AISGetCouponDefinitionResponse:
		return "EMsg_AISGetCouponDefinitionResponse"
	case EMsg_BaseAM:
		return "EMsg_BaseAM"
	case EMsg_AMUpdateUserBanRequest:
		return "EMsg_AMUpdateUserBanRequest"
	case EMsg_AMAddLicense:
		return "EMsg_AMAddLicense"
	case EMsg_AMBeginProcessingLicenses:
		return "EMsg_AMBeginProcessingLicenses"
	case EMsg_AMSendSystemIMToUser:
		return "EMsg_AMSendSystemIMToUser"
	case EMsg_AMExtendLicense:
		return "EMsg_AMExtendLicense"
	case EMsg_AMAddMinutesToLicense:
		return "EMsg_AMAddMinutesToLicense"
	case EMsg_AMCancelLicense:
		return "EMsg_AMCancelLicense"
	case EMsg_AMInitPurchase:
		return "EMsg_AMInitPurchase"
	case EMsg_AMPurchaseResponse:
		return "EMsg_AMPurchaseResponse"
	case EMsg_AMGetFinalPrice:
		return "EMsg_AMGetFinalPrice"
	case EMsg_AMGetFinalPriceResponse:
		return "EMsg_AMGetFinalPriceResponse"
	case EMsg_AMGetLegacyGameKey:
		return "EMsg_AMGetLegacyGameKey"
	case EMsg_AMGetLegacyGameKeyResponse:
		return "EMsg_AMGetLegacyGameKeyResponse"
	case EMsg_AMFindHungTransactions:
		return "EMsg_AMFindHungTransactions"
	case EMsg_AMSetAccountTrustedRequest:
		return "EMsg_AMSetAccountTrustedRequest"
	case EMsg_AMCompletePurchase:
		return "EMsg_AMCompletePurchase"
	case EMsg_AMCancelPurchase:
		return "EMsg_AMCancelPurchase"
	case EMsg_AMNewChallenge:
		return "EMsg_AMNewChallenge"
	case EMsg_AMFixPendingPurchaseResponse:
		return "EMsg_AMFixPendingPurchaseResponse"
	case EMsg_AMIsUserBanned:
		return "EMsg_AMIsUserBanned"
	case EMsg_AMRegisterKey:
		return "EMsg_AMRegisterKey"
	case EMsg_AMLoadActivationCodes:
		return "EMsg_AMLoadActivationCodes"
	case EMsg_AMLoadActivationCodesResponse:
		return "EMsg_AMLoadActivationCodesResponse"
	case EMsg_AMLookupKeyResponse:
		return "EMsg_AMLookupKeyResponse"
	case EMsg_AMLookupKey:
		return "EMsg_AMLookupKey"
	case EMsg_AMChatCleanup:
		return "EMsg_AMChatCleanup"
	case EMsg_AMClanCleanup:
		return "EMsg_AMClanCleanup"
	case EMsg_AMFixPendingRefund:
		return "EMsg_AMFixPendingRefund"
	case EMsg_AMReverseChargeback:
		return "EMsg_AMReverseChargeback"
	case EMsg_AMReverseChargebackResponse:
		return "EMsg_AMReverseChargebackResponse"
	case EMsg_AMClanCleanupList:
		return "EMsg_AMClanCleanupList"
	case EMsg_AMGetLicenses:
		return "EMsg_AMGetLicenses"
	case EMsg_AMGetLicensesResponse:
		return "EMsg_AMGetLicensesResponse"
	case EMsg_AllowUserToPlayQuery:
		return "EMsg_AllowUserToPlayQuery"
	case EMsg_AllowUserToPlayResponse:
		return "EMsg_AllowUserToPlayResponse"
	case EMsg_AMVerfiyUser:
		return "EMsg_AMVerfiyUser"
	case EMsg_AMClientNotPlaying:
		return "EMsg_AMClientNotPlaying"
	case EMsg_ClientRequestFriendship:
		return "EMsg_ClientRequestFriendship"
	case EMsg_AMRelayPublishStatus:
		return "EMsg_AMRelayPublishStatus"
	case EMsg_AMResetCommunityContent:
		return "EMsg_AMResetCommunityContent"
	case EMsg_AMPrimePersonaStateCache:
		return "EMsg_AMPrimePersonaStateCache"
	case EMsg_AMAllowUserContentQuery:
		return "EMsg_AMAllowUserContentQuery"
	case EMsg_AMAllowUserContentResponse:
		return "EMsg_AMAllowUserContentResponse"
	case EMsg_AMInitPurchaseResponse:
		return "EMsg_AMInitPurchaseResponse"
	case EMsg_AMRevokePurchaseResponse:
		return "EMsg_AMRevokePurchaseResponse"
	case EMsg_AMLockProfile:
		return "EMsg_AMLockProfile"
	case EMsg_AMRefreshGuestPasses:
		return "EMsg_AMRefreshGuestPasses"
	case EMsg_AMInviteUserToClan:
		return "EMsg_AMInviteUserToClan"
	case EMsg_AMAcknowledgeClanInvite:
		return "EMsg_AMAcknowledgeClanInvite"
	case EMsg_AMGrantGuestPasses:
		return "EMsg_AMGrantGuestPasses"
	case EMsg_AMClanDataUpdated:
		return "EMsg_AMClanDataUpdated"
	case EMsg_AMReloadAccount:
		return "EMsg_AMReloadAccount"
	case EMsg_AMClientChatMsgRelay:
		return "EMsg_AMClientChatMsgRelay"
	case EMsg_AMChatMulti:
		return "EMsg_AMChatMulti"
	case EMsg_AMClientChatInviteRelay:
		return "EMsg_AMClientChatInviteRelay"
	case EMsg_AMChatInvite:
		return "EMsg_AMChatInvite"
	case EMsg_AMClientJoinChatRelay:
		return "EMsg_AMClientJoinChatRelay"
	case EMsg_AMClientChatMemberInfoRelay:
		return "EMsg_AMClientChatMemberInfoRelay"
	case EMsg_AMPublishChatMemberInfo:
		return "EMsg_AMPublishChatMemberInfo"
	case EMsg_AMClientAcceptFriendInvite:
		return "EMsg_AMClientAcceptFriendInvite"
	case EMsg_AMChatEnter:
		return "EMsg_AMChatEnter"
	case EMsg_AMClientPublishRemovalFromSource:
		return "EMsg_AMClientPublishRemovalFromSource"
	case EMsg_AMChatActionResult:
		return "EMsg_AMChatActionResult"
	case EMsg_AMFindAccounts:
		return "EMsg_AMFindAccounts"
	case EMsg_AMFindAccountsResponse:
		return "EMsg_AMFindAccountsResponse"
	case EMsg_AMSetAccountFlags:
		return "EMsg_AMSetAccountFlags"
	case EMsg_AMCreateClan:
		return "EMsg_AMCreateClan"
	case EMsg_AMCreateClanResponse:
		return "EMsg_AMCreateClanResponse"
	case EMsg_AMGetClanDetails:
		return "EMsg_AMGetClanDetails"
	case EMsg_AMGetClanDetailsResponse:
		return "EMsg_AMGetClanDetailsResponse"
	case EMsg_AMSetPersonaName:
		return "EMsg_AMSetPersonaName"
	case EMsg_AMSetAvatar:
		return "EMsg_AMSetAvatar"
	case EMsg_AMAuthenticateUser:
		return "EMsg_AMAuthenticateUser"
	case EMsg_AMAuthenticateUserResponse:
		return "EMsg_AMAuthenticateUserResponse"
	case EMsg_AMGetAccountFriendsCount:
		return "EMsg_AMGetAccountFriendsCount"
	case EMsg_AMGetAccountFriendsCountResponse:
		return "EMsg_AMGetAccountFriendsCountResponse"
	case EMsg_AMP2PIntroducerMessage:
		return "EMsg_AMP2PIntroducerMessage"
	case EMsg_ClientChatAction:
		return "EMsg_ClientChatAction"
	case EMsg_AMClientChatActionRelay:
		return "EMsg_AMClientChatActionRelay"
	case EMsg_BaseVS:
		return "EMsg_BaseVS"
	case EMsg_VACResponse:
		return "EMsg_VACResponse"
	case EMsg_ReqChallengeTest:
		return "EMsg_ReqChallengeTest"
	case EMsg_VSMarkCheat:
		return "EMsg_VSMarkCheat"
	case EMsg_VSAddCheat:
		return "EMsg_VSAddCheat"
	case EMsg_VSPurgeCodeModDB:
		return "EMsg_VSPurgeCodeModDB"
	case EMsg_VSGetChallengeResults:
		return "EMsg_VSGetChallengeResults"
	case EMsg_VSChallengeResultText:
		return "EMsg_VSChallengeResultText"
	case EMsg_VSReportLingerer:
		return "EMsg_VSReportLingerer"
	case EMsg_VSRequestManagedChallenge:
		return "EMsg_VSRequestManagedChallenge"
	case EMsg_VSLoadDBFinished:
		return "EMsg_VSLoadDBFinished"
	case EMsg_BaseDRMS:
		return "EMsg_BaseDRMS"
	case EMsg_DRMBuildBlobRequest:
		return "EMsg_DRMBuildBlobRequest"
	case EMsg_DRMBuildBlobResponse:
		return "EMsg_DRMBuildBlobResponse"
	case EMsg_DRMResolveGuidRequest:
		return "EMsg_DRMResolveGuidRequest"
	case EMsg_DRMResolveGuidResponse:
		return "EMsg_DRMResolveGuidResponse"
	case EMsg_DRMVariabilityReport:
		return "EMsg_DRMVariabilityReport"
	case EMsg_DRMVariabilityReportResponse:
		return "EMsg_DRMVariabilityReportResponse"
	case EMsg_DRMStabilityReport:
		return "EMsg_DRMStabilityReport"
	case EMsg_DRMStabilityReportResponse:
		return "EMsg_DRMStabilityReportResponse"
	case EMsg_DRMDetailsReportRequest:
		return "EMsg_DRMDetailsReportRequest"
	case EMsg_DRMDetailsReportResponse:
		return "EMsg_DRMDetailsReportResponse"
	case EMsg_DRMProcessFile:
		return "EMsg_DRMProcessFile"
	case EMsg_DRMAdminUpdate:
		return "EMsg_DRMAdminUpdate"
	case EMsg_DRMAdminUpdateResponse:
		return "EMsg_DRMAdminUpdateResponse"
	case EMsg_DRMSync:
		return "EMsg_DRMSync"
	case EMsg_DRMSyncResponse:
		return "EMsg_DRMSyncResponse"
	case EMsg_DRMProcessFileResponse:
		return "EMsg_DRMProcessFileResponse"
	case EMsg_DRMEmptyGuidCache:
		return "EMsg_DRMEmptyGuidCache"
	case EMsg_DRMEmptyGuidCacheResponse:
		return "EMsg_DRMEmptyGuidCacheResponse"
	case EMsg_BaseCS:
		return "EMsg_BaseCS"
	case EMsg_CSUserContentRequest:
		return "EMsg_CSUserContentRequest"
	case EMsg_BaseClient:
		return "EMsg_BaseClient"
	case EMsg_ClientLogOn_Deprecated:
		return "EMsg_ClientLogOn_Deprecated"
	case EMsg_ClientAnonLogOn_Deprecated:
		return "EMsg_ClientAnonLogOn_Deprecated"
	case EMsg_ClientHeartBeat:
		return "EMsg_ClientHeartBeat"
	case EMsg_ClientVACResponse:
		return "EMsg_ClientVACResponse"
	case EMsg_ClientGamesPlayed_obsolete:
		return "EMsg_ClientGamesPlayed_obsolete"
	case EMsg_ClientLogOff:
		return "EMsg_ClientLogOff"
	case EMsg_ClientNoUDPConnectivity:
		return "EMsg_ClientNoUDPConnectivity"
	case EMsg_ClientInformOfCreateAccount:
		return "EMsg_ClientInformOfCreateAccount"
	case EMsg_ClientAckVACBan:
		return "EMsg_ClientAckVACBan"
	case EMsg_ClientConnectionStats:
		return "EMsg_ClientConnectionStats"
	case EMsg_ClientInitPurchase:
		return "EMsg_ClientInitPurchase"
	case EMsg_ClientPingResponse:
		return "EMsg_ClientPingResponse"
	case EMsg_ClientRemoveFriend:
		return "EMsg_ClientRemoveFriend"
	case EMsg_ClientGamesPlayedNoDataBlob:
		return "EMsg_ClientGamesPlayedNoDataBlob"
	case EMsg_ClientChangeStatus:
		return "EMsg_ClientChangeStatus"
	case EMsg_ClientVacStatusResponse:
		return "EMsg_ClientVacStatusResponse"
	case EMsg_ClientFriendMsg:
		return "EMsg_ClientFriendMsg"
	case EMsg_ClientGameConnect_obsolete:
		return "EMsg_ClientGameConnect_obsolete"
	case EMsg_ClientGamesPlayed2_obsolete:
		return "EMsg_ClientGamesPlayed2_obsolete"
	case EMsg_ClientGameEnded_obsolete:
		return "EMsg_ClientGameEnded_obsolete"
	case EMsg_ClientGetFinalPrice:
		return "EMsg_ClientGetFinalPrice"
	case EMsg_ClientSystemIM:
		return "EMsg_ClientSystemIM"
	case EMsg_ClientSystemIMAck:
		return "EMsg_ClientSystemIMAck"
	case EMsg_ClientGetLicenses:
		return "EMsg_ClientGetLicenses"
	case EMsg_ClientCancelLicense:
		return "EMsg_ClientCancelLicense"
	case EMsg_ClientGetLegacyGameKey:
		return "EMsg_ClientGetLegacyGameKey"
	case EMsg_ClientContentServerLogOn_Deprecated:
		return "EMsg_ClientContentServerLogOn_Deprecated"
	case EMsg_ClientAckVACBan2:
		return "EMsg_ClientAckVACBan2"
	case EMsg_ClientAckMessageByGID:
		return "EMsg_ClientAckMessageByGID"
	case EMsg_ClientGetPurchaseReceipts:
		return "EMsg_ClientGetPurchaseReceipts"
	case EMsg_ClientAckPurchaseReceipt:
		return "EMsg_ClientAckPurchaseReceipt"
	case EMsg_ClientGamesPlayed3_obsolete:
		return "EMsg_ClientGamesPlayed3_obsolete"
	case EMsg_ClientSendGuestPass:
		return "EMsg_ClientSendGuestPass"
	case EMsg_ClientAckGuestPass:
		return "EMsg_ClientAckGuestPass"
	case EMsg_ClientRedeemGuestPass:
		return "EMsg_ClientRedeemGuestPass"
	case EMsg_ClientGamesPlayed:
		return "EMsg_ClientGamesPlayed"
	case EMsg_ClientRegisterKey:
		return "EMsg_ClientRegisterKey"
	case EMsg_ClientInviteUserToClan:
		return "EMsg_ClientInviteUserToClan"
	case EMsg_ClientAcknowledgeClanInvite:
		return "EMsg_ClientAcknowledgeClanInvite"
	case EMsg_ClientPurchaseWithMachineID:
		return "EMsg_ClientPurchaseWithMachineID"
	case EMsg_ClientAppUsageEvent:
		return "EMsg_ClientAppUsageEvent"
	case EMsg_ClientGetGiftTargetList:
		return "EMsg_ClientGetGiftTargetList"
	case EMsg_ClientGetGiftTargetListResponse:
		return "EMsg_ClientGetGiftTargetListResponse"
	case EMsg_ClientLogOnResponse:
		return "EMsg_ClientLogOnResponse"
	case EMsg_ClientVACChallenge:
		return "EMsg_ClientVACChallenge"
	case EMsg_ClientSetHeartbeatRate:
		return "EMsg_ClientSetHeartbeatRate"
	case EMsg_ClientNotLoggedOnDeprecated:
		return "EMsg_ClientNotLoggedOnDeprecated"
	case EMsg_ClientLoggedOff:
		return "EMsg_ClientLoggedOff"
	case EMsg_GSApprove:
		return "EMsg_GSApprove"
	case EMsg_GSDeny:
		return "EMsg_GSDeny"
	case EMsg_GSKick:
		return "EMsg_GSKick"
	case EMsg_ClientCreateAcctResponse:
		return "EMsg_ClientCreateAcctResponse"
	case EMsg_ClientPurchaseResponse:
		return "EMsg_ClientPurchaseResponse"
	case EMsg_ClientPing:
		return "EMsg_ClientPing"
	case EMsg_ClientNOP:
		return "EMsg_ClientNOP"
	case EMsg_ClientPersonaState:
		return "EMsg_ClientPersonaState"
	case EMsg_ClientFriendsList:
		return "EMsg_ClientFriendsList"
	case EMsg_ClientAccountInfo:
		return "EMsg_ClientAccountInfo"
	case EMsg_ClientVacStatusQuery:
		return "EMsg_ClientVacStatusQuery"
	case EMsg_ClientNewsUpdate:
		return "EMsg_ClientNewsUpdate"
	case EMsg_ClientGameConnectDeny:
		return "EMsg_ClientGameConnectDeny"
	case EMsg_GSStatusReply:
		return "EMsg_GSStatusReply"
	case EMsg_ClientGetFinalPriceResponse:
		return "EMsg_ClientGetFinalPriceResponse"
	case EMsg_ClientGameConnectTokens:
		return "EMsg_ClientGameConnectTokens"
	case EMsg_ClientLicenseList:
		return "EMsg_ClientLicenseList"
	case EMsg_ClientCancelLicenseResponse:
		return "EMsg_ClientCancelLicenseResponse"
	case EMsg_ClientVACBanStatus:
		return "EMsg_ClientVACBanStatus"
	case EMsg_ClientCMList:
		return "EMsg_ClientCMList"
	case EMsg_ClientEncryptPct:
		return "EMsg_ClientEncryptPct"
	case EMsg_ClientGetLegacyGameKeyResponse:
		return "EMsg_ClientGetLegacyGameKeyResponse"
	case EMsg_ClientFavoritesList:
		return "EMsg_ClientFavoritesList"
	case EMsg_CSUserContentApprove:
		return "EMsg_CSUserContentApprove"
	case EMsg_CSUserContentDeny:
		return "EMsg_CSUserContentDeny"
	case EMsg_ClientInitPurchaseResponse:
		return "EMsg_ClientInitPurchaseResponse"
	case EMsg_ClientAddFriend:
		return "EMsg_ClientAddFriend"
	case EMsg_ClientAddFriendResponse:
		return "EMsg_ClientAddFriendResponse"
	case EMsg_ClientInviteFriend:
		return "EMsg_ClientInviteFriend"
	case EMsg_ClientInviteFriendResponse:
		return "EMsg_ClientInviteFriendResponse"
	case EMsg_ClientSendGuestPassResponse:
		return "EMsg_ClientSendGuestPassResponse"
	case EMsg_ClientAckGuestPassResponse:
		return "EMsg_ClientAckGuestPassResponse"
	case EMsg_ClientRedeemGuestPassResponse:
		return "EMsg_ClientRedeemGuestPassResponse"
	case EMsg_ClientUpdateGuestPassesList:
		return "EMsg_ClientUpdateGuestPassesList"
	case EMsg_ClientChatMsg:
		return "EMsg_ClientChatMsg"
	case EMsg_ClientChatInvite:
		return "EMsg_ClientChatInvite"
	case EMsg_ClientJoinChat:
		return "EMsg_ClientJoinChat"
	case EMsg_ClientChatMemberInfo:
		return "EMsg_ClientChatMemberInfo"
	case EMsg_ClientLogOnWithCredentials_Deprecated:
		return "EMsg_ClientLogOnWithCredentials_Deprecated"
	case EMsg_ClientPasswordChangeResponse:
		return "EMsg_ClientPasswordChangeResponse"
	case EMsg_ClientChatEnter:
		return "EMsg_ClientChatEnter"
	case EMsg_ClientFriendRemovedFromSource:
		return "EMsg_ClientFriendRemovedFromSource"
	case EMsg_ClientCreateChat:
		return "EMsg_ClientCreateChat"
	case EMsg_ClientCreateChatResponse:
		return "EMsg_ClientCreateChatResponse"
	case EMsg_ClientUpdateChatMetadata:
		return "EMsg_ClientUpdateChatMetadata"
	case EMsg_ClientP2PIntroducerMessage:
		return "EMsg_ClientP2PIntroducerMessage"
	case EMsg_ClientChatActionResult:
		return "EMsg_ClientChatActionResult"
	case EMsg_ClientRequestFriendData:
		return "EMsg_ClientRequestFriendData"
	case EMsg_ClientGetUserStats:
		return "EMsg_ClientGetUserStats"
	case EMsg_ClientGetUserStatsResponse:
		return "EMsg_ClientGetUserStatsResponse"
	case EMsg_ClientStoreUserStats:
		return "EMsg_ClientStoreUserStats"
	case EMsg_ClientStoreUserStatsResponse:
		return "EMsg_ClientStoreUserStatsResponse"
	case EMsg_ClientClanState:
		return "EMsg_ClientClanState"
	case EMsg_ClientServiceModule:
		return "EMsg_ClientServiceModule"
	case EMsg_ClientServiceCall:
		return "EMsg_ClientServiceCall"
	case EMsg_ClientServiceCallResponse:
		return "EMsg_ClientServiceCallResponse"
	case EMsg_ClientPackageInfoRequest:
		return "EMsg_ClientPackageInfoRequest"
	case EMsg_ClientPackageInfoResponse:
		return "EMsg_ClientPackageInfoResponse"
	case EMsg_ClientNatTraversalStatEvent:
		return "EMsg_ClientNatTraversalStatEvent"
	case EMsg_ClientAppInfoRequest:
		return "EMsg_ClientAppInfoRequest"
	case EMsg_ClientAppInfoResponse:
		return "EMsg_ClientAppInfoResponse"
	case EMsg_ClientSteamUsageEvent:
		return "EMsg_ClientSteamUsageEvent"
	case EMsg_ClientCheckPassword:
		return "EMsg_ClientCheckPassword"
	case EMsg_ClientResetPassword:
		return "EMsg_ClientResetPassword"
	case EMsg_ClientCheckPasswordResponse:
		return "EMsg_ClientCheckPasswordResponse"
	case EMsg_ClientResetPasswordResponse:
		return "EMsg_ClientResetPasswordResponse"
	case EMsg_ClientSessionToken:
		return "EMsg_ClientSessionToken"
	case EMsg_ClientDRMProblemReport:
		return "EMsg_ClientDRMProblemReport"
	case EMsg_ClientSetIgnoreFriend:
		return "EMsg_ClientSetIgnoreFriend"
	case EMsg_ClientSetIgnoreFriendResponse:
		return "EMsg_ClientSetIgnoreFriendResponse"
	case EMsg_ClientGetAppOwnershipTicket:
		return "EMsg_ClientGetAppOwnershipTicket"
	case EMsg_ClientGetAppOwnershipTicketResponse:
		return "EMsg_ClientGetAppOwnershipTicketResponse"
	case EMsg_ClientGetLobbyListResponse:
		return "EMsg_ClientGetLobbyListResponse"
	case EMsg_ClientGetLobbyMetadata:
		return "EMsg_ClientGetLobbyMetadata"
	case EMsg_ClientGetLobbyMetadataResponse:
		return "EMsg_ClientGetLobbyMetadataResponse"
	case EMsg_ClientVTTCert:
		return "EMsg_ClientVTTCert"
	case EMsg_ClientAppInfoUpdate:
		return "EMsg_ClientAppInfoUpdate"
	case EMsg_ClientAppInfoChanges:
		return "EMsg_ClientAppInfoChanges"
	case EMsg_ClientServerList:
		return "EMsg_ClientServerList"
	case EMsg_ClientEmailChangeResponse:
		return "EMsg_ClientEmailChangeResponse"
	case EMsg_ClientSecretQAChangeResponse:
		return "EMsg_ClientSecretQAChangeResponse"
	case EMsg_ClientDRMBlobRequest:
		return "EMsg_ClientDRMBlobRequest"
	case EMsg_ClientDRMBlobResponse:
		return "EMsg_ClientDRMBlobResponse"
	case EMsg_ClientLookupKey:
		return "EMsg_ClientLookupKey"
	case EMsg_ClientLookupKeyResponse:
		return "EMsg_ClientLookupKeyResponse"
	case EMsg_BaseGameServer:
		return "EMsg_BaseGameServer"
	case EMsg_GSDisconnectNotice:
		return "EMsg_GSDisconnectNotice"
	case EMsg_GSStatus:
		return "EMsg_GSStatus"
	case EMsg_GSUserPlaying:
		return "EMsg_GSUserPlaying"
	case EMsg_GSStatus2:
		return "EMsg_GSStatus2"
	case EMsg_GSStatusUpdate_Unused:
		return "EMsg_GSStatusUpdate_Unused"
	case EMsg_GSServerType:
		return "EMsg_GSServerType"
	case EMsg_GSPlayerList:
		return "EMsg_GSPlayerList"
	case EMsg_GSGetUserAchievementStatus:
		return "EMsg_GSGetUserAchievementStatus"
	case EMsg_GSGetUserAchievementStatusResponse:
		return "EMsg_GSGetUserAchievementStatusResponse"
	case EMsg_GSGetPlayStats:
		return "EMsg_GSGetPlayStats"
	case EMsg_GSGetPlayStatsResponse:
		return "EMsg_GSGetPlayStatsResponse"
	case EMsg_GSGetUserGroupStatus:
		return "EMsg_GSGetUserGroupStatus"
	case EMsg_AMGetUserGroupStatus:
		return "EMsg_AMGetUserGroupStatus"
	case EMsg_AMGetUserGroupStatusResponse:
		return "EMsg_AMGetUserGroupStatusResponse"
	case EMsg_GSGetUserGroupStatusResponse:
		return "EMsg_GSGetUserGroupStatusResponse"
	case EMsg_GSGetReputation:
		return "EMsg_GSGetReputation"
	case EMsg_GSGetReputationResponse:
		return "EMsg_GSGetReputationResponse"
	case EMsg_GSAssociateWithClan:
		return "EMsg_GSAssociateWithClan"
	case EMsg_GSAssociateWithClanResponse:
		return "EMsg_GSAssociateWithClanResponse"
	case EMsg_GSComputeNewPlayerCompatibility:
		return "EMsg_GSComputeNewPlayerCompatibility"
	case EMsg_GSComputeNewPlayerCompatibilityResponse:
		return "EMsg_GSComputeNewPlayerCompatibilityResponse"
	case EMsg_BaseAdmin:
		return "EMsg_BaseAdmin"
	case EMsg_AdminCmdResponse:
		return "EMsg_AdminCmdResponse"
	case EMsg_AdminLogListenRequest:
		return "EMsg_AdminLogListenRequest"
	case EMsg_AdminLogEvent:
		return "EMsg_AdminLogEvent"
	case EMsg_LogSearchRequest:
		return "EMsg_LogSearchRequest"
	case EMsg_LogSearchResponse:
		return "EMsg_LogSearchResponse"
	case EMsg_LogSearchCancel:
		return "EMsg_LogSearchCancel"
	case EMsg_UniverseData:
		return "EMsg_UniverseData"
	case EMsg_RequestStatHistory:
		return "EMsg_RequestStatHistory"
	case EMsg_StatHistory:
		return "EMsg_StatHistory"
	case EMsg_AdminPwLogon:
		return "EMsg_AdminPwLogon"
	case EMsg_AdminPwLogonResponse:
		return "EMsg_AdminPwLogonResponse"
	case EMsg_AdminSpew:
		return "EMsg_AdminSpew"
	case EMsg_AdminConsoleTitle:
		return "EMsg_AdminConsoleTitle"
	case EMsg_AdminGCSpew:
		return "EMsg_AdminGCSpew"
	case EMsg_AdminGCCommand:
		return "EMsg_AdminGCCommand"
	case EMsg_AdminGCGetCommandList:
		return "EMsg_AdminGCGetCommandList"
	case EMsg_AdminGCGetCommandListResponse:
		return "EMsg_AdminGCGetCommandListResponse"
	case EMsg_FBSConnectionData:
		return "EMsg_FBSConnectionData"
	case EMsg_AdminMsgSpew:
		return "EMsg_AdminMsgSpew"
	case EMsg_BaseFBS:
		return "EMsg_BaseFBS"
	case EMsg_FBSVersionInfo:
		return "EMsg_FBSVersionInfo"
	case EMsg_FBSForceRefresh:
		return "EMsg_FBSForceRefresh"
	case EMsg_FBSForceBounce:
		return "EMsg_FBSForceBounce"
	case EMsg_FBSDeployPackage:
		return "EMsg_FBSDeployPackage"
	case EMsg_FBSDeployResponse:
		return "EMsg_FBSDeployResponse"
	case EMsg_FBSUpdateBootstrapper:
		return "EMsg_FBSUpdateBootstrapper"
	case EMsg_FBSSetState:
		return "EMsg_FBSSetState"
	case EMsg_FBSApplyOSUpdates:
		return "EMsg_FBSApplyOSUpdates"
	case EMsg_FBSRunCMDScript:
		return "EMsg_FBSRunCMDScript"
	case EMsg_FBSRebootBox:
		return "EMsg_FBSRebootBox"
	case EMsg_FBSSetBigBrotherMode:
		return "EMsg_FBSSetBigBrotherMode"
	case EMsg_FBSMinidumpServer:
		return "EMsg_FBSMinidumpServer"
	case EMsg_FBSSetShellCount_obsolete:
		return "EMsg_FBSSetShellCount_obsolete"
	case EMsg_FBSDeployHotFixPackage:
		return "EMsg_FBSDeployHotFixPackage"
	case EMsg_FBSDeployHotFixResponse:
		return "EMsg_FBSDeployHotFixResponse"
	case EMsg_FBSDownloadHotFix:
		return "EMsg_FBSDownloadHotFix"
	case EMsg_FBSDownloadHotFixResponse:
		return "EMsg_FBSDownloadHotFixResponse"
	case EMsg_FBSUpdateTargetConfigFile:
		return "EMsg_FBSUpdateTargetConfigFile"
	case EMsg_FBSApplyAccountCred:
		return "EMsg_FBSApplyAccountCred"
	case EMsg_FBSApplyAccountCredResponse:
		return "EMsg_FBSApplyAccountCredResponse"
	case EMsg_FBSSetShellCount:
		return "EMsg_FBSSetShellCount"
	case EMsg_FBSTerminateShell:
		return "EMsg_FBSTerminateShell"
	case EMsg_FBSQueryGMForRequest:
		return "EMsg_FBSQueryGMForRequest"
	case EMsg_FBSQueryGMResponse:
		return "EMsg_FBSQueryGMResponse"
	case EMsg_FBSTerminateZombies:
		return "EMsg_FBSTerminateZombies"
	case EMsg_FBSInfoFromBootstrapper:
		return "EMsg_FBSInfoFromBootstrapper"
	case EMsg_FBSRebootBoxResponse:
		return "EMsg_FBSRebootBoxResponse"
	case EMsg_FBSBootstrapperPackageRequest:
		return "EMsg_FBSBootstrapperPackageRequest"
	case EMsg_FBSBootstrapperPackageResponse:
		return "EMsg_FBSBootstrapperPackageResponse"
	case EMsg_FBSBootstrapperGetPackageChunk:
		return "EMsg_FBSBootstrapperGetPackageChunk"
	case EMsg_FBSBootstrapperGetPackageChunkResponse:
		return "EMsg_FBSBootstrapperGetPackageChunkResponse"
	case EMsg_FBSBootstrapperPackageTransferProgress:
		return "EMsg_FBSBootstrapperPackageTransferProgress"
	case EMsg_FBSRestartBootstrapper:
		return "EMsg_FBSRestartBootstrapper"
	case EMsg_BaseFileXfer:
		return "EMsg_BaseFileXfer"
	case EMsg_FileXferResponse:
		return "EMsg_FileXferResponse"
	case EMsg_FileXferData:
		return "EMsg_FileXferData"
	case EMsg_FileXferEnd:
		return "EMsg_FileXferEnd"
	case EMsg_FileXferDataAck:
		return "EMsg_FileXferDataAck"
	case EMsg_BaseChannelAuth:
		return "EMsg_BaseChannelAuth"
	case EMsg_ChannelAuthResponse:
		return "EMsg_ChannelAuthResponse"
	case EMsg_ChannelAuthResult:
		return "EMsg_ChannelAuthResult"
	case EMsg_ChannelEncryptRequest:
		return "EMsg_ChannelEncryptRequest"
	case EMsg_ChannelEncryptResponse:
		return "EMsg_ChannelEncryptResponse"
	case EMsg_ChannelEncryptResult:
		return "EMsg_ChannelEncryptResult"
	case EMsg_BaseBS:
		return "EMsg_BaseBS"
	case EMsg_BSPurchaseStart:
		return "EMsg_BSPurchaseStart"
	case EMsg_BSPurchaseResponse:
		return "EMsg_BSPurchaseResponse"
	case EMsg_BSSettleNOVA:
		return "EMsg_BSSettleNOVA"
	case EMsg_BSSettleComplete:
		return "EMsg_BSSettleComplete"
	case EMsg_BSBannedRequest:
		return "EMsg_BSBannedRequest"
	case EMsg_BSInitPayPalTxn:
		return "EMsg_BSInitPayPalTxn"
	case EMsg_BSInitPayPalTxnResponse:
		return "EMsg_BSInitPayPalTxnResponse"
	case EMsg_BSGetPayPalUserInfo:
		return "EMsg_BSGetPayPalUserInfo"
	case EMsg_BSGetPayPalUserInfoResponse:
		return "EMsg_BSGetPayPalUserInfoResponse"
	case EMsg_BSRefundTxn:
		return "EMsg_BSRefundTxn"
	case EMsg_BSRefundTxnResponse:
		return "EMsg_BSRefundTxnResponse"
	case EMsg_BSGetEvents:
		return "EMsg_BSGetEvents"
	case EMsg_BSChaseRFRRequest:
		return "EMsg_BSChaseRFRRequest"
	case EMsg_BSPaymentInstrBan:
		return "EMsg_BSPaymentInstrBan"
	case EMsg_BSPaymentInstrBanResponse:
		return "EMsg_BSPaymentInstrBanResponse"
	case EMsg_BSProcessGCReports:
		return "EMsg_BSProcessGCReports"
	case EMsg_BSProcessPPReports:
		return "EMsg_BSProcessPPReports"
	case EMsg_BSInitGCBankXferTxn:
		return "EMsg_BSInitGCBankXferTxn"
	case EMsg_BSInitGCBankXferTxnResponse:
		return "EMsg_BSInitGCBankXferTxnResponse"
	case EMsg_BSQueryGCBankXferTxn:
		return "EMsg_BSQueryGCBankXferTxn"
	case EMsg_BSQueryGCBankXferTxnResponse:
		return "EMsg_BSQueryGCBankXferTxnResponse"
	case EMsg_BSCommitGCTxn:
		return "EMsg_BSCommitGCTxn"
	case EMsg_BSQueryTransactionStatus:
		return "EMsg_BSQueryTransactionStatus"
	case EMsg_BSQueryTransactionStatusResponse:
		return "EMsg_BSQueryTransactionStatusResponse"
	case EMsg_BSQueryCBOrderStatus:
		return "EMsg_BSQueryCBOrderStatus"
	case EMsg_BSQueryCBOrderStatusResponse:
		return "EMsg_BSQueryCBOrderStatusResponse"
	case EMsg_BSRunRedFlagReport:
		return "EMsg_BSRunRedFlagReport"
	case EMsg_BSQueryPaymentInstUsage:
		return "EMsg_BSQueryPaymentInstUsage"
	case EMsg_BSQueryPaymentInstResponse:
		return "EMsg_BSQueryPaymentInstResponse"
	case EMsg_BSQueryTxnExtendedInfo:
		return "EMsg_BSQueryTxnExtendedInfo"
	case EMsg_BSQueryTxnExtendedInfoResponse:
		return "EMsg_BSQueryTxnExtendedInfoResponse"
	case EMsg_BSUpdateConversionRates:
		return "EMsg_BSUpdateConversionRates"
	case EMsg_BSProcessUSBankReports:
		return "EMsg_BSProcessUSBankReports"
	case EMsg_BSPurchaseRunFraudChecks:
		return "EMsg_BSPurchaseRunFraudChecks"
	case EMsg_BSPurchaseRunFraudChecksResponse:
		return "EMsg_BSPurchaseRunFraudChecksResponse"
	case EMsg_BSStartShippingJobs:
		return "EMsg_BSStartShippingJobs"
	case EMsg_BSQueryBankInformation:
		return "EMsg_BSQueryBankInformation"
	case EMsg_BSQueryBankInformationResponse:
		return "EMsg_BSQueryBankInformationResponse"
	case EMsg_BSValidateXsollaSignature:
		return "EMsg_BSValidateXsollaSignature"
	case EMsg_BSValidateXsollaSignatureResponse:
		return "EMsg_BSValidateXsollaSignatureResponse"
	case EMsg_BSQiwiWalletInvoice:
		return "EMsg_BSQiwiWalletInvoice"
	case EMsg_BSQiwiWalletInvoiceResponse:
		return "EMsg_BSQiwiWalletInvoiceResponse"
	case EMsg_BSUpdateInventoryFromProPack:
		return "EMsg_BSUpdateInventoryFromProPack"
	case EMsg_BSUpdateInventoryFromProPackResponse:
		return "EMsg_BSUpdateInventoryFromProPackResponse"
	case EMsg_BSSendShippingRequest:
		return "EMsg_BSSendShippingRequest"
	case EMsg_BSSendShippingRequestResponse:
		return "EMsg_BSSendShippingRequestResponse"
	case EMsg_BSGetProPackOrderStatus:
		return "EMsg_BSGetProPackOrderStatus"
	case EMsg_BSGetProPackOrderStatusResponse:
		return "EMsg_BSGetProPackOrderStatusResponse"
	case EMsg_BSCheckJobRunning:
		return "EMsg_BSCheckJobRunning"
	case EMsg_BSCheckJobRunningResponse:
		return "EMsg_BSCheckJobRunningResponse"
	case EMsg_BSResetPackagePurchaseRateLimit:
		return "EMsg_BSResetPackagePurchaseRateLimit"
	case EMsg_BSResetPackagePurchaseRateLimitResponse:
		return "EMsg_BSResetPackagePurchaseRateLimitResponse"
	case EMsg_BSUpdatePaymentData:
		return "EMsg_BSUpdatePaymentData"
	case EMsg_BSUpdatePaymentDataResponse:
		return "EMsg_BSUpdatePaymentDataResponse"
	case EMsg_BSGetBillingAddress:
		return "EMsg_BSGetBillingAddress"
	case EMsg_BSGetBillingAddressResponse:
		return "EMsg_BSGetBillingAddressResponse"
	case EMsg_BSGetCreditCardInfo:
		return "EMsg_BSGetCreditCardInfo"
	case EMsg_BSGetCreditCardInfoResponse:
		return "EMsg_BSGetCreditCardInfoResponse"
	case EMsg_BSRemoveExpiredPaymentData:
		return "EMsg_BSRemoveExpiredPaymentData"
	case EMsg_BSRemoveExpiredPaymentDataResponse:
		return "EMsg_BSRemoveExpiredPaymentDataResponse"
	case EMsg_BSConvertToCurrentKeys:
		return "EMsg_BSConvertToCurrentKeys"
	case EMsg_BSConvertToCurrentKeysResponse:
		return "EMsg_BSConvertToCurrentKeysResponse"
	case EMsg_BSInitPurchase:
		return "EMsg_BSInitPurchase"
	case EMsg_BSInitPurchaseResponse:
		return "EMsg_BSInitPurchaseResponse"
	case EMsg_BSCompletePurchase:
		return "EMsg_BSCompletePurchase"
	case EMsg_BSCompletePurchaseResponse:
		return "EMsg_BSCompletePurchaseResponse"
	case EMsg_BSPruneCardUsageStats:
		return "EMsg_BSPruneCardUsageStats"
	case EMsg_BSPruneCardUsageStatsResponse:
		return "EMsg_BSPruneCardUsageStatsResponse"
	case EMsg_BSStoreBankInformation:
		return "EMsg_BSStoreBankInformation"
	case EMsg_BSStoreBankInformationResponse:
		return "EMsg_BSStoreBankInformationResponse"
	case EMsg_BSVerifyPOSAKey:
		return "EMsg_BSVerifyPOSAKey"
	case EMsg_BSVerifyPOSAKeyResponse:
		return "EMsg_BSVerifyPOSAKeyResponse"
	case EMsg_BSReverseRedeemPOSAKey:
		return "EMsg_BSReverseRedeemPOSAKey"
	case EMsg_BSReverseRedeemPOSAKeyResponse:
		return "EMsg_BSReverseRedeemPOSAKeyResponse"
	case EMsg_BSQueryFindCreditCard:
		return "EMsg_BSQueryFindCreditCard"
	case EMsg_BSQueryFindCreditCardResponse:
		return "EMsg_BSQueryFindCreditCardResponse"
	case EMsg_BSStatusInquiryPOSAKey:
		return "EMsg_BSStatusInquiryPOSAKey"
	case EMsg_BSStatusInquiryPOSAKeyResponse:
		return "EMsg_BSStatusInquiryPOSAKeyResponse"
	case EMsg_BSValidateMoPaySignature:
		return "EMsg_BSValidateMoPaySignature"
	case EMsg_BSValidateMoPaySignatureResponse:
		return "EMsg_BSValidateMoPaySignatureResponse"
	case EMsg_BSMoPayConfirmProductDelivery:
		return "EMsg_BSMoPayConfirmProductDelivery"
	case EMsg_BSMoPayConfirmProductDeliveryResponse:
		return "EMsg_BSMoPayConfirmProductDeliveryResponse"
	case EMsg_BSGenerateMoPayMD5:
		return "EMsg_BSGenerateMoPayMD5"
	case EMsg_BSGenerateMoPayMD5Response:
		return "EMsg_BSGenerateMoPayMD5Response"
	case EMsg_BSBoaCompraConfirmProductDelivery:
		return "EMsg_BSBoaCompraConfirmProductDelivery"
	case EMsg_BSBoaCompraConfirmProductDeliveryResponse:
		return "EMsg_BSBoaCompraConfirmProductDeliveryResponse"
	case EMsg_BSGenerateBoaCompraMD5:
		return "EMsg_BSGenerateBoaCompraMD5"
	case EMsg_BSGenerateBoaCompraMD5Response:
		return "EMsg_BSGenerateBoaCompraMD5Response"
	case EMsg_BaseATS:
		return "EMsg_BaseATS"
	case EMsg_ATSStartStressTest:
		return "EMsg_ATSStartStressTest"
	case EMsg_ATSStopStressTest:
		return "EMsg_ATSStopStressTest"
	case EMsg_ATSRunFailServerTest:
		return "EMsg_ATSRunFailServerTest"
	case EMsg_ATSUFSPerfTestTask:
		return "EMsg_ATSUFSPerfTestTask"
	case EMsg_ATSUFSPerfTestResponse:
		return "EMsg_ATSUFSPerfTestResponse"
	case EMsg_ATSCycleTCM:
		return "EMsg_ATSCycleTCM"
	case EMsg_ATSInitDRMSStressTest:
		return "EMsg_ATSInitDRMSStressTest"
	case EMsg_ATSCallTest:
		return "EMsg_ATSCallTest"
	case EMsg_ATSCallTestReply:
		return "EMsg_ATSCallTestReply"
	case EMsg_ATSStartExternalStress:
		return "EMsg_ATSStartExternalStress"
	case EMsg_ATSExternalStressJobStart:
		return "EMsg_ATSExternalStressJobStart"
	case EMsg_ATSExternalStressJobQueued:
		return "EMsg_ATSExternalStressJobQueued"
	case EMsg_ATSExternalStressJobRunning:
		return "EMsg_ATSExternalStressJobRunning"
	case EMsg_ATSExternalStressJobStopped:
		return "EMsg_ATSExternalStressJobStopped"
	case EMsg_ATSExternalStressJobStopAll:
		return "EMsg_ATSExternalStressJobStopAll"
	case EMsg_ATSExternalStressActionResult:
		return "EMsg_ATSExternalStressActionResult"
	case EMsg_ATSStarted:
		return "EMsg_ATSStarted"
	case EMsg_ATSCSPerfTestTask:
		return "EMsg_ATSCSPerfTestTask"
	case EMsg_ATSCSPerfTestResponse:
		return "EMsg_ATSCSPerfTestResponse"
	case EMsg_BaseDP:
		return "EMsg_BaseDP"
	case EMsg_DPSetPublishingState:
		return "EMsg_DPSetPublishingState"
	case EMsg_DPGamePlayedStats:
		return "EMsg_DPGamePlayedStats"
	case EMsg_DPUniquePlayersStat:
		return "EMsg_DPUniquePlayersStat"
	case EMsg_DPVacInfractionStats:
		return "EMsg_DPVacInfractionStats"
	case EMsg_DPVacBanStats:
		return "EMsg_DPVacBanStats"
	case EMsg_DPBlockingStats:
		return "EMsg_DPBlockingStats"
	case EMsg_DPNatTraversalStats:
		return "EMsg_DPNatTraversalStats"
	case EMsg_DPSteamUsageEvent:
		return "EMsg_DPSteamUsageEvent"
	case EMsg_DPVacCertBanStats:
		return "EMsg_DPVacCertBanStats"
	case EMsg_DPVacCafeBanStats:
		return "EMsg_DPVacCafeBanStats"
	case EMsg_DPCloudStats:
		return "EMsg_DPCloudStats"
	case EMsg_DPAchievementStats:
		return "EMsg_DPAchievementStats"
	case EMsg_DPAccountCreationStats:
		return "EMsg_DPAccountCreationStats"
	case EMsg_DPGetPlayerCount:
		return "EMsg_DPGetPlayerCount"
	case EMsg_DPGetPlayerCountResponse:
		return "EMsg_DPGetPlayerCountResponse"
	case EMsg_DPGameServersPlayersStats:
		return "EMsg_DPGameServersPlayersStats"
	case EMsg_DPDownloadRateStatistics:
		return "EMsg_DPDownloadRateStatistics"
	case EMsg_DPFacebookStatistics:
		return "EMsg_DPFacebookStatistics"
	case EMsg_ClientDPCheckSpecialSurvey:
		return "EMsg_ClientDPCheckSpecialSurvey"
	case EMsg_ClientDPCheckSpecialSurveyResponse:
		return "EMsg_ClientDPCheckSpecialSurveyResponse"
	case EMsg_ClientDPSendSpecialSurveyResponse:
		return "EMsg_ClientDPSendSpecialSurveyResponse"
	case EMsg_ClientDPSendSpecialSurveyResponseReply:
		return "EMsg_ClientDPSendSpecialSurveyResponseReply"
	case EMsg_DPStoreSaleStatistics:
		return "EMsg_DPStoreSaleStatistics"
	case EMsg_ClientDPUpdateAppJobReport:
		return "EMsg_ClientDPUpdateAppJobReport"
	case EMsg_ClientDPSteam2AppStarted:
		return "EMsg_ClientDPSteam2AppStarted"
	case EMsg_DPUpdateContentEvent:
		return "EMsg_DPUpdateContentEvent"
	case EMsg_BaseCM:
		return "EMsg_BaseCM"
	case EMsg_CMSetAllowState:
		return "EMsg_CMSetAllowState"
	case EMsg_CMSpewAllowState:
		return "EMsg_CMSpewAllowState"
	case EMsg_CMAppInfoResponseDeprecated:
		return "EMsg_CMAppInfoResponseDeprecated"
	case EMsg_BaseDSS:
		return "EMsg_BaseDSS"
	case EMsg_DSSNewFile:
		return "EMsg_DSSNewFile"
	case EMsg_DSSCurrentFileList:
		return "EMsg_DSSCurrentFileList"
	case EMsg_DSSSynchList:
		return "EMsg_DSSSynchList"
	case EMsg_DSSSynchListResponse:
		return "EMsg_DSSSynchListResponse"
	case EMsg_DSSSynchSubscribe:
		return "EMsg_DSSSynchSubscribe"
	case EMsg_DSSSynchUnsubscribe:
		return "EMsg_DSSSynchUnsubscribe"
	case EMsg_BaseEPM:
		return "EMsg_BaseEPM"
	case EMsg_EPMStartProcess:
		return "EMsg_EPMStartProcess"
	case EMsg_EPMStopProcess:
		return "EMsg_EPMStopProcess"
	case EMsg_EPMRestartProcess:
		return "EMsg_EPMRestartProcess"
	case EMsg_BaseGC:
		return "EMsg_BaseGC"
	case EMsg_AMRelayToGC:
		return "EMsg_AMRelayToGC"
	case EMsg_GCUpdatePlayedState:
		return "EMsg_GCUpdatePlayedState"
	case EMsg_GCCmdRevive:
		return "EMsg_GCCmdRevive"
	case EMsg_GCCmdBounce:
		return "EMsg_GCCmdBounce"
	case EMsg_GCCmdForceBounce:
		return "EMsg_GCCmdForceBounce"
	case EMsg_GCCmdDown:
		return "EMsg_GCCmdDown"
	case EMsg_GCCmdDeploy:
		return "EMsg_GCCmdDeploy"
	case EMsg_GCCmdDeployResponse:
		return "EMsg_GCCmdDeployResponse"
	case EMsg_GCCmdSwitch:
		return "EMsg_GCCmdSwitch"
	case EMsg_AMRefreshSessions:
		return "EMsg_AMRefreshSessions"
	case EMsg_GCUpdateGSState:
		return "EMsg_GCUpdateGSState"
	case EMsg_GCAchievementAwarded:
		return "EMsg_GCAchievementAwarded"
	case EMsg_GCSystemMessage:
		return "EMsg_GCSystemMessage"
	case EMsg_GCValidateSession:
		return "EMsg_GCValidateSession"
	case EMsg_GCValidateSessionResponse:
		return "EMsg_GCValidateSessionResponse"
	case EMsg_GCCmdStatus:
		return "EMsg_GCCmdStatus"
	case EMsg_GCRegisterWebInterfaces:
		return "EMsg_GCRegisterWebInterfaces"
	case EMsg_GCGetAccountDetails:
		return "EMsg_GCGetAccountDetails"
	case EMsg_GCInterAppMessage:
		return "EMsg_GCInterAppMessage"
	case EMsg_GCGetEmailTemplate:
		return "EMsg_GCGetEmailTemplate"
	case EMsg_GCGetEmailTemplateResponse:
		return "EMsg_GCGetEmailTemplateResponse"
	case EMsg_ISRelayToGCH:
		return "EMsg_ISRelayToGCH"
	case EMsg_GCHRelayClientToIS:
		return "EMsg_GCHRelayClientToIS"
	case EMsg_GCHUpdateSession:
		return "EMsg_GCHUpdateSession"
	case EMsg_GCHRequestUpdateSession:
		return "EMsg_GCHRequestUpdateSession"
	case EMsg_GCHRequestStatus:
		return "EMsg_GCHRequestStatus"
	case EMsg_GCHRequestStatusResponse:
		return "EMsg_GCHRequestStatusResponse"
	case EMsg_BaseP2P:
		return "EMsg_BaseP2P"
	case EMsg_P2PIntroducerMessage:
		return "EMsg_P2PIntroducerMessage"
	case EMsg_BaseSM:
		return "EMsg_BaseSM"
	case EMsg_SMExpensiveReport:
		return "EMsg_SMExpensiveReport"
	case EMsg_SMHourlyReport:
		return "EMsg_SMHourlyReport"
	case EMsg_SMFishingReport:
		return "EMsg_SMFishingReport"
	case EMsg_SMPartitionRenames:
		return "EMsg_SMPartitionRenames"
	case EMsg_SMMonitorSpace:
		return "EMsg_SMMonitorSpace"
	case EMsg_SMGetSchemaConversionResults:
		return "EMsg_SMGetSchemaConversionResults"
	case EMsg_SMGetSchemaConversionResultsResponse:
		return "EMsg_SMGetSchemaConversionResultsResponse"
	case EMsg_BaseTest:
		return "EMsg_BaseTest"
	case EMsg_JobHeartbeatTest:
		return "EMsg_JobHeartbeatTest"
	case EMsg_JobHeartbeatTestResponse:
		return "EMsg_JobHeartbeatTestResponse"
	case EMsg_BaseFTSRange:
		return "EMsg_BaseFTSRange"
	case EMsg_FTSGetBrowseCounts:
		return "EMsg_FTSGetBrowseCounts"
	case EMsg_FTSGetBrowseCountsResponse:
		return "EMsg_FTSGetBrowseCountsResponse"
	case EMsg_FTSBrowseClans:
		return "EMsg_FTSBrowseClans"
	case EMsg_FTSBrowseClansResponse:
		return "EMsg_FTSBrowseClansResponse"
	case EMsg_FTSSearchClansByLocation:
		return "EMsg_FTSSearchClansByLocation"
	case EMsg_FTSSearchClansByLocationResponse:
		return "EMsg_FTSSearchClansByLocationResponse"
	case EMsg_FTSSearchPlayersByLocation:
		return "EMsg_FTSSearchPlayersByLocation"
	case EMsg_FTSSearchPlayersByLocationResponse:
		return "EMsg_FTSSearchPlayersByLocationResponse"
	case EMsg_FTSClanDeleted:
		return "EMsg_FTSClanDeleted"
	case EMsg_FTSSearch:
		return "EMsg_FTSSearch"
	case EMsg_FTSSearchResponse:
		return "EMsg_FTSSearchResponse"
	case EMsg_FTSSearchStatus:
		return "EMsg_FTSSearchStatus"
	case EMsg_FTSSearchStatusResponse:
		return "EMsg_FTSSearchStatusResponse"
	case EMsg_FTSGetGSPlayStats:
		return "EMsg_FTSGetGSPlayStats"
	case EMsg_FTSGetGSPlayStatsResponse:
		return "EMsg_FTSGetGSPlayStatsResponse"
	case EMsg_FTSGetGSPlayStatsForServer:
		return "EMsg_FTSGetGSPlayStatsForServer"
	case EMsg_FTSGetGSPlayStatsForServerResponse:
		return "EMsg_FTSGetGSPlayStatsForServerResponse"
	case EMsg_FTSReportIPUpdates:
		return "EMsg_FTSReportIPUpdates"
	case EMsg_BaseCCSRange:
		return "EMsg_BaseCCSRange"
	case EMsg_CCSGetComments:
		return "EMsg_CCSGetComments"
	case EMsg_CCSGetCommentsResponse:
		return "EMsg_CCSGetCommentsResponse"
	case EMsg_CCSAddComment:
		return "EMsg_CCSAddComment"
	case EMsg_CCSAddCommentResponse:
		return "EMsg_CCSAddCommentResponse"
	case EMsg_CCSDeleteComment:
		return "EMsg_CCSDeleteComment"
	case EMsg_CCSDeleteCommentResponse:
		return "EMsg_CCSDeleteCommentResponse"
	case EMsg_CCSPreloadComments:
		return "EMsg_CCSPreloadComments"
	case EMsg_CCSNotifyCommentCount:
		return "EMsg_CCSNotifyCommentCount"
	case EMsg_CCSGetCommentsForNews:
		return "EMsg_CCSGetCommentsForNews"
	case EMsg_CCSGetCommentsForNewsResponse:
		return "EMsg_CCSGetCommentsForNewsResponse"
	case EMsg_CCSDeleteAllCommentsByAuthor:
		return "EMsg_CCSDeleteAllCommentsByAuthor"
	case EMsg_CCSDeleteAllCommentsByAuthorResponse:
		return "EMsg_CCSDeleteAllCommentsByAuthorResponse"
	case EMsg_BaseLBSRange:
		return "EMsg_BaseLBSRange"
	case EMsg_LBSSetScore:
		return "EMsg_LBSSetScore"
	case EMsg_LBSSetScoreResponse:
		return "EMsg_LBSSetScoreResponse"
	case EMsg_LBSFindOrCreateLB:
		return "EMsg_LBSFindOrCreateLB"
	case EMsg_LBSFindOrCreateLBResponse:
		return "EMsg_LBSFindOrCreateLBResponse"
	case EMsg_LBSGetLBEntries:
		return "EMsg_LBSGetLBEntries"
	case EMsg_LBSGetLBEntriesResponse:
		return "EMsg_LBSGetLBEntriesResponse"
	case EMsg_LBSGetLBList:
		return "EMsg_LBSGetLBList"
	case EMsg_LBSGetLBListResponse:
		return "EMsg_LBSGetLBListResponse"
	case EMsg_LBSSetLBDetails:
		return "EMsg_LBSSetLBDetails"
	case EMsg_LBSDeleteLB:
		return "EMsg_LBSDeleteLB"
	case EMsg_LBSDeleteLBEntry:
		return "EMsg_LBSDeleteLBEntry"
	case EMsg_LBSResetLB:
		return "EMsg_LBSResetLB"
	case EMsg_BaseOGS:
		return "EMsg_BaseOGS"
	case EMsg_OGSBeginSession:
		return "EMsg_OGSBeginSession"
	case EMsg_OGSBeginSessionResponse:
		return "EMsg_OGSBeginSessionResponse"
	case EMsg_OGSEndSession:
		return "EMsg_OGSEndSession"
	case EMsg_OGSEndSessionResponse:
		return "EMsg_OGSEndSessionResponse"
	case EMsg_OGSWriteAppSessionRow:
		return "EMsg_OGSWriteAppSessionRow"
	case EMsg_BaseBRP:
		return "EMsg_BaseBRP"
	case EMsg_BRPStartShippingJobs:
		return "EMsg_BRPStartShippingJobs"
	case EMsg_BRPProcessUSBankReports:
		return "EMsg_BRPProcessUSBankReports"
	case EMsg_BRPProcessGCReports:
		return "EMsg_BRPProcessGCReports"
	case EMsg_BRPProcessPPReports:
		return "EMsg_BRPProcessPPReports"
	case EMsg_BRPSettleNOVA:
		return "EMsg_BRPSettleNOVA"
	case EMsg_BRPSettleCB:
		return "EMsg_BRPSettleCB"
	case EMsg_BRPCommitGC:
		return "EMsg_BRPCommitGC"
	case EMsg_BRPCommitGCResponse:
		return "EMsg_BRPCommitGCResponse"
	case EMsg_BRPFindHungTransactions:
		return "EMsg_BRPFindHungTransactions"
	case EMsg_BRPCheckFinanceCloseOutDate:
		return "EMsg_BRPCheckFinanceCloseOutDate"
	case EMsg_BRPProcessLicenses:
		return "EMsg_BRPProcessLicenses"
	case EMsg_BRPProcessLicensesResponse:
		return "EMsg_BRPProcessLicensesResponse"
	case EMsg_BRPRemoveExpiredPaymentData:
		return "EMsg_BRPRemoveExpiredPaymentData"
	case EMsg_BRPRemoveExpiredPaymentDataResponse:
		return "EMsg_BRPRemoveExpiredPaymentDataResponse"
	case EMsg_BRPConvertToCurrentKeys:
		return "EMsg_BRPConvertToCurrentKeys"
	case EMsg_BRPConvertToCurrentKeysResponse:
		return "EMsg_BRPConvertToCurrentKeysResponse"
	case EMsg_BRPPruneCardUsageStats:
		return "EMsg_BRPPruneCardUsageStats"
	case EMsg_BRPPruneCardUsageStatsResponse:
		return "EMsg_BRPPruneCardUsageStatsResponse"
	case EMsg_BRPCheckActivationCodes:
		return "EMsg_BRPCheckActivationCodes"
	case EMsg_BRPCheckActivationCodesResponse:
		return "EMsg_BRPCheckActivationCodesResponse"
	case EMsg_BaseAMRange2:
		return "EMsg_BaseAMRange2"
	case EMsg_AMCreateChat:
		return "EMsg_AMCreateChat"
	case EMsg_AMCreateChatResponse:
		return "EMsg_AMCreateChatResponse"
	case EMsg_AMUpdateChatMetadata:
		return "EMsg_AMUpdateChatMetadata"
	case EMsg_AMPublishChatMetadata:
		return "EMsg_AMPublishChatMetadata"
	case EMsg_AMSetProfileURL:
		return "EMsg_AMSetProfileURL"
	case EMsg_AMGetAccountEmailAddress:
		return "EMsg_AMGetAccountEmailAddress"
	case EMsg_AMGetAccountEmailAddressResponse:
		return "EMsg_AMGetAccountEmailAddressResponse"
	case EMsg_AMRequestFriendData:
		return "EMsg_AMRequestFriendData"
	case EMsg_AMRouteToClients:
		return "EMsg_AMRouteToClients"
	case EMsg_AMLeaveClan:
		return "EMsg_AMLeaveClan"
	case EMsg_AMClanPermissions:
		return "EMsg_AMClanPermissions"
	case EMsg_AMClanPermissionsResponse:
		return "EMsg_AMClanPermissionsResponse"
	case EMsg_AMCreateClanEvent:
		return "EMsg_AMCreateClanEvent"
	case EMsg_AMCreateClanEventResponse:
		return "EMsg_AMCreateClanEventResponse"
	case EMsg_AMUpdateClanEvent:
		return "EMsg_AMUpdateClanEvent"
	case EMsg_AMUpdateClanEventResponse:
		return "EMsg_AMUpdateClanEventResponse"
	case EMsg_AMGetClanEvents:
		return "EMsg_AMGetClanEvents"
	case EMsg_AMGetClanEventsResponse:
		return "EMsg_AMGetClanEventsResponse"
	case EMsg_AMDeleteClanEvent:
		return "EMsg_AMDeleteClanEvent"
	case EMsg_AMDeleteClanEventResponse:
		return "EMsg_AMDeleteClanEventResponse"
	case EMsg_AMSetClanPermissionSettings:
		return "EMsg_AMSetClanPermissionSettings"
	case EMsg_AMSetClanPermissionSettingsResponse:
		return "EMsg_AMSetClanPermissionSettingsResponse"
	case EMsg_AMGetClanPermissionSettings:
		return "EMsg_AMGetClanPermissionSettings"
	case EMsg_AMGetClanPermissionSettingsResponse:
		return "EMsg_AMGetClanPermissionSettingsResponse"
	case EMsg_AMPublishChatRoomInfo:
		return "EMsg_AMPublishChatRoomInfo"
	case EMsg_ClientChatRoomInfo:
		return "EMsg_ClientChatRoomInfo"
	case EMsg_AMCreateClanAnnouncement:
		return "EMsg_AMCreateClanAnnouncement"
	case EMsg_AMCreateClanAnnouncementResponse:
		return "EMsg_AMCreateClanAnnouncementResponse"
	case EMsg_AMUpdateClanAnnouncement:
		return "EMsg_AMUpdateClanAnnouncement"
	case EMsg_AMUpdateClanAnnouncementResponse:
		return "EMsg_AMUpdateClanAnnouncementResponse"
	case EMsg_AMGetClanAnnouncementsCount:
		return "EMsg_AMGetClanAnnouncementsCount"
	case EMsg_AMGetClanAnnouncementsCountResponse:
		return "EMsg_AMGetClanAnnouncementsCountResponse"
	case EMsg_AMGetClanAnnouncements:
		return "EMsg_AMGetClanAnnouncements"
	case EMsg_AMGetClanAnnouncementsResponse:
		return "EMsg_AMGetClanAnnouncementsResponse"
	case EMsg_AMDeleteClanAnnouncement:
		return "EMsg_AMDeleteClanAnnouncement"
	case EMsg_AMDeleteClanAnnouncementResponse:
		return "EMsg_AMDeleteClanAnnouncementResponse"
	case EMsg_AMGetSingleClanAnnouncement:
		return "EMsg_AMGetSingleClanAnnouncement"
	case EMsg_AMGetSingleClanAnnouncementResponse:
		return "EMsg_AMGetSingleClanAnnouncementResponse"
	case EMsg_AMGetClanHistory:
		return "EMsg_AMGetClanHistory"
	case EMsg_AMGetClanHistoryResponse:
		return "EMsg_AMGetClanHistoryResponse"
	case EMsg_AMGetClanPermissionBits:
		return "EMsg_AMGetClanPermissionBits"
	case EMsg_AMGetClanPermissionBitsResponse:
		return "EMsg_AMGetClanPermissionBitsResponse"
	case EMsg_AMSetClanPermissionBits:
		return "EMsg_AMSetClanPermissionBits"
	case EMsg_AMSetClanPermissionBitsResponse:
		return "EMsg_AMSetClanPermissionBitsResponse"
	case EMsg_AMSessionInfoRequest:
		return "EMsg_AMSessionInfoRequest"
	case EMsg_AMSessionInfoResponse:
		return "EMsg_AMSessionInfoResponse"
	case EMsg_AMValidateWGToken:
		return "EMsg_AMValidateWGToken"
	case EMsg_AMGetSingleClanEvent:
		return "EMsg_AMGetSingleClanEvent"
	case EMsg_AMGetSingleClanEventResponse:
		return "EMsg_AMGetSingleClanEventResponse"
	case EMsg_AMGetClanRank:
		return "EMsg_AMGetClanRank"
	case EMsg_AMGetClanRankResponse:
		return "EMsg_AMGetClanRankResponse"
	case EMsg_AMSetClanRank:
		return "EMsg_AMSetClanRank"
	case EMsg_AMSetClanRankResponse:
		return "EMsg_AMSetClanRankResponse"
	case EMsg_AMGetClanPOTW:
		return "EMsg_AMGetClanPOTW"
	case EMsg_AMGetClanPOTWResponse:
		return "EMsg_AMGetClanPOTWResponse"
	case EMsg_AMSetClanPOTW:
		return "EMsg_AMSetClanPOTW"
	case EMsg_AMSetClanPOTWResponse:
		return "EMsg_AMSetClanPOTWResponse"
	case EMsg_AMRequestChatMetadata:
		return "EMsg_AMRequestChatMetadata"
	case EMsg_AMDumpUser:
		return "EMsg_AMDumpUser"
	case EMsg_AMKickUserFromClan:
		return "EMsg_AMKickUserFromClan"
	case EMsg_AMAddFounderToClan:
		return "EMsg_AMAddFounderToClan"
	case EMsg_AMValidateWGTokenResponse:
		return "EMsg_AMValidateWGTokenResponse"
	case EMsg_AMSetCommunityState:
		return "EMsg_AMSetCommunityState"
	case EMsg_AMSetAccountDetails:
		return "EMsg_AMSetAccountDetails"
	case EMsg_AMGetChatBanList:
		return "EMsg_AMGetChatBanList"
	case EMsg_AMGetChatBanListResponse:
		return "EMsg_AMGetChatBanListResponse"
	case EMsg_AMUnBanFromChat:
		return "EMsg_AMUnBanFromChat"
	case EMsg_AMSetClanDetails:
		return "EMsg_AMSetClanDetails"
	case EMsg_AMGetAccountLinks:
		return "EMsg_AMGetAccountLinks"
	case EMsg_AMGetAccountLinksResponse:
		return "EMsg_AMGetAccountLinksResponse"
	case EMsg_AMSetAccountLinks:
		return "EMsg_AMSetAccountLinks"
	case EMsg_AMSetAccountLinksResponse:
		return "EMsg_AMSetAccountLinksResponse"
	case EMsg_AMGetUserGameStats:
		return "EMsg_AMGetUserGameStats"
	case EMsg_AMGetUserGameStatsResponse:
		return "EMsg_AMGetUserGameStatsResponse"
	case EMsg_AMCheckClanMembership:
		return "EMsg_AMCheckClanMembership"
	case EMsg_AMGetClanMembers:
		return "EMsg_AMGetClanMembers"
	case EMsg_AMGetClanMembersResponse:
		return "EMsg_AMGetClanMembersResponse"
	case EMsg_AMJoinPublicClan:
		return "EMsg_AMJoinPublicClan"
	case EMsg_AMNotifyChatOfClanChange:
		return "EMsg_AMNotifyChatOfClanChange"
	case EMsg_AMResubmitPurchase:
		return "EMsg_AMResubmitPurchase"
	case EMsg_AMAddFriend:
		return "EMsg_AMAddFriend"
	case EMsg_AMAddFriendResponse:
		return "EMsg_AMAddFriendResponse"
	case EMsg_AMRemoveFriend:
		return "EMsg_AMRemoveFriend"
	case EMsg_AMDumpClan:
		return "EMsg_AMDumpClan"
	case EMsg_AMChangeClanOwner:
		return "EMsg_AMChangeClanOwner"
	case EMsg_AMCancelEasyCollect:
		return "EMsg_AMCancelEasyCollect"
	case EMsg_AMCancelEasyCollectResponse:
		return "EMsg_AMCancelEasyCollectResponse"
	case EMsg_AMGetClanMembershipList:
		return "EMsg_AMGetClanMembershipList"
	case EMsg_AMGetClanMembershipListResponse:
		return "EMsg_AMGetClanMembershipListResponse"
	case EMsg_AMClansInCommon:
		return "EMsg_AMClansInCommon"
	case EMsg_AMClansInCommonResponse:
		return "EMsg_AMClansInCommonResponse"
	case EMsg_AMIsValidAccountID:
		return "EMsg_AMIsValidAccountID"
	case EMsg_AMConvertClan:
		return "EMsg_AMConvertClan"
	case EMsg_AMGetGiftTargetListRelay:
		return "EMsg_AMGetGiftTargetListRelay"
	case EMsg_AMWipeFriendsList:
		return "EMsg_AMWipeFriendsList"
	case EMsg_AMSetIgnored:
		return "EMsg_AMSetIgnored"
	case EMsg_AMClansInCommonCountResponse:
		return "EMsg_AMClansInCommonCountResponse"
	case EMsg_AMFriendsList:
		return "EMsg_AMFriendsList"
	case EMsg_AMFriendsListResponse:
		return "EMsg_AMFriendsListResponse"
	case EMsg_AMFriendsInCommon:
		return "EMsg_AMFriendsInCommon"
	case EMsg_AMFriendsInCommonResponse:
		return "EMsg_AMFriendsInCommonResponse"
	case EMsg_AMFriendsInCommonCountResponse:
		return "EMsg_AMFriendsInCommonCountResponse"
	case EMsg_AMClansInCommonCount:
		return "EMsg_AMClansInCommonCount"
	case EMsg_AMChallengeVerdict:
		return "EMsg_AMChallengeVerdict"
	case EMsg_AMChallengeNotification:
		return "EMsg_AMChallengeNotification"
	case EMsg_AMFindGSByIP:
		return "EMsg_AMFindGSByIP"
	case EMsg_AMFoundGSByIP:
		return "EMsg_AMFoundGSByIP"
	case EMsg_AMGiftRevoked:
		return "EMsg_AMGiftRevoked"
	case EMsg_AMCreateAccountRecord:
		return "EMsg_AMCreateAccountRecord"
	case EMsg_AMUserClanList:
		return "EMsg_AMUserClanList"
	case EMsg_AMUserClanListResponse:
		return "EMsg_AMUserClanListResponse"
	case EMsg_AMGetAccountDetails2:
		return "EMsg_AMGetAccountDetails2"
	case EMsg_AMGetAccountDetailsResponse2:
		return "EMsg_AMGetAccountDetailsResponse2"
	case EMsg_AMSetCommunityProfileSettings:
		return "EMsg_AMSetCommunityProfileSettings"
	case EMsg_AMSetCommunityProfileSettingsResponse:
		return "EMsg_AMSetCommunityProfileSettingsResponse"
	case EMsg_AMGetCommunityPrivacyState:
		return "EMsg_AMGetCommunityPrivacyState"
	case EMsg_AMGetCommunityPrivacyStateResponse:
		return "EMsg_AMGetCommunityPrivacyStateResponse"
	case EMsg_AMCheckClanInviteRateLimiting:
		return "EMsg_AMCheckClanInviteRateLimiting"
	case EMsg_AMGetUserAchievementStatus:
		return "EMsg_AMGetUserAchievementStatus"
	case EMsg_AMGetIgnored:
		return "EMsg_AMGetIgnored"
	case EMsg_AMGetIgnoredResponse:
		return "EMsg_AMGetIgnoredResponse"
	case EMsg_AMSetIgnoredResponse:
		return "EMsg_AMSetIgnoredResponse"
	case EMsg_AMSetFriendRelationshipNone:
		return "EMsg_AMSetFriendRelationshipNone"
	case EMsg_AMGetFriendRelationship:
		return "EMsg_AMGetFriendRelationship"
	case EMsg_AMGetFriendRelationshipResponse:
		return "EMsg_AMGetFriendRelationshipResponse"
	case EMsg_AMServiceModulesCache:
		return "EMsg_AMServiceModulesCache"
	case EMsg_AMServiceModulesCall:
		return "EMsg_AMServiceModulesCall"
	case EMsg_AMServiceModulesCallResponse:
		return "EMsg_AMServiceModulesCallResponse"
	case EMsg_AMGetCaptchaDataForIP:
		return "EMsg_AMGetCaptchaDataForIP"
	case EMsg_AMGetCaptchaDataForIPResponse:
		return "EMsg_AMGetCaptchaDataForIPResponse"
	case EMsg_AMValidateCaptchaDataForIP:
		return "EMsg_AMValidateCaptchaDataForIP"
	case EMsg_AMValidateCaptchaDataForIPResponse:
		return "EMsg_AMValidateCaptchaDataForIPResponse"
	case EMsg_AMTrackFailedAuthByIP:
		return "EMsg_AMTrackFailedAuthByIP"
	case EMsg_AMGetCaptchaDataByGID:
		return "EMsg_AMGetCaptchaDataByGID"
	case EMsg_AMGetCaptchaDataByGIDResponse:
		return "EMsg_AMGetCaptchaDataByGIDResponse"
	case EMsg_AMGetLobbyList:
		return "EMsg_AMGetLobbyList"
	case EMsg_AMGetLobbyListResponse:
		return "EMsg_AMGetLobbyListResponse"
	case EMsg_AMGetLobbyMetadata:
		return "EMsg_AMGetLobbyMetadata"
	case EMsg_AMGetLobbyMetadataResponse:
		return "EMsg_AMGetLobbyMetadataResponse"
	case EMsg_CommunityAddFriendNews:
		return "EMsg_CommunityAddFriendNews"
	case EMsg_AMAddClanNews:
		return "EMsg_AMAddClanNews"
	case EMsg_AMWriteNews:
		return "EMsg_AMWriteNews"
	case EMsg_AMFindClanUser:
		return "EMsg_AMFindClanUser"
	case EMsg_AMFindClanUserResponse:
		return "EMsg_AMFindClanUserResponse"
	case EMsg_AMBanFromChat:
		return "EMsg_AMBanFromChat"
	case EMsg_AMGetUserHistoryResponse:
		return "EMsg_AMGetUserHistoryResponse"
	case EMsg_AMGetUserNewsSubscriptions:
		return "EMsg_AMGetUserNewsSubscriptions"
	case EMsg_AMGetUserNewsSubscriptionsResponse:
		return "EMsg_AMGetUserNewsSubscriptionsResponse"
	case EMsg_AMSetUserNewsSubscriptions:
		return "EMsg_AMSetUserNewsSubscriptions"
	case EMsg_AMGetUserNews:
		return "EMsg_AMGetUserNews"
	case EMsg_AMGetUserNewsResponse:
		return "EMsg_AMGetUserNewsResponse"
	case EMsg_AMSendQueuedEmails:
		return "EMsg_AMSendQueuedEmails"
	case EMsg_AMSetLicenseFlags:
		return "EMsg_AMSetLicenseFlags"
	case EMsg_AMGetUserHistory:
		return "EMsg_AMGetUserHistory"
	case EMsg_CommunityDeleteUserNews:
		return "EMsg_CommunityDeleteUserNews"
	case EMsg_AMAllowUserFilesRequest:
		return "EMsg_AMAllowUserFilesRequest"
	case EMsg_AMAllowUserFilesResponse:
		return "EMsg_AMAllowUserFilesResponse"
	case EMsg_AMGetAccountStatus:
		return "EMsg_AMGetAccountStatus"
	case EMsg_AMGetAccountStatusResponse:
		return "EMsg_AMGetAccountStatusResponse"
	case EMsg_AMEditBanReason:
		return "EMsg_AMEditBanReason"
	case EMsg_AMCheckClanMembershipResponse:
		return "EMsg_AMCheckClanMembershipResponse"
	case EMsg_AMProbeClanMembershipList:
		return "EMsg_AMProbeClanMembershipList"
	case EMsg_AMProbeClanMembershipListResponse:
		return "EMsg_AMProbeClanMembershipListResponse"
	case EMsg_AMGetFriendsLobbies:
		return "EMsg_AMGetFriendsLobbies"
	case EMsg_AMGetFriendsLobbiesResponse:
		return "EMsg_AMGetFriendsLobbiesResponse"
	case EMsg_AMGetUserFriendNewsResponse:
		return "EMsg_AMGetUserFriendNewsResponse"
	case EMsg_CommunityGetUserFriendNews:
		return "EMsg_CommunityGetUserFriendNews"
	case EMsg_AMGetUserClansNewsResponse:
		return "EMsg_AMGetUserClansNewsResponse"
	case EMsg_AMGetUserClansNews:
		return "EMsg_AMGetUserClansNews"
	case EMsg_AMStoreInitPurchase:
		return "EMsg_AMStoreInitPurchase"
	case EMsg_AMStoreInitPurchaseResponse:
		return "EMsg_AMStoreInitPurchaseResponse"
	case EMsg_AMStoreGetFinalPrice:
		return "EMsg_AMStoreGetFinalPrice"
	case EMsg_AMStoreGetFinalPriceResponse:
		return "EMsg_AMStoreGetFinalPriceResponse"
	case EMsg_AMStoreCompletePurchase:
		return "EMsg_AMStoreCompletePurchase"
	case EMsg_AMStoreCancelPurchase:
		return "EMsg_AMStoreCancelPurchase"
	case EMsg_AMStorePurchaseResponse:
		return "EMsg_AMStorePurchaseResponse"
	case EMsg_AMCreateAccountRecordInSteam3:
		return "EMsg_AMCreateAccountRecordInSteam3"
	case EMsg_AMGetPreviousCBAccount:
		return "EMsg_AMGetPreviousCBAccount"
	case EMsg_AMGetPreviousCBAccountResponse:
		return "EMsg_AMGetPreviousCBAccountResponse"
	case EMsg_AMUpdateBillingAddress:
		return "EMsg_AMUpdateBillingAddress"
	case EMsg_AMUpdateBillingAddressResponse:
		return "EMsg_AMUpdateBillingAddressResponse"
	case EMsg_AMGetBillingAddress:
		return "EMsg_AMGetBillingAddress"
	case EMsg_AMGetBillingAddressResponse:
		return "EMsg_AMGetBillingAddressResponse"
	case EMsg_AMGetUserLicenseHistory:
		return "EMsg_AMGetUserLicenseHistory"
	case EMsg_AMGetUserLicenseHistoryResponse:
		return "EMsg_AMGetUserLicenseHistoryResponse"
	case EMsg_AMSupportChangePassword:
		return "EMsg_AMSupportChangePassword"
	case EMsg_AMSupportChangeEmail:
		return "EMsg_AMSupportChangeEmail"
	case EMsg_AMSupportChangeSecretQA:
		return "EMsg_AMSupportChangeSecretQA"
	case EMsg_AMResetUserVerificationGSByIP:
		return "EMsg_AMResetUserVerificationGSByIP"
	case EMsg_AMUpdateGSPlayStats:
		return "EMsg_AMUpdateGSPlayStats"
	case EMsg_AMSupportEnableOrDisable:
		return "EMsg_AMSupportEnableOrDisable"
	case EMsg_AMGetComments:
		return "EMsg_AMGetComments"
	case EMsg_AMGetCommentsResponse:
		return "EMsg_AMGetCommentsResponse"
	case EMsg_AMAddComment:
		return "EMsg_AMAddComment"
	case EMsg_AMAddCommentResponse:
		return "EMsg_AMAddCommentResponse"
	case EMsg_AMDeleteComment:
		return "EMsg_AMDeleteComment"
	case EMsg_AMDeleteCommentResponse:
		return "EMsg_AMDeleteCommentResponse"
	case EMsg_AMGetPurchaseStatus:
		return "EMsg_AMGetPurchaseStatus"
	case EMsg_AMSupportIsAccountEnabled:
		return "EMsg_AMSupportIsAccountEnabled"
	case EMsg_AMSupportIsAccountEnabledResponse:
		return "EMsg_AMSupportIsAccountEnabledResponse"
	case EMsg_AMGetUserStats:
		return "EMsg_AMGetUserStats"
	case EMsg_AMSupportKickSession:
		return "EMsg_AMSupportKickSession"
	case EMsg_AMGSSearch:
		return "EMsg_AMGSSearch"
	case EMsg_MarketingMessageUpdate:
		return "EMsg_MarketingMessageUpdate"
	case EMsg_AMRouteFriendMsg:
		return "EMsg_AMRouteFriendMsg"
	case EMsg_AMTicketAuthRequestOrResponse:
		return "EMsg_AMTicketAuthRequestOrResponse"
	case EMsg_AMVerifyDepotManagementRights:
		return "EMsg_AMVerifyDepotManagementRights"
	case EMsg_AMVerifyDepotManagementRightsResponse:
		return "EMsg_AMVerifyDepotManagementRightsResponse"
	case EMsg_AMAddFreeLicense:
		return "EMsg_AMAddFreeLicense"
	case EMsg_AMGetUserFriendsMinutesPlayed:
		return "EMsg_AMGetUserFriendsMinutesPlayed"
	case EMsg_AMGetUserFriendsMinutesPlayedResponse:
		return "EMsg_AMGetUserFriendsMinutesPlayedResponse"
	case EMsg_AMGetUserMinutesPlayed:
		return "EMsg_AMGetUserMinutesPlayed"
	case EMsg_AMGetUserMinutesPlayedResponse:
		return "EMsg_AMGetUserMinutesPlayedResponse"
	case EMsg_AMValidateEmailLink:
		return "EMsg_AMValidateEmailLink"
	case EMsg_AMValidateEmailLinkResponse:
		return "EMsg_AMValidateEmailLinkResponse"
	case EMsg_AMAddUsersToMarketingTreatment:
		return "EMsg_AMAddUsersToMarketingTreatment"
	case EMsg_AMStoreUserStats:
		return "EMsg_AMStoreUserStats"
	case EMsg_AMGetUserGameplayInfo:
		return "EMsg_AMGetUserGameplayInfo"
	case EMsg_AMGetUserGameplayInfoResponse:
		return "EMsg_AMGetUserGameplayInfoResponse"
	case EMsg_AMGetCardList:
		return "EMsg_AMGetCardList"
	case EMsg_AMGetCardListResponse:
		return "EMsg_AMGetCardListResponse"
	case EMsg_AMDeleteStoredCard:
		return "EMsg_AMDeleteStoredCard"
	case EMsg_AMRevokeLegacyGameKeys:
		return "EMsg_AMRevokeLegacyGameKeys"
	case EMsg_AMGetWalletDetails:
		return "EMsg_AMGetWalletDetails"
	case EMsg_AMGetWalletDetailsResponse:
		return "EMsg_AMGetWalletDetailsResponse"
	case EMsg_AMDeleteStoredPaymentInfo:
		return "EMsg_AMDeleteStoredPaymentInfo"
	case EMsg_AMGetStoredPaymentSummary:
		return "EMsg_AMGetStoredPaymentSummary"
	case EMsg_AMGetStoredPaymentSummaryResponse:
		return "EMsg_AMGetStoredPaymentSummaryResponse"
	case EMsg_AMGetWalletConversionRate:
		return "EMsg_AMGetWalletConversionRate"
	case EMsg_AMGetWalletConversionRateResponse:
		return "EMsg_AMGetWalletConversionRateResponse"
	case EMsg_AMConvertWallet:
		return "EMsg_AMConvertWallet"
	case EMsg_AMConvertWalletResponse:
		return "EMsg_AMConvertWalletResponse"
	case EMsg_AMRelayGetFriendsWhoPlayGame:
		return "EMsg_AMRelayGetFriendsWhoPlayGame"
	case EMsg_AMRelayGetFriendsWhoPlayGameResponse:
		return "EMsg_AMRelayGetFriendsWhoPlayGameResponse"
	case EMsg_AMSetPreApproval:
		return "EMsg_AMSetPreApproval"
	case EMsg_AMSetPreApprovalResponse:
		return "EMsg_AMSetPreApprovalResponse"
	case EMsg_AMMarketingTreatmentUpdate:
		return "EMsg_AMMarketingTreatmentUpdate"
	case EMsg_AMCreateRefund:
		return "EMsg_AMCreateRefund"
	case EMsg_AMCreateRefundResponse:
		return "EMsg_AMCreateRefundResponse"
	case EMsg_AMCreateChargeback:
		return "EMsg_AMCreateChargeback"
	case EMsg_AMCreateChargebackResponse:
		return "EMsg_AMCreateChargebackResponse"
	case EMsg_AMCreateDispute:
		return "EMsg_AMCreateDispute"
	case EMsg_AMCreateDisputeResponse:
		return "EMsg_AMCreateDisputeResponse"
	case EMsg_AMClearDispute:
		return "EMsg_AMClearDispute"
	case EMsg_AMClearDisputeResponse:
		return "EMsg_AMClearDisputeResponse"
	case EMsg_AMPlayerNicknameList:
		return "EMsg_AMPlayerNicknameList"
	case EMsg_AMPlayerNicknameListResponse:
		return "EMsg_AMPlayerNicknameListResponse"
	case EMsg_AMSetDRMTestConfig:
		return "EMsg_AMSetDRMTestConfig"
	case EMsg_AMGetUserCurrentGameInfo:
		return "EMsg_AMGetUserCurrentGameInfo"
	case EMsg_AMGetUserCurrentGameInfoResponse:
		return "EMsg_AMGetUserCurrentGameInfoResponse"
	case EMsg_AMGetGSPlayerList:
		return "EMsg_AMGetGSPlayerList"
	case EMsg_AMGetGSPlayerListResponse:
		return "EMsg_AMGetGSPlayerListResponse"
	case EMsg_AMUpdatePersonaStateCache:
		return "EMsg_AMUpdatePersonaStateCache"
	case EMsg_AMGetGameMembers:
		return "EMsg_AMGetGameMembers"
	case EMsg_AMGetGameMembersResponse:
		return "EMsg_AMGetGameMembersResponse"
	case EMsg_AMGetSteamIDForMicroTxn:
		return "EMsg_AMGetSteamIDForMicroTxn"
	case EMsg_AMGetSteamIDForMicroTxnResponse:
		return "EMsg_AMGetSteamIDForMicroTxnResponse"
	case EMsg_AMAddPublisherUser:
		return "EMsg_AMAddPublisherUser"
	case EMsg_AMRemovePublisherUser:
		return "EMsg_AMRemovePublisherUser"
	case EMsg_AMGetUserLicenseList:
		return "EMsg_AMGetUserLicenseList"
	case EMsg_AMGetUserLicenseListResponse:
		return "EMsg_AMGetUserLicenseListResponse"
	case EMsg_AMReloadGameGroupPolicy:
		return "EMsg_AMReloadGameGroupPolicy"
	case EMsg_AMAddFreeLicenseResponse:
		return "EMsg_AMAddFreeLicenseResponse"
	case EMsg_AMVACStatusUpdate:
		return "EMsg_AMVACStatusUpdate"
	case EMsg_AMGetAccountDetails:
		return "EMsg_AMGetAccountDetails"
	case EMsg_AMGetAccountDetailsResponse:
		return "EMsg_AMGetAccountDetailsResponse"
	case EMsg_AMGetPlayerLinkDetails:
		return "EMsg_AMGetPlayerLinkDetails"
	case EMsg_AMGetPlayerLinkDetailsResponse:
		return "EMsg_AMGetPlayerLinkDetailsResponse"
	case EMsg_AMSubscribeToPersonaFeed:
		return "EMsg_AMSubscribeToPersonaFeed"
	case EMsg_AMGetUserVacBanList:
		return "EMsg_AMGetUserVacBanList"
	case EMsg_AMGetUserVacBanListResponse:
		return "EMsg_AMGetUserVacBanListResponse"
	case EMsg_AMGetAccountFlagsForWGSpoofing:
		return "EMsg_AMGetAccountFlagsForWGSpoofing"
	case EMsg_AMGetAccountFlagsForWGSpoofingResponse:
		return "EMsg_AMGetAccountFlagsForWGSpoofingResponse"
	case EMsg_AMGetFriendsWishlistInfo:
		return "EMsg_AMGetFriendsWishlistInfo"
	case EMsg_AMGetFriendsWishlistInfoResponse:
		return "EMsg_AMGetFriendsWishlistInfoResponse"
	case EMsg_AMGetClanOfficers:
		return "EMsg_AMGetClanOfficers"
	case EMsg_AMGetClanOfficersResponse:
		return "EMsg_AMGetClanOfficersResponse"
	case EMsg_AMNameChange:
		return "EMsg_AMNameChange"
	case EMsg_AMGetNameHistory:
		return "EMsg_AMGetNameHistory"
	case EMsg_AMGetNameHistoryResponse:
		return "EMsg_AMGetNameHistoryResponse"
	case EMsg_AMUpdateProviderStatus:
		return "EMsg_AMUpdateProviderStatus"
	case EMsg_AMClearPersonaMetadataBlob:
		return "EMsg_AMClearPersonaMetadataBlob"
	case EMsg_AMSupportRemoveAccountSecurity:
		return "EMsg_AMSupportRemoveAccountSecurity"
	case EMsg_AMIsAccountInCaptchaGracePeriod:
		return "EMsg_AMIsAccountInCaptchaGracePeriod"
	case EMsg_AMIsAccountInCaptchaGracePeriodResponse:
		return "EMsg_AMIsAccountInCaptchaGracePeriodResponse"
	case EMsg_AMAccountPS3Unlink:
		return "EMsg_AMAccountPS3Unlink"
	case EMsg_AMAccountPS3UnlinkResponse:
		return "EMsg_AMAccountPS3UnlinkResponse"
	case EMsg_AMStoreUserStatsResponse:
		return "EMsg_AMStoreUserStatsResponse"
	case EMsg_AMGetAccountPSNInfo:
		return "EMsg_AMGetAccountPSNInfo"
	case EMsg_AMGetAccountPSNInfoResponse:
		return "EMsg_AMGetAccountPSNInfoResponse"
	case EMsg_AMAuthenticatedPlayerList:
		return "EMsg_AMAuthenticatedPlayerList"
	case EMsg_AMGetUserGifts:
		return "EMsg_AMGetUserGifts"
	case EMsg_AMGetUserGiftsResponse:
		return "EMsg_AMGetUserGiftsResponse"
	case EMsg_AMTransferLockedGifts:
		return "EMsg_AMTransferLockedGifts"
	case EMsg_AMTransferLockedGiftsResponse:
		return "EMsg_AMTransferLockedGiftsResponse"
	case EMsg_AMPlayerHostedOnGameServer:
		return "EMsg_AMPlayerHostedOnGameServer"
	case EMsg_AMGetAccountBanInfo:
		return "EMsg_AMGetAccountBanInfo"
	case EMsg_AMGetAccountBanInfoResponse:
		return "EMsg_AMGetAccountBanInfoResponse"
	case EMsg_AMRecordBanEnforcement:
		return "EMsg_AMRecordBanEnforcement"
	case EMsg_AMRollbackGiftTransfer:
		return "EMsg_AMRollbackGiftTransfer"
	case EMsg_AMRollbackGiftTransferResponse:
		return "EMsg_AMRollbackGiftTransferResponse"
	case EMsg_AMHandlePendingTransaction:
		return "EMsg_AMHandlePendingTransaction"
	case EMsg_AMRequestClanDetails:
		return "EMsg_AMRequestClanDetails"
	case EMsg_AMDeleteStoredPaypalAgreement:
		return "EMsg_AMDeleteStoredPaypalAgreement"
	case EMsg_AMGameServerUpdate:
		return "EMsg_AMGameServerUpdate"
	case EMsg_AMGameServerRemove:
		return "EMsg_AMGameServerRemove"
	case EMsg_AMGetPaypalAgreements:
		return "EMsg_AMGetPaypalAgreements"
	case EMsg_AMGetPaypalAgreementsResponse:
		return "EMsg_AMGetPaypalAgreementsResponse"
	case EMsg_AMGameServerPlayerCompatibilityCheck:
		return "EMsg_AMGameServerPlayerCompatibilityCheck"
	case EMsg_AMGameServerPlayerCompatibilityCheckResponse:
		return "EMsg_AMGameServerPlayerCompatibilityCheckResponse"
	case EMsg_AMRenewLicense:
		return "EMsg_AMRenewLicense"
	case EMsg_AMGetAccountCommunityBanInfo:
		return "EMsg_AMGetAccountCommunityBanInfo"
	case EMsg_AMGetAccountCommunityBanInfoResponse:
		return "EMsg_AMGetAccountCommunityBanInfoResponse"
	case EMsg_AMGameServerAccountChangePassword:
		return "EMsg_AMGameServerAccountChangePassword"
	case EMsg_AMGameServerAccountDeleteAccount:
		return "EMsg_AMGameServerAccountDeleteAccount"
	case EMsg_AMRenewAgreement:
		return "EMsg_AMRenewAgreement"
	case EMsg_AMSendEmail:
		return "EMsg_AMSendEmail"
	case EMsg_AMXsollaPayment:
		return "EMsg_AMXsollaPayment"
	case EMsg_AMXsollaPaymentResponse:
		return "EMsg_AMXsollaPaymentResponse"
	case EMsg_AMAcctAllowedToPurchase:
		return "EMsg_AMAcctAllowedToPurchase"
	case EMsg_AMAcctAllowedToPurchaseResponse:
		return "EMsg_AMAcctAllowedToPurchaseResponse"
	case EMsg_AMSwapKioskDeposit:
		return "EMsg_AMSwapKioskDeposit"
	case EMsg_AMSwapKioskDepositResponse:
		return "EMsg_AMSwapKioskDepositResponse"
	case EMsg_AMSetUserGiftUnowned:
		return "EMsg_AMSetUserGiftUnowned"
	case EMsg_AMSetUserGiftUnownedResponse:
		return "EMsg_AMSetUserGiftUnownedResponse"
	case EMsg_AMClaimUnownedUserGift:
		return "EMsg_AMClaimUnownedUserGift"
	case EMsg_AMClaimUnownedUserGiftResponse:
		return "EMsg_AMClaimUnownedUserGiftResponse"
	case EMsg_AMSetClanName:
		return "EMsg_AMSetClanName"
	case EMsg_AMSetClanNameResponse:
		return "EMsg_AMSetClanNameResponse"
	case EMsg_AMGrantCoupon:
		return "EMsg_AMGrantCoupon"
	case EMsg_AMGrantCouponResponse:
		return "EMsg_AMGrantCouponResponse"
	case EMsg_AMIsPackageRestrictedInUserCountry:
		return "EMsg_AMIsPackageRestrictedInUserCountry"
	case EMsg_AMIsPackageRestrictedInUserCountryResponse:
		return "EMsg_AMIsPackageRestrictedInUserCountryResponse"
	case EMsg_AMHandlePendingTransactionResponse:
		return "EMsg_AMHandlePendingTransactionResponse"
	case EMsg_AMGrantGuestPasses2:
		return "EMsg_AMGrantGuestPasses2"
	case EMsg_AMGrantGuestPasses2Response:
		return "EMsg_AMGrantGuestPasses2Response"
	case EMsg_AMSessionQuery:
		return "EMsg_AMSessionQuery"
	case EMsg_AMSessionQueryResponse:
		return "EMsg_AMSessionQueryResponse"
	case EMsg_AMGetPlayerBanDetails:
		return "EMsg_AMGetPlayerBanDetails"
	case EMsg_AMGetPlayerBanDetailsResponse:
		return "EMsg_AMGetPlayerBanDetailsResponse"
	case EMsg_AMFinalizePurchase:
		return "EMsg_AMFinalizePurchase"
	case EMsg_AMFinalizePurchaseResponse:
		return "EMsg_AMFinalizePurchaseResponse"
	case EMsg_AMPersonaChangeResponse:
		return "EMsg_AMPersonaChangeResponse"
	case EMsg_AMGetClanDetailsForForumCreation:
		return "EMsg_AMGetClanDetailsForForumCreation"
	case EMsg_AMGetClanDetailsForForumCreationResponse:
		return "EMsg_AMGetClanDetailsForForumCreationResponse"
	case EMsg_AMGetPendingNotificationCount:
		return "EMsg_AMGetPendingNotificationCount"
	case EMsg_AMGetPendingNotificationCountResponse:
		return "EMsg_AMGetPendingNotificationCountResponse"
	case EMsg_AMPasswordHashUpgrade:
		return "EMsg_AMPasswordHashUpgrade"
	case EMsg_AMMoPayPayment:
		return "EMsg_AMMoPayPayment"
	case EMsg_AMMoPayPaymentResponse:
		return "EMsg_AMMoPayPaymentResponse"
	case EMsg_AMBoaCompraPayment:
		return "EMsg_AMBoaCompraPayment"
	case EMsg_AMBoaCompraPaymentResponse:
		return "EMsg_AMBoaCompraPaymentResponse"
	case EMsg_AMExpireCaptchaByGID:
		return "EMsg_AMExpireCaptchaByGID"
	case EMsg_AMCompleteExternalPurchase:
		return "EMsg_AMCompleteExternalPurchase"
	case EMsg_AMCompleteExternalPurchaseResponse:
		return "EMsg_AMCompleteExternalPurchaseResponse"
	case EMsg_AMResolveNegativeWalletCredits:
		return "EMsg_AMResolveNegativeWalletCredits"
	case EMsg_AMResolveNegativeWalletCreditsResponse:
		return "EMsg_AMResolveNegativeWalletCreditsResponse"
	case EMsg_AMPayelpPayment:
		return "EMsg_AMPayelpPayment"
	case EMsg_AMPayelpPaymentResponse:
		return "EMsg_AMPayelpPaymentResponse"
	case EMsg_AMPlayerGetClanBasicDetails:
		return "EMsg_AMPlayerGetClanBasicDetails"
	case EMsg_AMPlayerGetClanBasicDetailsResponse:
		return "EMsg_AMPlayerGetClanBasicDetailsResponse"
	case EMsg_BasePSRange:
		return "EMsg_BasePSRange"
	case EMsg_PSCreateShoppingCart:
		return "EMsg_PSCreateShoppingCart"
	case EMsg_PSCreateShoppingCartResponse:
		return "EMsg_PSCreateShoppingCartResponse"
	case EMsg_PSIsValidShoppingCart:
		return "EMsg_PSIsValidShoppingCart"
	case EMsg_PSIsValidShoppingCartResponse:
		return "EMsg_PSIsValidShoppingCartResponse"
	case EMsg_PSAddPackageToShoppingCart:
		return "EMsg_PSAddPackageToShoppingCart"
	case EMsg_PSAddPackageToShoppingCartResponse:
		return "EMsg_PSAddPackageToShoppingCartResponse"
	case EMsg_PSRemoveLineItemFromShoppingCart:
		return "EMsg_PSRemoveLineItemFromShoppingCart"
	case EMsg_PSRemoveLineItemFromShoppingCartResponse:
		return "EMsg_PSRemoveLineItemFromShoppingCartResponse"
	case EMsg_PSGetShoppingCartContents:
		return "EMsg_PSGetShoppingCartContents"
	case EMsg_PSGetShoppingCartContentsResponse:
		return "EMsg_PSGetShoppingCartContentsResponse"
	case EMsg_PSAddWalletCreditToShoppingCart:
		return "EMsg_PSAddWalletCreditToShoppingCart"
	case EMsg_PSAddWalletCreditToShoppingCartResponse:
		return "EMsg_PSAddWalletCreditToShoppingCartResponse"
	case EMsg_BaseUFSRange:
		return "EMsg_BaseUFSRange"
	case EMsg_ClientUFSUploadFileRequest:
		return "EMsg_ClientUFSUploadFileRequest"
	case EMsg_ClientUFSUploadFileResponse:
		return "EMsg_ClientUFSUploadFileResponse"
	case EMsg_ClientUFSUploadFileChunk:
		return "EMsg_ClientUFSUploadFileChunk"
	case EMsg_ClientUFSUploadFileFinished:
		return "EMsg_ClientUFSUploadFileFinished"
	case EMsg_ClientUFSGetFileListForApp:
		return "EMsg_ClientUFSGetFileListForApp"
	case EMsg_ClientUFSGetFileListForAppResponse:
		return "EMsg_ClientUFSGetFileListForAppResponse"
	case EMsg_ClientUFSDownloadRequest:
		return "EMsg_ClientUFSDownloadRequest"
	case EMsg_ClientUFSDownloadResponse:
		return "EMsg_ClientUFSDownloadResponse"
	case EMsg_ClientUFSDownloadChunk:
		return "EMsg_ClientUFSDownloadChunk"
	case EMsg_ClientUFSLoginRequest:
		return "EMsg_ClientUFSLoginRequest"
	case EMsg_ClientUFSLoginResponse:
		return "EMsg_ClientUFSLoginResponse"
	case EMsg_UFSReloadPartitionInfo:
		return "EMsg_UFSReloadPartitionInfo"
	case EMsg_ClientUFSTransferHeartbeat:
		return "EMsg_ClientUFSTransferHeartbeat"
	case EMsg_UFSSynchronizeFile:
		return "EMsg_UFSSynchronizeFile"
	case EMsg_UFSSynchronizeFileResponse:
		return "EMsg_UFSSynchronizeFileResponse"
	case EMsg_ClientUFSDeleteFileRequest:
		return "EMsg_ClientUFSDeleteFileRequest"
	case EMsg_ClientUFSDeleteFileResponse:
		return "EMsg_ClientUFSDeleteFileResponse"
	case EMsg_UFSDownloadRequest:
		return "EMsg_UFSDownloadRequest"
	case EMsg_UFSDownloadResponse:
		return "EMsg_UFSDownloadResponse"
	case EMsg_UFSDownloadChunk:
		return "EMsg_UFSDownloadChunk"
	case EMsg_ClientUFSGetUGCDetails:
		return "EMsg_ClientUFSGetUGCDetails"
	case EMsg_ClientUFSGetUGCDetailsResponse:
		return "EMsg_ClientUFSGetUGCDetailsResponse"
	case EMsg_UFSUpdateFileFlags:
		return "EMsg_UFSUpdateFileFlags"
	case EMsg_UFSUpdateFileFlagsResponse:
		return "EMsg_UFSUpdateFileFlagsResponse"
	case EMsg_ClientUFSGetSingleFileInfo:
		return "EMsg_ClientUFSGetSingleFileInfo"
	case EMsg_ClientUFSGetSingleFileInfoResponse:
		return "EMsg_ClientUFSGetSingleFileInfoResponse"
	case EMsg_ClientUFSShareFile:
		return "EMsg_ClientUFSShareFile"
	case EMsg_ClientUFSShareFileResponse:
		return "EMsg_ClientUFSShareFileResponse"
	case EMsg_UFSReloadAccount:
		return "EMsg_UFSReloadAccount"
	case EMsg_UFSReloadAccountResponse:
		return "EMsg_UFSReloadAccountResponse"
	case EMsg_UFSUpdateRecordBatched:
		return "EMsg_UFSUpdateRecordBatched"
	case EMsg_UFSUpdateRecordBatchedResponse:
		return "EMsg_UFSUpdateRecordBatchedResponse"
	case EMsg_UFSMigrateFile:
		return "EMsg_UFSMigrateFile"
	case EMsg_UFSMigrateFileResponse:
		return "EMsg_UFSMigrateFileResponse"
	case EMsg_UFSGetUGCURLs:
		return "EMsg_UFSGetUGCURLs"
	case EMsg_UFSGetUGCURLsResponse:
		return "EMsg_UFSGetUGCURLsResponse"
	case EMsg_UFSHttpUploadFileFinishRequest:
		return "EMsg_UFSHttpUploadFileFinishRequest"
	case EMsg_UFSHttpUploadFileFinishResponse:
		return "EMsg_UFSHttpUploadFileFinishResponse"
	case EMsg_UFSDownloadStartRequest:
		return "EMsg_UFSDownloadStartRequest"
	case EMsg_UFSDownloadStartResponse:
		return "EMsg_UFSDownloadStartResponse"
	case EMsg_UFSDownloadChunkRequest:
		return "EMsg_UFSDownloadChunkRequest"
	case EMsg_UFSDownloadChunkResponse:
		return "EMsg_UFSDownloadChunkResponse"
	case EMsg_UFSDownloadFinishRequest:
		return "EMsg_UFSDownloadFinishRequest"
	case EMsg_UFSDownloadFinishResponse:
		return "EMsg_UFSDownloadFinishResponse"
	case EMsg_UFSFlushURLCache:
		return "EMsg_UFSFlushURLCache"
	case EMsg_UFSUploadCommit:
		return "EMsg_UFSUploadCommit"
	case EMsg_UFSUploadCommitResponse:
		return "EMsg_UFSUploadCommitResponse"
	case EMsg_BaseClient2:
		return "EMsg_BaseClient2"
	case EMsg_ClientRequestForgottenPasswordEmail:
		return "EMsg_ClientRequestForgottenPasswordEmail"
	case EMsg_ClientRequestForgottenPasswordEmailResponse:
		return "EMsg_ClientRequestForgottenPasswordEmailResponse"
	case EMsg_ClientCreateAccountResponse:
		return "EMsg_ClientCreateAccountResponse"
	case EMsg_ClientResetForgottenPassword:
		return "EMsg_ClientResetForgottenPassword"
	case EMsg_ClientResetForgottenPasswordResponse:
		return "EMsg_ClientResetForgottenPasswordResponse"
	case EMsg_ClientCreateAccount2:
		return "EMsg_ClientCreateAccount2"
	case EMsg_ClientInformOfResetForgottenPassword:
		return "EMsg_ClientInformOfResetForgottenPassword"
	case EMsg_ClientInformOfResetForgottenPasswordResponse:
		return "EMsg_ClientInformOfResetForgottenPasswordResponse"
	case EMsg_ClientAnonUserLogOn_Deprecated:
		return "EMsg_ClientAnonUserLogOn_Deprecated"
	case EMsg_ClientGamesPlayedWithDataBlob:
		return "EMsg_ClientGamesPlayedWithDataBlob"
	case EMsg_ClientUpdateUserGameInfo:
		return "EMsg_ClientUpdateUserGameInfo"
	case EMsg_ClientFileToDownload:
		return "EMsg_ClientFileToDownload"
	case EMsg_ClientFileToDownloadResponse:
		return "EMsg_ClientFileToDownloadResponse"
	case EMsg_ClientLBSSetScore:
		return "EMsg_ClientLBSSetScore"
	case EMsg_ClientLBSSetScoreResponse:
		return "EMsg_ClientLBSSetScoreResponse"
	case EMsg_ClientLBSFindOrCreateLB:
		return "EMsg_ClientLBSFindOrCreateLB"
	case EMsg_ClientLBSFindOrCreateLBResponse:
		return "EMsg_ClientLBSFindOrCreateLBResponse"
	case EMsg_ClientLBSGetLBEntries:
		return "EMsg_ClientLBSGetLBEntries"
	case EMsg_ClientLBSGetLBEntriesResponse:
		return "EMsg_ClientLBSGetLBEntriesResponse"
	case EMsg_ClientMarketingMessageUpdate:
		return "EMsg_ClientMarketingMessageUpdate"
	case EMsg_ClientChatDeclined:
		return "EMsg_ClientChatDeclined"
	case EMsg_ClientFriendMsgIncoming:
		return "EMsg_ClientFriendMsgIncoming"
	case EMsg_ClientAuthList_Deprecated:
		return "EMsg_ClientAuthList_Deprecated"
	case EMsg_ClientTicketAuthComplete:
		return "EMsg_ClientTicketAuthComplete"
	case EMsg_ClientIsLimitedAccount:
		return "EMsg_ClientIsLimitedAccount"
	case EMsg_ClientRequestAuthList:
		return "EMsg_ClientRequestAuthList"
	case EMsg_ClientAuthList:
		return "EMsg_ClientAuthList"
	case EMsg_ClientStat:
		return "EMsg_ClientStat"
	case EMsg_ClientP2PConnectionInfo:
		return "EMsg_ClientP2PConnectionInfo"
	case EMsg_ClientP2PConnectionFailInfo:
		return "EMsg_ClientP2PConnectionFailInfo"
	case EMsg_ClientGetNumberOfCurrentPlayers:
		return "EMsg_ClientGetNumberOfCurrentPlayers"
	case EMsg_ClientGetNumberOfCurrentPlayersResponse:
		return "EMsg_ClientGetNumberOfCurrentPlayersResponse"
	case EMsg_ClientGetDepotDecryptionKey:
		return "EMsg_ClientGetDepotDecryptionKey"
	case EMsg_ClientGetDepotDecryptionKeyResponse:
		return "EMsg_ClientGetDepotDecryptionKeyResponse"
	case EMsg_GSPerformHardwareSurvey:
		return "EMsg_GSPerformHardwareSurvey"
	case EMsg_ClientGetAppBetaPasswords:
		return "EMsg_ClientGetAppBetaPasswords"
	case EMsg_ClientGetAppBetaPasswordsResponse:
		return "EMsg_ClientGetAppBetaPasswordsResponse"
	case EMsg_ClientEnableTestLicense:
		return "EMsg_ClientEnableTestLicense"
	case EMsg_ClientEnableTestLicenseResponse:
		return "EMsg_ClientEnableTestLicenseResponse"
	case EMsg_ClientDisableTestLicense:
		return "EMsg_ClientDisableTestLicense"
	case EMsg_ClientDisableTestLicenseResponse:
		return "EMsg_ClientDisableTestLicenseResponse"
	case EMsg_ClientRequestValidationMail:
		return "EMsg_ClientRequestValidationMail"
	case EMsg_ClientRequestValidationMailResponse:
		return "EMsg_ClientRequestValidationMailResponse"
	case EMsg_ClientToGC:
		return "EMsg_ClientToGC"
	case EMsg_ClientFromGC:
		return "EMsg_ClientFromGC"
	case EMsg_ClientRequestChangeMail:
		return "EMsg_ClientRequestChangeMail"
	case EMsg_ClientRequestChangeMailResponse:
		return "EMsg_ClientRequestChangeMailResponse"
	case EMsg_ClientEmailAddrInfo:
		return "EMsg_ClientEmailAddrInfo"
	case EMsg_ClientPasswordChange3:
		return "EMsg_ClientPasswordChange3"
	case EMsg_ClientEmailChange3:
		return "EMsg_ClientEmailChange3"
	case EMsg_ClientPersonalQAChange3:
		return "EMsg_ClientPersonalQAChange3"
	case EMsg_ClientResetForgottenPassword3:
		return "EMsg_ClientResetForgottenPassword3"
	case EMsg_ClientRequestForgottenPasswordEmail3:
		return "EMsg_ClientRequestForgottenPasswordEmail3"
	case EMsg_ClientCreateAccount3:
		return "EMsg_ClientCreateAccount3"
	case EMsg_ClientNewLoginKey:
		return "EMsg_ClientNewLoginKey"
	case EMsg_ClientNewLoginKeyAccepted:
		return "EMsg_ClientNewLoginKeyAccepted"
	case EMsg_ClientLogOnWithHash_Deprecated:
		return "EMsg_ClientLogOnWithHash_Deprecated"
	case EMsg_ClientStoreUserStats2:
		return "EMsg_ClientStoreUserStats2"
	case EMsg_ClientStatsUpdated:
		return "EMsg_ClientStatsUpdated"
	case EMsg_ClientActivateOEMLicense:
		return "EMsg_ClientActivateOEMLicense"
	case EMsg_ClientRegisterOEMMachine:
		return "EMsg_ClientRegisterOEMMachine"
	case EMsg_ClientRegisterOEMMachineResponse:
		return "EMsg_ClientRegisterOEMMachineResponse"
	case EMsg_ClientRequestedClientStats:
		return "EMsg_ClientRequestedClientStats"
	case EMsg_ClientStat2Int32:
		return "EMsg_ClientStat2Int32"
	case EMsg_ClientStat2:
		return "EMsg_ClientStat2"
	case EMsg_ClientVerifyPassword:
		return "EMsg_ClientVerifyPassword"
	case EMsg_ClientVerifyPasswordResponse:
		return "EMsg_ClientVerifyPasswordResponse"
	case EMsg_ClientDRMDownloadRequest:
		return "EMsg_ClientDRMDownloadRequest"
	case EMsg_ClientDRMDownloadResponse:
		return "EMsg_ClientDRMDownloadResponse"
	case EMsg_ClientDRMFinalResult:
		return "EMsg_ClientDRMFinalResult"
	case EMsg_ClientGetFriendsWhoPlayGame:
		return "EMsg_ClientGetFriendsWhoPlayGame"
	case EMsg_ClientGetFriendsWhoPlayGameResponse:
		return "EMsg_ClientGetFriendsWhoPlayGameResponse"
	case EMsg_ClientOGSBeginSession:
		return "EMsg_ClientOGSBeginSession"
	case EMsg_ClientOGSBeginSessionResponse:
		return "EMsg_ClientOGSBeginSessionResponse"
	case EMsg_ClientOGSEndSession:
		return "EMsg_ClientOGSEndSession"
	case EMsg_ClientOGSEndSessionResponse:
		return "EMsg_ClientOGSEndSessionResponse"
	case EMsg_ClientOGSWriteRow:
		return "EMsg_ClientOGSWriteRow"
	case EMsg_ClientDRMTest:
		return "EMsg_ClientDRMTest"
	case EMsg_ClientDRMTestResult:
		return "EMsg_ClientDRMTestResult"
	case EMsg_ClientServerUnavailable:
		return "EMsg_ClientServerUnavailable"
	case EMsg_ClientServersAvailable:
		return "EMsg_ClientServersAvailable"
	case EMsg_ClientRegisterAuthTicketWithCM:
		return "EMsg_ClientRegisterAuthTicketWithCM"
	case EMsg_ClientGCMsgFailed:
		return "EMsg_ClientGCMsgFailed"
	case EMsg_ClientMicroTxnAuthRequest:
		return "EMsg_ClientMicroTxnAuthRequest"
	case EMsg_ClientMicroTxnAuthorize:
		return "EMsg_ClientMicroTxnAuthorize"
	case EMsg_ClientMicroTxnAuthorizeResponse:
		return "EMsg_ClientMicroTxnAuthorizeResponse"
	case EMsg_ClientAppMinutesPlayedData:
		return "EMsg_ClientAppMinutesPlayedData"
	case EMsg_ClientGetMicroTxnInfo:
		return "EMsg_ClientGetMicroTxnInfo"
	case EMsg_ClientGetMicroTxnInfoResponse:
		return "EMsg_ClientGetMicroTxnInfoResponse"
	case EMsg_ClientMarketingMessageUpdate2:
		return "EMsg_ClientMarketingMessageUpdate2"
	case EMsg_ClientDeregisterWithServer:
		return "EMsg_ClientDeregisterWithServer"
	case EMsg_ClientSubscribeToPersonaFeed:
		return "EMsg_ClientSubscribeToPersonaFeed"
	case EMsg_ClientLogon:
		return "EMsg_ClientLogon"
	case EMsg_ClientGetClientDetails:
		return "EMsg_ClientGetClientDetails"
	case EMsg_ClientGetClientDetailsResponse:
		return "EMsg_ClientGetClientDetailsResponse"
	case EMsg_ClientReportOverlayDetourFailure:
		return "EMsg_ClientReportOverlayDetourFailure"
	case EMsg_ClientGetClientAppList:
		return "EMsg_ClientGetClientAppList"
	case EMsg_ClientGetClientAppListResponse:
		return "EMsg_ClientGetClientAppListResponse"
	case EMsg_ClientInstallClientApp:
		return "EMsg_ClientInstallClientApp"
	case EMsg_ClientInstallClientAppResponse:
		return "EMsg_ClientInstallClientAppResponse"
	case EMsg_ClientUninstallClientApp:
		return "EMsg_ClientUninstallClientApp"
	case EMsg_ClientUninstallClientAppResponse:
		return "EMsg_ClientUninstallClientAppResponse"
	case EMsg_ClientSetClientAppUpdateState:
		return "EMsg_ClientSetClientAppUpdateState"
	case EMsg_ClientSetClientAppUpdateStateResponse:
		return "EMsg_ClientSetClientAppUpdateStateResponse"
	case EMsg_ClientRequestEncryptedAppTicket:
		return "EMsg_ClientRequestEncryptedAppTicket"
	case EMsg_ClientRequestEncryptedAppTicketResponse:
		return "EMsg_ClientRequestEncryptedAppTicketResponse"
	case EMsg_ClientWalletInfoUpdate:
		return "EMsg_ClientWalletInfoUpdate"
	case EMsg_ClientLBSSetUGC:
		return "EMsg_ClientLBSSetUGC"
	case EMsg_ClientLBSSetUGCResponse:
		return "EMsg_ClientLBSSetUGCResponse"
	case EMsg_ClientAMGetClanOfficers:
		return "EMsg_ClientAMGetClanOfficers"
	case EMsg_ClientAMGetClanOfficersResponse:
		return "EMsg_ClientAMGetClanOfficersResponse"
	case EMsg_ClientCheckFileSignature:
		return "EMsg_ClientCheckFileSignature"
	case EMsg_ClientCheckFileSignatureResponse:
		return "EMsg_ClientCheckFileSignatureResponse"
	case EMsg_ClientFriendProfileInfo:
		return "EMsg_ClientFriendProfileInfo"
	case EMsg_ClientFriendProfileInfoResponse:
		return "EMsg_ClientFriendProfileInfoResponse"
	case EMsg_ClientUpdateMachineAuth:
		return "EMsg_ClientUpdateMachineAuth"
	case EMsg_ClientUpdateMachineAuthResponse:
		return "EMsg_ClientUpdateMachineAuthResponse"
	case EMsg_ClientReadMachineAuth:
		return "EMsg_ClientReadMachineAuth"
	case EMsg_ClientReadMachineAuthResponse:
		return "EMsg_ClientReadMachineAuthResponse"
	case EMsg_ClientRequestMachineAuth:
		return "EMsg_ClientRequestMachineAuth"
	case EMsg_ClientRequestMachineAuthResponse:
		return "EMsg_ClientRequestMachineAuthResponse"
	case EMsg_ClientScreenshotsChanged:
		return "EMsg_ClientScreenshotsChanged"
	case EMsg_ClientEmailChange4:
		return "EMsg_ClientEmailChange4"
	case EMsg_ClientEmailChangeResponse4:
		return "EMsg_ClientEmailChangeResponse4"
	case EMsg_ClientGetCDNAuthToken:
		return "EMsg_ClientGetCDNAuthToken"
	case EMsg_ClientGetCDNAuthTokenResponse:
		return "EMsg_ClientGetCDNAuthTokenResponse"
	case EMsg_ClientDownloadRateStatistics:
		return "EMsg_ClientDownloadRateStatistics"
	case EMsg_ClientRequestAccountData:
		return "EMsg_ClientRequestAccountData"
	case EMsg_ClientRequestAccountDataResponse:
		return "EMsg_ClientRequestAccountDataResponse"
	case EMsg_ClientResetForgottenPassword4:
		return "EMsg_ClientResetForgottenPassword4"
	case EMsg_ClientHideFriend:
		return "EMsg_ClientHideFriend"
	case EMsg_ClientFriendsGroupsList:
		return "EMsg_ClientFriendsGroupsList"
	case EMsg_ClientGetClanActivityCounts:
		return "EMsg_ClientGetClanActivityCounts"
	case EMsg_ClientGetClanActivityCountsResponse:
		return "EMsg_ClientGetClanActivityCountsResponse"
	case EMsg_ClientOGSReportString:
		return "EMsg_ClientOGSReportString"
	case EMsg_ClientOGSReportBug:
		return "EMsg_ClientOGSReportBug"
	case EMsg_ClientSentLogs:
		return "EMsg_ClientSentLogs"
	case EMsg_ClientLogonGameServer:
		return "EMsg_ClientLogonGameServer"
	case EMsg_AMClientCreateFriendsGroup:
		return "EMsg_AMClientCreateFriendsGroup"
	case EMsg_AMClientCreateFriendsGroupResponse:
		return "EMsg_AMClientCreateFriendsGroupResponse"
	case EMsg_AMClientDeleteFriendsGroup:
		return "EMsg_AMClientDeleteFriendsGroup"
	case EMsg_AMClientDeleteFriendsGroupResponse:
		return "EMsg_AMClientDeleteFriendsGroupResponse"
	case EMsg_AMClientRenameFriendsGroup:
		return "EMsg_AMClientRenameFriendsGroup"
	case EMsg_AMClientRenameFriendsGroupResponse:
		return "EMsg_AMClientRenameFriendsGroupResponse"
	case EMsg_AMClientAddFriendToGroup:
		return "EMsg_AMClientAddFriendToGroup"
	case EMsg_AMClientAddFriendToGroupResponse:
		return "EMsg_AMClientAddFriendToGroupResponse"
	case EMsg_AMClientRemoveFriendFromGroup:
		return "EMsg_AMClientRemoveFriendFromGroup"
	case EMsg_AMClientRemoveFriendFromGroupResponse:
		return "EMsg_AMClientRemoveFriendFromGroupResponse"
	case EMsg_ClientAMGetPersonaNameHistory:
		return "EMsg_ClientAMGetPersonaNameHistory"
	case EMsg_ClientAMGetPersonaNameHistoryResponse:
		return "EMsg_ClientAMGetPersonaNameHistoryResponse"
	case EMsg_ClientRequestFreeLicense:
		return "EMsg_ClientRequestFreeLicense"
	case EMsg_ClientRequestFreeLicenseResponse:
		return "EMsg_ClientRequestFreeLicenseResponse"
	case EMsg_ClientDRMDownloadRequestWithCrashData:
		return "EMsg_ClientDRMDownloadRequestWithCrashData"
	case EMsg_ClientAuthListAck:
		return "EMsg_ClientAuthListAck"
	case EMsg_ClientItemAnnouncements:
		return "EMsg_ClientItemAnnouncements"
	case EMsg_ClientRequestItemAnnouncements:
		return "EMsg_ClientRequestItemAnnouncements"
	case EMsg_ClientFriendMsgEchoToSender:
		return "EMsg_ClientFriendMsgEchoToSender"
	case EMsg_ClientChangeSteamGuardOptions:
		return "EMsg_ClientChangeSteamGuardOptions"
	case EMsg_ClientChangeSteamGuardOptionsResponse:
		return "EMsg_ClientChangeSteamGuardOptionsResponse"
	case EMsg_ClientOGSGameServerPingSample:
		return "EMsg_ClientOGSGameServerPingSample"
	case EMsg_ClientCommentNotifications:
		return "EMsg_ClientCommentNotifications"
	case EMsg_ClientRequestCommentNotifications:
		return "EMsg_ClientRequestCommentNotifications"
	case EMsg_ClientPersonaChangeResponse:
		return "EMsg_ClientPersonaChangeResponse"
	case EMsg_ClientRequestWebAPIAuthenticateUserNonce:
		return "EMsg_ClientRequestWebAPIAuthenticateUserNonce"
	case EMsg_ClientRequestWebAPIAuthenticateUserNonceResponse:
		return "EMsg_ClientRequestWebAPIAuthenticateUserNonceResponse"
	case EMsg_ClientPlayerNicknameList:
		return "EMsg_ClientPlayerNicknameList"
	case EMsg_AMClientSetPlayerNickname:
		return "EMsg_AMClientSetPlayerNickname"
	case EMsg_AMClientSetPlayerNicknameResponse:
		return "EMsg_AMClientSetPlayerNicknameResponse"
	case EMsg_ClientRequestOAuthTokenForApp:
		return "EMsg_ClientRequestOAuthTokenForApp"
	case EMsg_ClientRequestOAuthTokenForAppResponse:
		return "EMsg_ClientRequestOAuthTokenForAppResponse"
	case EMsg_ClientGetNumberOfCurrentPlayersDP:
		return "EMsg_ClientGetNumberOfCurrentPlayersDP"
	case EMsg_ClientGetNumberOfCurrentPlayersDPResponse:
		return "EMsg_ClientGetNumberOfCurrentPlayersDPResponse"
	case EMsg_ClientServiceMethod:
		return "EMsg_ClientServiceMethod"
	case EMsg_ClientServiceMethodResponse:
		return "EMsg_ClientServiceMethodResponse"
	case EMsg_ClientFriendUserStatusPublished:
		return "EMsg_ClientFriendUserStatusPublished"
	case EMsg_ClientCurrentUIMode:
		return "EMsg_ClientCurrentUIMode"
	case EMsg_ClientVanityURLChangedNotification:
		return "EMsg_ClientVanityURLChangedNotification"
	case EMsg_ClientUserNotifications:
		return "EMsg_ClientUserNotifications"
	case EMsg_BaseDFS:
		return "EMsg_BaseDFS"
	case EMsg_DFSGetFile:
		return "EMsg_DFSGetFile"
	case EMsg_DFSInstallLocalFile:
		return "EMsg_DFSInstallLocalFile"
	case EMsg_DFSConnection:
		return "EMsg_DFSConnection"
	case EMsg_DFSConnectionReply:
		return "EMsg_DFSConnectionReply"
	case EMsg_ClientDFSAuthenticateRequest:
		return "EMsg_ClientDFSAuthenticateRequest"
	case EMsg_ClientDFSAuthenticateResponse:
		return "EMsg_ClientDFSAuthenticateResponse"
	case EMsg_ClientDFSEndSession:
		return "EMsg_ClientDFSEndSession"
	case EMsg_DFSPurgeFile:
		return "EMsg_DFSPurgeFile"
	case EMsg_DFSRouteFile:
		return "EMsg_DFSRouteFile"
	case EMsg_DFSGetFileFromServer:
		return "EMsg_DFSGetFileFromServer"
	case EMsg_DFSAcceptedResponse:
		return "EMsg_DFSAcceptedResponse"
	case EMsg_DFSRequestPingback:
		return "EMsg_DFSRequestPingback"
	case EMsg_DFSRecvTransmitFile:
		return "EMsg_DFSRecvTransmitFile"
	case EMsg_DFSSendTransmitFile:
		return "EMsg_DFSSendTransmitFile"
	case EMsg_DFSRequestPingback2:
		return "EMsg_DFSRequestPingback2"
	case EMsg_DFSResponsePingback2:
		return "EMsg_DFSResponsePingback2"
	case EMsg_ClientDFSDownloadStatus:
		return "EMsg_ClientDFSDownloadStatus"
	case EMsg_DFSStartTransfer:
		return "EMsg_DFSStartTransfer"
	case EMsg_DFSTransferComplete:
		return "EMsg_DFSTransferComplete"
	case EMsg_BaseMDS:
		return "EMsg_BaseMDS"
	case EMsg_ClientMDSLoginRequest:
		return "EMsg_ClientMDSLoginRequest"
	case EMsg_ClientMDSLoginResponse:
		return "EMsg_ClientMDSLoginResponse"
	case EMsg_ClientMDSUploadManifestRequest:
		return "EMsg_ClientMDSUploadManifestRequest"
	case EMsg_ClientMDSUploadManifestResponse:
		return "EMsg_ClientMDSUploadManifestResponse"
	case EMsg_ClientMDSTransmitManifestDataChunk:
		return "EMsg_ClientMDSTransmitManifestDataChunk"
	case EMsg_ClientMDSHeartbeat:
		return "EMsg_ClientMDSHeartbeat"
	case EMsg_ClientMDSUploadDepotChunks:
		return "EMsg_ClientMDSUploadDepotChunks"
	case EMsg_ClientMDSUploadDepotChunksResponse:
		return "EMsg_ClientMDSUploadDepotChunksResponse"
	case EMsg_ClientMDSInitDepotBuildRequest:
		return "EMsg_ClientMDSInitDepotBuildRequest"
	case EMsg_ClientMDSInitDepotBuildResponse:
		return "EMsg_ClientMDSInitDepotBuildResponse"
	case EMsg_AMToMDSGetDepotDecryptionKey:
		return "EMsg_AMToMDSGetDepotDecryptionKey"
	case EMsg_MDSToAMGetDepotDecryptionKeyResponse:
		return "EMsg_MDSToAMGetDepotDecryptionKeyResponse"
	case EMsg_MDSGetVersionsForDepot:
		return "EMsg_MDSGetVersionsForDepot"
	case EMsg_MDSGetVersionsForDepotResponse:
		return "EMsg_MDSGetVersionsForDepotResponse"
	case EMsg_MDSSetPublicVersionForDepot:
		return "EMsg_MDSSetPublicVersionForDepot"
	case EMsg_MDSSetPublicVersionForDepotResponse:
		return "EMsg_MDSSetPublicVersionForDepotResponse"
	case EMsg_ClientMDSGetDepotManifest:
		return "EMsg_ClientMDSGetDepotManifest"
	case EMsg_ClientMDSGetDepotManifestResponse:
		return "EMsg_ClientMDSGetDepotManifestResponse"
	case EMsg_ClientMDSGetDepotManifestChunk:
		return "EMsg_ClientMDSGetDepotManifestChunk"
	case EMsg_ClientMDSUploadRateTest:
		return "EMsg_ClientMDSUploadRateTest"
	case EMsg_ClientMDSUploadRateTestResponse:
		return "EMsg_ClientMDSUploadRateTestResponse"
	case EMsg_MDSDownloadDepotChunksAck:
		return "EMsg_MDSDownloadDepotChunksAck"
	case EMsg_MDSContentServerStatsBroadcast:
		return "EMsg_MDSContentServerStatsBroadcast"
	case EMsg_MDSContentServerConfigRequest:
		return "EMsg_MDSContentServerConfigRequest"
	case EMsg_MDSContentServerConfig:
		return "EMsg_MDSContentServerConfig"
	case EMsg_MDSGetDepotManifest:
		return "EMsg_MDSGetDepotManifest"
	case EMsg_MDSGetDepotManifestResponse:
		return "EMsg_MDSGetDepotManifestResponse"
	case EMsg_MDSGetDepotManifestChunk:
		return "EMsg_MDSGetDepotManifestChunk"
	case EMsg_MDSGetDepotChunk:
		return "EMsg_MDSGetDepotChunk"
	case EMsg_MDSGetDepotChunkResponse:
		return "EMsg_MDSGetDepotChunkResponse"
	case EMsg_MDSGetDepotChunkChunk:
		return "EMsg_MDSGetDepotChunkChunk"
	case EMsg_MDSUpdateContentServerConfig:
		return "EMsg_MDSUpdateContentServerConfig"
	case EMsg_MDSGetServerListForUser:
		return "EMsg_MDSGetServerListForUser"
	case EMsg_MDSGetServerListForUserResponse:
		return "EMsg_MDSGetServerListForUserResponse"
	case EMsg_ClientMDSRegisterAppBuild:
		return "EMsg_ClientMDSRegisterAppBuild"
	case EMsg_ClientMDSRegisterAppBuildResponse:
		return "EMsg_ClientMDSRegisterAppBuildResponse"
	case EMsg_ClientMDSSetAppBuildLive:
		return "EMsg_ClientMDSSetAppBuildLive"
	case EMsg_ClientMDSSetAppBuildLiveResponse:
		return "EMsg_ClientMDSSetAppBuildLiveResponse"
	case EMsg_ClientMDSGetPrevDepotBuild:
		return "EMsg_ClientMDSGetPrevDepotBuild"
	case EMsg_ClientMDSGetPrevDepotBuildResponse:
		return "EMsg_ClientMDSGetPrevDepotBuildResponse"
	case EMsg_MDSToCSFlushChunk:
		return "EMsg_MDSToCSFlushChunk"
	case EMsg_ClientMDSSignInstallScript:
		return "EMsg_ClientMDSSignInstallScript"
	case EMsg_ClientMDSSignInstallScriptResponse:
		return "EMsg_ClientMDSSignInstallScriptResponse"
	case EMsg_CSBase:
		return "EMsg_CSBase"
	case EMsg_CSPing:
		return "EMsg_CSPing"
	case EMsg_CSPingResponse:
		return "EMsg_CSPingResponse"
	case EMsg_GMSBase:
		return "EMsg_GMSBase"
	case EMsg_GMSGameServerReplicate:
		return "EMsg_GMSGameServerReplicate"
	case EMsg_ClientGMSServerQuery:
		return "EMsg_ClientGMSServerQuery"
	case EMsg_GMSClientServerQueryResponse:
		return "EMsg_GMSClientServerQueryResponse"
	case EMsg_AMGMSGameServerUpdate:
		return "EMsg_AMGMSGameServerUpdate"
	case EMsg_AMGMSGameServerRemove:
		return "EMsg_AMGMSGameServerRemove"
	case EMsg_GameServerOutOfDate:
		return "EMsg_GameServerOutOfDate"
	case EMsg_ClientAuthorizeLocalDeviceRequest:
		return "EMsg_ClientAuthorizeLocalDeviceRequest"
	case EMsg_ClientAuthorizeLocalDevice:
		return "EMsg_ClientAuthorizeLocalDevice"
	case EMsg_ClientDeauthorizeDeviceRequest:
		return "EMsg_ClientDeauthorizeDeviceRequest"
	case EMsg_ClientDeauthorizeDevice:
		return "EMsg_ClientDeauthorizeDevice"
	case EMsg_ClientUseLocalDeviceAuthorizations:
		return "EMsg_ClientUseLocalDeviceAuthorizations"
	case EMsg_ClientGetAuthorizedDevices:
		return "EMsg_ClientGetAuthorizedDevices"
	case EMsg_ClientGetAuthorizedDevicesResponse:
		return "EMsg_ClientGetAuthorizedDevicesResponse"
	case EMsg_MMSBase:
		return "EMsg_MMSBase"
	case EMsg_ClientMMSCreateLobby:
		return "EMsg_ClientMMSCreateLobby"
	case EMsg_ClientMMSCreateLobbyResponse:
		return "EMsg_ClientMMSCreateLobbyResponse"
	case EMsg_ClientMMSJoinLobby:
		return "EMsg_ClientMMSJoinLobby"
	case EMsg_ClientMMSJoinLobbyResponse:
		return "EMsg_ClientMMSJoinLobbyResponse"
	case EMsg_ClientMMSLeaveLobby:
		return "EMsg_ClientMMSLeaveLobby"
	case EMsg_ClientMMSLeaveLobbyResponse:
		return "EMsg_ClientMMSLeaveLobbyResponse"
	case EMsg_ClientMMSGetLobbyList:
		return "EMsg_ClientMMSGetLobbyList"
	case EMsg_ClientMMSGetLobbyListResponse:
		return "EMsg_ClientMMSGetLobbyListResponse"
	case EMsg_ClientMMSSetLobbyData:
		return "EMsg_ClientMMSSetLobbyData"
	case EMsg_ClientMMSSetLobbyDataResponse:
		return "EMsg_ClientMMSSetLobbyDataResponse"
	case EMsg_ClientMMSGetLobbyData:
		return "EMsg_ClientMMSGetLobbyData"
	case EMsg_ClientMMSLobbyData:
		return "EMsg_ClientMMSLobbyData"
	case EMsg_ClientMMSSendLobbyChatMsg:
		return "EMsg_ClientMMSSendLobbyChatMsg"
	case EMsg_ClientMMSLobbyChatMsg:
		return "EMsg_ClientMMSLobbyChatMsg"
	case EMsg_ClientMMSSetLobbyOwner:
		return "EMsg_ClientMMSSetLobbyOwner"
	case EMsg_ClientMMSSetLobbyOwnerResponse:
		return "EMsg_ClientMMSSetLobbyOwnerResponse"
	case EMsg_ClientMMSSetLobbyGameServer:
		return "EMsg_ClientMMSSetLobbyGameServer"
	case EMsg_ClientMMSLobbyGameServerSet:
		return "EMsg_ClientMMSLobbyGameServerSet"
	case EMsg_ClientMMSUserJoinedLobby:
		return "EMsg_ClientMMSUserJoinedLobby"
	case EMsg_ClientMMSUserLeftLobby:
		return "EMsg_ClientMMSUserLeftLobby"
	case EMsg_ClientMMSInviteToLobby:
		return "EMsg_ClientMMSInviteToLobby"
	case EMsg_ClientMMSFlushFrenemyListCache:
		return "EMsg_ClientMMSFlushFrenemyListCache"
	case EMsg_ClientMMSFlushFrenemyListCacheResponse:
		return "EMsg_ClientMMSFlushFrenemyListCacheResponse"
	case EMsg_ClientMMSSetLobbyLinked:
		return "EMsg_ClientMMSSetLobbyLinked"
	case EMsg_NonStdMsgBase:
		return "EMsg_NonStdMsgBase"
	case EMsg_NonStdMsgMemcached:
		return "EMsg_NonStdMsgMemcached"
	case EMsg_NonStdMsgHTTPServer:
		return "EMsg_NonStdMsgHTTPServer"
	case EMsg_NonStdMsgHTTPClient:
		return "EMsg_NonStdMsgHTTPClient"
	case EMsg_NonStdMsgWGResponse:
		return "EMsg_NonStdMsgWGResponse"
	case EMsg_NonStdMsgPHPSimulator:
		return "EMsg_NonStdMsgPHPSimulator"
	case EMsg_NonStdMsgChase:
		return "EMsg_NonStdMsgChase"
	case EMsg_NonStdMsgDFSTransfer:
		return "EMsg_NonStdMsgDFSTransfer"
	case EMsg_NonStdMsgTests:
		return "EMsg_NonStdMsgTests"
	case EMsg_NonStdMsgUMQpipeAAPL:
		return "EMsg_NonStdMsgUMQpipeAAPL"
	case EMsg_NonStdMsgSyslog:
		return "EMsg_NonStdMsgSyslog"
	case EMsg_NonStdMsgLogsink:
		return "EMsg_NonStdMsgLogsink"
	case EMsg_UDSBase:
		return "EMsg_UDSBase"
	case EMsg_ClientUDSP2PSessionStarted:
		return "EMsg_ClientUDSP2PSessionStarted"
	case EMsg_ClientUDSP2PSessionEnded:
		return "EMsg_ClientUDSP2PSessionEnded"
	case EMsg_UDSRenderUserAuth:
		return "EMsg_UDSRenderUserAuth"
	case EMsg_UDSRenderUserAuthResponse:
		return "EMsg_UDSRenderUserAuthResponse"
	case EMsg_ClientUDSInviteToGame:
		return "EMsg_ClientUDSInviteToGame"
	case EMsg_UDSFindSession:
		return "EMsg_UDSFindSession"
	case EMsg_UDSFindSessionResponse:
		return "EMsg_UDSFindSessionResponse"
	case EMsg_MPASBase:
		return "EMsg_MPASBase"
	case EMsg_MPASVacBanReset:
		return "EMsg_MPASVacBanReset"
	case EMsg_KGSBase:
		return "EMsg_KGSBase"
	case EMsg_KGSAllocateKeyRange:
		return "EMsg_KGSAllocateKeyRange"
	case EMsg_KGSAllocateKeyRangeResponse:
		return "EMsg_KGSAllocateKeyRangeResponse"
	case EMsg_KGSGenerateKeys:
		return "EMsg_KGSGenerateKeys"
	case EMsg_KGSGenerateKeysResponse:
		return "EMsg_KGSGenerateKeysResponse"
	case EMsg_KGSRemapKeys:
		return "EMsg_KGSRemapKeys"
	case EMsg_KGSRemapKeysResponse:
		return "EMsg_KGSRemapKeysResponse"
	case EMsg_KGSGenerateGameStopWCKeys:
		return "EMsg_KGSGenerateGameStopWCKeys"
	case EMsg_KGSGenerateGameStopWCKeysResponse:
		return "EMsg_KGSGenerateGameStopWCKeysResponse"
	case EMsg_UCMBase:
		return "EMsg_UCMBase"
	case EMsg_ClientUCMAddScreenshot:
		return "EMsg_ClientUCMAddScreenshot"
	case EMsg_ClientUCMAddScreenshotResponse:
		return "EMsg_ClientUCMAddScreenshotResponse"
	case EMsg_UCMValidateObjectExists:
		return "EMsg_UCMValidateObjectExists"
	case EMsg_UCMValidateObjectExistsResponse:
		return "EMsg_UCMValidateObjectExistsResponse"
	case EMsg_UCMResetCommunityContent:
		return "EMsg_UCMResetCommunityContent"
	case EMsg_UCMResetCommunityContentResponse:
		return "EMsg_UCMResetCommunityContentResponse"
	case EMsg_ClientUCMDeleteScreenshot:
		return "EMsg_ClientUCMDeleteScreenshot"
	case EMsg_ClientUCMDeleteScreenshotResponse:
		return "EMsg_ClientUCMDeleteScreenshotResponse"
	case EMsg_ClientUCMPublishFile:
		return "EMsg_ClientUCMPublishFile"
	case EMsg_ClientUCMPublishFileResponse:
		return "EMsg_ClientUCMPublishFileResponse"
	case EMsg_ClientUCMGetPublishedFileDetails:
		return "EMsg_ClientUCMGetPublishedFileDetails"
	case EMsg_ClientUCMGetPublishedFileDetailsResponse:
		return "EMsg_ClientUCMGetPublishedFileDetailsResponse"
	case EMsg_ClientUCMDeletePublishedFile:
		return "EMsg_ClientUCMDeletePublishedFile"
	case EMsg_ClientUCMDeletePublishedFileResponse:
		return "EMsg_ClientUCMDeletePublishedFileResponse"
	case EMsg_ClientUCMEnumerateUserPublishedFiles:
		return "EMsg_ClientUCMEnumerateUserPublishedFiles"
	case EMsg_ClientUCMEnumerateUserPublishedFilesResponse:
		return "EMsg_ClientUCMEnumerateUserPublishedFilesResponse"
	case EMsg_ClientUCMSubscribePublishedFile:
		return "EMsg_ClientUCMSubscribePublishedFile"
	case EMsg_ClientUCMSubscribePublishedFileResponse:
		return "EMsg_ClientUCMSubscribePublishedFileResponse"
	case EMsg_ClientUCMEnumerateUserSubscribedFiles:
		return "EMsg_ClientUCMEnumerateUserSubscribedFiles"
	case EMsg_ClientUCMEnumerateUserSubscribedFilesResponse:
		return "EMsg_ClientUCMEnumerateUserSubscribedFilesResponse"
	case EMsg_ClientUCMUnsubscribePublishedFile:
		return "EMsg_ClientUCMUnsubscribePublishedFile"
	case EMsg_ClientUCMUnsubscribePublishedFileResponse:
		return "EMsg_ClientUCMUnsubscribePublishedFileResponse"
	case EMsg_ClientUCMUpdatePublishedFile:
		return "EMsg_ClientUCMUpdatePublishedFile"
	case EMsg_ClientUCMUpdatePublishedFileResponse:
		return "EMsg_ClientUCMUpdatePublishedFileResponse"
	case EMsg_UCMUpdatePublishedFile:
		return "EMsg_UCMUpdatePublishedFile"
	case EMsg_UCMUpdatePublishedFileResponse:
		return "EMsg_UCMUpdatePublishedFileResponse"
	case EMsg_UCMDeletePublishedFile:
		return "EMsg_UCMDeletePublishedFile"
	case EMsg_UCMDeletePublishedFileResponse:
		return "EMsg_UCMDeletePublishedFileResponse"
	case EMsg_UCMUpdatePublishedFileStat:
		return "EMsg_UCMUpdatePublishedFileStat"
	case EMsg_UCMUpdatePublishedFileBan:
		return "EMsg_UCMUpdatePublishedFileBan"
	case EMsg_UCMUpdatePublishedFileBanResponse:
		return "EMsg_UCMUpdatePublishedFileBanResponse"
	case EMsg_UCMUpdateTaggedScreenshot:
		return "EMsg_UCMUpdateTaggedScreenshot"
	case EMsg_UCMAddTaggedScreenshot:
		return "EMsg_UCMAddTaggedScreenshot"
	case EMsg_UCMRemoveTaggedScreenshot:
		return "EMsg_UCMRemoveTaggedScreenshot"
	case EMsg_UCMReloadPublishedFile:
		return "EMsg_UCMReloadPublishedFile"
	case EMsg_UCMReloadUserFileListCaches:
		return "EMsg_UCMReloadUserFileListCaches"
	case EMsg_UCMPublishedFileReported:
		return "EMsg_UCMPublishedFileReported"
	case EMsg_UCMUpdatePublishedFileIncompatibleStatus:
		return "EMsg_UCMUpdatePublishedFileIncompatibleStatus"
	case EMsg_UCMPublishedFilePreviewAdd:
		return "EMsg_UCMPublishedFilePreviewAdd"
	case EMsg_UCMPublishedFilePreviewAddResponse:
		return "EMsg_UCMPublishedFilePreviewAddResponse"
	case EMsg_UCMPublishedFilePreviewRemove:
		return "EMsg_UCMPublishedFilePreviewRemove"
	case EMsg_UCMPublishedFilePreviewRemoveResponse:
		return "EMsg_UCMPublishedFilePreviewRemoveResponse"
	case EMsg_UCMPublishedFilePreviewChangeSortOrder:
		return "EMsg_UCMPublishedFilePreviewChangeSortOrder"
	case EMsg_UCMPublishedFilePreviewChangeSortOrderResponse:
		return "EMsg_UCMPublishedFilePreviewChangeSortOrderResponse"
	case EMsg_ClientUCMPublishedFileSubscribed:
		return "EMsg_ClientUCMPublishedFileSubscribed"
	case EMsg_ClientUCMPublishedFileUnsubscribed:
		return "EMsg_ClientUCMPublishedFileUnsubscribed"
	case EMsg_UCMPublishedFileSubscribed:
		return "EMsg_UCMPublishedFileSubscribed"
	case EMsg_UCMPublishedFileUnsubscribed:
		return "EMsg_UCMPublishedFileUnsubscribed"
	case EMsg_UCMPublishFile:
		return "EMsg_UCMPublishFile"
	case EMsg_UCMPublishFileResponse:
		return "EMsg_UCMPublishFileResponse"
	case EMsg_UCMPublishedFileChildAdd:
		return "EMsg_UCMPublishedFileChildAdd"
	case EMsg_UCMPublishedFileChildAddResponse:
		return "EMsg_UCMPublishedFileChildAddResponse"
	case EMsg_UCMPublishedFileChildRemove:
		return "EMsg_UCMPublishedFileChildRemove"
	case EMsg_UCMPublishedFileChildRemoveResponse:
		return "EMsg_UCMPublishedFileChildRemoveResponse"
	case EMsg_UCMPublishedFileChildChangeSortOrder:
		return "EMsg_UCMPublishedFileChildChangeSortOrder"
	case EMsg_UCMPublishedFileChildChangeSortOrderResponse:
		return "EMsg_UCMPublishedFileChildChangeSortOrderResponse"
	case EMsg_UCMPublishedFileParentChanged:
		return "EMsg_UCMPublishedFileParentChanged"
	case EMsg_ClientUCMGetPublishedFilesForUser:
		return "EMsg_ClientUCMGetPublishedFilesForUser"
	case EMsg_ClientUCMGetPublishedFilesForUserResponse:
		return "EMsg_ClientUCMGetPublishedFilesForUserResponse"
	case EMsg_UCMGetPublishedFilesForUser:
		return "EMsg_UCMGetPublishedFilesForUser"
	case EMsg_UCMGetPublishedFilesForUserResponse:
		return "EMsg_UCMGetPublishedFilesForUserResponse"
	case EMsg_ClientUCMSetUserPublishedFileAction:
		return "EMsg_ClientUCMSetUserPublishedFileAction"
	case EMsg_ClientUCMSetUserPublishedFileActionResponse:
		return "EMsg_ClientUCMSetUserPublishedFileActionResponse"
	case EMsg_ClientUCMEnumeratePublishedFilesByUserAction:
		return "EMsg_ClientUCMEnumeratePublishedFilesByUserAction"
	case EMsg_ClientUCMEnumeratePublishedFilesByUserActionResponse:
		return "EMsg_ClientUCMEnumeratePublishedFilesByUserActionResponse"
	case EMsg_ClientUCMPublishedFileDeleted:
		return "EMsg_ClientUCMPublishedFileDeleted"
	case EMsg_UCMGetUserSubscribedFiles:
		return "EMsg_UCMGetUserSubscribedFiles"
	case EMsg_UCMGetUserSubscribedFilesResponse:
		return "EMsg_UCMGetUserSubscribedFilesResponse"
	case EMsg_UCMFixStatsPublishedFile:
		return "EMsg_UCMFixStatsPublishedFile"
	case EMsg_UCMDeleteOldScreenshot:
		return "EMsg_UCMDeleteOldScreenshot"
	case EMsg_UCMDeleteOldScreenshotResponse:
		return "EMsg_UCMDeleteOldScreenshotResponse"
	case EMsg_UCMDeleteOldVideo:
		return "EMsg_UCMDeleteOldVideo"
	case EMsg_UCMDeleteOldVideoResponse:
		return "EMsg_UCMDeleteOldVideoResponse"
	case EMsg_UCMUpdateOldScreenshotPrivacy:
		return "EMsg_UCMUpdateOldScreenshotPrivacy"
	case EMsg_UCMUpdateOldScreenshotPrivacyResponse:
		return "EMsg_UCMUpdateOldScreenshotPrivacyResponse"
	case EMsg_ClientUCMEnumerateUserSubscribedFilesWithUpdates:
		return "EMsg_ClientUCMEnumerateUserSubscribedFilesWithUpdates"
	case EMsg_ClientUCMEnumerateUserSubscribedFilesWithUpdatesResponse:
		return "EMsg_ClientUCMEnumerateUserSubscribedFilesWithUpdatesResponse"
	case EMsg_UCMPublishedFileContentUpdated:
		return "EMsg_UCMPublishedFileContentUpdated"
	case EMsg_UCMPublishedFileUpdated:
		return "EMsg_UCMPublishedFileUpdated"
	case EMsg_FSBase:
		return "EMsg_FSBase"
	case EMsg_ClientRichPresenceUpload:
		return "EMsg_ClientRichPresenceUpload"
	case EMsg_ClientRichPresenceRequest:
		return "EMsg_ClientRichPresenceRequest"
	case EMsg_ClientRichPresenceInfo:
		return "EMsg_ClientRichPresenceInfo"
	case EMsg_FSRichPresenceRequest:
		return "EMsg_FSRichPresenceRequest"
	case EMsg_FSRichPresenceResponse:
		return "EMsg_FSRichPresenceResponse"
	case EMsg_FSComputeFrenematrix:
		return "EMsg_FSComputeFrenematrix"
	case EMsg_FSComputeFrenematrixResponse:
		return "EMsg_FSComputeFrenematrixResponse"
	case EMsg_FSPlayStatusNotification:
		return "EMsg_FSPlayStatusNotification"
	case EMsg_FSPublishPersonaStatus:
		return "EMsg_FSPublishPersonaStatus"
	case EMsg_FSAddOrRemoveFollower:
		return "EMsg_FSAddOrRemoveFollower"
	case EMsg_FSAddOrRemoveFollowerResponse:
		return "EMsg_FSAddOrRemoveFollowerResponse"
	case EMsg_FSUpdateFollowingList:
		return "EMsg_FSUpdateFollowingList"
	case EMsg_FSCommentNotification:
		return "EMsg_FSCommentNotification"
	case EMsg_FSCommentNotificationViewed:
		return "EMsg_FSCommentNotificationViewed"
	case EMsg_ClientFSGetFollowerCount:
		return "EMsg_ClientFSGetFollowerCount"
	case EMsg_ClientFSGetFollowerCountResponse:
		return "EMsg_ClientFSGetFollowerCountResponse"
	case EMsg_ClientFSGetIsFollowing:
		return "EMsg_ClientFSGetIsFollowing"
	case EMsg_ClientFSGetIsFollowingResponse:
		return "EMsg_ClientFSGetIsFollowingResponse"
	case EMsg_ClientFSEnumerateFollowingList:
		return "EMsg_ClientFSEnumerateFollowingList"
	case EMsg_ClientFSEnumerateFollowingListResponse:
		return "EMsg_ClientFSEnumerateFollowingListResponse"
	case EMsg_FSGetPendingNotificationCount:
		return "EMsg_FSGetPendingNotificationCount"
	case EMsg_FSGetPendingNotificationCountResponse:
		return "EMsg_FSGetPendingNotificationCountResponse"
	case EMsg_ClientFSOfflineMessageNotification:
		return "EMsg_ClientFSOfflineMessageNotification"
	case EMsg_ClientFSRequestOfflineMessageCount:
		return "EMsg_ClientFSRequestOfflineMessageCount"
	case EMsg_ClientFSGetFriendMessageHistory:
		return "EMsg_ClientFSGetFriendMessageHistory"
	case EMsg_ClientFSGetFriendMessageHistoryResponse:
		return "EMsg_ClientFSGetFriendMessageHistoryResponse"
	case EMsg_ClientFSGetFriendMessageHistoryForOfflineMessages:
		return "EMsg_ClientFSGetFriendMessageHistoryForOfflineMessages"
	case EMsg_ClientFSGetFriendsSteamLevels:
		return "EMsg_ClientFSGetFriendsSteamLevels"
	case EMsg_ClientFSGetFriendsSteamLevelsResponse:
		return "EMsg_ClientFSGetFriendsSteamLevelsResponse"
	case EMsg_DRMRange2:
		return "EMsg_DRMRange2"
	case EMsg_CEGVersionSetEnableDisableResponse:
		return "EMsg_CEGVersionSetEnableDisableResponse"
	case EMsg_CEGPropStatusDRMSRequest:
		return "EMsg_CEGPropStatusDRMSRequest"
	case EMsg_CEGPropStatusDRMSResponse:
		return "EMsg_CEGPropStatusDRMSResponse"
	case EMsg_CEGWhackFailureReportRequest:
		return "EMsg_CEGWhackFailureReportRequest"
	case EMsg_CEGWhackFailureReportResponse:
		return "EMsg_CEGWhackFailureReportResponse"
	case EMsg_DRMSFetchVersionSet:
		return "EMsg_DRMSFetchVersionSet"
	case EMsg_DRMSFetchVersionSetResponse:
		return "EMsg_DRMSFetchVersionSetResponse"
	case EMsg_EconBase:
		return "EMsg_EconBase"
	case EMsg_EconTrading_InitiateTradeRequest:
		return "EMsg_EconTrading_InitiateTradeRequest"
	case EMsg_EconTrading_InitiateTradeProposed:
		return "EMsg_EconTrading_InitiateTradeProposed"
	case EMsg_EconTrading_InitiateTradeResponse:
		return "EMsg_EconTrading_InitiateTradeResponse"
	case EMsg_EconTrading_InitiateTradeResult:
		return "EMsg_EconTrading_InitiateTradeResult"
	case EMsg_EconTrading_StartSession:
		return "EMsg_EconTrading_StartSession"
	case EMsg_EconTrading_CancelTradeRequest:
		return "EMsg_EconTrading_CancelTradeRequest"
	case EMsg_EconFlushInventoryCache:
		return "EMsg_EconFlushInventoryCache"
	case EMsg_EconFlushInventoryCacheResponse:
		return "EMsg_EconFlushInventoryCacheResponse"
	case EMsg_EconCDKeyProcessTransaction:
		return "EMsg_EconCDKeyProcessTransaction"
	case EMsg_EconCDKeyProcessTransactionResponse:
		return "EMsg_EconCDKeyProcessTransactionResponse"
	case EMsg_EconGetErrorLogs:
		return "EMsg_EconGetErrorLogs"
	case EMsg_EconGetErrorLogsResponse:
		return "EMsg_EconGetErrorLogsResponse"
	case EMsg_RMRange:
		return "EMsg_RMRange"
	case EMsg_RMTestVerisignOTPResponse:
		return "EMsg_RMTestVerisignOTPResponse"
	case EMsg_RMDeleteMemcachedKeys:
		return "EMsg_RMDeleteMemcachedKeys"
	case EMsg_RMRemoteInvoke:
		return "EMsg_RMRemoteInvoke"
	case EMsg_BadLoginIPList:
		return "EMsg_BadLoginIPList"
	case EMsg_UGSBase:
		return "EMsg_UGSBase"
	case EMsg_ClientUGSGetGlobalStats:
		return "EMsg_ClientUGSGetGlobalStats"
	case EMsg_ClientUGSGetGlobalStatsResponse:
		return "EMsg_ClientUGSGetGlobalStatsResponse"
	case EMsg_StoreBase:
		return "EMsg_StoreBase"
	case EMsg_UMQBase:
		return "EMsg_UMQBase"
	case EMsg_UMQLogonResponse:
		return "EMsg_UMQLogonResponse"
	case EMsg_UMQLogoffRequest:
		return "EMsg_UMQLogoffRequest"
	case EMsg_UMQLogoffResponse:
		return "EMsg_UMQLogoffResponse"
	case EMsg_UMQSendChatMessage:
		return "EMsg_UMQSendChatMessage"
	case EMsg_UMQIncomingChatMessage:
		return "EMsg_UMQIncomingChatMessage"
	case EMsg_UMQPoll:
		return "EMsg_UMQPoll"
	case EMsg_UMQPollResults:
		return "EMsg_UMQPollResults"
	case EMsg_UMQ2AM_ClientMsgBatch:
		return "EMsg_UMQ2AM_ClientMsgBatch"
	case EMsg_UMQEnqueueMobileSalePromotions:
		return "EMsg_UMQEnqueueMobileSalePromotions"
	case EMsg_UMQEnqueueMobileAnnouncements:
		return "EMsg_UMQEnqueueMobileAnnouncements"
	case EMsg_WorkshopBase:
		return "EMsg_WorkshopBase"
	case EMsg_WorkshopAcceptTOSResponse:
		return "EMsg_WorkshopAcceptTOSResponse"
	case EMsg_WebAPIBase:
		return "EMsg_WebAPIBase"
	case EMsg_WebAPIValidateOAuth2TokenResponse:
		return "EMsg_WebAPIValidateOAuth2TokenResponse"
	case EMsg_WebAPIInvalidateTokensForAccount:
		return "EMsg_WebAPIInvalidateTokensForAccount"
	case EMsg_WebAPIRegisterGCInterfaces:
		return "EMsg_WebAPIRegisterGCInterfaces"
	case EMsg_WebAPIInvalidateOAuthClientCache:
		return "EMsg_WebAPIInvalidateOAuthClientCache"
	case EMsg_WebAPIInvalidateOAuthTokenCache:
		return "EMsg_WebAPIInvalidateOAuthTokenCache"
	case EMsg_BackpackBase:
		return "EMsg_BackpackBase"
	case EMsg_BackpackAddToCurrency:
		return "EMsg_BackpackAddToCurrency"
	case EMsg_BackpackAddToCurrencyResponse:
		return "EMsg_BackpackAddToCurrencyResponse"
	case EMsg_CREBase:
		return "EMsg_CREBase"
	case EMsg_CRERankByTrend:
		return "EMsg_CRERankByTrend"
	case EMsg_CRERankByTrendResponse:
		return "EMsg_CRERankByTrendResponse"
	case EMsg_CREItemVoteSummary:
		return "EMsg_CREItemVoteSummary"
	case EMsg_CREItemVoteSummaryResponse:
		return "EMsg_CREItemVoteSummaryResponse"
	case EMsg_CRERankByVote:
		return "EMsg_CRERankByVote"
	case EMsg_CRERankByVoteResponse:
		return "EMsg_CRERankByVoteResponse"
	case EMsg_CREUpdateUserPublishedItemVote:
		return "EMsg_CREUpdateUserPublishedItemVote"
	case EMsg_CREUpdateUserPublishedItemVoteResponse:
		return "EMsg_CREUpdateUserPublishedItemVoteResponse"
	case EMsg_CREGetUserPublishedItemVoteDetails:
		return "EMsg_CREGetUserPublishedItemVoteDetails"
	case EMsg_CREGetUserPublishedItemVoteDetailsResponse:
		return "EMsg_CREGetUserPublishedItemVoteDetailsResponse"
	case EMsg_CREEnumeratePublishedFiles:
		return "EMsg_CREEnumeratePublishedFiles"
	case EMsg_CREEnumeratePublishedFilesResponse:
		return "EMsg_CREEnumeratePublishedFilesResponse"
	case EMsg_CREPublishedFileVoteAdded:
		return "EMsg_CREPublishedFileVoteAdded"
	case EMsg_SecretsBase:
		return "EMsg_SecretsBase"
	case EMsg_SecretsCredentialPairResponse:
		return "EMsg_SecretsCredentialPairResponse"
	case EMsg_SecretsRequestServerIdentity:
		return "EMsg_SecretsRequestServerIdentity"
	case EMsg_SecretsServerIdentityResponse:
		return "EMsg_SecretsServerIdentityResponse"
	case EMsg_SecretsUpdateServerIdentities:
		return "EMsg_SecretsUpdateServerIdentities"
	case EMsg_BoxMonitorBase:
		return "EMsg_BoxMonitorBase"
	case EMsg_BoxMonitorReportResponse:
		return "EMsg_BoxMonitorReportResponse"
	case EMsg_LogsinkBase:
		return "EMsg_LogsinkBase"
	case EMsg_PICSBase:
		return "EMsg_PICSBase"
	case EMsg_ClientPICSChangesSinceRequest:
		return "EMsg_ClientPICSChangesSinceRequest"
	case EMsg_ClientPICSChangesSinceResponse:
		return "EMsg_ClientPICSChangesSinceResponse"
	case EMsg_ClientPICSProductInfoRequest:
		return "EMsg_ClientPICSProductInfoRequest"
	case EMsg_ClientPICSProductInfoResponse:
		return "EMsg_ClientPICSProductInfoResponse"
	case EMsg_ClientPICSAccessTokenRequest:
		return "EMsg_ClientPICSAccessTokenRequest"
	case EMsg_ClientPICSAccessTokenResponse:
		return "EMsg_ClientPICSAccessTokenResponse"
	case EMsg_WorkerProcess:
		return "EMsg_WorkerProcess"
	case EMsg_WorkerProcessPingResponse:
		return "EMsg_WorkerProcessPingResponse"
	case EMsg_WorkerProcessShutdown:
		return "EMsg_WorkerProcessShutdown"
	case EMsg_DRMWorkerProcess:
		return "EMsg_DRMWorkerProcess"
	case EMsg_DRMWorkerProcessDRMAndSignResponse:
		return "EMsg_DRMWorkerProcessDRMAndSignResponse"
	case EMsg_DRMWorkerProcessSteamworksInfoRequest:
		return "EMsg_DRMWorkerProcessSteamworksInfoRequest"
	case EMsg_DRMWorkerProcessSteamworksInfoResponse:
		return "EMsg_DRMWorkerProcessSteamworksInfoResponse"
	case EMsg_DRMWorkerProcessInstallDRMDLLRequest:
		return "EMsg_DRMWorkerProcessInstallDRMDLLRequest"
	case EMsg_DRMWorkerProcessInstallDRMDLLResponse:
		return "EMsg_DRMWorkerProcessInstallDRMDLLResponse"
	case EMsg_DRMWorkerProcessSecretIdStringRequest:
		return "EMsg_DRMWorkerProcessSecretIdStringRequest"
	case EMsg_DRMWorkerProcessSecretIdStringResponse:
		return "EMsg_DRMWorkerProcessSecretIdStringResponse"
	case EMsg_DRMWorkerProcessGetDRMGuidsFromFileRequest:
		return "EMsg_DRMWorkerProcessGetDRMGuidsFromFileRequest"
	case EMsg_DRMWorkerProcessGetDRMGuidsFromFileResponse:
		return "EMsg_DRMWorkerProcessGetDRMGuidsFromFileResponse"
	case EMsg_DRMWorkerProcessInstallProcessedFilesRequest:
		return "EMsg_DRMWorkerProcessInstallProcessedFilesRequest"
	case EMsg_DRMWorkerProcessInstallProcessedFilesResponse:
		return "EMsg_DRMWorkerProcessInstallProcessedFilesResponse"
	case EMsg_DRMWorkerProcessExamineBlobRequest:
		return "EMsg_DRMWorkerProcessExamineBlobRequest"
	case EMsg_DRMWorkerProcessExamineBlobResponse:
		return "EMsg_DRMWorkerProcessExamineBlobResponse"
	case EMsg_DRMWorkerProcessDescribeSecretRequest:
		return "EMsg_DRMWorkerProcessDescribeSecretRequest"
	case EMsg_DRMWorkerProcessDescribeSecretResponse:
		return "EMsg_DRMWorkerProcessDescribeSecretResponse"
	case EMsg_DRMWorkerProcessBackfillOriginalRequest:
		return "EMsg_DRMWorkerProcessBackfillOriginalRequest"
	case EMsg_DRMWorkerProcessBackfillOriginalResponse:
		return "EMsg_DRMWorkerProcessBackfillOriginalResponse"
	case EMsg_DRMWorkerProcessValidateDRMDLLRequest:
		return "EMsg_DRMWorkerProcessValidateDRMDLLRequest"
	case EMsg_DRMWorkerProcessValidateDRMDLLResponse:
		return "EMsg_DRMWorkerProcessValidateDRMDLLResponse"
	case EMsg_DRMWorkerProcessValidateFileRequest:
		return "EMsg_DRMWorkerProcessValidateFileRequest"
	case EMsg_DRMWorkerProcessValidateFileResponse:
		return "EMsg_DRMWorkerProcessValidateFileResponse"
	case EMsg_DRMWorkerProcessSplitAndInstallRequest:
		return "EMsg_DRMWorkerProcessSplitAndInstallRequest"
	case EMsg_DRMWorkerProcessSplitAndInstallResponse:
		return "EMsg_DRMWorkerProcessSplitAndInstallResponse"
	case EMsg_DRMWorkerProcessGetBlobRequest:
		return "EMsg_DRMWorkerProcessGetBlobRequest"
	case EMsg_DRMWorkerProcessGetBlobResponse:
		return "EMsg_DRMWorkerProcessGetBlobResponse"
	case EMsg_DRMWorkerProcessEvaluateCrashRequest:
		return "EMsg_DRMWorkerProcessEvaluateCrashRequest"
	case EMsg_DRMWorkerProcessEvaluateCrashResponse:
		return "EMsg_DRMWorkerProcessEvaluateCrashResponse"
	case EMsg_DRMWorkerProcessAnalyzeFileRequest:
		return "EMsg_DRMWorkerProcessAnalyzeFileRequest"
	case EMsg_DRMWorkerProcessAnalyzeFileResponse:
		return "EMsg_DRMWorkerProcessAnalyzeFileResponse"
	case EMsg_DRMWorkerProcessUnpackBlobRequest:
		return "EMsg_DRMWorkerProcessUnpackBlobRequest"
	case EMsg_DRMWorkerProcessUnpackBlobResponse:
		return "EMsg_DRMWorkerProcessUnpackBlobResponse"
	case EMsg_DRMWorkerProcessInstallAllRequest:
		return "EMsg_DRMWorkerProcessInstallAllRequest"
	case EMsg_DRMWorkerProcessInstallAllResponse:
		return "EMsg_DRMWorkerProcessInstallAllResponse"
	case EMsg_TestWorkerProcess:
		return "EMsg_TestWorkerProcess"
	case EMsg_TestWorkerProcessLoadUnloadModuleResponse:
		return "EMsg_TestWorkerProcessLoadUnloadModuleResponse"
	case EMsg_TestWorkerProcessServiceModuleCallRequest:
		return "EMsg_TestWorkerProcessServiceModuleCallRequest"
	case EMsg_TestWorkerProcessServiceModuleCallResponse:
		return "EMsg_TestWorkerProcessServiceModuleCallResponse"
	case EMsg_ClientGetEmoticonList:
		return "EMsg_ClientGetEmoticonList"
	case EMsg_ClientEmoticonList:
		return "EMsg_ClientEmoticonList"
	case EMsg_ClientSharedLibraryBase:
		return "EMsg_ClientSharedLibraryBase"
	case EMsg_ClientSharedLicensesLockStatus:
		return "EMsg_ClientSharedLicensesLockStatus"
	case EMsg_ClientSharedLicensesStopPlaying:
		return "EMsg_ClientSharedLicensesStopPlaying"
	case EMsg_ClientSharedLibraryLockStatus:
		return "EMsg_ClientSharedLibraryLockStatus"
	case EMsg_ClientSharedLibraryStopPlaying:
		return "EMsg_ClientSharedLibraryStopPlaying"
	case EMsg_ClientUnlockStreaming:
		return "EMsg_ClientUnlockStreaming"
	case EMsg_ClientUnlockStreamingResponse:
		return "EMsg_ClientUnlockStreamingResponse"
	case EMsg_ClientPlayingSessionState:
		return "EMsg_ClientPlayingSessionState"
	case EMsg_ClientKickPlayingSession:
		return "EMsg_ClientKickPlayingSession"
	default:
		return "INVALID"
	}
}

type EResult int32

const (
	EResult_Invalid                                 EResult = 0
	EResult_OK                                              = 1
	EResult_Fail                                            = 2
	EResult_NoConnection                                    = 3
	EResult_InvalidPassword                                 = 5
	EResult_LoggedInElsewhere                               = 6
	EResult_InvalidProtocolVer                              = 7
	EResult_InvalidParam                                    = 8
	EResult_FileNotFound                                    = 9
	EResult_Busy                                            = 10
	EResult_InvalidState                                    = 11
	EResult_InvalidName                                     = 12
	EResult_InvalidEmail                                    = 13
	EResult_DuplicateName                                   = 14
	EResult_AccessDenied                                    = 15
	EResult_Timeout                                         = 16
	EResult_Banned                                          = 17
	EResult_AccountNotFound                                 = 18
	EResult_InvalidSteamID                                  = 19
	EResult_ServiceUnavailable                              = 20
	EResult_NotLoggedOn                                     = 21
	EResult_Pending                                         = 22
	EResult_EncryptionFailure                               = 23
	EResult_InsufficientPrivilege                           = 24
	EResult_LimitExceeded                                   = 25
	EResult_Revoked                                         = 26
	EResult_Expired                                         = 27
	EResult_AlreadyRedeemed                                 = 28
	EResult_DuplicateRequest                                = 29
	EResult_AlreadyOwned                                    = 30
	EResult_IPNotFound                                      = 31
	EResult_PersistFailed                                   = 32
	EResult_LockingFailed                                   = 33
	EResult_LogonSessionReplaced                            = 34
	EResult_ConnectFailed                                   = 35
	EResult_HandshakeFailed                                 = 36
	EResult_IOFailure                                       = 37
	EResult_RemoteDisconnect                                = 38
	EResult_ShoppingCartNotFound                            = 39
	EResult_Blocked                                         = 40
	EResult_Ignored                                         = 41
	EResult_NoMatch                                         = 42
	EResult_AccountDisabled                                 = 43
	EResult_ServiceReadOnly                                 = 44
	EResult_AccountNotFeatured                              = 45
	EResult_AdministratorOK                                 = 46
	EResult_ContentVersion                                  = 47
	EResult_TryAnotherCM                                    = 48
	EResult_PasswordRequiredToKickSession                   = 49
	EResult_AlreadyLoggedInElsewhere                        = 50
	EResult_Suspended                                       = 51
	EResult_Cancelled                                       = 52
	EResult_DataCorruption                                  = 53
	EResult_DiskFull                                        = 54
	EResult_RemoteCallFailed                                = 55
	EResult_PasswordNotSet                                  = 56
	EResult_ExternalAccountUnlinked                         = 57
	EResult_PSNTicketInvalid                                = 58
	EResult_ExternalAccountAlreadyLinked                    = 59
	EResult_RemoteFileConflict                              = 60
	EResult_IllegalPassword                                 = 61
	EResult_SameAsPreviousValue                             = 62
	EResult_AccountLogonDenied                              = 63
	EResult_CannotUseOldPassword                            = 64
	EResult_InvalidLoginAuthCode                            = 65
	EResult_AccountLogonDeniedNoMailSent                    = 66
	EResult_HardwareNotCapableOfIPT                         = 67
	EResult_IPTInitError                                    = 68
	EResult_ParentalControlRestricted                       = 69
	EResult_FacebookQueryError                              = 70
	EResult_ExpiredLoginAuthCode                            = 71
	EResult_IPLoginRestrictionFailed                        = 72
	EResult_AccountLocked                                   = 73
	EResult_AccountLogonDeniedVerifiedEmailRequired         = 74
	EResult_NoMatchingURL                                   = 75
	EResult_BadResponse                                     = 76
	EResult_RequirePasswordReEntry                          = 77
	EResult_ValueOutOfRange                                 = 78
	EResult_UnexpectedError                                 = 79
	EResult_Disabled                                        = 80
	EResult_InvalidCEGSubmission                            = 81
	EResult_RestrictedDevice                                = 82
	EResult_RegionLocked                                    = 83
	EResult_RateLimitExceeded                               = 84
	EResult_AccountLogonDeniedNeedTwoFactorCode             = 85
	EResult_ItemOrEntryHasBeenDeleted                       = 86
)

func (e EResult) String() string {
	switch e {
	case EResult_Invalid:
		return "EResult_Invalid"
	case EResult_OK:
		return "EResult_OK"
	case EResult_Fail:
		return "EResult_Fail"
	case EResult_NoConnection:
		return "EResult_NoConnection"
	case EResult_InvalidPassword:
		return "EResult_InvalidPassword"
	case EResult_LoggedInElsewhere:
		return "EResult_LoggedInElsewhere"
	case EResult_InvalidProtocolVer:
		return "EResult_InvalidProtocolVer"
	case EResult_InvalidParam:
		return "EResult_InvalidParam"
	case EResult_FileNotFound:
		return "EResult_FileNotFound"
	case EResult_Busy:
		return "EResult_Busy"
	case EResult_InvalidState:
		return "EResult_InvalidState"
	case EResult_InvalidName:
		return "EResult_InvalidName"
	case EResult_InvalidEmail:
		return "EResult_InvalidEmail"
	case EResult_DuplicateName:
		return "EResult_DuplicateName"
	case EResult_AccessDenied:
		return "EResult_AccessDenied"
	case EResult_Timeout:
		return "EResult_Timeout"
	case EResult_Banned:
		return "EResult_Banned"
	case EResult_AccountNotFound:
		return "EResult_AccountNotFound"
	case EResult_InvalidSteamID:
		return "EResult_InvalidSteamID"
	case EResult_ServiceUnavailable:
		return "EResult_ServiceUnavailable"
	case EResult_NotLoggedOn:
		return "EResult_NotLoggedOn"
	case EResult_Pending:
		return "EResult_Pending"
	case EResult_EncryptionFailure:
		return "EResult_EncryptionFailure"
	case EResult_InsufficientPrivilege:
		return "EResult_InsufficientPrivilege"
	case EResult_LimitExceeded:
		return "EResult_LimitExceeded"
	case EResult_Revoked:
		return "EResult_Revoked"
	case EResult_Expired:
		return "EResult_Expired"
	case EResult_AlreadyRedeemed:
		return "EResult_AlreadyRedeemed"
	case EResult_DuplicateRequest:
		return "EResult_DuplicateRequest"
	case EResult_AlreadyOwned:
		return "EResult_AlreadyOwned"
	case EResult_IPNotFound:
		return "EResult_IPNotFound"
	case EResult_PersistFailed:
		return "EResult_PersistFailed"
	case EResult_LockingFailed:
		return "EResult_LockingFailed"
	case EResult_LogonSessionReplaced:
		return "EResult_LogonSessionReplaced"
	case EResult_ConnectFailed:
		return "EResult_ConnectFailed"
	case EResult_HandshakeFailed:
		return "EResult_HandshakeFailed"
	case EResult_IOFailure:
		return "EResult_IOFailure"
	case EResult_RemoteDisconnect:
		return "EResult_RemoteDisconnect"
	case EResult_ShoppingCartNotFound:
		return "EResult_ShoppingCartNotFound"
	case EResult_Blocked:
		return "EResult_Blocked"
	case EResult_Ignored:
		return "EResult_Ignored"
	case EResult_NoMatch:
		return "EResult_NoMatch"
	case EResult_AccountDisabled:
		return "EResult_AccountDisabled"
	case EResult_ServiceReadOnly:
		return "EResult_ServiceReadOnly"
	case EResult_AccountNotFeatured:
		return "EResult_AccountNotFeatured"
	case EResult_AdministratorOK:
		return "EResult_AdministratorOK"
	case EResult_ContentVersion:
		return "EResult_ContentVersion"
	case EResult_TryAnotherCM:
		return "EResult_TryAnotherCM"
	case EResult_PasswordRequiredToKickSession:
		return "EResult_PasswordRequiredToKickSession"
	case EResult_AlreadyLoggedInElsewhere:
		return "EResult_AlreadyLoggedInElsewhere"
	case EResult_Suspended:
		return "EResult_Suspended"
	case EResult_Cancelled:
		return "EResult_Cancelled"
	case EResult_DataCorruption:
		return "EResult_DataCorruption"
	case EResult_DiskFull:
		return "EResult_DiskFull"
	case EResult_RemoteCallFailed:
		return "EResult_RemoteCallFailed"
	case EResult_PasswordNotSet:
		return "EResult_PasswordNotSet"
	case EResult_ExternalAccountUnlinked:
		return "EResult_ExternalAccountUnlinked"
	case EResult_PSNTicketInvalid:
		return "EResult_PSNTicketInvalid"
	case EResult_ExternalAccountAlreadyLinked:
		return "EResult_ExternalAccountAlreadyLinked"
	case EResult_RemoteFileConflict:
		return "EResult_RemoteFileConflict"
	case EResult_IllegalPassword:
		return "EResult_IllegalPassword"
	case EResult_SameAsPreviousValue:
		return "EResult_SameAsPreviousValue"
	case EResult_AccountLogonDenied:
		return "EResult_AccountLogonDenied"
	case EResult_CannotUseOldPassword:
		return "EResult_CannotUseOldPassword"
	case EResult_InvalidLoginAuthCode:
		return "EResult_InvalidLoginAuthCode"
	case EResult_AccountLogonDeniedNoMailSent:
		return "EResult_AccountLogonDeniedNoMailSent"
	case EResult_HardwareNotCapableOfIPT:
		return "EResult_HardwareNotCapableOfIPT"
	case EResult_IPTInitError:
		return "EResult_IPTInitError"
	case EResult_ParentalControlRestricted:
		return "EResult_ParentalControlRestricted"
	case EResult_FacebookQueryError:
		return "EResult_FacebookQueryError"
	case EResult_ExpiredLoginAuthCode:
		return "EResult_ExpiredLoginAuthCode"
	case EResult_IPLoginRestrictionFailed:
		return "EResult_IPLoginRestrictionFailed"
	case EResult_AccountLocked:
		return "EResult_AccountLocked"
	case EResult_AccountLogonDeniedVerifiedEmailRequired:
		return "EResult_AccountLogonDeniedVerifiedEmailRequired"
	case EResult_NoMatchingURL:
		return "EResult_NoMatchingURL"
	case EResult_BadResponse:
		return "EResult_BadResponse"
	case EResult_RequirePasswordReEntry:
		return "EResult_RequirePasswordReEntry"
	case EResult_ValueOutOfRange:
		return "EResult_ValueOutOfRange"
	case EResult_UnexpectedError:
		return "EResult_UnexpectedError"
	case EResult_Disabled:
		return "EResult_Disabled"
	case EResult_InvalidCEGSubmission:
		return "EResult_InvalidCEGSubmission"
	case EResult_RestrictedDevice:
		return "EResult_RestrictedDevice"
	case EResult_RegionLocked:
		return "EResult_RegionLocked"
	case EResult_RateLimitExceeded:
		return "EResult_RateLimitExceeded"
	case EResult_AccountLogonDeniedNeedTwoFactorCode:
		return "EResult_AccountLogonDeniedNeedTwoFactorCode"
	case EResult_ItemOrEntryHasBeenDeleted:
		return "EResult_ItemOrEntryHasBeenDeleted"
	default:
		return "INVALID"
	}
}

type EUniverse int32

const (
	EUniverse_Invalid  EUniverse = 0
	EUniverse_Public             = 1
	EUniverse_Beta               = 2
	EUniverse_Internal           = 3
	EUniverse_Dev                = 4
	EUniverse_RC                 = 5 // Deprecated: Universe no longer exists
	EUniverse_Max                = 5
)

func (e EUniverse) String() string {
	switch e {
	case EUniverse_Invalid:
		return "EUniverse_Invalid"
	case EUniverse_Public:
		return "EUniverse_Public"
	case EUniverse_Beta:
		return "EUniverse_Beta"
	case EUniverse_Internal:
		return "EUniverse_Internal"
	case EUniverse_Dev:
		return "EUniverse_Dev"
	case EUniverse_RC:
		return "EUniverse_RC"
	default:
		return "INVALID"
	}
}

type EChatEntryType int32

const (
	EChatEntryType_Invalid          EChatEntryType = 0
	EChatEntryType_ChatMsg                         = 1
	EChatEntryType_Typing                          = 2
	EChatEntryType_InviteGame                      = 3
	EChatEntryType_Emote                           = 4 // Deprecated: No longer supported by clients
	EChatEntryType_LobbyGameStart                  = 5
	EChatEntryType_LeftConversation                = 6
	EChatEntryType_Entered                         = 7
	EChatEntryType_WasKicked                       = 8
	EChatEntryType_WasBanned                       = 9
	EChatEntryType_Disconnected                    = 10
	EChatEntryType_HistoricalChat                  = 11
)

func (e EChatEntryType) String() string {
	switch e {
	case EChatEntryType_Invalid:
		return "EChatEntryType_Invalid"
	case EChatEntryType_ChatMsg:
		return "EChatEntryType_ChatMsg"
	case EChatEntryType_Typing:
		return "EChatEntryType_Typing"
	case EChatEntryType_InviteGame:
		return "EChatEntryType_InviteGame"
	case EChatEntryType_Emote:
		return "EChatEntryType_Emote"
	case EChatEntryType_LobbyGameStart:
		return "EChatEntryType_LobbyGameStart"
	case EChatEntryType_LeftConversation:
		return "EChatEntryType_LeftConversation"
	case EChatEntryType_Entered:
		return "EChatEntryType_Entered"
	case EChatEntryType_WasKicked:
		return "EChatEntryType_WasKicked"
	case EChatEntryType_WasBanned:
		return "EChatEntryType_WasBanned"
	case EChatEntryType_Disconnected:
		return "EChatEntryType_Disconnected"
	case EChatEntryType_HistoricalChat:
		return "EChatEntryType_HistoricalChat"
	default:
		return "INVALID"
	}
}

type EPersonaState int32

const (
	EPersonaState_Offline        EPersonaState = 0
	EPersonaState_Online                       = 1
	EPersonaState_Busy                         = 2
	EPersonaState_Away                         = 3
	EPersonaState_Snooze                       = 4
	EPersonaState_LookingToTrade               = 5
	EPersonaState_LookingToPlay                = 6
	EPersonaState_Max                          = 7
)

func (e EPersonaState) String() string {
	switch e {
	case EPersonaState_Offline:
		return "EPersonaState_Offline"
	case EPersonaState_Online:
		return "EPersonaState_Online"
	case EPersonaState_Busy:
		return "EPersonaState_Busy"
	case EPersonaState_Away:
		return "EPersonaState_Away"
	case EPersonaState_Snooze:
		return "EPersonaState_Snooze"
	case EPersonaState_LookingToTrade:
		return "EPersonaState_LookingToTrade"
	case EPersonaState_LookingToPlay:
		return "EPersonaState_LookingToPlay"
	case EPersonaState_Max:
		return "EPersonaState_Max"
	default:
		return "INVALID"
	}
}

type EAccountType int32

const (
	EAccountType_Invalid        EAccountType = 0
	EAccountType_Individual                  = 1
	EAccountType_Multiseat                   = 2
	EAccountType_GameServer                  = 3
	EAccountType_AnonGameServer              = 4
	EAccountType_Pending                     = 5
	EAccountType_ContentServer               = 6
	EAccountType_Clan                        = 7
	EAccountType_Chat                        = 8
	EAccountType_ConsoleUser                 = 9
	EAccountType_AnonUser                    = 10
	EAccountType_Max                         = 11
)

func (e EAccountType) String() string {
	switch e {
	case EAccountType_Invalid:
		return "EAccountType_Invalid"
	case EAccountType_Individual:
		return "EAccountType_Individual"
	case EAccountType_Multiseat:
		return "EAccountType_Multiseat"
	case EAccountType_GameServer:
		return "EAccountType_GameServer"
	case EAccountType_AnonGameServer:
		return "EAccountType_AnonGameServer"
	case EAccountType_Pending:
		return "EAccountType_Pending"
	case EAccountType_ContentServer:
		return "EAccountType_ContentServer"
	case EAccountType_Clan:
		return "EAccountType_Clan"
	case EAccountType_Chat:
		return "EAccountType_Chat"
	case EAccountType_ConsoleUser:
		return "EAccountType_ConsoleUser"
	case EAccountType_AnonUser:
		return "EAccountType_AnonUser"
	case EAccountType_Max:
		return "EAccountType_Max"
	default:
		return "INVALID"
	}
}

type EFriendRelationship int32

const (
	EFriendRelationship_None             EFriendRelationship = 0
	EFriendRelationship_Blocked                              = 1
	EFriendRelationship_PendingInvitee                       = 2 // Deprecated: renamed to RequestRecipient
	EFriendRelationship_RequestRecipient                     = 2
	EFriendRelationship_Friend                               = 3
	EFriendRelationship_RequestInitiator                     = 4
	EFriendRelationship_PendingInviter                       = 4 // Deprecated: renamed to RequestInitiator
	EFriendRelationship_Ignored                              = 5
	EFriendRelationship_IgnoredFriend                        = 6
	EFriendRelationship_SuggestedFriend                      = 7
	EFriendRelationship_Max                                  = 8
)

func (e EFriendRelationship) String() string {
	switch e {
	case EFriendRelationship_None:
		return "EFriendRelationship_None"
	case EFriendRelationship_Blocked:
		return "EFriendRelationship_Blocked"
	case EFriendRelationship_PendingInvitee:
		return "EFriendRelationship_PendingInvitee"
	case EFriendRelationship_Friend:
		return "EFriendRelationship_Friend"
	case EFriendRelationship_RequestInitiator:
		return "EFriendRelationship_RequestInitiator"
	case EFriendRelationship_Ignored:
		return "EFriendRelationship_Ignored"
	case EFriendRelationship_IgnoredFriend:
		return "EFriendRelationship_IgnoredFriend"
	case EFriendRelationship_SuggestedFriend:
		return "EFriendRelationship_SuggestedFriend"
	case EFriendRelationship_Max:
		return "EFriendRelationship_Max"
	default:
		return "INVALID"
	}
}

type EAccountFlags int32

const (
	EAccountFlags_NormalUser                 EAccountFlags = 0
	EAccountFlags_PersonaNameSet                           = 1
	EAccountFlags_Unbannable                               = 2
	EAccountFlags_PasswordSet                              = 4
	EAccountFlags_Support                                  = 8
	EAccountFlags_Admin                                    = 16
	EAccountFlags_Supervisor                               = 32
	EAccountFlags_AppEditor                                = 64
	EAccountFlags_HWIDSet                                  = 128
	EAccountFlags_PersonalQASet                            = 256
	EAccountFlags_VacBeta                                  = 512
	EAccountFlags_Debug                                    = 1024
	EAccountFlags_Disabled                                 = 2048
	EAccountFlags_LimitedUser                              = 4096
	EAccountFlags_LimitedUserForce                         = 8192
	EAccountFlags_EmailValidated                           = 16384
	EAccountFlags_MarketingTreatment                       = 32768
	EAccountFlags_OGGInviteOptOut                          = 65536
	EAccountFlags_ForcePasswordChange                      = 131072
	EAccountFlags_ForceEmailVerification                   = 262144
	EAccountFlags_LogonExtraSecurity                       = 524288
	EAccountFlags_LogonExtraSecurityDisabled               = 1048576
	EAccountFlags_Steam2MigrationComplete                  = 2097152
	EAccountFlags_NeedLogs                                 = 4194304
	EAccountFlags_Lockdown                                 = 8388608
	EAccountFlags_MasterAppEditor                          = 16777216
	EAccountFlags_BannedFromWebAPI                         = 33554432
	EAccountFlags_ClansOnlyFromFriends                     = 67108864
	EAccountFlags_GlobalModerator                          = 134217728
)

func (e EAccountFlags) String() string {
	switch e {
	case EAccountFlags_NormalUser:
		return "EAccountFlags_NormalUser"
	case EAccountFlags_PersonaNameSet:
		return "EAccountFlags_PersonaNameSet"
	case EAccountFlags_Unbannable:
		return "EAccountFlags_Unbannable"
	case EAccountFlags_PasswordSet:
		return "EAccountFlags_PasswordSet"
	case EAccountFlags_Support:
		return "EAccountFlags_Support"
	case EAccountFlags_Admin:
		return "EAccountFlags_Admin"
	case EAccountFlags_Supervisor:
		return "EAccountFlags_Supervisor"
	case EAccountFlags_AppEditor:
		return "EAccountFlags_AppEditor"
	case EAccountFlags_HWIDSet:
		return "EAccountFlags_HWIDSet"
	case EAccountFlags_PersonalQASet:
		return "EAccountFlags_PersonalQASet"
	case EAccountFlags_VacBeta:
		return "EAccountFlags_VacBeta"
	case EAccountFlags_Debug:
		return "EAccountFlags_Debug"
	case EAccountFlags_Disabled:
		return "EAccountFlags_Disabled"
	case EAccountFlags_LimitedUser:
		return "EAccountFlags_LimitedUser"
	case EAccountFlags_LimitedUserForce:
		return "EAccountFlags_LimitedUserForce"
	case EAccountFlags_EmailValidated:
		return "EAccountFlags_EmailValidated"
	case EAccountFlags_MarketingTreatment:
		return "EAccountFlags_MarketingTreatment"
	case EAccountFlags_OGGInviteOptOut:
		return "EAccountFlags_OGGInviteOptOut"
	case EAccountFlags_ForcePasswordChange:
		return "EAccountFlags_ForcePasswordChange"
	case EAccountFlags_ForceEmailVerification:
		return "EAccountFlags_ForceEmailVerification"
	case EAccountFlags_LogonExtraSecurity:
		return "EAccountFlags_LogonExtraSecurity"
	case EAccountFlags_LogonExtraSecurityDisabled:
		return "EAccountFlags_LogonExtraSecurityDisabled"
	case EAccountFlags_Steam2MigrationComplete:
		return "EAccountFlags_Steam2MigrationComplete"
	case EAccountFlags_NeedLogs:
		return "EAccountFlags_NeedLogs"
	case EAccountFlags_Lockdown:
		return "EAccountFlags_Lockdown"
	case EAccountFlags_MasterAppEditor:
		return "EAccountFlags_MasterAppEditor"
	case EAccountFlags_BannedFromWebAPI:
		return "EAccountFlags_BannedFromWebAPI"
	case EAccountFlags_ClansOnlyFromFriends:
		return "EAccountFlags_ClansOnlyFromFriends"
	case EAccountFlags_GlobalModerator:
		return "EAccountFlags_GlobalModerator"
	default:
		return "INVALID"
	}
}

type EClanPermission int32

const (
	EClanPermission_Nobody                EClanPermission = 0
	EClanPermission_Owner                                 = 1
	EClanPermission_Officer                               = 2
	EClanPermission_OwnerAndOfficer                       = 3
	EClanPermission_Member                                = 4
	EClanPermission_Moderator                             = 8
	EClanPermission_OwnerOfficerModerator                 = EClanPermission_Owner | EClanPermission_Officer | EClanPermission_Moderator
	EClanPermission_AllMembers                            = EClanPermission_Owner | EClanPermission_Officer | EClanPermission_Moderator | EClanPermission_Member
	EClanPermission_OGGGameOwner                          = 16
	EClanPermission_NonMember                             = 128
	EClanPermission_MemberAllowed                         = EClanPermission_NonMember | EClanPermission_Member
	EClanPermission_ModeratorAllowed                      = EClanPermission_NonMember | EClanPermission_Member | EClanPermission_Moderator
	EClanPermission_OfficerAllowed                        = EClanPermission_NonMember | EClanPermission_Member | EClanPermission_Moderator | EClanPermission_Officer
	EClanPermission_OwnerAllowed                          = EClanPermission_NonMember | EClanPermission_Member | EClanPermission_Moderator | EClanPermission_Officer | EClanPermission_Owner
	EClanPermission_Anybody                               = EClanPermission_NonMember | EClanPermission_Member | EClanPermission_Moderator | EClanPermission_Officer | EClanPermission_Owner
)

func (e EClanPermission) String() string {
	switch e {
	case EClanPermission_Nobody:
		return "EClanPermission_Nobody"
	case EClanPermission_Owner:
		return "EClanPermission_Owner"
	case EClanPermission_Officer:
		return "EClanPermission_Officer"
	case EClanPermission_OwnerAndOfficer:
		return "EClanPermission_OwnerAndOfficer"
	case EClanPermission_Member:
		return "EClanPermission_Member"
	case EClanPermission_Moderator:
		return "EClanPermission_Moderator"
	case EClanPermission_OwnerOfficerModerator:
		return "EClanPermission_OwnerOfficerModerator"
	case EClanPermission_AllMembers:
		return "EClanPermission_AllMembers"
	case EClanPermission_OGGGameOwner:
		return "EClanPermission_OGGGameOwner"
	case EClanPermission_NonMember:
		return "EClanPermission_NonMember"
	case EClanPermission_MemberAllowed:
		return "EClanPermission_MemberAllowed"
	case EClanPermission_ModeratorAllowed:
		return "EClanPermission_ModeratorAllowed"
	case EClanPermission_OfficerAllowed:
		return "EClanPermission_OfficerAllowed"
	case EClanPermission_OwnerAllowed:
		return "EClanPermission_OwnerAllowed"
	default:
		return "INVALID"
	}
}

type EChatPermission int32

const (
	EChatPermission_Close                    EChatPermission = 1
	EChatPermission_Invite                                   = 2
	EChatPermission_Talk                                     = 8
	EChatPermission_Kick                                     = 16
	EChatPermission_Mute                                     = 32
	EChatPermission_SetMetadata                              = 64
	EChatPermission_ChangePermissions                        = 128
	EChatPermission_Ban                                      = 256
	EChatPermission_ChangeAccess                             = 512
	EChatPermission_EveryoneNotInClanDefault                 = EChatPermission_Talk
	EChatPermission_EveryoneDefault                          = EChatPermission_Talk | EChatPermission_Invite
	EChatPermission_MemberDefault                            = EChatPermission_Ban | EChatPermission_Kick | EChatPermission_Talk | EChatPermission_Invite
	EChatPermission_OfficerDefault                           = EChatPermission_Ban | EChatPermission_Kick | EChatPermission_Talk | EChatPermission_Invite
	EChatPermission_OwnerDefault                             = EChatPermission_ChangeAccess | EChatPermission_Ban | EChatPermission_SetMetadata | EChatPermission_Mute | EChatPermission_Kick | EChatPermission_Talk | EChatPermission_Invite | EChatPermission_Close
	EChatPermission_Mask                                     = 1019
)

func (e EChatPermission) String() string {
	switch e {
	case EChatPermission_Close:
		return "EChatPermission_Close"
	case EChatPermission_Invite:
		return "EChatPermission_Invite"
	case EChatPermission_Talk:
		return "EChatPermission_Talk"
	case EChatPermission_Kick:
		return "EChatPermission_Kick"
	case EChatPermission_Mute:
		return "EChatPermission_Mute"
	case EChatPermission_SetMetadata:
		return "EChatPermission_SetMetadata"
	case EChatPermission_ChangePermissions:
		return "EChatPermission_ChangePermissions"
	case EChatPermission_Ban:
		return "EChatPermission_Ban"
	case EChatPermission_ChangeAccess:
		return "EChatPermission_ChangeAccess"
	case EChatPermission_EveryoneDefault:
		return "EChatPermission_EveryoneDefault"
	case EChatPermission_MemberDefault:
		return "EChatPermission_MemberDefault"
	case EChatPermission_OwnerDefault:
		return "EChatPermission_OwnerDefault"
	case EChatPermission_Mask:
		return "EChatPermission_Mask"
	default:
		return "INVALID"
	}
}

type EFriendFlags int32

const (
	EFriendFlags_None                 EFriendFlags = 0
	EFriendFlags_Blocked                           = 1
	EFriendFlags_FriendshipRequested               = 2
	EFriendFlags_Immediate                         = 4
	EFriendFlags_ClanMember                        = 8
	EFriendFlags_GameServer                        = 16 // Deprecated: renamed to OnGameServer
	EFriendFlags_OnGameServer                      = 16
	EFriendFlags_RequestingFriendship              = 128
	EFriendFlags_RequestingInfo                    = 256
	EFriendFlags_Ignored                           = 512
	EFriendFlags_IgnoredFriend                     = 1024
	EFriendFlags_Suggested                         = 2048
	EFriendFlags_FlagAll                           = 65535
)

func (e EFriendFlags) String() string {
	switch e {
	case EFriendFlags_None:
		return "EFriendFlags_None"
	case EFriendFlags_Blocked:
		return "EFriendFlags_Blocked"
	case EFriendFlags_FriendshipRequested:
		return "EFriendFlags_FriendshipRequested"
	case EFriendFlags_Immediate:
		return "EFriendFlags_Immediate"
	case EFriendFlags_ClanMember:
		return "EFriendFlags_ClanMember"
	case EFriendFlags_GameServer:
		return "EFriendFlags_GameServer"
	case EFriendFlags_RequestingFriendship:
		return "EFriendFlags_RequestingFriendship"
	case EFriendFlags_RequestingInfo:
		return "EFriendFlags_RequestingInfo"
	case EFriendFlags_Ignored:
		return "EFriendFlags_Ignored"
	case EFriendFlags_IgnoredFriend:
		return "EFriendFlags_IgnoredFriend"
	case EFriendFlags_Suggested:
		return "EFriendFlags_Suggested"
	case EFriendFlags_FlagAll:
		return "EFriendFlags_FlagAll"
	default:
		return "INVALID"
	}
}

type EPersonaStateFlag int32

const (
	EPersonaStateFlag_HasRichPresence       EPersonaStateFlag = 1
	EPersonaStateFlag_InJoinableGame                          = 2
	EPersonaStateFlag_OnlineUsingWeb                          = 256
	EPersonaStateFlag_OnlineUsingMobile                       = 512
	EPersonaStateFlag_OnlineUsingBigPicture                   = 1024
)

func (e EPersonaStateFlag) String() string {
	switch e {
	case EPersonaStateFlag_HasRichPresence:
		return "EPersonaStateFlag_HasRichPresence"
	case EPersonaStateFlag_InJoinableGame:
		return "EPersonaStateFlag_InJoinableGame"
	case EPersonaStateFlag_OnlineUsingWeb:
		return "EPersonaStateFlag_OnlineUsingWeb"
	case EPersonaStateFlag_OnlineUsingMobile:
		return "EPersonaStateFlag_OnlineUsingMobile"
	case EPersonaStateFlag_OnlineUsingBigPicture:
		return "EPersonaStateFlag_OnlineUsingBigPicture"
	default:
		return "INVALID"
	}
}

type EClientPersonaStateFlag int32

const (
	EClientPersonaStateFlag_Status        EClientPersonaStateFlag = 1
	EClientPersonaStateFlag_PlayerName                            = 2
	EClientPersonaStateFlag_QueryPort                             = 4
	EClientPersonaStateFlag_SourceID                              = 8
	EClientPersonaStateFlag_Presence                              = 16
	EClientPersonaStateFlag_Metadata                              = 32
	EClientPersonaStateFlag_LastSeen                              = 64
	EClientPersonaStateFlag_ClanInfo                              = 128
	EClientPersonaStateFlag_GameExtraInfo                         = 256
	EClientPersonaStateFlag_GameDataBlob                          = 512
	EClientPersonaStateFlag_ClanTag                               = 1024
	EClientPersonaStateFlag_Facebook                              = 2048
)

func (e EClientPersonaStateFlag) String() string {
	switch e {
	case EClientPersonaStateFlag_Status:
		return "EClientPersonaStateFlag_Status"
	case EClientPersonaStateFlag_PlayerName:
		return "EClientPersonaStateFlag_PlayerName"
	case EClientPersonaStateFlag_QueryPort:
		return "EClientPersonaStateFlag_QueryPort"
	case EClientPersonaStateFlag_SourceID:
		return "EClientPersonaStateFlag_SourceID"
	case EClientPersonaStateFlag_Presence:
		return "EClientPersonaStateFlag_Presence"
	case EClientPersonaStateFlag_Metadata:
		return "EClientPersonaStateFlag_Metadata"
	case EClientPersonaStateFlag_LastSeen:
		return "EClientPersonaStateFlag_LastSeen"
	case EClientPersonaStateFlag_ClanInfo:
		return "EClientPersonaStateFlag_ClanInfo"
	case EClientPersonaStateFlag_GameExtraInfo:
		return "EClientPersonaStateFlag_GameExtraInfo"
	case EClientPersonaStateFlag_GameDataBlob:
		return "EClientPersonaStateFlag_GameDataBlob"
	case EClientPersonaStateFlag_ClanTag:
		return "EClientPersonaStateFlag_ClanTag"
	case EClientPersonaStateFlag_Facebook:
		return "EClientPersonaStateFlag_Facebook"
	default:
		return "INVALID"
	}
}

type EAppUsageEvent int32

const (
	EAppUsageEvent_GameLaunch            EAppUsageEvent = 1
	EAppUsageEvent_GameLaunchTrial                      = 2
	EAppUsageEvent_Media                                = 3
	EAppUsageEvent_PreloadStart                         = 4
	EAppUsageEvent_PreloadFinish                        = 5
	EAppUsageEvent_MarketingMessageView                 = 6
	EAppUsageEvent_InGameAdViewed                       = 7
	EAppUsageEvent_GameLaunchFreeWeekend                = 8
)

func (e EAppUsageEvent) String() string {
	switch e {
	case EAppUsageEvent_GameLaunch:
		return "EAppUsageEvent_GameLaunch"
	case EAppUsageEvent_GameLaunchTrial:
		return "EAppUsageEvent_GameLaunchTrial"
	case EAppUsageEvent_Media:
		return "EAppUsageEvent_Media"
	case EAppUsageEvent_PreloadStart:
		return "EAppUsageEvent_PreloadStart"
	case EAppUsageEvent_PreloadFinish:
		return "EAppUsageEvent_PreloadFinish"
	case EAppUsageEvent_MarketingMessageView:
		return "EAppUsageEvent_MarketingMessageView"
	case EAppUsageEvent_InGameAdViewed:
		return "EAppUsageEvent_InGameAdViewed"
	case EAppUsageEvent_GameLaunchFreeWeekend:
		return "EAppUsageEvent_GameLaunchFreeWeekend"
	default:
		return "INVALID"
	}
}

type ELicenseFlags int32

const (
	ELicenseFlags_None               ELicenseFlags = 0
	ELicenseFlags_Renew                            = 0x01
	ELicenseFlags_RenewalFailed                    = 0x02
	ELicenseFlags_Pending                          = 0x04
	ELicenseFlags_Expired                          = 0x08
	ELicenseFlags_CancelledByUser                  = 0x10
	ELicenseFlags_CancelledByAdmin                 = 0x20
	ELicenseFlags_LowViolenceContent               = 0x40
	ELicenseFlags_ImportedFromSteam2               = 0x80
)

func (e ELicenseFlags) String() string {
	switch e {
	case ELicenseFlags_None:
		return "ELicenseFlags_None"
	case ELicenseFlags_Renew:
		return "ELicenseFlags_Renew"
	case ELicenseFlags_RenewalFailed:
		return "ELicenseFlags_RenewalFailed"
	case ELicenseFlags_Pending:
		return "ELicenseFlags_Pending"
	case ELicenseFlags_Expired:
		return "ELicenseFlags_Expired"
	case ELicenseFlags_CancelledByUser:
		return "ELicenseFlags_CancelledByUser"
	case ELicenseFlags_CancelledByAdmin:
		return "ELicenseFlags_CancelledByAdmin"
	case ELicenseFlags_LowViolenceContent:
		return "ELicenseFlags_LowViolenceContent"
	case ELicenseFlags_ImportedFromSteam2:
		return "ELicenseFlags_ImportedFromSteam2"
	default:
		return "INVALID"
	}
}

type ELicenseType int32

const (
	ELicenseType_NoLicense                             ELicenseType = 0
	ELicenseType_SinglePurchase                                     = 1
	ELicenseType_SinglePurchaseLimitedUse                           = 2
	ELicenseType_RecurringCharge                                    = 3
	ELicenseType_RecurringChargeLimitedUse                          = 4
	ELicenseType_RecurringChargeLimitedUseWithOverages              = 5
	ELicenseType_RecurringOption                                    = 6
)

func (e ELicenseType) String() string {
	switch e {
	case ELicenseType_NoLicense:
		return "ELicenseType_NoLicense"
	case ELicenseType_SinglePurchase:
		return "ELicenseType_SinglePurchase"
	case ELicenseType_SinglePurchaseLimitedUse:
		return "ELicenseType_SinglePurchaseLimitedUse"
	case ELicenseType_RecurringCharge:
		return "ELicenseType_RecurringCharge"
	case ELicenseType_RecurringChargeLimitedUse:
		return "ELicenseType_RecurringChargeLimitedUse"
	case ELicenseType_RecurringChargeLimitedUseWithOverages:
		return "ELicenseType_RecurringChargeLimitedUseWithOverages"
	case ELicenseType_RecurringOption:
		return "ELicenseType_RecurringOption"
	default:
		return "INVALID"
	}
}

type EPaymentMethod int32

const (
	EPaymentMethod_None                EPaymentMethod = 0
	EPaymentMethod_ActivationCode                     = 1
	EPaymentMethod_CreditCard                         = 2
	EPaymentMethod_Giropay                            = 3
	EPaymentMethod_PayPal                             = 4
	EPaymentMethod_Ideal                              = 5
	EPaymentMethod_PaySafeCard                        = 6
	EPaymentMethod_Sofort                             = 7
	EPaymentMethod_GuestPass                          = 8
	EPaymentMethod_WebMoney                           = 9
	EPaymentMethod_MoneyBookers                       = 10
	EPaymentMethod_AliPay                             = 11
	EPaymentMethod_Yandex                             = 12
	EPaymentMethod_Kiosk                              = 13
	EPaymentMethod_Qiwi                               = 14
	EPaymentMethod_GameStop                           = 15
	EPaymentMethod_HardwarePromo                      = 16
	EPaymentMethod_MoPay                              = 17
	EPaymentMethod_BoletoBancario                     = 18
	EPaymentMethod_BoaCompraGold                      = 19
	EPaymentMethod_BancoDoBrasilOnline                = 20
	EPaymentMethod_ItauOnline                         = 21
	EPaymentMethod_BradescoOnline                     = 22
	EPaymentMethod_Pagseguro                          = 23
	EPaymentMethod_VisaBrazil                         = 24
	EPaymentMethod_AmexBrazil                         = 25
	EPaymentMethod_Aura                               = 26
	EPaymentMethod_Hipercard                          = 27
	EPaymentMethod_MastercardBrazil                   = 28
	EPaymentMethod_DinersCardBrazil                   = 29
	EPaymentMethod_ClickAndBuy                        = 32
	EPaymentMethod_AutoGrant                          = 64
	EPaymentMethod_Wallet                             = 128
	EPaymentMethod_Valve                              = 129
	EPaymentMethod_OEMTicket                          = 256
	EPaymentMethod_Split                              = 512
	EPaymentMethod_Complimentary                      = 1024
)

func (e EPaymentMethod) String() string {
	switch e {
	case EPaymentMethod_None:
		return "EPaymentMethod_None"
	case EPaymentMethod_ActivationCode:
		return "EPaymentMethod_ActivationCode"
	case EPaymentMethod_CreditCard:
		return "EPaymentMethod_CreditCard"
	case EPaymentMethod_Giropay:
		return "EPaymentMethod_Giropay"
	case EPaymentMethod_PayPal:
		return "EPaymentMethod_PayPal"
	case EPaymentMethod_Ideal:
		return "EPaymentMethod_Ideal"
	case EPaymentMethod_PaySafeCard:
		return "EPaymentMethod_PaySafeCard"
	case EPaymentMethod_Sofort:
		return "EPaymentMethod_Sofort"
	case EPaymentMethod_GuestPass:
		return "EPaymentMethod_GuestPass"
	case EPaymentMethod_WebMoney:
		return "EPaymentMethod_WebMoney"
	case EPaymentMethod_MoneyBookers:
		return "EPaymentMethod_MoneyBookers"
	case EPaymentMethod_AliPay:
		return "EPaymentMethod_AliPay"
	case EPaymentMethod_Yandex:
		return "EPaymentMethod_Yandex"
	case EPaymentMethod_Kiosk:
		return "EPaymentMethod_Kiosk"
	case EPaymentMethod_Qiwi:
		return "EPaymentMethod_Qiwi"
	case EPaymentMethod_GameStop:
		return "EPaymentMethod_GameStop"
	case EPaymentMethod_HardwarePromo:
		return "EPaymentMethod_HardwarePromo"
	case EPaymentMethod_MoPay:
		return "EPaymentMethod_MoPay"
	case EPaymentMethod_BoletoBancario:
		return "EPaymentMethod_BoletoBancario"
	case EPaymentMethod_BoaCompraGold:
		return "EPaymentMethod_BoaCompraGold"
	case EPaymentMethod_BancoDoBrasilOnline:
		return "EPaymentMethod_BancoDoBrasilOnline"
	case EPaymentMethod_ItauOnline:
		return "EPaymentMethod_ItauOnline"
	case EPaymentMethod_BradescoOnline:
		return "EPaymentMethod_BradescoOnline"
	case EPaymentMethod_Pagseguro:
		return "EPaymentMethod_Pagseguro"
	case EPaymentMethod_VisaBrazil:
		return "EPaymentMethod_VisaBrazil"
	case EPaymentMethod_AmexBrazil:
		return "EPaymentMethod_AmexBrazil"
	case EPaymentMethod_Aura:
		return "EPaymentMethod_Aura"
	case EPaymentMethod_Hipercard:
		return "EPaymentMethod_Hipercard"
	case EPaymentMethod_MastercardBrazil:
		return "EPaymentMethod_MastercardBrazil"
	case EPaymentMethod_DinersCardBrazil:
		return "EPaymentMethod_DinersCardBrazil"
	case EPaymentMethod_ClickAndBuy:
		return "EPaymentMethod_ClickAndBuy"
	case EPaymentMethod_AutoGrant:
		return "EPaymentMethod_AutoGrant"
	case EPaymentMethod_Wallet:
		return "EPaymentMethod_Wallet"
	case EPaymentMethod_Valve:
		return "EPaymentMethod_Valve"
	case EPaymentMethod_OEMTicket:
		return "EPaymentMethod_OEMTicket"
	case EPaymentMethod_Split:
		return "EPaymentMethod_Split"
	case EPaymentMethod_Complimentary:
		return "EPaymentMethod_Complimentary"
	default:
		return "INVALID"
	}
}

type EIntroducerRouting int32

const (
	EIntroducerRouting_FileShare     EIntroducerRouting = 0 // Deprecated
	EIntroducerRouting_P2PVoiceChat                     = 1
	EIntroducerRouting_P2PNetworking                    = 2
)

func (e EIntroducerRouting) String() string {
	switch e {
	case EIntroducerRouting_FileShare:
		return "EIntroducerRouting_FileShare"
	case EIntroducerRouting_P2PVoiceChat:
		return "EIntroducerRouting_P2PVoiceChat"
	case EIntroducerRouting_P2PNetworking:
		return "EIntroducerRouting_P2PNetworking"
	default:
		return "INVALID"
	}
}

type EServerFlags int32

const (
	EServerFlags_None       EServerFlags = 0
	EServerFlags_Active                  = 1
	EServerFlags_Secure                  = 2
	EServerFlags_Dedicated               = 4
	EServerFlags_Linux                   = 8
	EServerFlags_Passworded              = 16
	EServerFlags_Private                 = 32
)

func (e EServerFlags) String() string {
	switch e {
	case EServerFlags_None:
		return "EServerFlags_None"
	case EServerFlags_Active:
		return "EServerFlags_Active"
	case EServerFlags_Secure:
		return "EServerFlags_Secure"
	case EServerFlags_Dedicated:
		return "EServerFlags_Dedicated"
	case EServerFlags_Linux:
		return "EServerFlags_Linux"
	case EServerFlags_Passworded:
		return "EServerFlags_Passworded"
	case EServerFlags_Private:
		return "EServerFlags_Private"
	default:
		return "INVALID"
	}
}

type EDenyReason int32

const (
	EDenyReason_InvalidVersion          EDenyReason = 1
	EDenyReason_Generic                             = 2
	EDenyReason_NotLoggedOn                         = 3
	EDenyReason_NoLicense                           = 4
	EDenyReason_Cheater                             = 5
	EDenyReason_LoggedInElseWhere                   = 6
	EDenyReason_UnknownText                         = 7
	EDenyReason_IncompatibleAnticheat               = 8
	EDenyReason_MemoryCorruption                    = 9
	EDenyReason_IncompatibleSoftware                = 10
	EDenyReason_SteamConnectionLost                 = 11
	EDenyReason_SteamConnectionError                = 12
	EDenyReason_SteamResponseTimedOut               = 13
	EDenyReason_SteamValidationStalled              = 14
	EDenyReason_SteamOwnerLeftGuestUser             = 15
)

func (e EDenyReason) String() string {
	switch e {
	case EDenyReason_InvalidVersion:
		return "EDenyReason_InvalidVersion"
	case EDenyReason_Generic:
		return "EDenyReason_Generic"
	case EDenyReason_NotLoggedOn:
		return "EDenyReason_NotLoggedOn"
	case EDenyReason_NoLicense:
		return "EDenyReason_NoLicense"
	case EDenyReason_Cheater:
		return "EDenyReason_Cheater"
	case EDenyReason_LoggedInElseWhere:
		return "EDenyReason_LoggedInElseWhere"
	case EDenyReason_UnknownText:
		return "EDenyReason_UnknownText"
	case EDenyReason_IncompatibleAnticheat:
		return "EDenyReason_IncompatibleAnticheat"
	case EDenyReason_MemoryCorruption:
		return "EDenyReason_MemoryCorruption"
	case EDenyReason_IncompatibleSoftware:
		return "EDenyReason_IncompatibleSoftware"
	case EDenyReason_SteamConnectionLost:
		return "EDenyReason_SteamConnectionLost"
	case EDenyReason_SteamConnectionError:
		return "EDenyReason_SteamConnectionError"
	case EDenyReason_SteamResponseTimedOut:
		return "EDenyReason_SteamResponseTimedOut"
	case EDenyReason_SteamValidationStalled:
		return "EDenyReason_SteamValidationStalled"
	case EDenyReason_SteamOwnerLeftGuestUser:
		return "EDenyReason_SteamOwnerLeftGuestUser"
	default:
		return "INVALID"
	}
}

type EClanRank int32

const (
	EClanRank_None      EClanRank = 0
	EClanRank_Owner               = 1
	EClanRank_Officer             = 2
	EClanRank_Member              = 3
	EClanRank_Moderator           = 4
)

func (e EClanRank) String() string {
	switch e {
	case EClanRank_None:
		return "EClanRank_None"
	case EClanRank_Owner:
		return "EClanRank_Owner"
	case EClanRank_Officer:
		return "EClanRank_Officer"
	case EClanRank_Member:
		return "EClanRank_Member"
	case EClanRank_Moderator:
		return "EClanRank_Moderator"
	default:
		return "INVALID"
	}
}

type EClanRelationship int32

const (
	EClanRelationship_None             EClanRelationship = 0
	EClanRelationship_Blocked                            = 1
	EClanRelationship_Invited                            = 2
	EClanRelationship_Member                             = 3
	EClanRelationship_Kicked                             = 4
	EClanRelationship_KickAcknowledged                   = 5
)

func (e EClanRelationship) String() string {
	switch e {
	case EClanRelationship_None:
		return "EClanRelationship_None"
	case EClanRelationship_Blocked:
		return "EClanRelationship_Blocked"
	case EClanRelationship_Invited:
		return "EClanRelationship_Invited"
	case EClanRelationship_Member:
		return "EClanRelationship_Member"
	case EClanRelationship_Kicked:
		return "EClanRelationship_Kicked"
	case EClanRelationship_KickAcknowledged:
		return "EClanRelationship_KickAcknowledged"
	default:
		return "INVALID"
	}
}

type EAuthSessionResponse int32

const (
	EAuthSessionResponse_OK                           EAuthSessionResponse = 0
	EAuthSessionResponse_UserNotConnectedToSteam                           = 1
	EAuthSessionResponse_NoLicenseOrExpired                                = 2
	EAuthSessionResponse_VACBanned                                         = 3
	EAuthSessionResponse_LoggedInElseWhere                                 = 4
	EAuthSessionResponse_VACCheckTimedOut                                  = 5
	EAuthSessionResponse_AuthTicketCanceled                                = 6
	EAuthSessionResponse_AuthTicketInvalidAlreadyUsed                      = 7
	EAuthSessionResponse_AuthTicketInvalid                                 = 8
)

func (e EAuthSessionResponse) String() string {
	switch e {
	case EAuthSessionResponse_OK:
		return "EAuthSessionResponse_OK"
	case EAuthSessionResponse_UserNotConnectedToSteam:
		return "EAuthSessionResponse_UserNotConnectedToSteam"
	case EAuthSessionResponse_NoLicenseOrExpired:
		return "EAuthSessionResponse_NoLicenseOrExpired"
	case EAuthSessionResponse_VACBanned:
		return "EAuthSessionResponse_VACBanned"
	case EAuthSessionResponse_LoggedInElseWhere:
		return "EAuthSessionResponse_LoggedInElseWhere"
	case EAuthSessionResponse_VACCheckTimedOut:
		return "EAuthSessionResponse_VACCheckTimedOut"
	case EAuthSessionResponse_AuthTicketCanceled:
		return "EAuthSessionResponse_AuthTicketCanceled"
	case EAuthSessionResponse_AuthTicketInvalidAlreadyUsed:
		return "EAuthSessionResponse_AuthTicketInvalidAlreadyUsed"
	case EAuthSessionResponse_AuthTicketInvalid:
		return "EAuthSessionResponse_AuthTicketInvalid"
	default:
		return "INVALID"
	}
}

type EChatRoomEnterResponse int32

const (
	EChatRoomEnterResponse_Success            EChatRoomEnterResponse = 1
	EChatRoomEnterResponse_DoesntExist                               = 2
	EChatRoomEnterResponse_NotAllowed                                = 3
	EChatRoomEnterResponse_Full                                      = 4
	EChatRoomEnterResponse_Error                                     = 5
	EChatRoomEnterResponse_Banned                                    = 6
	EChatRoomEnterResponse_Limited                                   = 7
	EChatRoomEnterResponse_ClanDisabled                              = 8
	EChatRoomEnterResponse_CommunityBan                              = 9
	EChatRoomEnterResponse_MemberBlockedYou                          = 10
	EChatRoomEnterResponse_YouBlockedMember                          = 11
	EChatRoomEnterResponse_NoRankingDataLobby                        = 12 // Deprecated
	EChatRoomEnterResponse_NoRankingDataUser                         = 13 // Deprecated
	EChatRoomEnterResponse_RankOutOfRange                            = 14 // Deprecated
)

func (e EChatRoomEnterResponse) String() string {
	switch e {
	case EChatRoomEnterResponse_Success:
		return "EChatRoomEnterResponse_Success"
	case EChatRoomEnterResponse_DoesntExist:
		return "EChatRoomEnterResponse_DoesntExist"
	case EChatRoomEnterResponse_NotAllowed:
		return "EChatRoomEnterResponse_NotAllowed"
	case EChatRoomEnterResponse_Full:
		return "EChatRoomEnterResponse_Full"
	case EChatRoomEnterResponse_Error:
		return "EChatRoomEnterResponse_Error"
	case EChatRoomEnterResponse_Banned:
		return "EChatRoomEnterResponse_Banned"
	case EChatRoomEnterResponse_Limited:
		return "EChatRoomEnterResponse_Limited"
	case EChatRoomEnterResponse_ClanDisabled:
		return "EChatRoomEnterResponse_ClanDisabled"
	case EChatRoomEnterResponse_CommunityBan:
		return "EChatRoomEnterResponse_CommunityBan"
	case EChatRoomEnterResponse_MemberBlockedYou:
		return "EChatRoomEnterResponse_MemberBlockedYou"
	case EChatRoomEnterResponse_YouBlockedMember:
		return "EChatRoomEnterResponse_YouBlockedMember"
	case EChatRoomEnterResponse_NoRankingDataLobby:
		return "EChatRoomEnterResponse_NoRankingDataLobby"
	case EChatRoomEnterResponse_NoRankingDataUser:
		return "EChatRoomEnterResponse_NoRankingDataUser"
	case EChatRoomEnterResponse_RankOutOfRange:
		return "EChatRoomEnterResponse_RankOutOfRange"
	default:
		return "INVALID"
	}
}

type EChatRoomType int32

const (
	EChatRoomType_Friend EChatRoomType = 1
	EChatRoomType_MUC                  = 2
	EChatRoomType_Lobby                = 3
)

func (e EChatRoomType) String() string {
	switch e {
	case EChatRoomType_Friend:
		return "EChatRoomType_Friend"
	case EChatRoomType_MUC:
		return "EChatRoomType_MUC"
	case EChatRoomType_Lobby:
		return "EChatRoomType_Lobby"
	default:
		return "INVALID"
	}
}

type EChatInfoType int32

const (
	EChatInfoType_StateChange       EChatInfoType = 1
	EChatInfoType_InfoUpdate                      = 2
	EChatInfoType_MemberLimitChange               = 3
)

func (e EChatInfoType) String() string {
	switch e {
	case EChatInfoType_StateChange:
		return "EChatInfoType_StateChange"
	case EChatInfoType_InfoUpdate:
		return "EChatInfoType_InfoUpdate"
	case EChatInfoType_MemberLimitChange:
		return "EChatInfoType_MemberLimitChange"
	default:
		return "INVALID"
	}
}

type EChatAction int32

const (
	EChatAction_InviteChat            EChatAction = 1
	EChatAction_Kick                              = 2
	EChatAction_Ban                               = 3
	EChatAction_UnBan                             = 4
	EChatAction_StartVoiceSpeak                   = 5
	EChatAction_EndVoiceSpeak                     = 6
	EChatAction_LockChat                          = 7
	EChatAction_UnlockChat                        = 8
	EChatAction_CloseChat                         = 9
	EChatAction_SetJoinable                       = 10
	EChatAction_SetUnjoinable                     = 11
	EChatAction_SetOwner                          = 12
	EChatAction_SetInvisibleToFriends             = 13
	EChatAction_SetVisibleToFriends               = 14
	EChatAction_SetModerated                      = 15
	EChatAction_SetUnmoderated                    = 16
)

func (e EChatAction) String() string {
	switch e {
	case EChatAction_InviteChat:
		return "EChatAction_InviteChat"
	case EChatAction_Kick:
		return "EChatAction_Kick"
	case EChatAction_Ban:
		return "EChatAction_Ban"
	case EChatAction_UnBan:
		return "EChatAction_UnBan"
	case EChatAction_StartVoiceSpeak:
		return "EChatAction_StartVoiceSpeak"
	case EChatAction_EndVoiceSpeak:
		return "EChatAction_EndVoiceSpeak"
	case EChatAction_LockChat:
		return "EChatAction_LockChat"
	case EChatAction_UnlockChat:
		return "EChatAction_UnlockChat"
	case EChatAction_CloseChat:
		return "EChatAction_CloseChat"
	case EChatAction_SetJoinable:
		return "EChatAction_SetJoinable"
	case EChatAction_SetUnjoinable:
		return "EChatAction_SetUnjoinable"
	case EChatAction_SetOwner:
		return "EChatAction_SetOwner"
	case EChatAction_SetInvisibleToFriends:
		return "EChatAction_SetInvisibleToFriends"
	case EChatAction_SetVisibleToFriends:
		return "EChatAction_SetVisibleToFriends"
	case EChatAction_SetModerated:
		return "EChatAction_SetModerated"
	case EChatAction_SetUnmoderated:
		return "EChatAction_SetUnmoderated"
	default:
		return "INVALID"
	}
}

type EChatActionResult int32

const (
	EChatActionResult_Success                EChatActionResult = 1
	EChatActionResult_Error                                    = 2
	EChatActionResult_NotPermitted                             = 3
	EChatActionResult_NotAllowedOnClanMember                   = 4
	EChatActionResult_NotAllowedOnBannedUser                   = 5
	EChatActionResult_NotAllowedOnChatOwner                    = 6
	EChatActionResult_NotAllowedOnSelf                         = 7
	EChatActionResult_ChatDoesntExist                          = 8
	EChatActionResult_ChatFull                                 = 9
	EChatActionResult_VoiceSlotsFull                           = 10
)

func (e EChatActionResult) String() string {
	switch e {
	case EChatActionResult_Success:
		return "EChatActionResult_Success"
	case EChatActionResult_Error:
		return "EChatActionResult_Error"
	case EChatActionResult_NotPermitted:
		return "EChatActionResult_NotPermitted"
	case EChatActionResult_NotAllowedOnClanMember:
		return "EChatActionResult_NotAllowedOnClanMember"
	case EChatActionResult_NotAllowedOnBannedUser:
		return "EChatActionResult_NotAllowedOnBannedUser"
	case EChatActionResult_NotAllowedOnChatOwner:
		return "EChatActionResult_NotAllowedOnChatOwner"
	case EChatActionResult_NotAllowedOnSelf:
		return "EChatActionResult_NotAllowedOnSelf"
	case EChatActionResult_ChatDoesntExist:
		return "EChatActionResult_ChatDoesntExist"
	case EChatActionResult_ChatFull:
		return "EChatActionResult_ChatFull"
	case EChatActionResult_VoiceSlotsFull:
		return "EChatActionResult_VoiceSlotsFull"
	default:
		return "INVALID"
	}
}

type EAppInfoSection int32

const (
	EAppInfoSection_Unknown     EAppInfoSection = 0
	EAppInfoSection_All                         = 1
	EAppInfoSection_First                       = 2
	EAppInfoSection_Common                      = 2
	EAppInfoSection_Extended                    = 3
	EAppInfoSection_Config                      = 4
	EAppInfoSection_Stats                       = 5
	EAppInfoSection_Install                     = 6
	EAppInfoSection_Depots                      = 7
	EAppInfoSection_VAC                         = 8
	EAppInfoSection_DRM                         = 9
	EAppInfoSection_UFS                         = 10
	EAppInfoSection_OGG                         = 11
	EAppInfoSection_Items                       = 12 // Deprecated
	EAppInfoSection_ItemsUNUSED                 = 12 // Deprecated
	EAppInfoSection_Policies                    = 13
	EAppInfoSection_SysReqs                     = 14
	EAppInfoSection_Community                   = 15
	EAppInfoSection_Max                         = 16
)

func (e EAppInfoSection) String() string {
	switch e {
	case EAppInfoSection_Unknown:
		return "EAppInfoSection_Unknown"
	case EAppInfoSection_All:
		return "EAppInfoSection_All"
	case EAppInfoSection_First:
		return "EAppInfoSection_First"
	case EAppInfoSection_Extended:
		return "EAppInfoSection_Extended"
	case EAppInfoSection_Config:
		return "EAppInfoSection_Config"
	case EAppInfoSection_Stats:
		return "EAppInfoSection_Stats"
	case EAppInfoSection_Install:
		return "EAppInfoSection_Install"
	case EAppInfoSection_Depots:
		return "EAppInfoSection_Depots"
	case EAppInfoSection_VAC:
		return "EAppInfoSection_VAC"
	case EAppInfoSection_DRM:
		return "EAppInfoSection_DRM"
	case EAppInfoSection_UFS:
		return "EAppInfoSection_UFS"
	case EAppInfoSection_OGG:
		return "EAppInfoSection_OGG"
	case EAppInfoSection_Items:
		return "EAppInfoSection_Items"
	case EAppInfoSection_Policies:
		return "EAppInfoSection_Policies"
	case EAppInfoSection_SysReqs:
		return "EAppInfoSection_SysReqs"
	case EAppInfoSection_Community:
		return "EAppInfoSection_Community"
	case EAppInfoSection_Max:
		return "EAppInfoSection_Max"
	default:
		return "INVALID"
	}
}

type EContentDownloadSourceType int32

const (
	EContentDownloadSourceType_Invalid    EContentDownloadSourceType = 0
	EContentDownloadSourceType_CS                                    = 1
	EContentDownloadSourceType_CDN                                   = 2
	EContentDownloadSourceType_LCS                                   = 3
	EContentDownloadSourceType_Proxy                                 = 4 // Deprecated: renamed to ProxyCache
	EContentDownloadSourceType_ProxyCache                            = 4
	EContentDownloadSourceType_Max                                   = 5
)

func (e EContentDownloadSourceType) String() string {
	switch e {
	case EContentDownloadSourceType_Invalid:
		return "EContentDownloadSourceType_Invalid"
	case EContentDownloadSourceType_CS:
		return "EContentDownloadSourceType_CS"
	case EContentDownloadSourceType_CDN:
		return "EContentDownloadSourceType_CDN"
	case EContentDownloadSourceType_LCS:
		return "EContentDownloadSourceType_LCS"
	case EContentDownloadSourceType_Proxy:
		return "EContentDownloadSourceType_Proxy"
	case EContentDownloadSourceType_Max:
		return "EContentDownloadSourceType_Max"
	default:
		return "INVALID"
	}
}

type EPlatformType int32

const (
	EPlatformType_Unknown EPlatformType = 0
	EPlatformType_Win32                 = 1
	EPlatformType_Win64                 = 2
	EPlatformType_Linux                 = 3
	EPlatformType_OSX                   = 4
	EPlatformType_PS3                   = 5
	EPlatformType_Max                   = 6
)

func (e EPlatformType) String() string {
	switch e {
	case EPlatformType_Unknown:
		return "EPlatformType_Unknown"
	case EPlatformType_Win32:
		return "EPlatformType_Win32"
	case EPlatformType_Win64:
		return "EPlatformType_Win64"
	case EPlatformType_Linux:
		return "EPlatformType_Linux"
	case EPlatformType_OSX:
		return "EPlatformType_OSX"
	case EPlatformType_PS3:
		return "EPlatformType_PS3"
	case EPlatformType_Max:
		return "EPlatformType_Max"
	default:
		return "INVALID"
	}
}

type EOSType int32

const (
	EOSType_Unknown        EOSType = -1
	EOSType_UMQ                    = -400
	EOSType_PS3                    = -300
	EOSType_MacOSUnknown           = -102
	EOSType_MacOS104               = -101
	EOSType_MacOS105               = -100
	EOSType_MacOS1058              = -99
	EOSType_MacOS106               = -95
	EOSType_MacOS1063              = -94
	EOSType_MacOS1064_slgu         = -93
	EOSType_MacOS1067              = -92
	EOSType_MacOS107               = -90
	EOSType_MacOS108               = -89
	EOSType_MacOS109               = -88
	EOSType_LinuxUnknown           = -203
	EOSType_Linux22                = -202
	EOSType_Linux24                = -201
	EOSType_Linux26                = -200
	EOSType_Linux32                = -199
	EOSType_Linux35                = -198
	EOSType_Linux36                = -197
	EOSType_Linux310               = -196
	EOSType_WinUnknown             = 0
	EOSType_Win311                 = 1
	EOSType_Win95                  = 2
	EOSType_Win98                  = 3
	EOSType_WinME                  = 4
	EOSType_WinNT                  = 5
	EOSType_Win200                 = 6
	EOSType_WinXP                  = 7
	EOSType_Win2003                = 8
	EOSType_WinVista               = 9
	EOSType_Win7                   = 10 // Deprecated: renamed to Windows7
	EOSType_Windows7               = 10
	EOSType_Win2008                = 11
	EOSType_Win2012                = 12
	EOSType_Windows8               = 13
	EOSType_Windows81              = 14
	EOSType_WinMAX                 = 15
	EOSType_Max                    = 26
)

func (e EOSType) String() string {
	switch e {
	case EOSType_Unknown:
		return "EOSType_Unknown"
	case EOSType_UMQ:
		return "EOSType_UMQ"
	case EOSType_PS3:
		return "EOSType_PS3"
	case EOSType_MacOSUnknown:
		return "EOSType_MacOSUnknown"
	case EOSType_MacOS104:
		return "EOSType_MacOS104"
	case EOSType_MacOS105:
		return "EOSType_MacOS105"
	case EOSType_MacOS1058:
		return "EOSType_MacOS1058"
	case EOSType_MacOS106:
		return "EOSType_MacOS106"
	case EOSType_MacOS1063:
		return "EOSType_MacOS1063"
	case EOSType_MacOS1064_slgu:
		return "EOSType_MacOS1064_slgu"
	case EOSType_MacOS1067:
		return "EOSType_MacOS1067"
	case EOSType_MacOS107:
		return "EOSType_MacOS107"
	case EOSType_MacOS108:
		return "EOSType_MacOS108"
	case EOSType_MacOS109:
		return "EOSType_MacOS109"
	case EOSType_LinuxUnknown:
		return "EOSType_LinuxUnknown"
	case EOSType_Linux22:
		return "EOSType_Linux22"
	case EOSType_Linux24:
		return "EOSType_Linux24"
	case EOSType_Linux26:
		return "EOSType_Linux26"
	case EOSType_Linux32:
		return "EOSType_Linux32"
	case EOSType_Linux35:
		return "EOSType_Linux35"
	case EOSType_Linux36:
		return "EOSType_Linux36"
	case EOSType_Linux310:
		return "EOSType_Linux310"
	case EOSType_WinUnknown:
		return "EOSType_WinUnknown"
	case EOSType_Win311:
		return "EOSType_Win311"
	case EOSType_Win95:
		return "EOSType_Win95"
	case EOSType_Win98:
		return "EOSType_Win98"
	case EOSType_WinME:
		return "EOSType_WinME"
	case EOSType_WinNT:
		return "EOSType_WinNT"
	case EOSType_Win200:
		return "EOSType_Win200"
	case EOSType_WinXP:
		return "EOSType_WinXP"
	case EOSType_Win2003:
		return "EOSType_Win2003"
	case EOSType_WinVista:
		return "EOSType_WinVista"
	case EOSType_Win7:
		return "EOSType_Win7"
	case EOSType_Win2008:
		return "EOSType_Win2008"
	case EOSType_Win2012:
		return "EOSType_Win2012"
	case EOSType_Windows8:
		return "EOSType_Windows8"
	case EOSType_Windows81:
		return "EOSType_Windows81"
	case EOSType_WinMAX:
		return "EOSType_WinMAX"
	case EOSType_Max:
		return "EOSType_Max"
	default:
		return "INVALID"
	}
}

type EServerType int32

const (
	EServerType_Invalid           EServerType = -1
	EServerType_First                         = 0
	EServerType_Shell                         = 0
	EServerType_GM                            = 1
	EServerType_BUM                           = 2
	EServerType_AM                            = 3
	EServerType_BS                            = 4
	EServerType_VS                            = 5
	EServerType_ATS                           = 6
	EServerType_CM                            = 7
	EServerType_FBS                           = 8
	EServerType_FG                            = 9 // Deprecated: renamed to BoxMonitor
	EServerType_BoxMonitor                    = 9
	EServerType_SS                            = 10
	EServerType_DRMS                          = 11
	EServerType_HubOBSOLETE                   = 12 // Deprecated
	EServerType_Console                       = 13
	EServerType_ASBOBSOLETE                   = 14 // Deprecated
	EServerType_PICS                          = 14
	EServerType_Client                        = 15
	EServerType_BootstrapOBSOLETE             = 16 // Deprecated
	EServerType_DP                            = 17
	EServerType_WG                            = 18
	EServerType_SM                            = 19
	EServerType_UFS                           = 21
	EServerType_Util                          = 23
	EServerType_DSS                           = 24 // Deprecated: renamed to Community
	EServerType_Community                     = 24
	EServerType_P2PRelayOBSOLETE              = 25 // Deprecated
	EServerType_AppInformation                = 26
	EServerType_Spare                         = 27
	EServerType_FTS                           = 28
	EServerType_EPM                           = 29
	EServerType_PS                            = 30
	EServerType_IS                            = 31
	EServerType_CCS                           = 32
	EServerType_DFS                           = 33
	EServerType_LBS                           = 34
	EServerType_MDS                           = 35
	EServerType_CS                            = 36
	EServerType_GC                            = 37
	EServerType_NS                            = 38
	EServerType_OGS                           = 39
	EServerType_WebAPI                        = 40
	EServerType_UDS                           = 41
	EServerType_MMS                           = 42
	EServerType_GMS                           = 43
	EServerType_KGS                           = 44
	EServerType_UCM                           = 45
	EServerType_RM                            = 46
	EServerType_FS                            = 47
	EServerType_Econ                          = 48
	EServerType_Backpack                      = 49
	EServerType_UGS                           = 50
	EServerType_Store                         = 51
	EServerType_MoneyStats                    = 52
	EServerType_CRE                           = 53
	EServerType_UMQ                           = 54
	EServerType_Workshop                      = 55
	EServerType_BRP                           = 56
	EServerType_GCH                           = 57
	EServerType_MPAS                          = 58
	EServerType_Trade                         = 59
	EServerType_Secrets                       = 60
	EServerType_Logsink                       = 61
	EServerType_Market                        = 62
	EServerType_Quest                         = 63
	EServerType_WDS                           = 64
	EServerType_ACS                           = 65
	EServerType_PNP                           = 66
	EServerType_Max                           = 67
)

func (e EServerType) String() string {
	switch e {
	case EServerType_Invalid:
		return "EServerType_Invalid"
	case EServerType_First:
		return "EServerType_First"
	case EServerType_GM:
		return "EServerType_GM"
	case EServerType_BUM:
		return "EServerType_BUM"
	case EServerType_AM:
		return "EServerType_AM"
	case EServerType_BS:
		return "EServerType_BS"
	case EServerType_VS:
		return "EServerType_VS"
	case EServerType_ATS:
		return "EServerType_ATS"
	case EServerType_CM:
		return "EServerType_CM"
	case EServerType_FBS:
		return "EServerType_FBS"
	case EServerType_FG:
		return "EServerType_FG"
	case EServerType_SS:
		return "EServerType_SS"
	case EServerType_DRMS:
		return "EServerType_DRMS"
	case EServerType_HubOBSOLETE:
		return "EServerType_HubOBSOLETE"
	case EServerType_Console:
		return "EServerType_Console"
	case EServerType_ASBOBSOLETE:
		return "EServerType_ASBOBSOLETE"
	case EServerType_Client:
		return "EServerType_Client"
	case EServerType_BootstrapOBSOLETE:
		return "EServerType_BootstrapOBSOLETE"
	case EServerType_DP:
		return "EServerType_DP"
	case EServerType_WG:
		return "EServerType_WG"
	case EServerType_SM:
		return "EServerType_SM"
	case EServerType_UFS:
		return "EServerType_UFS"
	case EServerType_Util:
		return "EServerType_Util"
	case EServerType_DSS:
		return "EServerType_DSS"
	case EServerType_P2PRelayOBSOLETE:
		return "EServerType_P2PRelayOBSOLETE"
	case EServerType_AppInformation:
		return "EServerType_AppInformation"
	case EServerType_Spare:
		return "EServerType_Spare"
	case EServerType_FTS:
		return "EServerType_FTS"
	case EServerType_EPM:
		return "EServerType_EPM"
	case EServerType_PS:
		return "EServerType_PS"
	case EServerType_IS:
		return "EServerType_IS"
	case EServerType_CCS:
		return "EServerType_CCS"
	case EServerType_DFS:
		return "EServerType_DFS"
	case EServerType_LBS:
		return "EServerType_LBS"
	case EServerType_MDS:
		return "EServerType_MDS"
	case EServerType_CS:
		return "EServerType_CS"
	case EServerType_GC:
		return "EServerType_GC"
	case EServerType_NS:
		return "EServerType_NS"
	case EServerType_OGS:
		return "EServerType_OGS"
	case EServerType_WebAPI:
		return "EServerType_WebAPI"
	case EServerType_UDS:
		return "EServerType_UDS"
	case EServerType_MMS:
		return "EServerType_MMS"
	case EServerType_GMS:
		return "EServerType_GMS"
	case EServerType_KGS:
		return "EServerType_KGS"
	case EServerType_UCM:
		return "EServerType_UCM"
	case EServerType_RM:
		return "EServerType_RM"
	case EServerType_FS:
		return "EServerType_FS"
	case EServerType_Econ:
		return "EServerType_Econ"
	case EServerType_Backpack:
		return "EServerType_Backpack"
	case EServerType_UGS:
		return "EServerType_UGS"
	case EServerType_Store:
		return "EServerType_Store"
	case EServerType_MoneyStats:
		return "EServerType_MoneyStats"
	case EServerType_CRE:
		return "EServerType_CRE"
	case EServerType_UMQ:
		return "EServerType_UMQ"
	case EServerType_Workshop:
		return "EServerType_Workshop"
	case EServerType_BRP:
		return "EServerType_BRP"
	case EServerType_GCH:
		return "EServerType_GCH"
	case EServerType_MPAS:
		return "EServerType_MPAS"
	case EServerType_Trade:
		return "EServerType_Trade"
	case EServerType_Secrets:
		return "EServerType_Secrets"
	case EServerType_Logsink:
		return "EServerType_Logsink"
	case EServerType_Market:
		return "EServerType_Market"
	case EServerType_Quest:
		return "EServerType_Quest"
	case EServerType_WDS:
		return "EServerType_WDS"
	case EServerType_ACS:
		return "EServerType_ACS"
	case EServerType_PNP:
		return "EServerType_PNP"
	case EServerType_Max:
		return "EServerType_Max"
	default:
		return "INVALID"
	}
}

type EBillingType int32

const (
	EBillingType_NoCost                 EBillingType = 0
	EBillingType_BillOnceOnly                        = 1
	EBillingType_BillMonthly                         = 2
	EBillingType_ProofOfPrepurchaseOnly              = 3
	EBillingType_GuestPass                           = 4
	EBillingType_HardwarePromo                       = 5
	EBillingType_Gift                                = 6
	EBillingType_AutoGrant                           = 7
	EBillingType_OEMTicket                           = 8
	EBillingType_RecurringOption                     = 9
	EBillingType_NumBillingTypes                     = 10
)

func (e EBillingType) String() string {
	switch e {
	case EBillingType_NoCost:
		return "EBillingType_NoCost"
	case EBillingType_BillOnceOnly:
		return "EBillingType_BillOnceOnly"
	case EBillingType_BillMonthly:
		return "EBillingType_BillMonthly"
	case EBillingType_ProofOfPrepurchaseOnly:
		return "EBillingType_ProofOfPrepurchaseOnly"
	case EBillingType_GuestPass:
		return "EBillingType_GuestPass"
	case EBillingType_HardwarePromo:
		return "EBillingType_HardwarePromo"
	case EBillingType_Gift:
		return "EBillingType_Gift"
	case EBillingType_AutoGrant:
		return "EBillingType_AutoGrant"
	case EBillingType_OEMTicket:
		return "EBillingType_OEMTicket"
	case EBillingType_RecurringOption:
		return "EBillingType_RecurringOption"
	case EBillingType_NumBillingTypes:
		return "EBillingType_NumBillingTypes"
	default:
		return "INVALID"
	}
}

type EActivationCodeClass uint32

const (
	EActivationCodeClass_WonCDKey     EActivationCodeClass = 0
	EActivationCodeClass_ValveCDKey                        = 1
	EActivationCodeClass_Doom3CDKey                        = 2
	EActivationCodeClass_DBLookup                          = 3
	EActivationCodeClass_Steam2010Key                      = 4
	EActivationCodeClass_Max                               = 5
	EActivationCodeClass_Test                              = 2147483647
	EActivationCodeClass_Invalid                           = 4294967295
)

func (e EActivationCodeClass) String() string {
	switch e {
	case EActivationCodeClass_WonCDKey:
		return "EActivationCodeClass_WonCDKey"
	case EActivationCodeClass_ValveCDKey:
		return "EActivationCodeClass_ValveCDKey"
	case EActivationCodeClass_Doom3CDKey:
		return "EActivationCodeClass_Doom3CDKey"
	case EActivationCodeClass_DBLookup:
		return "EActivationCodeClass_DBLookup"
	case EActivationCodeClass_Steam2010Key:
		return "EActivationCodeClass_Steam2010Key"
	case EActivationCodeClass_Max:
		return "EActivationCodeClass_Max"
	case EActivationCodeClass_Test:
		return "EActivationCodeClass_Test"
	case EActivationCodeClass_Invalid:
		return "EActivationCodeClass_Invalid"
	default:
		return "INVALID"
	}
}

type EChatMemberStateChange int32

const (
	EChatMemberStateChange_Entered           EChatMemberStateChange = 0x01
	EChatMemberStateChange_Left                                     = 0x02
	EChatMemberStateChange_Disconnected                             = 0x04
	EChatMemberStateChange_Kicked                                   = 0x08
	EChatMemberStateChange_Banned                                   = 0x10
	EChatMemberStateChange_VoiceSpeaking                            = 0x1000
	EChatMemberStateChange_VoiceDoneSpeaking                        = 0x2000
)

func (e EChatMemberStateChange) String() string {
	switch e {
	case EChatMemberStateChange_Entered:
		return "EChatMemberStateChange_Entered"
	case EChatMemberStateChange_Left:
		return "EChatMemberStateChange_Left"
	case EChatMemberStateChange_Disconnected:
		return "EChatMemberStateChange_Disconnected"
	case EChatMemberStateChange_Kicked:
		return "EChatMemberStateChange_Kicked"
	case EChatMemberStateChange_Banned:
		return "EChatMemberStateChange_Banned"
	case EChatMemberStateChange_VoiceSpeaking:
		return "EChatMemberStateChange_VoiceSpeaking"
	case EChatMemberStateChange_VoiceDoneSpeaking:
		return "EChatMemberStateChange_VoiceDoneSpeaking"
	default:
		return "INVALID"
	}
}

type ERegionCode uint8

const (
	ERegionCode_USEast       ERegionCode = 0x00
	ERegionCode_USWest                   = 0x01
	ERegionCode_SouthAmerica             = 0x02
	ERegionCode_Europe                   = 0x03
	ERegionCode_Asia                     = 0x04
	ERegionCode_Australia                = 0x05
	ERegionCode_MiddleEast               = 0x06
	ERegionCode_Africa                   = 0x07
	ERegionCode_World                    = 0xFF
)

func (e ERegionCode) String() string {
	switch e {
	case ERegionCode_USEast:
		return "ERegionCode_USEast"
	case ERegionCode_USWest:
		return "ERegionCode_USWest"
	case ERegionCode_SouthAmerica:
		return "ERegionCode_SouthAmerica"
	case ERegionCode_Europe:
		return "ERegionCode_Europe"
	case ERegionCode_Asia:
		return "ERegionCode_Asia"
	case ERegionCode_Australia:
		return "ERegionCode_Australia"
	case ERegionCode_MiddleEast:
		return "ERegionCode_MiddleEast"
	case ERegionCode_Africa:
		return "ERegionCode_Africa"
	case ERegionCode_World:
		return "ERegionCode_World"
	default:
		return "INVALID"
	}
}

type ECurrencyCode int32

const (
	ECurrencyCode_Invalid ECurrencyCode = 0
	ECurrencyCode_USD                   = 1
	ECurrencyCode_GBP                   = 2
	ECurrencyCode_EUR                   = 3
	ECurrencyCode_CHF                   = 4
	ECurrencyCode_RUB                   = 5
	ECurrencyCode_PLN                   = 6
	ECurrencyCode_BRL                   = 7
	ECurrencyCode_NOK                   = 9
	ECurrencyCode_Max                   = 10
)

func (e ECurrencyCode) String() string {
	switch e {
	case ECurrencyCode_Invalid:
		return "ECurrencyCode_Invalid"
	case ECurrencyCode_USD:
		return "ECurrencyCode_USD"
	case ECurrencyCode_GBP:
		return "ECurrencyCode_GBP"
	case ECurrencyCode_EUR:
		return "ECurrencyCode_EUR"
	case ECurrencyCode_CHF:
		return "ECurrencyCode_CHF"
	case ECurrencyCode_RUB:
		return "ECurrencyCode_RUB"
	case ECurrencyCode_PLN:
		return "ECurrencyCode_PLN"
	case ECurrencyCode_BRL:
		return "ECurrencyCode_BRL"
	case ECurrencyCode_NOK:
		return "ECurrencyCode_NOK"
	case ECurrencyCode_Max:
		return "ECurrencyCode_Max"
	default:
		return "INVALID"
	}
}

type EDepotFileFlag int32

const (
	EDepotFileFlag_UserConfig          EDepotFileFlag = 1
	EDepotFileFlag_VersionedUserConfig                = 2
	EDepotFileFlag_Encrypted                          = 4
	EDepotFileFlag_ReadOnly                           = 8
	EDepotFileFlag_Hidden                             = 16
	EDepotFileFlag_Executable                         = 32
	EDepotFileFlag_Directory                          = 64
	EDepotFileFlag_CustomExecutable                   = 128
	EDepotFileFlag_InstallScript                      = 256
)

func (e EDepotFileFlag) String() string {
	switch e {
	case EDepotFileFlag_UserConfig:
		return "EDepotFileFlag_UserConfig"
	case EDepotFileFlag_VersionedUserConfig:
		return "EDepotFileFlag_VersionedUserConfig"
	case EDepotFileFlag_Encrypted:
		return "EDepotFileFlag_Encrypted"
	case EDepotFileFlag_ReadOnly:
		return "EDepotFileFlag_ReadOnly"
	case EDepotFileFlag_Hidden:
		return "EDepotFileFlag_Hidden"
	case EDepotFileFlag_Executable:
		return "EDepotFileFlag_Executable"
	case EDepotFileFlag_Directory:
		return "EDepotFileFlag_Directory"
	case EDepotFileFlag_CustomExecutable:
		return "EDepotFileFlag_CustomExecutable"
	case EDepotFileFlag_InstallScript:
		return "EDepotFileFlag_InstallScript"
	default:
		return "INVALID"
	}
}

type EWorkshopEnumerationType int32

const (
	EWorkshopEnumerationType_RankedByVote            EWorkshopEnumerationType = 0
	EWorkshopEnumerationType_Recent                                           = 1
	EWorkshopEnumerationType_Trending                                         = 2
	EWorkshopEnumerationType_FavoriteOfFriends                                = 3
	EWorkshopEnumerationType_VotedByFriends                                   = 4
	EWorkshopEnumerationType_ContentByFriends                                 = 5
	EWorkshopEnumerationType_RecentFromFollowedUsers                          = 6
)

func (e EWorkshopEnumerationType) String() string {
	switch e {
	case EWorkshopEnumerationType_RankedByVote:
		return "EWorkshopEnumerationType_RankedByVote"
	case EWorkshopEnumerationType_Recent:
		return "EWorkshopEnumerationType_Recent"
	case EWorkshopEnumerationType_Trending:
		return "EWorkshopEnumerationType_Trending"
	case EWorkshopEnumerationType_FavoriteOfFriends:
		return "EWorkshopEnumerationType_FavoriteOfFriends"
	case EWorkshopEnumerationType_VotedByFriends:
		return "EWorkshopEnumerationType_VotedByFriends"
	case EWorkshopEnumerationType_ContentByFriends:
		return "EWorkshopEnumerationType_ContentByFriends"
	case EWorkshopEnumerationType_RecentFromFollowedUsers:
		return "EWorkshopEnumerationType_RecentFromFollowedUsers"
	default:
		return "INVALID"
	}
}

type EPublishedFileVisibility int32

const (
	EPublishedFileVisibility_Public      EPublishedFileVisibility = 0
	EPublishedFileVisibility_FriendsOnly                          = 1
	EPublishedFileVisibility_Private                              = 2
)

func (e EPublishedFileVisibility) String() string {
	switch e {
	case EPublishedFileVisibility_Public:
		return "EPublishedFileVisibility_Public"
	case EPublishedFileVisibility_FriendsOnly:
		return "EPublishedFileVisibility_FriendsOnly"
	case EPublishedFileVisibility_Private:
		return "EPublishedFileVisibility_Private"
	default:
		return "INVALID"
	}
}

type EWorkshopFileType int32

const (
	EWorkshopFileType_First                  EWorkshopFileType = 0
	EWorkshopFileType_Community                                = 0
	EWorkshopFileType_Microtransaction                         = 1
	EWorkshopFileType_Collection                               = 2
	EWorkshopFileType_Art                                      = 3
	EWorkshopFileType_Video                                    = 4
	EWorkshopFileType_Screenshot                               = 5
	EWorkshopFileType_Game                                     = 6
	EWorkshopFileType_Software                                 = 7
	EWorkshopFileType_Concept                                  = 8
	EWorkshopFileType_WebGuide                                 = 9
	EWorkshopFileType_IntegratedGuide                          = 10
	EWorkshopFileType_Merch                                    = 11
	EWorkshopFileType_ControllerBinding                        = 12
	EWorkshopFileType_SteamworksAccessInvite                   = 13
	EWorkshopFileType_Max                                      = 14
)

func (e EWorkshopFileType) String() string {
	switch e {
	case EWorkshopFileType_First:
		return "EWorkshopFileType_First"
	case EWorkshopFileType_Microtransaction:
		return "EWorkshopFileType_Microtransaction"
	case EWorkshopFileType_Collection:
		return "EWorkshopFileType_Collection"
	case EWorkshopFileType_Art:
		return "EWorkshopFileType_Art"
	case EWorkshopFileType_Video:
		return "EWorkshopFileType_Video"
	case EWorkshopFileType_Screenshot:
		return "EWorkshopFileType_Screenshot"
	case EWorkshopFileType_Game:
		return "EWorkshopFileType_Game"
	case EWorkshopFileType_Software:
		return "EWorkshopFileType_Software"
	case EWorkshopFileType_Concept:
		return "EWorkshopFileType_Concept"
	case EWorkshopFileType_WebGuide:
		return "EWorkshopFileType_WebGuide"
	case EWorkshopFileType_IntegratedGuide:
		return "EWorkshopFileType_IntegratedGuide"
	case EWorkshopFileType_Merch:
		return "EWorkshopFileType_Merch"
	case EWorkshopFileType_ControllerBinding:
		return "EWorkshopFileType_ControllerBinding"
	case EWorkshopFileType_SteamworksAccessInvite:
		return "EWorkshopFileType_SteamworksAccessInvite"
	case EWorkshopFileType_Max:
		return "EWorkshopFileType_Max"
	default:
		return "INVALID"
	}
}

type EWorkshopFileAction int32

const (
	EWorkshopFileAction_Played    EWorkshopFileAction = 0
	EWorkshopFileAction_Completed                     = 1
)

func (e EWorkshopFileAction) String() string {
	switch e {
	case EWorkshopFileAction_Played:
		return "EWorkshopFileAction_Played"
	case EWorkshopFileAction_Completed:
		return "EWorkshopFileAction_Completed"
	default:
		return "INVALID"
	}
}

type EEconTradeResponse int32

const (
	EEconTradeResponse_Accepted                        EEconTradeResponse = 0
	EEconTradeResponse_Declined                                           = 1
	EEconTradeResponse_VacBannedInitiator                                 = 2 // Deprecated: renamed to TradeBannedInitiator
	EEconTradeResponse_TradeBannedInitiator                               = 2
	EEconTradeResponse_VacBannedTarget                                    = 3 // Deprecated: renamed to TradeBannedTarget
	EEconTradeResponse_TradeBannedTarget                                  = 3
	EEconTradeResponse_TargetAlreadyTrading                               = 4
	EEconTradeResponse_Disabled                                           = 5
	EEconTradeResponse_NotLoggedIn                                        = 6
	EEconTradeResponse_Cancel                                             = 7
	EEconTradeResponse_TooSoon                                            = 8
	EEconTradeResponse_TooSoonPenalty                                     = 9
	EEconTradeResponse_ConnectionFailed                                   = 10
	EEconTradeResponse_InitiatorAlreadyTrading                            = 11 // Deprecated: renamed to AlreadyTrading
	EEconTradeResponse_AlreadyTrading                                     = 11
	EEconTradeResponse_Error                                              = 12 // Deprecated: renamed to AlreadyHasTradeRequest
	EEconTradeResponse_AlreadyHasTradeRequest                             = 12
	EEconTradeResponse_Timeout                                            = 13 // Deprecated: renamed to NoResponse
	EEconTradeResponse_NoResponse                                         = 13
	EEconTradeResponse_CyberCafeInitiator                                 = 14
	EEconTradeResponse_CyberCafeTarget                                    = 15
	EEconTradeResponse_SchoolLabInitiator                                 = 16
	EEconTradeResponse_SchoolLabTarget                                    = 16
	EEconTradeResponse_InitiatorBlockedTarget                             = 18
	EEconTradeResponse_InitiatorNeedsVerifiedEmail                        = 20
	EEconTradeResponse_InitiatorNeedsSteamGuard                           = 21
	EEconTradeResponse_TargetAccountCannotTrade                           = 22
	EEconTradeResponse_InitiatorSteamGuardDuration                        = 23
	EEconTradeResponse_InitiatorPasswordResetProbation                    = 24
	EEconTradeResponse_InitiatorNewDeviceCooldown                         = 25
	EEconTradeResponse_OKToDeliver                                        = 50
)

func (e EEconTradeResponse) String() string {
	switch e {
	case EEconTradeResponse_Accepted:
		return "EEconTradeResponse_Accepted"
	case EEconTradeResponse_Declined:
		return "EEconTradeResponse_Declined"
	case EEconTradeResponse_VacBannedInitiator:
		return "EEconTradeResponse_VacBannedInitiator"
	case EEconTradeResponse_VacBannedTarget:
		return "EEconTradeResponse_VacBannedTarget"
	case EEconTradeResponse_TargetAlreadyTrading:
		return "EEconTradeResponse_TargetAlreadyTrading"
	case EEconTradeResponse_Disabled:
		return "EEconTradeResponse_Disabled"
	case EEconTradeResponse_NotLoggedIn:
		return "EEconTradeResponse_NotLoggedIn"
	case EEconTradeResponse_Cancel:
		return "EEconTradeResponse_Cancel"
	case EEconTradeResponse_TooSoon:
		return "EEconTradeResponse_TooSoon"
	case EEconTradeResponse_TooSoonPenalty:
		return "EEconTradeResponse_TooSoonPenalty"
	case EEconTradeResponse_ConnectionFailed:
		return "EEconTradeResponse_ConnectionFailed"
	case EEconTradeResponse_InitiatorAlreadyTrading:
		return "EEconTradeResponse_InitiatorAlreadyTrading"
	case EEconTradeResponse_Error:
		return "EEconTradeResponse_Error"
	case EEconTradeResponse_Timeout:
		return "EEconTradeResponse_Timeout"
	case EEconTradeResponse_CyberCafeInitiator:
		return "EEconTradeResponse_CyberCafeInitiator"
	case EEconTradeResponse_CyberCafeTarget:
		return "EEconTradeResponse_CyberCafeTarget"
	case EEconTradeResponse_SchoolLabInitiator:
		return "EEconTradeResponse_SchoolLabInitiator"
	case EEconTradeResponse_InitiatorBlockedTarget:
		return "EEconTradeResponse_InitiatorBlockedTarget"
	case EEconTradeResponse_InitiatorNeedsVerifiedEmail:
		return "EEconTradeResponse_InitiatorNeedsVerifiedEmail"
	case EEconTradeResponse_InitiatorNeedsSteamGuard:
		return "EEconTradeResponse_InitiatorNeedsSteamGuard"
	case EEconTradeResponse_TargetAccountCannotTrade:
		return "EEconTradeResponse_TargetAccountCannotTrade"
	case EEconTradeResponse_InitiatorSteamGuardDuration:
		return "EEconTradeResponse_InitiatorSteamGuardDuration"
	case EEconTradeResponse_InitiatorPasswordResetProbation:
		return "EEconTradeResponse_InitiatorPasswordResetProbation"
	case EEconTradeResponse_InitiatorNewDeviceCooldown:
		return "EEconTradeResponse_InitiatorNewDeviceCooldown"
	case EEconTradeResponse_OKToDeliver:
		return "EEconTradeResponse_OKToDeliver"
	default:
		return "INVALID"
	}
}

type EMarketingMessageFlags int32

const (
	EMarketingMessageFlags_None                 EMarketingMessageFlags = 0
	EMarketingMessageFlags_HighPriority                                = 1
	EMarketingMessageFlags_PlatformWindows                             = 2
	EMarketingMessageFlags_PlatformMac                                 = 4
	EMarketingMessageFlags_PlatformLinux                               = 8
	EMarketingMessageFlags_PlatformRestrictions                        = EMarketingMessageFlags_PlatformWindows | EMarketingMessageFlags_PlatformMac | EMarketingMessageFlags_PlatformLinux
)

func (e EMarketingMessageFlags) String() string {
	switch e {
	case EMarketingMessageFlags_None:
		return "EMarketingMessageFlags_None"
	case EMarketingMessageFlags_HighPriority:
		return "EMarketingMessageFlags_HighPriority"
	case EMarketingMessageFlags_PlatformWindows:
		return "EMarketingMessageFlags_PlatformWindows"
	case EMarketingMessageFlags_PlatformMac:
		return "EMarketingMessageFlags_PlatformMac"
	case EMarketingMessageFlags_PlatformLinux:
		return "EMarketingMessageFlags_PlatformLinux"
	case EMarketingMessageFlags_PlatformRestrictions:
		return "EMarketingMessageFlags_PlatformRestrictions"
	default:
		return "INVALID"
	}
}

type ENewsUpdateType int32

const (
	ENewsUpdateType_AppNews      ENewsUpdateType = 0
	ENewsUpdateType_SteamAds                     = 1
	ENewsUpdateType_SteamNews                    = 2
	ENewsUpdateType_CDDBUpdate                   = 3
	ENewsUpdateType_ClientUpdate                 = 4
)

func (e ENewsUpdateType) String() string {
	switch e {
	case ENewsUpdateType_AppNews:
		return "ENewsUpdateType_AppNews"
	case ENewsUpdateType_SteamAds:
		return "ENewsUpdateType_SteamAds"
	case ENewsUpdateType_SteamNews:
		return "ENewsUpdateType_SteamNews"
	case ENewsUpdateType_CDDBUpdate:
		return "ENewsUpdateType_CDDBUpdate"
	case ENewsUpdateType_ClientUpdate:
		return "ENewsUpdateType_ClientUpdate"
	default:
		return "INVALID"
	}
}

type ESystemIMType int32

const (
	ESystemIMType_RawText                  ESystemIMType = 0
	ESystemIMType_InvalidCard                            = 1
	ESystemIMType_RecurringPurchaseFailed                = 2
	ESystemIMType_CardWillExpire                         = 3
	ESystemIMType_SubscriptionExpired                    = 4
	ESystemIMType_GuestPassReceived                      = 5
	ESystemIMType_GuestPassGranted                       = 6
	ESystemIMType_GiftRevoked                            = 7
	ESystemIMType_SupportMessage                         = 8
	ESystemIMType_SupportMessageClearAlert               = 9
	ESystemIMType_Max                                    = 10
)

func (e ESystemIMType) String() string {
	switch e {
	case ESystemIMType_RawText:
		return "ESystemIMType_RawText"
	case ESystemIMType_InvalidCard:
		return "ESystemIMType_InvalidCard"
	case ESystemIMType_RecurringPurchaseFailed:
		return "ESystemIMType_RecurringPurchaseFailed"
	case ESystemIMType_CardWillExpire:
		return "ESystemIMType_CardWillExpire"
	case ESystemIMType_SubscriptionExpired:
		return "ESystemIMType_SubscriptionExpired"
	case ESystemIMType_GuestPassReceived:
		return "ESystemIMType_GuestPassReceived"
	case ESystemIMType_GuestPassGranted:
		return "ESystemIMType_GuestPassGranted"
	case ESystemIMType_GiftRevoked:
		return "ESystemIMType_GiftRevoked"
	case ESystemIMType_SupportMessage:
		return "ESystemIMType_SupportMessage"
	case ESystemIMType_SupportMessageClearAlert:
		return "ESystemIMType_SupportMessageClearAlert"
	case ESystemIMType_Max:
		return "ESystemIMType_Max"
	default:
		return "INVALID"
	}
}

type EChatFlags int32

const (
	EChatFlags_Locked             EChatFlags = 1
	EChatFlags_InvisibleToFriends            = 2
	EChatFlags_Moderated                     = 4
	EChatFlags_Unjoinable                    = 8
)

func (e EChatFlags) String() string {
	switch e {
	case EChatFlags_Locked:
		return "EChatFlags_Locked"
	case EChatFlags_InvisibleToFriends:
		return "EChatFlags_InvisibleToFriends"
	case EChatFlags_Moderated:
		return "EChatFlags_Moderated"
	case EChatFlags_Unjoinable:
		return "EChatFlags_Unjoinable"
	default:
		return "INVALID"
	}
}

type ERemoteStoragePlatform int32

const (
	ERemoteStoragePlatform_None      ERemoteStoragePlatform = 0
	ERemoteStoragePlatform_Windows                          = 1
	ERemoteStoragePlatform_OSX                              = 2
	ERemoteStoragePlatform_PS3                              = 4
	ERemoteStoragePlatform_Reserved1                        = 8
	ERemoteStoragePlatform_Reserved2                        = 16
	ERemoteStoragePlatform_All                              = -1
)

func (e ERemoteStoragePlatform) String() string {
	switch e {
	case ERemoteStoragePlatform_None:
		return "ERemoteStoragePlatform_None"
	case ERemoteStoragePlatform_Windows:
		return "ERemoteStoragePlatform_Windows"
	case ERemoteStoragePlatform_OSX:
		return "ERemoteStoragePlatform_OSX"
	case ERemoteStoragePlatform_PS3:
		return "ERemoteStoragePlatform_PS3"
	case ERemoteStoragePlatform_Reserved1:
		return "ERemoteStoragePlatform_Reserved1"
	case ERemoteStoragePlatform_Reserved2:
		return "ERemoteStoragePlatform_Reserved2"
	case ERemoteStoragePlatform_All:
		return "ERemoteStoragePlatform_All"
	default:
		return "INVALID"
	}
}

type EDRMBlobDownloadType int32

const (
	EDRMBlobDownloadType_Error        EDRMBlobDownloadType = 0
	EDRMBlobDownloadType_File                              = 1
	EDRMBlobDownloadType_Parts                             = 2
	EDRMBlobDownloadType_Compressed                        = 4
	EDRMBlobDownloadType_AllMask                           = 7
	EDRMBlobDownloadType_IsJob                             = 8
	EDRMBlobDownloadType_HighPriority                      = 16
	EDRMBlobDownloadType_AddTimestamp                      = 32
	EDRMBlobDownloadType_LowPriority                       = 64
)

func (e EDRMBlobDownloadType) String() string {
	switch e {
	case EDRMBlobDownloadType_Error:
		return "EDRMBlobDownloadType_Error"
	case EDRMBlobDownloadType_File:
		return "EDRMBlobDownloadType_File"
	case EDRMBlobDownloadType_Parts:
		return "EDRMBlobDownloadType_Parts"
	case EDRMBlobDownloadType_Compressed:
		return "EDRMBlobDownloadType_Compressed"
	case EDRMBlobDownloadType_AllMask:
		return "EDRMBlobDownloadType_AllMask"
	case EDRMBlobDownloadType_IsJob:
		return "EDRMBlobDownloadType_IsJob"
	case EDRMBlobDownloadType_HighPriority:
		return "EDRMBlobDownloadType_HighPriority"
	case EDRMBlobDownloadType_AddTimestamp:
		return "EDRMBlobDownloadType_AddTimestamp"
	case EDRMBlobDownloadType_LowPriority:
		return "EDRMBlobDownloadType_LowPriority"
	default:
		return "INVALID"
	}
}

type EDRMBlobDownloadErrorDetail int32

const (
	EDRMBlobDownloadErrorDetail_None                      EDRMBlobDownloadErrorDetail = 0
	EDRMBlobDownloadErrorDetail_DownloadFailed                                        = 1
	EDRMBlobDownloadErrorDetail_TargetLocked                                          = 2
	EDRMBlobDownloadErrorDetail_OpenZip                                               = 3
	EDRMBlobDownloadErrorDetail_ReadZipDirectory                                      = 4
	EDRMBlobDownloadErrorDetail_UnexpectedZipEntry                                    = 5
	EDRMBlobDownloadErrorDetail_UnzipFullFile                                         = 6
	EDRMBlobDownloadErrorDetail_UnknownBlobType                                       = 7
	EDRMBlobDownloadErrorDetail_UnzipStrips                                           = 8
	EDRMBlobDownloadErrorDetail_UnzipMergeGuid                                        = 9
	EDRMBlobDownloadErrorDetail_UnzipSignature                                        = 10
	EDRMBlobDownloadErrorDetail_ApplyStrips                                           = 11
	EDRMBlobDownloadErrorDetail_ApplyMergeGuid                                        = 12
	EDRMBlobDownloadErrorDetail_ApplySignature                                        = 13
	EDRMBlobDownloadErrorDetail_AppIdMismatch                                         = 14
	EDRMBlobDownloadErrorDetail_AppIdUnexpected                                       = 15
	EDRMBlobDownloadErrorDetail_AppliedSignatureCorrupt                               = 16
	EDRMBlobDownloadErrorDetail_ApplyValveSignatureHeader                             = 17
	EDRMBlobDownloadErrorDetail_UnzipValveSignatureHeader                             = 18
	EDRMBlobDownloadErrorDetail_PathManipulationError                                 = 19
	EDRMBlobDownloadErrorDetail_TargetLocked_Base                                     = 65536
	EDRMBlobDownloadErrorDetail_TargetLocked_Max                                      = 131071
	EDRMBlobDownloadErrorDetail_NextBase                                              = 131072
)

func (e EDRMBlobDownloadErrorDetail) String() string {
	switch e {
	case EDRMBlobDownloadErrorDetail_None:
		return "EDRMBlobDownloadErrorDetail_None"
	case EDRMBlobDownloadErrorDetail_DownloadFailed:
		return "EDRMBlobDownloadErrorDetail_DownloadFailed"
	case EDRMBlobDownloadErrorDetail_TargetLocked:
		return "EDRMBlobDownloadErrorDetail_TargetLocked"
	case EDRMBlobDownloadErrorDetail_OpenZip:
		return "EDRMBlobDownloadErrorDetail_OpenZip"
	case EDRMBlobDownloadErrorDetail_ReadZipDirectory:
		return "EDRMBlobDownloadErrorDetail_ReadZipDirectory"
	case EDRMBlobDownloadErrorDetail_UnexpectedZipEntry:
		return "EDRMBlobDownloadErrorDetail_UnexpectedZipEntry"
	case EDRMBlobDownloadErrorDetail_UnzipFullFile:
		return "EDRMBlobDownloadErrorDetail_UnzipFullFile"
	case EDRMBlobDownloadErrorDetail_UnknownBlobType:
		return "EDRMBlobDownloadErrorDetail_UnknownBlobType"
	case EDRMBlobDownloadErrorDetail_UnzipStrips:
		return "EDRMBlobDownloadErrorDetail_UnzipStrips"
	case EDRMBlobDownloadErrorDetail_UnzipMergeGuid:
		return "EDRMBlobDownloadErrorDetail_UnzipMergeGuid"
	case EDRMBlobDownloadErrorDetail_UnzipSignature:
		return "EDRMBlobDownloadErrorDetail_UnzipSignature"
	case EDRMBlobDownloadErrorDetail_ApplyStrips:
		return "EDRMBlobDownloadErrorDetail_ApplyStrips"
	case EDRMBlobDownloadErrorDetail_ApplyMergeGuid:
		return "EDRMBlobDownloadErrorDetail_ApplyMergeGuid"
	case EDRMBlobDownloadErrorDetail_ApplySignature:
		return "EDRMBlobDownloadErrorDetail_ApplySignature"
	case EDRMBlobDownloadErrorDetail_AppIdMismatch:
		return "EDRMBlobDownloadErrorDetail_AppIdMismatch"
	case EDRMBlobDownloadErrorDetail_AppIdUnexpected:
		return "EDRMBlobDownloadErrorDetail_AppIdUnexpected"
	case EDRMBlobDownloadErrorDetail_AppliedSignatureCorrupt:
		return "EDRMBlobDownloadErrorDetail_AppliedSignatureCorrupt"
	case EDRMBlobDownloadErrorDetail_ApplyValveSignatureHeader:
		return "EDRMBlobDownloadErrorDetail_ApplyValveSignatureHeader"
	case EDRMBlobDownloadErrorDetail_UnzipValveSignatureHeader:
		return "EDRMBlobDownloadErrorDetail_UnzipValveSignatureHeader"
	case EDRMBlobDownloadErrorDetail_PathManipulationError:
		return "EDRMBlobDownloadErrorDetail_PathManipulationError"
	case EDRMBlobDownloadErrorDetail_TargetLocked_Base:
		return "EDRMBlobDownloadErrorDetail_TargetLocked_Base"
	case EDRMBlobDownloadErrorDetail_TargetLocked_Max:
		return "EDRMBlobDownloadErrorDetail_TargetLocked_Max"
	case EDRMBlobDownloadErrorDetail_NextBase:
		return "EDRMBlobDownloadErrorDetail_NextBase"
	default:
		return "INVALID"
	}
}

type EClientStat int32

const (
	EClientStat_P2PConnectionsUDP   EClientStat = 0
	EClientStat_P2PConnectionsRelay             = 1
	EClientStat_P2PGameConnections              = 2
	EClientStat_P2PVoiceConnections             = 3
	EClientStat_BytesDownloaded                 = 4
	EClientStat_Max                             = 5
)

func (e EClientStat) String() string {
	switch e {
	case EClientStat_P2PConnectionsUDP:
		return "EClientStat_P2PConnectionsUDP"
	case EClientStat_P2PConnectionsRelay:
		return "EClientStat_P2PConnectionsRelay"
	case EClientStat_P2PGameConnections:
		return "EClientStat_P2PGameConnections"
	case EClientStat_P2PVoiceConnections:
		return "EClientStat_P2PVoiceConnections"
	case EClientStat_BytesDownloaded:
		return "EClientStat_BytesDownloaded"
	case EClientStat_Max:
		return "EClientStat_Max"
	default:
		return "INVALID"
	}
}

type EClientStatAggregateMethod int32

const (
	EClientStatAggregateMethod_LatestOnly EClientStatAggregateMethod = 0
	EClientStatAggregateMethod_Sum                                   = 1
	EClientStatAggregateMethod_Event                                 = 2
	EClientStatAggregateMethod_Scalar                                = 3
)

func (e EClientStatAggregateMethod) String() string {
	switch e {
	case EClientStatAggregateMethod_LatestOnly:
		return "EClientStatAggregateMethod_LatestOnly"
	case EClientStatAggregateMethod_Sum:
		return "EClientStatAggregateMethod_Sum"
	case EClientStatAggregateMethod_Event:
		return "EClientStatAggregateMethod_Event"
	case EClientStatAggregateMethod_Scalar:
		return "EClientStatAggregateMethod_Scalar"
	default:
		return "INVALID"
	}
}

type ELeaderboardDataRequest int32

const (
	ELeaderboardDataRequest_Global           ELeaderboardDataRequest = 0
	ELeaderboardDataRequest_GlobalAroundUser                         = 1
	ELeaderboardDataRequest_Friends                                  = 2
	ELeaderboardDataRequest_Users                                    = 3
)

func (e ELeaderboardDataRequest) String() string {
	switch e {
	case ELeaderboardDataRequest_Global:
		return "ELeaderboardDataRequest_Global"
	case ELeaderboardDataRequest_GlobalAroundUser:
		return "ELeaderboardDataRequest_GlobalAroundUser"
	case ELeaderboardDataRequest_Friends:
		return "ELeaderboardDataRequest_Friends"
	case ELeaderboardDataRequest_Users:
		return "ELeaderboardDataRequest_Users"
	default:
		return "INVALID"
	}
}

type ELeaderboardSortMethod int32

const (
	ELeaderboardSortMethod_None       ELeaderboardSortMethod = 0
	ELeaderboardSortMethod_Ascending                         = 1
	ELeaderboardSortMethod_Descending                        = 2
)

func (e ELeaderboardSortMethod) String() string {
	switch e {
	case ELeaderboardSortMethod_None:
		return "ELeaderboardSortMethod_None"
	case ELeaderboardSortMethod_Ascending:
		return "ELeaderboardSortMethod_Ascending"
	case ELeaderboardSortMethod_Descending:
		return "ELeaderboardSortMethod_Descending"
	default:
		return "INVALID"
	}
}

type ELeaderboardUploadScoreMethod int32

const (
	ELeaderboardUploadScoreMethod_None        ELeaderboardUploadScoreMethod = 0
	ELeaderboardUploadScoreMethod_KeepBest                                  = 1
	ELeaderboardUploadScoreMethod_ForceUpdate                               = 2
)

func (e ELeaderboardUploadScoreMethod) String() string {
	switch e {
	case ELeaderboardUploadScoreMethod_None:
		return "ELeaderboardUploadScoreMethod_None"
	case ELeaderboardUploadScoreMethod_KeepBest:
		return "ELeaderboardUploadScoreMethod_KeepBest"
	case ELeaderboardUploadScoreMethod_ForceUpdate:
		return "ELeaderboardUploadScoreMethod_ForceUpdate"
	default:
		return "INVALID"
	}
}

type EUCMFilePrivacyState int32

const (
	EUCMFilePrivacyState_Invalid     EUCMFilePrivacyState = -1
	EUCMFilePrivacyState_Private                          = 2
	EUCMFilePrivacyState_FriendsOnly                      = 4
	EUCMFilePrivacyState_Public                           = 8
	EUCMFilePrivacyState_All                              = EUCMFilePrivacyState_Public | EUCMFilePrivacyState_FriendsOnly | EUCMFilePrivacyState_Private
)

func (e EUCMFilePrivacyState) String() string {
	switch e {
	case EUCMFilePrivacyState_Invalid:
		return "EUCMFilePrivacyState_Invalid"
	case EUCMFilePrivacyState_Private:
		return "EUCMFilePrivacyState_Private"
	case EUCMFilePrivacyState_FriendsOnly:
		return "EUCMFilePrivacyState_FriendsOnly"
	case EUCMFilePrivacyState_Public:
		return "EUCMFilePrivacyState_Public"
	case EUCMFilePrivacyState_All:
		return "EUCMFilePrivacyState_All"
	default:
		return "INVALID"
	}
}

type EUdpPacketType uint8

const (
	EUdpPacketType_Invalid      EUdpPacketType = 0
	EUdpPacketType_ChallengeReq                = 1
	EUdpPacketType_Challenge                   = 2
	EUdpPacketType_Connect                     = 3
	EUdpPacketType_Accept                      = 4
	EUdpPacketType_Disconnect                  = 5
	EUdpPacketType_Data                        = 6
	EUdpPacketType_Datagram                    = 7
	EUdpPacketType_Max                         = 8
)

func (e EUdpPacketType) String() string {
	switch e {
	case EUdpPacketType_Invalid:
		return "EUdpPacketType_Invalid"
	case EUdpPacketType_ChallengeReq:
		return "EUdpPacketType_ChallengeReq"
	case EUdpPacketType_Challenge:
		return "EUdpPacketType_Challenge"
	case EUdpPacketType_Connect:
		return "EUdpPacketType_Connect"
	case EUdpPacketType_Accept:
		return "EUdpPacketType_Accept"
	case EUdpPacketType_Disconnect:
		return "EUdpPacketType_Disconnect"
	case EUdpPacketType_Data:
		return "EUdpPacketType_Data"
	case EUdpPacketType_Datagram:
		return "EUdpPacketType_Datagram"
	case EUdpPacketType_Max:
		return "EUdpPacketType_Max"
	default:
		return "INVALID"
	}
}
