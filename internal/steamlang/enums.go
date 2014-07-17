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

var EMsg_name = map[EMsg]string{
	0:    "EMsg_Invalid",
	1:    "EMsg_Multi",
	100:  "EMsg_BaseGeneral",
	113:  "EMsg_DestJobFailed",
	115:  "EMsg_Alert",
	120:  "EMsg_SCIDRequest",
	121:  "EMsg_SCIDResponse",
	123:  "EMsg_JobHeartbeat",
	124:  "EMsg_HubConnect",
	126:  "EMsg_Subscribe",
	127:  "EMsg_RouteMessage",
	128:  "EMsg_RemoteSysID",
	129:  "EMsg_AMCreateAccountResponse",
	130:  "EMsg_WGRequest",
	131:  "EMsg_WGResponse",
	132:  "EMsg_KeepAlive",
	133:  "EMsg_WebAPIJobRequest",
	134:  "EMsg_WebAPIJobResponse",
	135:  "EMsg_ClientSessionStart",
	136:  "EMsg_ClientSessionEnd",
	137:  "EMsg_ClientSessionUpdateAuthTicket",
	138:  "EMsg_StatsDeprecated",
	139:  "EMsg_Ping",
	140:  "EMsg_PingResponse",
	141:  "EMsg_Stats",
	142:  "EMsg_RequestFullStatsBlock",
	143:  "EMsg_LoadDBOCacheItem",
	144:  "EMsg_LoadDBOCacheItemResponse",
	145:  "EMsg_InvalidateDBOCacheItems",
	146:  "EMsg_ServiceMethod",
	147:  "EMsg_ServiceMethodResponse",
	200:  "EMsg_BaseShell",
	201:  "EMsg_Exit",
	202:  "EMsg_DirRequest",
	203:  "EMsg_DirResponse",
	204:  "EMsg_ZipRequest",
	205:  "EMsg_ZipResponse",
	215:  "EMsg_UpdateRecordResponse",
	221:  "EMsg_UpdateCreditCardRequest",
	225:  "EMsg_UpdateUserBanResponse",
	226:  "EMsg_PrepareToExit",
	227:  "EMsg_ContentDescriptionUpdate",
	228:  "EMsg_TestResetServer",
	229:  "EMsg_UniverseChanged",
	230:  "EMsg_ShellConfigInfoUpdate",
	233:  "EMsg_RequestWindowsEventLogEntries",
	234:  "EMsg_ProvideWindowsEventLogEntries",
	235:  "EMsg_ShellSearchLogs",
	236:  "EMsg_ShellSearchLogsResponse",
	237:  "EMsg_ShellCheckWindowsUpdates",
	238:  "EMsg_ShellCheckWindowsUpdatesResponse",
	239:  "EMsg_ShellFlushUserLicenseCache",
	300:  "EMsg_BaseGM",
	301:  "EMsg_ShellFailed",
	307:  "EMsg_ExitShells",
	308:  "EMsg_ExitShell",
	309:  "EMsg_GracefulExitShell",
	314:  "EMsg_NotifyWatchdog",
	316:  "EMsg_LicenseProcessingComplete",
	317:  "EMsg_SetTestFlag",
	318:  "EMsg_QueuedEmailsComplete",
	319:  "EMsg_GMReportPHPError",
	320:  "EMsg_GMDRMSync",
	321:  "EMsg_PhysicalBoxInventory",
	322:  "EMsg_UpdateConfigFile",
	323:  "EMsg_TestInitDB",
	324:  "EMsg_GMWriteConfigToSQL",
	325:  "EMsg_GMLoadActivationCodes",
	326:  "EMsg_GMQueueForFBS",
	327:  "EMsg_GMSchemaConversionResults",
	328:  "EMsg_GMSchemaConversionResultsResponse",
	329:  "EMsg_GMWriteShellFailureToSQL",
	400:  "EMsg_BaseAIS",
	401:  "EMsg_AISRefreshContentDescription",
	402:  "EMsg_AISRequestContentDescription",
	403:  "EMsg_AISUpdateAppInfo",
	404:  "EMsg_AISUpdatePackageInfo",
	405:  "EMsg_AISGetPackageChangeNumber",
	406:  "EMsg_AISGetPackageChangeNumberResponse",
	407:  "EMsg_AISAppInfoTableChanged",
	408:  "EMsg_AISUpdatePackageInfoResponse",
	409:  "EMsg_AISCreateMarketingMessage",
	410:  "EMsg_AISCreateMarketingMessageResponse",
	411:  "EMsg_AISGetMarketingMessage",
	412:  "EMsg_AISGetMarketingMessageResponse",
	413:  "EMsg_AISUpdateMarketingMessage",
	414:  "EMsg_AISUpdateMarketingMessageResponse",
	415:  "EMsg_AISRequestMarketingMessageUpdate",
	416:  "EMsg_AISDeleteMarketingMessage",
	419:  "EMsg_AISGetMarketingTreatments",
	420:  "EMsg_AISGetMarketingTreatmentsResponse",
	421:  "EMsg_AISRequestMarketingTreatmentUpdate",
	422:  "EMsg_AISTestAddPackage",
	423:  "EMsg_AIGetAppGCFlags",
	424:  "EMsg_AIGetAppGCFlagsResponse",
	425:  "EMsg_AIGetAppList",
	426:  "EMsg_AIGetAppListResponse",
	427:  "EMsg_AIGetAppInfo",
	428:  "EMsg_AIGetAppInfoResponse",
	429:  "EMsg_AISGetCouponDefinition",
	430:  "EMsg_AISGetCouponDefinitionResponse",
	500:  "EMsg_BaseAM",
	504:  "EMsg_AMUpdateUserBanRequest",
	505:  "EMsg_AMAddLicense",
	507:  "EMsg_AMBeginProcessingLicenses",
	508:  "EMsg_AMSendSystemIMToUser",
	509:  "EMsg_AMExtendLicense",
	510:  "EMsg_AMAddMinutesToLicense",
	511:  "EMsg_AMCancelLicense",
	512:  "EMsg_AMInitPurchase",
	513:  "EMsg_AMPurchaseResponse",
	514:  "EMsg_AMGetFinalPrice",
	515:  "EMsg_AMGetFinalPriceResponse",
	516:  "EMsg_AMGetLegacyGameKey",
	517:  "EMsg_AMGetLegacyGameKeyResponse",
	518:  "EMsg_AMFindHungTransactions",
	519:  "EMsg_AMSetAccountTrustedRequest",
	521:  "EMsg_AMCompletePurchase",
	522:  "EMsg_AMCancelPurchase",
	523:  "EMsg_AMNewChallenge",
	526:  "EMsg_AMFixPendingPurchaseResponse",
	527:  "EMsg_AMIsUserBanned",
	528:  "EMsg_AMRegisterKey",
	529:  "EMsg_AMLoadActivationCodes",
	530:  "EMsg_AMLoadActivationCodesResponse",
	531:  "EMsg_AMLookupKeyResponse",
	532:  "EMsg_AMLookupKey",
	533:  "EMsg_AMChatCleanup",
	534:  "EMsg_AMClanCleanup",
	535:  "EMsg_AMFixPendingRefund",
	536:  "EMsg_AMReverseChargeback",
	537:  "EMsg_AMReverseChargebackResponse",
	538:  "EMsg_AMClanCleanupList",
	539:  "EMsg_AMGetLicenses",
	540:  "EMsg_AMGetLicensesResponse",
	550:  "EMsg_AllowUserToPlayQuery",
	551:  "EMsg_AllowUserToPlayResponse",
	552:  "EMsg_AMVerfiyUser",
	553:  "EMsg_AMClientNotPlaying",
	554:  "EMsg_ClientRequestFriendship",
	555:  "EMsg_AMRelayPublishStatus",
	556:  "EMsg_AMResetCommunityContent",
	557:  "EMsg_AMPrimePersonaStateCache",
	558:  "EMsg_AMAllowUserContentQuery",
	559:  "EMsg_AMAllowUserContentResponse",
	560:  "EMsg_AMInitPurchaseResponse",
	561:  "EMsg_AMRevokePurchaseResponse",
	562:  "EMsg_AMLockProfile",
	563:  "EMsg_AMRefreshGuestPasses",
	564:  "EMsg_AMInviteUserToClan",
	565:  "EMsg_AMAcknowledgeClanInvite",
	566:  "EMsg_AMGrantGuestPasses",
	567:  "EMsg_AMClanDataUpdated",
	568:  "EMsg_AMReloadAccount",
	569:  "EMsg_AMClientChatMsgRelay",
	570:  "EMsg_AMChatMulti",
	571:  "EMsg_AMClientChatInviteRelay",
	572:  "EMsg_AMChatInvite",
	573:  "EMsg_AMClientJoinChatRelay",
	574:  "EMsg_AMClientChatMemberInfoRelay",
	575:  "EMsg_AMPublishChatMemberInfo",
	576:  "EMsg_AMClientAcceptFriendInvite",
	577:  "EMsg_AMChatEnter",
	578:  "EMsg_AMClientPublishRemovalFromSource",
	579:  "EMsg_AMChatActionResult",
	580:  "EMsg_AMFindAccounts",
	581:  "EMsg_AMFindAccountsResponse",
	584:  "EMsg_AMSetAccountFlags",
	586:  "EMsg_AMCreateClan",
	587:  "EMsg_AMCreateClanResponse",
	588:  "EMsg_AMGetClanDetails",
	589:  "EMsg_AMGetClanDetailsResponse",
	590:  "EMsg_AMSetPersonaName",
	591:  "EMsg_AMSetAvatar",
	592:  "EMsg_AMAuthenticateUser",
	593:  "EMsg_AMAuthenticateUserResponse",
	594:  "EMsg_AMGetAccountFriendsCount",
	595:  "EMsg_AMGetAccountFriendsCountResponse",
	596:  "EMsg_AMP2PIntroducerMessage",
	597:  "EMsg_ClientChatAction",
	598:  "EMsg_AMClientChatActionRelay",
	600:  "EMsg_BaseVS",
	601:  "EMsg_VACResponse",
	602:  "EMsg_ReqChallengeTest",
	604:  "EMsg_VSMarkCheat",
	605:  "EMsg_VSAddCheat",
	606:  "EMsg_VSPurgeCodeModDB",
	607:  "EMsg_VSGetChallengeResults",
	608:  "EMsg_VSChallengeResultText",
	609:  "EMsg_VSReportLingerer",
	610:  "EMsg_VSRequestManagedChallenge",
	611:  "EMsg_VSLoadDBFinished",
	625:  "EMsg_BaseDRMS",
	628:  "EMsg_DRMBuildBlobRequest",
	629:  "EMsg_DRMBuildBlobResponse",
	630:  "EMsg_DRMResolveGuidRequest",
	631:  "EMsg_DRMResolveGuidResponse",
	633:  "EMsg_DRMVariabilityReport",
	634:  "EMsg_DRMVariabilityReportResponse",
	635:  "EMsg_DRMStabilityReport",
	636:  "EMsg_DRMStabilityReportResponse",
	637:  "EMsg_DRMDetailsReportRequest",
	638:  "EMsg_DRMDetailsReportResponse",
	639:  "EMsg_DRMProcessFile",
	640:  "EMsg_DRMAdminUpdate",
	641:  "EMsg_DRMAdminUpdateResponse",
	642:  "EMsg_DRMSync",
	643:  "EMsg_DRMSyncResponse",
	644:  "EMsg_DRMProcessFileResponse",
	645:  "EMsg_DRMEmptyGuidCache",
	646:  "EMsg_DRMEmptyGuidCacheResponse",
	650:  "EMsg_BaseCS",
	652:  "EMsg_CSUserContentRequest",
	700:  "EMsg_BaseClient",
	701:  "EMsg_ClientLogOn_Deprecated",
	702:  "EMsg_ClientAnonLogOn_Deprecated",
	703:  "EMsg_ClientHeartBeat",
	704:  "EMsg_ClientVACResponse",
	705:  "EMsg_ClientGamesPlayed_obsolete",
	706:  "EMsg_ClientLogOff",
	707:  "EMsg_ClientNoUDPConnectivity",
	708:  "EMsg_ClientInformOfCreateAccount",
	709:  "EMsg_ClientAckVACBan",
	710:  "EMsg_ClientConnectionStats",
	711:  "EMsg_ClientInitPurchase",
	712:  "EMsg_ClientPingResponse",
	714:  "EMsg_ClientRemoveFriend",
	715:  "EMsg_ClientGamesPlayedNoDataBlob",
	716:  "EMsg_ClientChangeStatus",
	717:  "EMsg_ClientVacStatusResponse",
	718:  "EMsg_ClientFriendMsg",
	719:  "EMsg_ClientGameConnect_obsolete",
	720:  "EMsg_ClientGamesPlayed2_obsolete",
	721:  "EMsg_ClientGameEnded_obsolete",
	722:  "EMsg_ClientGetFinalPrice",
	726:  "EMsg_ClientSystemIM",
	727:  "EMsg_ClientSystemIMAck",
	728:  "EMsg_ClientGetLicenses",
	729:  "EMsg_ClientCancelLicense",
	730:  "EMsg_ClientGetLegacyGameKey",
	731:  "EMsg_ClientContentServerLogOn_Deprecated",
	732:  "EMsg_ClientAckVACBan2",
	735:  "EMsg_ClientAckMessageByGID",
	736:  "EMsg_ClientGetPurchaseReceipts",
	737:  "EMsg_ClientAckPurchaseReceipt",
	738:  "EMsg_ClientGamesPlayed3_obsolete",
	739:  "EMsg_ClientSendGuestPass",
	740:  "EMsg_ClientAckGuestPass",
	741:  "EMsg_ClientRedeemGuestPass",
	742:  "EMsg_ClientGamesPlayed",
	743:  "EMsg_ClientRegisterKey",
	744:  "EMsg_ClientInviteUserToClan",
	745:  "EMsg_ClientAcknowledgeClanInvite",
	746:  "EMsg_ClientPurchaseWithMachineID",
	747:  "EMsg_ClientAppUsageEvent",
	748:  "EMsg_ClientGetGiftTargetList",
	749:  "EMsg_ClientGetGiftTargetListResponse",
	751:  "EMsg_ClientLogOnResponse",
	753:  "EMsg_ClientVACChallenge",
	755:  "EMsg_ClientSetHeartbeatRate",
	756:  "EMsg_ClientNotLoggedOnDeprecated",
	757:  "EMsg_ClientLoggedOff",
	758:  "EMsg_GSApprove",
	759:  "EMsg_GSDeny",
	760:  "EMsg_GSKick",
	761:  "EMsg_ClientCreateAcctResponse",
	763:  "EMsg_ClientPurchaseResponse",
	764:  "EMsg_ClientPing",
	765:  "EMsg_ClientNOP",
	766:  "EMsg_ClientPersonaState",
	767:  "EMsg_ClientFriendsList",
	768:  "EMsg_ClientAccountInfo",
	770:  "EMsg_ClientVacStatusQuery",
	771:  "EMsg_ClientNewsUpdate",
	773:  "EMsg_ClientGameConnectDeny",
	774:  "EMsg_GSStatusReply",
	775:  "EMsg_ClientGetFinalPriceResponse",
	779:  "EMsg_ClientGameConnectTokens",
	780:  "EMsg_ClientLicenseList",
	781:  "EMsg_ClientCancelLicenseResponse",
	782:  "EMsg_ClientVACBanStatus",
	783:  "EMsg_ClientCMList",
	784:  "EMsg_ClientEncryptPct",
	785:  "EMsg_ClientGetLegacyGameKeyResponse",
	786:  "EMsg_ClientFavoritesList",
	787:  "EMsg_CSUserContentApprove",
	788:  "EMsg_CSUserContentDeny",
	789:  "EMsg_ClientInitPurchaseResponse",
	791:  "EMsg_ClientAddFriend",
	792:  "EMsg_ClientAddFriendResponse",
	793:  "EMsg_ClientInviteFriend",
	794:  "EMsg_ClientInviteFriendResponse",
	795:  "EMsg_ClientSendGuestPassResponse",
	796:  "EMsg_ClientAckGuestPassResponse",
	797:  "EMsg_ClientRedeemGuestPassResponse",
	798:  "EMsg_ClientUpdateGuestPassesList",
	799:  "EMsg_ClientChatMsg",
	800:  "EMsg_ClientChatInvite",
	801:  "EMsg_ClientJoinChat",
	802:  "EMsg_ClientChatMemberInfo",
	803:  "EMsg_ClientLogOnWithCredentials_Deprecated",
	805:  "EMsg_ClientPasswordChangeResponse",
	807:  "EMsg_ClientChatEnter",
	808:  "EMsg_ClientFriendRemovedFromSource",
	809:  "EMsg_ClientCreateChat",
	810:  "EMsg_ClientCreateChatResponse",
	811:  "EMsg_ClientUpdateChatMetadata",
	813:  "EMsg_ClientP2PIntroducerMessage",
	814:  "EMsg_ClientChatActionResult",
	815:  "EMsg_ClientRequestFriendData",
	818:  "EMsg_ClientGetUserStats",
	819:  "EMsg_ClientGetUserStatsResponse",
	820:  "EMsg_ClientStoreUserStats",
	821:  "EMsg_ClientStoreUserStatsResponse",
	822:  "EMsg_ClientClanState",
	830:  "EMsg_ClientServiceModule",
	831:  "EMsg_ClientServiceCall",
	832:  "EMsg_ClientServiceCallResponse",
	833:  "EMsg_ClientPackageInfoRequest",
	834:  "EMsg_ClientPackageInfoResponse",
	839:  "EMsg_ClientNatTraversalStatEvent",
	840:  "EMsg_ClientAppInfoRequest",
	841:  "EMsg_ClientAppInfoResponse",
	842:  "EMsg_ClientSteamUsageEvent",
	845:  "EMsg_ClientCheckPassword",
	846:  "EMsg_ClientResetPassword",
	848:  "EMsg_ClientCheckPasswordResponse",
	849:  "EMsg_ClientResetPasswordResponse",
	850:  "EMsg_ClientSessionToken",
	851:  "EMsg_ClientDRMProblemReport",
	855:  "EMsg_ClientSetIgnoreFriend",
	856:  "EMsg_ClientSetIgnoreFriendResponse",
	857:  "EMsg_ClientGetAppOwnershipTicket",
	858:  "EMsg_ClientGetAppOwnershipTicketResponse",
	860:  "EMsg_ClientGetLobbyListResponse",
	861:  "EMsg_ClientGetLobbyMetadata",
	862:  "EMsg_ClientGetLobbyMetadataResponse",
	863:  "EMsg_ClientVTTCert",
	866:  "EMsg_ClientAppInfoUpdate",
	867:  "EMsg_ClientAppInfoChanges",
	880:  "EMsg_ClientServerList",
	891:  "EMsg_ClientEmailChangeResponse",
	892:  "EMsg_ClientSecretQAChangeResponse",
	896:  "EMsg_ClientDRMBlobRequest",
	897:  "EMsg_ClientDRMBlobResponse",
	898:  "EMsg_ClientLookupKey",
	899:  "EMsg_ClientLookupKeyResponse",
	900:  "EMsg_BaseGameServer",
	901:  "EMsg_GSDisconnectNotice",
	903:  "EMsg_GSStatus",
	905:  "EMsg_GSUserPlaying",
	906:  "EMsg_GSStatus2",
	907:  "EMsg_GSStatusUpdate_Unused",
	908:  "EMsg_GSServerType",
	909:  "EMsg_GSPlayerList",
	910:  "EMsg_GSGetUserAchievementStatus",
	911:  "EMsg_GSGetUserAchievementStatusResponse",
	918:  "EMsg_GSGetPlayStats",
	919:  "EMsg_GSGetPlayStatsResponse",
	920:  "EMsg_GSGetUserGroupStatus",
	921:  "EMsg_AMGetUserGroupStatus",
	922:  "EMsg_AMGetUserGroupStatusResponse",
	923:  "EMsg_GSGetUserGroupStatusResponse",
	936:  "EMsg_GSGetReputation",
	937:  "EMsg_GSGetReputationResponse",
	938:  "EMsg_GSAssociateWithClan",
	939:  "EMsg_GSAssociateWithClanResponse",
	940:  "EMsg_GSComputeNewPlayerCompatibility",
	941:  "EMsg_GSComputeNewPlayerCompatibilityResponse",
	1000: "EMsg_BaseAdmin",
	1004: "EMsg_AdminCmdResponse",
	1005: "EMsg_AdminLogListenRequest",
	1006: "EMsg_AdminLogEvent",
	1007: "EMsg_LogSearchRequest",
	1008: "EMsg_LogSearchResponse",
	1009: "EMsg_LogSearchCancel",
	1010: "EMsg_UniverseData",
	1014: "EMsg_RequestStatHistory",
	1015: "EMsg_StatHistory",
	1017: "EMsg_AdminPwLogon",
	1018: "EMsg_AdminPwLogonResponse",
	1019: "EMsg_AdminSpew",
	1020: "EMsg_AdminConsoleTitle",
	1023: "EMsg_AdminGCSpew",
	1024: "EMsg_AdminGCCommand",
	1025: "EMsg_AdminGCGetCommandList",
	1026: "EMsg_AdminGCGetCommandListResponse",
	1027: "EMsg_FBSConnectionData",
	1028: "EMsg_AdminMsgSpew",
	1100: "EMsg_BaseFBS",
	1101: "EMsg_FBSVersionInfo",
	1102: "EMsg_FBSForceRefresh",
	1103: "EMsg_FBSForceBounce",
	1104: "EMsg_FBSDeployPackage",
	1105: "EMsg_FBSDeployResponse",
	1106: "EMsg_FBSUpdateBootstrapper",
	1107: "EMsg_FBSSetState",
	1108: "EMsg_FBSApplyOSUpdates",
	1109: "EMsg_FBSRunCMDScript",
	1110: "EMsg_FBSRebootBox",
	1111: "EMsg_FBSSetBigBrotherMode",
	1112: "EMsg_FBSMinidumpServer",
	1113: "EMsg_FBSSetShellCount_obsolete",
	1114: "EMsg_FBSDeployHotFixPackage",
	1115: "EMsg_FBSDeployHotFixResponse",
	1116: "EMsg_FBSDownloadHotFix",
	1117: "EMsg_FBSDownloadHotFixResponse",
	1118: "EMsg_FBSUpdateTargetConfigFile",
	1119: "EMsg_FBSApplyAccountCred",
	1120: "EMsg_FBSApplyAccountCredResponse",
	1121: "EMsg_FBSSetShellCount",
	1122: "EMsg_FBSTerminateShell",
	1123: "EMsg_FBSQueryGMForRequest",
	1124: "EMsg_FBSQueryGMResponse",
	1125: "EMsg_FBSTerminateZombies",
	1126: "EMsg_FBSInfoFromBootstrapper",
	1127: "EMsg_FBSRebootBoxResponse",
	1128: "EMsg_FBSBootstrapperPackageRequest",
	1129: "EMsg_FBSBootstrapperPackageResponse",
	1130: "EMsg_FBSBootstrapperGetPackageChunk",
	1131: "EMsg_FBSBootstrapperGetPackageChunkResponse",
	1132: "EMsg_FBSBootstrapperPackageTransferProgress",
	1133: "EMsg_FBSRestartBootstrapper",
	1200: "EMsg_BaseFileXfer",
	1201: "EMsg_FileXferResponse",
	1202: "EMsg_FileXferData",
	1203: "EMsg_FileXferEnd",
	1204: "EMsg_FileXferDataAck",
	1300: "EMsg_BaseChannelAuth",
	1301: "EMsg_ChannelAuthResponse",
	1302: "EMsg_ChannelAuthResult",
	1303: "EMsg_ChannelEncryptRequest",
	1304: "EMsg_ChannelEncryptResponse",
	1305: "EMsg_ChannelEncryptResult",
	1400: "EMsg_BaseBS",
	1401: "EMsg_BSPurchaseStart",
	1402: "EMsg_BSPurchaseResponse",
	1404: "EMsg_BSSettleNOVA",
	1406: "EMsg_BSSettleComplete",
	1407: "EMsg_BSBannedRequest",
	1408: "EMsg_BSInitPayPalTxn",
	1409: "EMsg_BSInitPayPalTxnResponse",
	1410: "EMsg_BSGetPayPalUserInfo",
	1411: "EMsg_BSGetPayPalUserInfoResponse",
	1413: "EMsg_BSRefundTxn",
	1414: "EMsg_BSRefundTxnResponse",
	1415: "EMsg_BSGetEvents",
	1416: "EMsg_BSChaseRFRRequest",
	1417: "EMsg_BSPaymentInstrBan",
	1418: "EMsg_BSPaymentInstrBanResponse",
	1419: "EMsg_BSProcessGCReports",
	1420: "EMsg_BSProcessPPReports",
	1421: "EMsg_BSInitGCBankXferTxn",
	1422: "EMsg_BSInitGCBankXferTxnResponse",
	1423: "EMsg_BSQueryGCBankXferTxn",
	1424: "EMsg_BSQueryGCBankXferTxnResponse",
	1425: "EMsg_BSCommitGCTxn",
	1426: "EMsg_BSQueryTransactionStatus",
	1427: "EMsg_BSQueryTransactionStatusResponse",
	1428: "EMsg_BSQueryCBOrderStatus",
	1429: "EMsg_BSQueryCBOrderStatusResponse",
	1430: "EMsg_BSRunRedFlagReport",
	1431: "EMsg_BSQueryPaymentInstUsage",
	1432: "EMsg_BSQueryPaymentInstResponse",
	1433: "EMsg_BSQueryTxnExtendedInfo",
	1434: "EMsg_BSQueryTxnExtendedInfoResponse",
	1435: "EMsg_BSUpdateConversionRates",
	1436: "EMsg_BSProcessUSBankReports",
	1437: "EMsg_BSPurchaseRunFraudChecks",
	1438: "EMsg_BSPurchaseRunFraudChecksResponse",
	1439: "EMsg_BSStartShippingJobs",
	1440: "EMsg_BSQueryBankInformation",
	1441: "EMsg_BSQueryBankInformationResponse",
	1445: "EMsg_BSValidateXsollaSignature",
	1446: "EMsg_BSValidateXsollaSignatureResponse",
	1448: "EMsg_BSQiwiWalletInvoice",
	1449: "EMsg_BSQiwiWalletInvoiceResponse",
	1450: "EMsg_BSUpdateInventoryFromProPack",
	1451: "EMsg_BSUpdateInventoryFromProPackResponse",
	1452: "EMsg_BSSendShippingRequest",
	1453: "EMsg_BSSendShippingRequestResponse",
	1454: "EMsg_BSGetProPackOrderStatus",
	1455: "EMsg_BSGetProPackOrderStatusResponse",
	1456: "EMsg_BSCheckJobRunning",
	1457: "EMsg_BSCheckJobRunningResponse",
	1458: "EMsg_BSResetPackagePurchaseRateLimit",
	1459: "EMsg_BSResetPackagePurchaseRateLimitResponse",
	1460: "EMsg_BSUpdatePaymentData",
	1461: "EMsg_BSUpdatePaymentDataResponse",
	1462: "EMsg_BSGetBillingAddress",
	1463: "EMsg_BSGetBillingAddressResponse",
	1464: "EMsg_BSGetCreditCardInfo",
	1465: "EMsg_BSGetCreditCardInfoResponse",
	1468: "EMsg_BSRemoveExpiredPaymentData",
	1469: "EMsg_BSRemoveExpiredPaymentDataResponse",
	1470: "EMsg_BSConvertToCurrentKeys",
	1471: "EMsg_BSConvertToCurrentKeysResponse",
	1472: "EMsg_BSInitPurchase",
	1473: "EMsg_BSInitPurchaseResponse",
	1474: "EMsg_BSCompletePurchase",
	1475: "EMsg_BSCompletePurchaseResponse",
	1476: "EMsg_BSPruneCardUsageStats",
	1477: "EMsg_BSPruneCardUsageStatsResponse",
	1478: "EMsg_BSStoreBankInformation",
	1479: "EMsg_BSStoreBankInformationResponse",
	1480: "EMsg_BSVerifyPOSAKey",
	1481: "EMsg_BSVerifyPOSAKeyResponse",
	1482: "EMsg_BSReverseRedeemPOSAKey",
	1483: "EMsg_BSReverseRedeemPOSAKeyResponse",
	1484: "EMsg_BSQueryFindCreditCard",
	1485: "EMsg_BSQueryFindCreditCardResponse",
	1486: "EMsg_BSStatusInquiryPOSAKey",
	1487: "EMsg_BSStatusInquiryPOSAKeyResponse",
	1488: "EMsg_BSValidateMoPaySignature",
	1489: "EMsg_BSValidateMoPaySignatureResponse",
	1490: "EMsg_BSMoPayConfirmProductDelivery",
	1491: "EMsg_BSMoPayConfirmProductDeliveryResponse",
	1492: "EMsg_BSGenerateMoPayMD5",
	1493: "EMsg_BSGenerateMoPayMD5Response",
	1494: "EMsg_BSBoaCompraConfirmProductDelivery",
	1495: "EMsg_BSBoaCompraConfirmProductDeliveryResponse",
	1496: "EMsg_BSGenerateBoaCompraMD5",
	1497: "EMsg_BSGenerateBoaCompraMD5Response",
	1500: "EMsg_BaseATS",
	1501: "EMsg_ATSStartStressTest",
	1502: "EMsg_ATSStopStressTest",
	1503: "EMsg_ATSRunFailServerTest",
	1504: "EMsg_ATSUFSPerfTestTask",
	1505: "EMsg_ATSUFSPerfTestResponse",
	1506: "EMsg_ATSCycleTCM",
	1507: "EMsg_ATSInitDRMSStressTest",
	1508: "EMsg_ATSCallTest",
	1509: "EMsg_ATSCallTestReply",
	1510: "EMsg_ATSStartExternalStress",
	1511: "EMsg_ATSExternalStressJobStart",
	1512: "EMsg_ATSExternalStressJobQueued",
	1513: "EMsg_ATSExternalStressJobRunning",
	1514: "EMsg_ATSExternalStressJobStopped",
	1515: "EMsg_ATSExternalStressJobStopAll",
	1516: "EMsg_ATSExternalStressActionResult",
	1517: "EMsg_ATSStarted",
	1518: "EMsg_ATSCSPerfTestTask",
	1519: "EMsg_ATSCSPerfTestResponse",
	1600: "EMsg_BaseDP",
	1601: "EMsg_DPSetPublishingState",
	1602: "EMsg_DPGamePlayedStats",
	1603: "EMsg_DPUniquePlayersStat",
	1605: "EMsg_DPVacInfractionStats",
	1606: "EMsg_DPVacBanStats",
	1607: "EMsg_DPBlockingStats",
	1608: "EMsg_DPNatTraversalStats",
	1609: "EMsg_DPSteamUsageEvent",
	1610: "EMsg_DPVacCertBanStats",
	1611: "EMsg_DPVacCafeBanStats",
	1612: "EMsg_DPCloudStats",
	1613: "EMsg_DPAchievementStats",
	1614: "EMsg_DPAccountCreationStats",
	1615: "EMsg_DPGetPlayerCount",
	1616: "EMsg_DPGetPlayerCountResponse",
	1617: "EMsg_DPGameServersPlayersStats",
	1618: "EMsg_DPDownloadRateStatistics",
	1619: "EMsg_DPFacebookStatistics",
	1620: "EMsg_ClientDPCheckSpecialSurvey",
	1621: "EMsg_ClientDPCheckSpecialSurveyResponse",
	1622: "EMsg_ClientDPSendSpecialSurveyResponse",
	1623: "EMsg_ClientDPSendSpecialSurveyResponseReply",
	1624: "EMsg_DPStoreSaleStatistics",
	1625: "EMsg_ClientDPUpdateAppJobReport",
	1627: "EMsg_ClientDPSteam2AppStarted",
	1626: "EMsg_DPUpdateContentEvent",
	1700: "EMsg_BaseCM",
	1701: "EMsg_CMSetAllowState",
	1702: "EMsg_CMSpewAllowState",
	1703: "EMsg_CMAppInfoResponseDeprecated",
	1800: "EMsg_BaseDSS",
	1801: "EMsg_DSSNewFile",
	1802: "EMsg_DSSCurrentFileList",
	1803: "EMsg_DSSSynchList",
	1804: "EMsg_DSSSynchListResponse",
	1805: "EMsg_DSSSynchSubscribe",
	1806: "EMsg_DSSSynchUnsubscribe",
	1900: "EMsg_BaseEPM",
	1901: "EMsg_EPMStartProcess",
	1902: "EMsg_EPMStopProcess",
	1903: "EMsg_EPMRestartProcess",
	2200: "EMsg_BaseGC",
	2201: "EMsg_AMRelayToGC",
	2202: "EMsg_GCUpdatePlayedState",
	2203: "EMsg_GCCmdRevive",
	2204: "EMsg_GCCmdBounce",
	2205: "EMsg_GCCmdForceBounce",
	2206: "EMsg_GCCmdDown",
	2207: "EMsg_GCCmdDeploy",
	2208: "EMsg_GCCmdDeployResponse",
	2209: "EMsg_GCCmdSwitch",
	2210: "EMsg_AMRefreshSessions",
	2211: "EMsg_GCUpdateGSState",
	2212: "EMsg_GCAchievementAwarded",
	2213: "EMsg_GCSystemMessage",
	2214: "EMsg_GCValidateSession",
	2215: "EMsg_GCValidateSessionResponse",
	2216: "EMsg_GCCmdStatus",
	2217: "EMsg_GCRegisterWebInterfaces",
	2218: "EMsg_GCGetAccountDetails",
	2219: "EMsg_GCInterAppMessage",
	2220: "EMsg_GCGetEmailTemplate",
	2221: "EMsg_GCGetEmailTemplateResponse",
	2222: "EMsg_ISRelayToGCH",
	2223: "EMsg_GCHRelayClientToIS",
	2224: "EMsg_GCHUpdateSession",
	2225: "EMsg_GCHRequestUpdateSession",
	2226: "EMsg_GCHRequestStatus",
	2227: "EMsg_GCHRequestStatusResponse",
	2500: "EMsg_BaseP2P",
	2502: "EMsg_P2PIntroducerMessage",
	2900: "EMsg_BaseSM",
	2902: "EMsg_SMExpensiveReport",
	2903: "EMsg_SMHourlyReport",
	2904: "EMsg_SMFishingReport",
	2905: "EMsg_SMPartitionRenames",
	2906: "EMsg_SMMonitorSpace",
	2907: "EMsg_SMGetSchemaConversionResults",
	2908: "EMsg_SMGetSchemaConversionResultsResponse",
	3000: "EMsg_BaseTest",
	3001: "EMsg_JobHeartbeatTest",
	3002: "EMsg_JobHeartbeatTestResponse",
	3100: "EMsg_BaseFTSRange",
	3101: "EMsg_FTSGetBrowseCounts",
	3102: "EMsg_FTSGetBrowseCountsResponse",
	3103: "EMsg_FTSBrowseClans",
	3104: "EMsg_FTSBrowseClansResponse",
	3105: "EMsg_FTSSearchClansByLocation",
	3106: "EMsg_FTSSearchClansByLocationResponse",
	3107: "EMsg_FTSSearchPlayersByLocation",
	3108: "EMsg_FTSSearchPlayersByLocationResponse",
	3109: "EMsg_FTSClanDeleted",
	3110: "EMsg_FTSSearch",
	3111: "EMsg_FTSSearchResponse",
	3112: "EMsg_FTSSearchStatus",
	3113: "EMsg_FTSSearchStatusResponse",
	3114: "EMsg_FTSGetGSPlayStats",
	3115: "EMsg_FTSGetGSPlayStatsResponse",
	3116: "EMsg_FTSGetGSPlayStatsForServer",
	3117: "EMsg_FTSGetGSPlayStatsForServerResponse",
	3118: "EMsg_FTSReportIPUpdates",
	3150: "EMsg_BaseCCSRange",
	3151: "EMsg_CCSGetComments",
	3152: "EMsg_CCSGetCommentsResponse",
	3153: "EMsg_CCSAddComment",
	3154: "EMsg_CCSAddCommentResponse",
	3155: "EMsg_CCSDeleteComment",
	3156: "EMsg_CCSDeleteCommentResponse",
	3157: "EMsg_CCSPreloadComments",
	3158: "EMsg_CCSNotifyCommentCount",
	3159: "EMsg_CCSGetCommentsForNews",
	3160: "EMsg_CCSGetCommentsForNewsResponse",
	3161: "EMsg_CCSDeleteAllCommentsByAuthor",
	3162: "EMsg_CCSDeleteAllCommentsByAuthorResponse",
	3200: "EMsg_BaseLBSRange",
	3201: "EMsg_LBSSetScore",
	3202: "EMsg_LBSSetScoreResponse",
	3203: "EMsg_LBSFindOrCreateLB",
	3204: "EMsg_LBSFindOrCreateLBResponse",
	3205: "EMsg_LBSGetLBEntries",
	3206: "EMsg_LBSGetLBEntriesResponse",
	3207: "EMsg_LBSGetLBList",
	3208: "EMsg_LBSGetLBListResponse",
	3209: "EMsg_LBSSetLBDetails",
	3210: "EMsg_LBSDeleteLB",
	3211: "EMsg_LBSDeleteLBEntry",
	3212: "EMsg_LBSResetLB",
	3400: "EMsg_BaseOGS",
	3401: "EMsg_OGSBeginSession",
	3402: "EMsg_OGSBeginSessionResponse",
	3403: "EMsg_OGSEndSession",
	3404: "EMsg_OGSEndSessionResponse",
	3406: "EMsg_OGSWriteAppSessionRow",
	3600: "EMsg_BaseBRP",
	3601: "EMsg_BRPStartShippingJobs",
	3602: "EMsg_BRPProcessUSBankReports",
	3603: "EMsg_BRPProcessGCReports",
	3604: "EMsg_BRPProcessPPReports",
	3605: "EMsg_BRPSettleNOVA",
	3606: "EMsg_BRPSettleCB",
	3607: "EMsg_BRPCommitGC",
	3608: "EMsg_BRPCommitGCResponse",
	3609: "EMsg_BRPFindHungTransactions",
	3610: "EMsg_BRPCheckFinanceCloseOutDate",
	3611: "EMsg_BRPProcessLicenses",
	3612: "EMsg_BRPProcessLicensesResponse",
	3613: "EMsg_BRPRemoveExpiredPaymentData",
	3614: "EMsg_BRPRemoveExpiredPaymentDataResponse",
	3615: "EMsg_BRPConvertToCurrentKeys",
	3616: "EMsg_BRPConvertToCurrentKeysResponse",
	3617: "EMsg_BRPPruneCardUsageStats",
	3618: "EMsg_BRPPruneCardUsageStatsResponse",
	3619: "EMsg_BRPCheckActivationCodes",
	3620: "EMsg_BRPCheckActivationCodesResponse",
	4000: "EMsg_BaseAMRange2",
	4001: "EMsg_AMCreateChat",
	4002: "EMsg_AMCreateChatResponse",
	4003: "EMsg_AMUpdateChatMetadata",
	4004: "EMsg_AMPublishChatMetadata",
	4005: "EMsg_AMSetProfileURL",
	4006: "EMsg_AMGetAccountEmailAddress",
	4007: "EMsg_AMGetAccountEmailAddressResponse",
	4008: "EMsg_AMRequestFriendData",
	4009: "EMsg_AMRouteToClients",
	4010: "EMsg_AMLeaveClan",
	4011: "EMsg_AMClanPermissions",
	4012: "EMsg_AMClanPermissionsResponse",
	4013: "EMsg_AMCreateClanEvent",
	4014: "EMsg_AMCreateClanEventResponse",
	4015: "EMsg_AMUpdateClanEvent",
	4016: "EMsg_AMUpdateClanEventResponse",
	4017: "EMsg_AMGetClanEvents",
	4018: "EMsg_AMGetClanEventsResponse",
	4019: "EMsg_AMDeleteClanEvent",
	4020: "EMsg_AMDeleteClanEventResponse",
	4021: "EMsg_AMSetClanPermissionSettings",
	4022: "EMsg_AMSetClanPermissionSettingsResponse",
	4023: "EMsg_AMGetClanPermissionSettings",
	4024: "EMsg_AMGetClanPermissionSettingsResponse",
	4025: "EMsg_AMPublishChatRoomInfo",
	4026: "EMsg_ClientChatRoomInfo",
	4027: "EMsg_AMCreateClanAnnouncement",
	4028: "EMsg_AMCreateClanAnnouncementResponse",
	4029: "EMsg_AMUpdateClanAnnouncement",
	4030: "EMsg_AMUpdateClanAnnouncementResponse",
	4031: "EMsg_AMGetClanAnnouncementsCount",
	4032: "EMsg_AMGetClanAnnouncementsCountResponse",
	4033: "EMsg_AMGetClanAnnouncements",
	4034: "EMsg_AMGetClanAnnouncementsResponse",
	4035: "EMsg_AMDeleteClanAnnouncement",
	4036: "EMsg_AMDeleteClanAnnouncementResponse",
	4037: "EMsg_AMGetSingleClanAnnouncement",
	4038: "EMsg_AMGetSingleClanAnnouncementResponse",
	4039: "EMsg_AMGetClanHistory",
	4040: "EMsg_AMGetClanHistoryResponse",
	4041: "EMsg_AMGetClanPermissionBits",
	4042: "EMsg_AMGetClanPermissionBitsResponse",
	4043: "EMsg_AMSetClanPermissionBits",
	4044: "EMsg_AMSetClanPermissionBitsResponse",
	4045: "EMsg_AMSessionInfoRequest",
	4046: "EMsg_AMSessionInfoResponse",
	4047: "EMsg_AMValidateWGToken",
	4048: "EMsg_AMGetSingleClanEvent",
	4049: "EMsg_AMGetSingleClanEventResponse",
	4050: "EMsg_AMGetClanRank",
	4051: "EMsg_AMGetClanRankResponse",
	4052: "EMsg_AMSetClanRank",
	4053: "EMsg_AMSetClanRankResponse",
	4054: "EMsg_AMGetClanPOTW",
	4055: "EMsg_AMGetClanPOTWResponse",
	4056: "EMsg_AMSetClanPOTW",
	4057: "EMsg_AMSetClanPOTWResponse",
	4058: "EMsg_AMRequestChatMetadata",
	4059: "EMsg_AMDumpUser",
	4060: "EMsg_AMKickUserFromClan",
	4061: "EMsg_AMAddFounderToClan",
	4062: "EMsg_AMValidateWGTokenResponse",
	4063: "EMsg_AMSetCommunityState",
	4064: "EMsg_AMSetAccountDetails",
	4065: "EMsg_AMGetChatBanList",
	4066: "EMsg_AMGetChatBanListResponse",
	4067: "EMsg_AMUnBanFromChat",
	4068: "EMsg_AMSetClanDetails",
	4069: "EMsg_AMGetAccountLinks",
	4070: "EMsg_AMGetAccountLinksResponse",
	4071: "EMsg_AMSetAccountLinks",
	4072: "EMsg_AMSetAccountLinksResponse",
	4073: "EMsg_AMGetUserGameStats",
	4074: "EMsg_AMGetUserGameStatsResponse",
	4075: "EMsg_AMCheckClanMembership",
	4076: "EMsg_AMGetClanMembers",
	4077: "EMsg_AMGetClanMembersResponse",
	4078: "EMsg_AMJoinPublicClan",
	4079: "EMsg_AMNotifyChatOfClanChange",
	4080: "EMsg_AMResubmitPurchase",
	4081: "EMsg_AMAddFriend",
	4082: "EMsg_AMAddFriendResponse",
	4083: "EMsg_AMRemoveFriend",
	4084: "EMsg_AMDumpClan",
	4085: "EMsg_AMChangeClanOwner",
	4086: "EMsg_AMCancelEasyCollect",
	4087: "EMsg_AMCancelEasyCollectResponse",
	4088: "EMsg_AMGetClanMembershipList",
	4089: "EMsg_AMGetClanMembershipListResponse",
	4090: "EMsg_AMClansInCommon",
	4091: "EMsg_AMClansInCommonResponse",
	4092: "EMsg_AMIsValidAccountID",
	4093: "EMsg_AMConvertClan",
	4094: "EMsg_AMGetGiftTargetListRelay",
	4095: "EMsg_AMWipeFriendsList",
	4096: "EMsg_AMSetIgnored",
	4097: "EMsg_AMClansInCommonCountResponse",
	4098: "EMsg_AMFriendsList",
	4099: "EMsg_AMFriendsListResponse",
	4100: "EMsg_AMFriendsInCommon",
	4101: "EMsg_AMFriendsInCommonResponse",
	4102: "EMsg_AMFriendsInCommonCountResponse",
	4103: "EMsg_AMClansInCommonCount",
	4104: "EMsg_AMChallengeVerdict",
	4105: "EMsg_AMChallengeNotification",
	4106: "EMsg_AMFindGSByIP",
	4107: "EMsg_AMFoundGSByIP",
	4108: "EMsg_AMGiftRevoked",
	4109: "EMsg_AMCreateAccountRecord",
	4110: "EMsg_AMUserClanList",
	4111: "EMsg_AMUserClanListResponse",
	4112: "EMsg_AMGetAccountDetails2",
	4113: "EMsg_AMGetAccountDetailsResponse2",
	4114: "EMsg_AMSetCommunityProfileSettings",
	4115: "EMsg_AMSetCommunityProfileSettingsResponse",
	4116: "EMsg_AMGetCommunityPrivacyState",
	4117: "EMsg_AMGetCommunityPrivacyStateResponse",
	4118: "EMsg_AMCheckClanInviteRateLimiting",
	4119: "EMsg_AMGetUserAchievementStatus",
	4120: "EMsg_AMGetIgnored",
	4121: "EMsg_AMGetIgnoredResponse",
	4122: "EMsg_AMSetIgnoredResponse",
	4123: "EMsg_AMSetFriendRelationshipNone",
	4124: "EMsg_AMGetFriendRelationship",
	4125: "EMsg_AMGetFriendRelationshipResponse",
	4126: "EMsg_AMServiceModulesCache",
	4127: "EMsg_AMServiceModulesCall",
	4128: "EMsg_AMServiceModulesCallResponse",
	4129: "EMsg_AMGetCaptchaDataForIP",
	4130: "EMsg_AMGetCaptchaDataForIPResponse",
	4131: "EMsg_AMValidateCaptchaDataForIP",
	4132: "EMsg_AMValidateCaptchaDataForIPResponse",
	4133: "EMsg_AMTrackFailedAuthByIP",
	4134: "EMsg_AMGetCaptchaDataByGID",
	4135: "EMsg_AMGetCaptchaDataByGIDResponse",
	4136: "EMsg_AMGetLobbyList",
	4137: "EMsg_AMGetLobbyListResponse",
	4138: "EMsg_AMGetLobbyMetadata",
	4139: "EMsg_AMGetLobbyMetadataResponse",
	4140: "EMsg_CommunityAddFriendNews",
	4141: "EMsg_AMAddClanNews",
	4142: "EMsg_AMWriteNews",
	4143: "EMsg_AMFindClanUser",
	4144: "EMsg_AMFindClanUserResponse",
	4145: "EMsg_AMBanFromChat",
	4146: "EMsg_AMGetUserHistoryResponse",
	4147: "EMsg_AMGetUserNewsSubscriptions",
	4148: "EMsg_AMGetUserNewsSubscriptionsResponse",
	4149: "EMsg_AMSetUserNewsSubscriptions",
	4150: "EMsg_AMGetUserNews",
	4151: "EMsg_AMGetUserNewsResponse",
	4152: "EMsg_AMSendQueuedEmails",
	4153: "EMsg_AMSetLicenseFlags",
	4154: "EMsg_AMGetUserHistory",
	4155: "EMsg_CommunityDeleteUserNews",
	4156: "EMsg_AMAllowUserFilesRequest",
	4157: "EMsg_AMAllowUserFilesResponse",
	4158: "EMsg_AMGetAccountStatus",
	4159: "EMsg_AMGetAccountStatusResponse",
	4160: "EMsg_AMEditBanReason",
	4161: "EMsg_AMCheckClanMembershipResponse",
	4162: "EMsg_AMProbeClanMembershipList",
	4163: "EMsg_AMProbeClanMembershipListResponse",
	4165: "EMsg_AMGetFriendsLobbies",
	4166: "EMsg_AMGetFriendsLobbiesResponse",
	4172: "EMsg_AMGetUserFriendNewsResponse",
	4173: "EMsg_CommunityGetUserFriendNews",
	4174: "EMsg_AMGetUserClansNewsResponse",
	4175: "EMsg_AMGetUserClansNews",
	4176: "EMsg_AMStoreInitPurchase",
	4177: "EMsg_AMStoreInitPurchaseResponse",
	4178: "EMsg_AMStoreGetFinalPrice",
	4179: "EMsg_AMStoreGetFinalPriceResponse",
	4180: "EMsg_AMStoreCompletePurchase",
	4181: "EMsg_AMStoreCancelPurchase",
	4182: "EMsg_AMStorePurchaseResponse",
	4183: "EMsg_AMCreateAccountRecordInSteam3",
	4184: "EMsg_AMGetPreviousCBAccount",
	4185: "EMsg_AMGetPreviousCBAccountResponse",
	4186: "EMsg_AMUpdateBillingAddress",
	4187: "EMsg_AMUpdateBillingAddressResponse",
	4188: "EMsg_AMGetBillingAddress",
	4189: "EMsg_AMGetBillingAddressResponse",
	4190: "EMsg_AMGetUserLicenseHistory",
	4191: "EMsg_AMGetUserLicenseHistoryResponse",
	4194: "EMsg_AMSupportChangePassword",
	4195: "EMsg_AMSupportChangeEmail",
	4196: "EMsg_AMSupportChangeSecretQA",
	4197: "EMsg_AMResetUserVerificationGSByIP",
	4198: "EMsg_AMUpdateGSPlayStats",
	4199: "EMsg_AMSupportEnableOrDisable",
	4200: "EMsg_AMGetComments",
	4201: "EMsg_AMGetCommentsResponse",
	4202: "EMsg_AMAddComment",
	4203: "EMsg_AMAddCommentResponse",
	4204: "EMsg_AMDeleteComment",
	4205: "EMsg_AMDeleteCommentResponse",
	4206: "EMsg_AMGetPurchaseStatus",
	4209: "EMsg_AMSupportIsAccountEnabled",
	4210: "EMsg_AMSupportIsAccountEnabledResponse",
	4211: "EMsg_AMGetUserStats",
	4212: "EMsg_AMSupportKickSession",
	4213: "EMsg_AMGSSearch",
	4216: "EMsg_MarketingMessageUpdate",
	4219: "EMsg_AMRouteFriendMsg",
	4220: "EMsg_AMTicketAuthRequestOrResponse",
	4222: "EMsg_AMVerifyDepotManagementRights",
	4223: "EMsg_AMVerifyDepotManagementRightsResponse",
	4224: "EMsg_AMAddFreeLicense",
	4225: "EMsg_AMGetUserFriendsMinutesPlayed",
	4226: "EMsg_AMGetUserFriendsMinutesPlayedResponse",
	4227: "EMsg_AMGetUserMinutesPlayed",
	4228: "EMsg_AMGetUserMinutesPlayedResponse",
	4231: "EMsg_AMValidateEmailLink",
	4232: "EMsg_AMValidateEmailLinkResponse",
	4234: "EMsg_AMAddUsersToMarketingTreatment",
	4236: "EMsg_AMStoreUserStats",
	4237: "EMsg_AMGetUserGameplayInfo",
	4238: "EMsg_AMGetUserGameplayInfoResponse",
	4239: "EMsg_AMGetCardList",
	4240: "EMsg_AMGetCardListResponse",
	4241: "EMsg_AMDeleteStoredCard",
	4242: "EMsg_AMRevokeLegacyGameKeys",
	4244: "EMsg_AMGetWalletDetails",
	4245: "EMsg_AMGetWalletDetailsResponse",
	4246: "EMsg_AMDeleteStoredPaymentInfo",
	4247: "EMsg_AMGetStoredPaymentSummary",
	4248: "EMsg_AMGetStoredPaymentSummaryResponse",
	4249: "EMsg_AMGetWalletConversionRate",
	4250: "EMsg_AMGetWalletConversionRateResponse",
	4251: "EMsg_AMConvertWallet",
	4252: "EMsg_AMConvertWalletResponse",
	4253: "EMsg_AMRelayGetFriendsWhoPlayGame",
	4254: "EMsg_AMRelayGetFriendsWhoPlayGameResponse",
	4255: "EMsg_AMSetPreApproval",
	4256: "EMsg_AMSetPreApprovalResponse",
	4257: "EMsg_AMMarketingTreatmentUpdate",
	4258: "EMsg_AMCreateRefund",
	4259: "EMsg_AMCreateRefundResponse",
	4260: "EMsg_AMCreateChargeback",
	4261: "EMsg_AMCreateChargebackResponse",
	4262: "EMsg_AMCreateDispute",
	4263: "EMsg_AMCreateDisputeResponse",
	4264: "EMsg_AMClearDispute",
	4265: "EMsg_AMClearDisputeResponse",
	4266: "EMsg_AMPlayerNicknameList",
	4267: "EMsg_AMPlayerNicknameListResponse",
	4268: "EMsg_AMSetDRMTestConfig",
	4269: "EMsg_AMGetUserCurrentGameInfo",
	4270: "EMsg_AMGetUserCurrentGameInfoResponse",
	4271: "EMsg_AMGetGSPlayerList",
	4272: "EMsg_AMGetGSPlayerListResponse",
	4275: "EMsg_AMUpdatePersonaStateCache",
	4276: "EMsg_AMGetGameMembers",
	4277: "EMsg_AMGetGameMembersResponse",
	4278: "EMsg_AMGetSteamIDForMicroTxn",
	4279: "EMsg_AMGetSteamIDForMicroTxnResponse",
	4280: "EMsg_AMAddPublisherUser",
	4281: "EMsg_AMRemovePublisherUser",
	4282: "EMsg_AMGetUserLicenseList",
	4283: "EMsg_AMGetUserLicenseListResponse",
	4284: "EMsg_AMReloadGameGroupPolicy",
	4285: "EMsg_AMAddFreeLicenseResponse",
	4286: "EMsg_AMVACStatusUpdate",
	4287: "EMsg_AMGetAccountDetails",
	4288: "EMsg_AMGetAccountDetailsResponse",
	4289: "EMsg_AMGetPlayerLinkDetails",
	4290: "EMsg_AMGetPlayerLinkDetailsResponse",
	4291: "EMsg_AMSubscribeToPersonaFeed",
	4292: "EMsg_AMGetUserVacBanList",
	4293: "EMsg_AMGetUserVacBanListResponse",
	4294: "EMsg_AMGetAccountFlagsForWGSpoofing",
	4295: "EMsg_AMGetAccountFlagsForWGSpoofingResponse",
	4296: "EMsg_AMGetFriendsWishlistInfo",
	4297: "EMsg_AMGetFriendsWishlistInfoResponse",
	4298: "EMsg_AMGetClanOfficers",
	4299: "EMsg_AMGetClanOfficersResponse",
	4300: "EMsg_AMNameChange",
	4301: "EMsg_AMGetNameHistory",
	4302: "EMsg_AMGetNameHistoryResponse",
	4305: "EMsg_AMUpdateProviderStatus",
	4306: "EMsg_AMClearPersonaMetadataBlob",
	4307: "EMsg_AMSupportRemoveAccountSecurity",
	4308: "EMsg_AMIsAccountInCaptchaGracePeriod",
	4309: "EMsg_AMIsAccountInCaptchaGracePeriodResponse",
	4310: "EMsg_AMAccountPS3Unlink",
	4311: "EMsg_AMAccountPS3UnlinkResponse",
	4312: "EMsg_AMStoreUserStatsResponse",
	4313: "EMsg_AMGetAccountPSNInfo",
	4314: "EMsg_AMGetAccountPSNInfoResponse",
	4315: "EMsg_AMAuthenticatedPlayerList",
	4316: "EMsg_AMGetUserGifts",
	4317: "EMsg_AMGetUserGiftsResponse",
	4320: "EMsg_AMTransferLockedGifts",
	4321: "EMsg_AMTransferLockedGiftsResponse",
	4322: "EMsg_AMPlayerHostedOnGameServer",
	4323: "EMsg_AMGetAccountBanInfo",
	4324: "EMsg_AMGetAccountBanInfoResponse",
	4325: "EMsg_AMRecordBanEnforcement",
	4326: "EMsg_AMRollbackGiftTransfer",
	4327: "EMsg_AMRollbackGiftTransferResponse",
	4328: "EMsg_AMHandlePendingTransaction",
	4329: "EMsg_AMRequestClanDetails",
	4330: "EMsg_AMDeleteStoredPaypalAgreement",
	4331: "EMsg_AMGameServerUpdate",
	4332: "EMsg_AMGameServerRemove",
	4333: "EMsg_AMGetPaypalAgreements",
	4334: "EMsg_AMGetPaypalAgreementsResponse",
	4335: "EMsg_AMGameServerPlayerCompatibilityCheck",
	4336: "EMsg_AMGameServerPlayerCompatibilityCheckResponse",
	4337: "EMsg_AMRenewLicense",
	4338: "EMsg_AMGetAccountCommunityBanInfo",
	4339: "EMsg_AMGetAccountCommunityBanInfoResponse",
	4340: "EMsg_AMGameServerAccountChangePassword",
	4341: "EMsg_AMGameServerAccountDeleteAccount",
	4342: "EMsg_AMRenewAgreement",
	4343: "EMsg_AMSendEmail",
	4344: "EMsg_AMXsollaPayment",
	4345: "EMsg_AMXsollaPaymentResponse",
	4346: "EMsg_AMAcctAllowedToPurchase",
	4347: "EMsg_AMAcctAllowedToPurchaseResponse",
	4348: "EMsg_AMSwapKioskDeposit",
	4349: "EMsg_AMSwapKioskDepositResponse",
	4350: "EMsg_AMSetUserGiftUnowned",
	4351: "EMsg_AMSetUserGiftUnownedResponse",
	4352: "EMsg_AMClaimUnownedUserGift",
	4353: "EMsg_AMClaimUnownedUserGiftResponse",
	4354: "EMsg_AMSetClanName",
	4355: "EMsg_AMSetClanNameResponse",
	4356: "EMsg_AMGrantCoupon",
	4357: "EMsg_AMGrantCouponResponse",
	4358: "EMsg_AMIsPackageRestrictedInUserCountry",
	4359: "EMsg_AMIsPackageRestrictedInUserCountryResponse",
	4360: "EMsg_AMHandlePendingTransactionResponse",
	4361: "EMsg_AMGrantGuestPasses2",
	4362: "EMsg_AMGrantGuestPasses2Response",
	4363: "EMsg_AMSessionQuery",
	4364: "EMsg_AMSessionQueryResponse",
	4365: "EMsg_AMGetPlayerBanDetails",
	4366: "EMsg_AMGetPlayerBanDetailsResponse",
	4367: "EMsg_AMFinalizePurchase",
	4368: "EMsg_AMFinalizePurchaseResponse",
	4372: "EMsg_AMPersonaChangeResponse",
	4373: "EMsg_AMGetClanDetailsForForumCreation",
	4374: "EMsg_AMGetClanDetailsForForumCreationResponse",
	4375: "EMsg_AMGetPendingNotificationCount",
	4376: "EMsg_AMGetPendingNotificationCountResponse",
	4377: "EMsg_AMPasswordHashUpgrade",
	4378: "EMsg_AMMoPayPayment",
	4379: "EMsg_AMMoPayPaymentResponse",
	4380: "EMsg_AMBoaCompraPayment",
	4381: "EMsg_AMBoaCompraPaymentResponse",
	4382: "EMsg_AMExpireCaptchaByGID",
	4383: "EMsg_AMCompleteExternalPurchase",
	4384: "EMsg_AMCompleteExternalPurchaseResponse",
	4385: "EMsg_AMResolveNegativeWalletCredits",
	4386: "EMsg_AMResolveNegativeWalletCreditsResponse",
	4387: "EMsg_AMPayelpPayment",
	4388: "EMsg_AMPayelpPaymentResponse",
	4389: "EMsg_AMPlayerGetClanBasicDetails",
	4390: "EMsg_AMPlayerGetClanBasicDetailsResponse",
	5000: "EMsg_BasePSRange",
	5001: "EMsg_PSCreateShoppingCart",
	5002: "EMsg_PSCreateShoppingCartResponse",
	5003: "EMsg_PSIsValidShoppingCart",
	5004: "EMsg_PSIsValidShoppingCartResponse",
	5005: "EMsg_PSAddPackageToShoppingCart",
	5006: "EMsg_PSAddPackageToShoppingCartResponse",
	5007: "EMsg_PSRemoveLineItemFromShoppingCart",
	5008: "EMsg_PSRemoveLineItemFromShoppingCartResponse",
	5009: "EMsg_PSGetShoppingCartContents",
	5010: "EMsg_PSGetShoppingCartContentsResponse",
	5011: "EMsg_PSAddWalletCreditToShoppingCart",
	5012: "EMsg_PSAddWalletCreditToShoppingCartResponse",
	5200: "EMsg_BaseUFSRange",
	5202: "EMsg_ClientUFSUploadFileRequest",
	5203: "EMsg_ClientUFSUploadFileResponse",
	5204: "EMsg_ClientUFSUploadFileChunk",
	5205: "EMsg_ClientUFSUploadFileFinished",
	5206: "EMsg_ClientUFSGetFileListForApp",
	5207: "EMsg_ClientUFSGetFileListForAppResponse",
	5210: "EMsg_ClientUFSDownloadRequest",
	5211: "EMsg_ClientUFSDownloadResponse",
	5212: "EMsg_ClientUFSDownloadChunk",
	5213: "EMsg_ClientUFSLoginRequest",
	5214: "EMsg_ClientUFSLoginResponse",
	5215: "EMsg_UFSReloadPartitionInfo",
	5216: "EMsg_ClientUFSTransferHeartbeat",
	5217: "EMsg_UFSSynchronizeFile",
	5218: "EMsg_UFSSynchronizeFileResponse",
	5219: "EMsg_ClientUFSDeleteFileRequest",
	5220: "EMsg_ClientUFSDeleteFileResponse",
	5221: "EMsg_UFSDownloadRequest",
	5222: "EMsg_UFSDownloadResponse",
	5223: "EMsg_UFSDownloadChunk",
	5226: "EMsg_ClientUFSGetUGCDetails",
	5227: "EMsg_ClientUFSGetUGCDetailsResponse",
	5228: "EMsg_UFSUpdateFileFlags",
	5229: "EMsg_UFSUpdateFileFlagsResponse",
	5230: "EMsg_ClientUFSGetSingleFileInfo",
	5231: "EMsg_ClientUFSGetSingleFileInfoResponse",
	5232: "EMsg_ClientUFSShareFile",
	5233: "EMsg_ClientUFSShareFileResponse",
	5234: "EMsg_UFSReloadAccount",
	5235: "EMsg_UFSReloadAccountResponse",
	5236: "EMsg_UFSUpdateRecordBatched",
	5237: "EMsg_UFSUpdateRecordBatchedResponse",
	5238: "EMsg_UFSMigrateFile",
	5239: "EMsg_UFSMigrateFileResponse",
	5240: "EMsg_UFSGetUGCURLs",
	5241: "EMsg_UFSGetUGCURLsResponse",
	5242: "EMsg_UFSHttpUploadFileFinishRequest",
	5243: "EMsg_UFSHttpUploadFileFinishResponse",
	5244: "EMsg_UFSDownloadStartRequest",
	5245: "EMsg_UFSDownloadStartResponse",
	5246: "EMsg_UFSDownloadChunkRequest",
	5247: "EMsg_UFSDownloadChunkResponse",
	5248: "EMsg_UFSDownloadFinishRequest",
	5249: "EMsg_UFSDownloadFinishResponse",
	5250: "EMsg_UFSFlushURLCache",
	5251: "EMsg_UFSUploadCommit",
	5252: "EMsg_UFSUploadCommitResponse",
	5400: "EMsg_BaseClient2",
	5401: "EMsg_ClientRequestForgottenPasswordEmail",
	5402: "EMsg_ClientRequestForgottenPasswordEmailResponse",
	5403: "EMsg_ClientCreateAccountResponse",
	5404: "EMsg_ClientResetForgottenPassword",
	5405: "EMsg_ClientResetForgottenPasswordResponse",
	5406: "EMsg_ClientCreateAccount2",
	5407: "EMsg_ClientInformOfResetForgottenPassword",
	5408: "EMsg_ClientInformOfResetForgottenPasswordResponse",
	5409: "EMsg_ClientAnonUserLogOn_Deprecated",
	5410: "EMsg_ClientGamesPlayedWithDataBlob",
	5411: "EMsg_ClientUpdateUserGameInfo",
	5412: "EMsg_ClientFileToDownload",
	5413: "EMsg_ClientFileToDownloadResponse",
	5414: "EMsg_ClientLBSSetScore",
	5415: "EMsg_ClientLBSSetScoreResponse",
	5416: "EMsg_ClientLBSFindOrCreateLB",
	5417: "EMsg_ClientLBSFindOrCreateLBResponse",
	5418: "EMsg_ClientLBSGetLBEntries",
	5419: "EMsg_ClientLBSGetLBEntriesResponse",
	5420: "EMsg_ClientMarketingMessageUpdate",
	5426: "EMsg_ClientChatDeclined",
	5427: "EMsg_ClientFriendMsgIncoming",
	5428: "EMsg_ClientAuthList_Deprecated",
	5429: "EMsg_ClientTicketAuthComplete",
	5430: "EMsg_ClientIsLimitedAccount",
	5431: "EMsg_ClientRequestAuthList",
	5432: "EMsg_ClientAuthList",
	5433: "EMsg_ClientStat",
	5434: "EMsg_ClientP2PConnectionInfo",
	5435: "EMsg_ClientP2PConnectionFailInfo",
	5436: "EMsg_ClientGetNumberOfCurrentPlayers",
	5437: "EMsg_ClientGetNumberOfCurrentPlayersResponse",
	5438: "EMsg_ClientGetDepotDecryptionKey",
	5439: "EMsg_ClientGetDepotDecryptionKeyResponse",
	5440: "EMsg_GSPerformHardwareSurvey",
	5441: "EMsg_ClientGetAppBetaPasswords",
	5442: "EMsg_ClientGetAppBetaPasswordsResponse",
	5443: "EMsg_ClientEnableTestLicense",
	5444: "EMsg_ClientEnableTestLicenseResponse",
	5445: "EMsg_ClientDisableTestLicense",
	5446: "EMsg_ClientDisableTestLicenseResponse",
	5448: "EMsg_ClientRequestValidationMail",
	5449: "EMsg_ClientRequestValidationMailResponse",
	5452: "EMsg_ClientToGC",
	5453: "EMsg_ClientFromGC",
	5454: "EMsg_ClientRequestChangeMail",
	5455: "EMsg_ClientRequestChangeMailResponse",
	5456: "EMsg_ClientEmailAddrInfo",
	5457: "EMsg_ClientPasswordChange3",
	5458: "EMsg_ClientEmailChange3",
	5459: "EMsg_ClientPersonalQAChange3",
	5460: "EMsg_ClientResetForgottenPassword3",
	5461: "EMsg_ClientRequestForgottenPasswordEmail3",
	5462: "EMsg_ClientCreateAccount3",
	5463: "EMsg_ClientNewLoginKey",
	5464: "EMsg_ClientNewLoginKeyAccepted",
	5465: "EMsg_ClientLogOnWithHash_Deprecated",
	5466: "EMsg_ClientStoreUserStats2",
	5467: "EMsg_ClientStatsUpdated",
	5468: "EMsg_ClientActivateOEMLicense",
	5469: "EMsg_ClientRegisterOEMMachine",
	5470: "EMsg_ClientRegisterOEMMachineResponse",
	5480: "EMsg_ClientRequestedClientStats",
	5481: "EMsg_ClientStat2Int32",
	5482: "EMsg_ClientStat2",
	5483: "EMsg_ClientVerifyPassword",
	5484: "EMsg_ClientVerifyPasswordResponse",
	5485: "EMsg_ClientDRMDownloadRequest",
	5486: "EMsg_ClientDRMDownloadResponse",
	5487: "EMsg_ClientDRMFinalResult",
	5488: "EMsg_ClientGetFriendsWhoPlayGame",
	5489: "EMsg_ClientGetFriendsWhoPlayGameResponse",
	5490: "EMsg_ClientOGSBeginSession",
	5491: "EMsg_ClientOGSBeginSessionResponse",
	5492: "EMsg_ClientOGSEndSession",
	5493: "EMsg_ClientOGSEndSessionResponse",
	5494: "EMsg_ClientOGSWriteRow",
	5495: "EMsg_ClientDRMTest",
	5496: "EMsg_ClientDRMTestResult",
	5500: "EMsg_ClientServerUnavailable",
	5501: "EMsg_ClientServersAvailable",
	5502: "EMsg_ClientRegisterAuthTicketWithCM",
	5503: "EMsg_ClientGCMsgFailed",
	5504: "EMsg_ClientMicroTxnAuthRequest",
	5505: "EMsg_ClientMicroTxnAuthorize",
	5506: "EMsg_ClientMicroTxnAuthorizeResponse",
	5507: "EMsg_ClientAppMinutesPlayedData",
	5508: "EMsg_ClientGetMicroTxnInfo",
	5509: "EMsg_ClientGetMicroTxnInfoResponse",
	5510: "EMsg_ClientMarketingMessageUpdate2",
	5511: "EMsg_ClientDeregisterWithServer",
	5512: "EMsg_ClientSubscribeToPersonaFeed",
	5514: "EMsg_ClientLogon",
	5515: "EMsg_ClientGetClientDetails",
	5516: "EMsg_ClientGetClientDetailsResponse",
	5517: "EMsg_ClientReportOverlayDetourFailure",
	5518: "EMsg_ClientGetClientAppList",
	5519: "EMsg_ClientGetClientAppListResponse",
	5520: "EMsg_ClientInstallClientApp",
	5521: "EMsg_ClientInstallClientAppResponse",
	5522: "EMsg_ClientUninstallClientApp",
	5523: "EMsg_ClientUninstallClientAppResponse",
	5524: "EMsg_ClientSetClientAppUpdateState",
	5525: "EMsg_ClientSetClientAppUpdateStateResponse",
	5526: "EMsg_ClientRequestEncryptedAppTicket",
	5527: "EMsg_ClientRequestEncryptedAppTicketResponse",
	5528: "EMsg_ClientWalletInfoUpdate",
	5529: "EMsg_ClientLBSSetUGC",
	5530: "EMsg_ClientLBSSetUGCResponse",
	5531: "EMsg_ClientAMGetClanOfficers",
	5532: "EMsg_ClientAMGetClanOfficersResponse",
	5533: "EMsg_ClientCheckFileSignature",
	5534: "EMsg_ClientCheckFileSignatureResponse",
	5535: "EMsg_ClientFriendProfileInfo",
	5536: "EMsg_ClientFriendProfileInfoResponse",
	5537: "EMsg_ClientUpdateMachineAuth",
	5538: "EMsg_ClientUpdateMachineAuthResponse",
	5539: "EMsg_ClientReadMachineAuth",
	5540: "EMsg_ClientReadMachineAuthResponse",
	5541: "EMsg_ClientRequestMachineAuth",
	5542: "EMsg_ClientRequestMachineAuthResponse",
	5543: "EMsg_ClientScreenshotsChanged",
	5544: "EMsg_ClientEmailChange4",
	5545: "EMsg_ClientEmailChangeResponse4",
	5546: "EMsg_ClientGetCDNAuthToken",
	5547: "EMsg_ClientGetCDNAuthTokenResponse",
	5548: "EMsg_ClientDownloadRateStatistics",
	5549: "EMsg_ClientRequestAccountData",
	5550: "EMsg_ClientRequestAccountDataResponse",
	5551: "EMsg_ClientResetForgottenPassword4",
	5552: "EMsg_ClientHideFriend",
	5553: "EMsg_ClientFriendsGroupsList",
	5554: "EMsg_ClientGetClanActivityCounts",
	5555: "EMsg_ClientGetClanActivityCountsResponse",
	5556: "EMsg_ClientOGSReportString",
	5557: "EMsg_ClientOGSReportBug",
	5558: "EMsg_ClientSentLogs",
	5559: "EMsg_ClientLogonGameServer",
	5560: "EMsg_AMClientCreateFriendsGroup",
	5561: "EMsg_AMClientCreateFriendsGroupResponse",
	5562: "EMsg_AMClientDeleteFriendsGroup",
	5563: "EMsg_AMClientDeleteFriendsGroupResponse",
	5564: "EMsg_AMClientRenameFriendsGroup",
	5565: "EMsg_AMClientRenameFriendsGroupResponse",
	5566: "EMsg_AMClientAddFriendToGroup",
	5567: "EMsg_AMClientAddFriendToGroupResponse",
	5568: "EMsg_AMClientRemoveFriendFromGroup",
	5569: "EMsg_AMClientRemoveFriendFromGroupResponse",
	5570: "EMsg_ClientAMGetPersonaNameHistory",
	5571: "EMsg_ClientAMGetPersonaNameHistoryResponse",
	5572: "EMsg_ClientRequestFreeLicense",
	5573: "EMsg_ClientRequestFreeLicenseResponse",
	5574: "EMsg_ClientDRMDownloadRequestWithCrashData",
	5575: "EMsg_ClientAuthListAck",
	5576: "EMsg_ClientItemAnnouncements",
	5577: "EMsg_ClientRequestItemAnnouncements",
	5578: "EMsg_ClientFriendMsgEchoToSender",
	5579: "EMsg_ClientChangeSteamGuardOptions",
	5580: "EMsg_ClientChangeSteamGuardOptionsResponse",
	5581: "EMsg_ClientOGSGameServerPingSample",
	5582: "EMsg_ClientCommentNotifications",
	5583: "EMsg_ClientRequestCommentNotifications",
	5584: "EMsg_ClientPersonaChangeResponse",
	5585: "EMsg_ClientRequestWebAPIAuthenticateUserNonce",
	5586: "EMsg_ClientRequestWebAPIAuthenticateUserNonceResponse",
	5587: "EMsg_ClientPlayerNicknameList",
	5588: "EMsg_AMClientSetPlayerNickname",
	5589: "EMsg_AMClientSetPlayerNicknameResponse",
	5590: "EMsg_ClientRequestOAuthTokenForApp",
	5591: "EMsg_ClientRequestOAuthTokenForAppResponse",
	5592: "EMsg_ClientGetNumberOfCurrentPlayersDP",
	5593: "EMsg_ClientGetNumberOfCurrentPlayersDPResponse",
	5594: "EMsg_ClientServiceMethod",
	5595: "EMsg_ClientServiceMethodResponse",
	5596: "EMsg_ClientFriendUserStatusPublished",
	5597: "EMsg_ClientCurrentUIMode",
	5598: "EMsg_ClientVanityURLChangedNotification",
	5599: "EMsg_ClientUserNotifications",
	5600: "EMsg_BaseDFS",
	5601: "EMsg_DFSGetFile",
	5602: "EMsg_DFSInstallLocalFile",
	5603: "EMsg_DFSConnection",
	5604: "EMsg_DFSConnectionReply",
	5605: "EMsg_ClientDFSAuthenticateRequest",
	5606: "EMsg_ClientDFSAuthenticateResponse",
	5607: "EMsg_ClientDFSEndSession",
	5608: "EMsg_DFSPurgeFile",
	5609: "EMsg_DFSRouteFile",
	5610: "EMsg_DFSGetFileFromServer",
	5611: "EMsg_DFSAcceptedResponse",
	5612: "EMsg_DFSRequestPingback",
	5613: "EMsg_DFSRecvTransmitFile",
	5614: "EMsg_DFSSendTransmitFile",
	5615: "EMsg_DFSRequestPingback2",
	5616: "EMsg_DFSResponsePingback2",
	5617: "EMsg_ClientDFSDownloadStatus",
	5618: "EMsg_DFSStartTransfer",
	5619: "EMsg_DFSTransferComplete",
	5800: "EMsg_BaseMDS",
	5801: "EMsg_ClientMDSLoginRequest",
	5802: "EMsg_ClientMDSLoginResponse",
	5803: "EMsg_ClientMDSUploadManifestRequest",
	5804: "EMsg_ClientMDSUploadManifestResponse",
	5805: "EMsg_ClientMDSTransmitManifestDataChunk",
	5806: "EMsg_ClientMDSHeartbeat",
	5807: "EMsg_ClientMDSUploadDepotChunks",
	5808: "EMsg_ClientMDSUploadDepotChunksResponse",
	5809: "EMsg_ClientMDSInitDepotBuildRequest",
	5810: "EMsg_ClientMDSInitDepotBuildResponse",
	5812: "EMsg_AMToMDSGetDepotDecryptionKey",
	5813: "EMsg_MDSToAMGetDepotDecryptionKeyResponse",
	5814: "EMsg_MDSGetVersionsForDepot",
	5815: "EMsg_MDSGetVersionsForDepotResponse",
	5816: "EMsg_MDSSetPublicVersionForDepot",
	5817: "EMsg_MDSSetPublicVersionForDepotResponse",
	5818: "EMsg_ClientMDSGetDepotManifest",
	5819: "EMsg_ClientMDSGetDepotManifestResponse",
	5820: "EMsg_ClientMDSGetDepotManifestChunk",
	5823: "EMsg_ClientMDSUploadRateTest",
	5824: "EMsg_ClientMDSUploadRateTestResponse",
	5825: "EMsg_MDSDownloadDepotChunksAck",
	5826: "EMsg_MDSContentServerStatsBroadcast",
	5827: "EMsg_MDSContentServerConfigRequest",
	5828: "EMsg_MDSContentServerConfig",
	5829: "EMsg_MDSGetDepotManifest",
	5830: "EMsg_MDSGetDepotManifestResponse",
	5831: "EMsg_MDSGetDepotManifestChunk",
	5832: "EMsg_MDSGetDepotChunk",
	5833: "EMsg_MDSGetDepotChunkResponse",
	5834: "EMsg_MDSGetDepotChunkChunk",
	5835: "EMsg_MDSUpdateContentServerConfig",
	5836: "EMsg_MDSGetServerListForUser",
	5837: "EMsg_MDSGetServerListForUserResponse",
	5838: "EMsg_ClientMDSRegisterAppBuild",
	5839: "EMsg_ClientMDSRegisterAppBuildResponse",
	5840: "EMsg_ClientMDSSetAppBuildLive",
	5841: "EMsg_ClientMDSSetAppBuildLiveResponse",
	5842: "EMsg_ClientMDSGetPrevDepotBuild",
	5843: "EMsg_ClientMDSGetPrevDepotBuildResponse",
	5844: "EMsg_MDSToCSFlushChunk",
	5845: "EMsg_ClientMDSSignInstallScript",
	5846: "EMsg_ClientMDSSignInstallScriptResponse",
	6200: "EMsg_CSBase",
	6201: "EMsg_CSPing",
	6202: "EMsg_CSPingResponse",
	6400: "EMsg_GMSBase",
	6401: "EMsg_GMSGameServerReplicate",
	6403: "EMsg_ClientGMSServerQuery",
	6404: "EMsg_GMSClientServerQueryResponse",
	6405: "EMsg_AMGMSGameServerUpdate",
	6406: "EMsg_AMGMSGameServerRemove",
	6407: "EMsg_GameServerOutOfDate",
	6501: "EMsg_ClientAuthorizeLocalDeviceRequest",
	6502: "EMsg_ClientAuthorizeLocalDevice",
	6503: "EMsg_ClientDeauthorizeDeviceRequest",
	6504: "EMsg_ClientDeauthorizeDevice",
	6505: "EMsg_ClientUseLocalDeviceAuthorizations",
	6506: "EMsg_ClientGetAuthorizedDevices",
	6507: "EMsg_ClientGetAuthorizedDevicesResponse",
	6600: "EMsg_MMSBase",
	6601: "EMsg_ClientMMSCreateLobby",
	6602: "EMsg_ClientMMSCreateLobbyResponse",
	6603: "EMsg_ClientMMSJoinLobby",
	6604: "EMsg_ClientMMSJoinLobbyResponse",
	6605: "EMsg_ClientMMSLeaveLobby",
	6606: "EMsg_ClientMMSLeaveLobbyResponse",
	6607: "EMsg_ClientMMSGetLobbyList",
	6608: "EMsg_ClientMMSGetLobbyListResponse",
	6609: "EMsg_ClientMMSSetLobbyData",
	6610: "EMsg_ClientMMSSetLobbyDataResponse",
	6611: "EMsg_ClientMMSGetLobbyData",
	6612: "EMsg_ClientMMSLobbyData",
	6613: "EMsg_ClientMMSSendLobbyChatMsg",
	6614: "EMsg_ClientMMSLobbyChatMsg",
	6615: "EMsg_ClientMMSSetLobbyOwner",
	6616: "EMsg_ClientMMSSetLobbyOwnerResponse",
	6617: "EMsg_ClientMMSSetLobbyGameServer",
	6618: "EMsg_ClientMMSLobbyGameServerSet",
	6619: "EMsg_ClientMMSUserJoinedLobby",
	6620: "EMsg_ClientMMSUserLeftLobby",
	6621: "EMsg_ClientMMSInviteToLobby",
	6622: "EMsg_ClientMMSFlushFrenemyListCache",
	6623: "EMsg_ClientMMSFlushFrenemyListCacheResponse",
	6624: "EMsg_ClientMMSSetLobbyLinked",
	6800: "EMsg_NonStdMsgBase",
	6801: "EMsg_NonStdMsgMemcached",
	6802: "EMsg_NonStdMsgHTTPServer",
	6803: "EMsg_NonStdMsgHTTPClient",
	6804: "EMsg_NonStdMsgWGResponse",
	6805: "EMsg_NonStdMsgPHPSimulator",
	6806: "EMsg_NonStdMsgChase",
	6807: "EMsg_NonStdMsgDFSTransfer",
	6808: "EMsg_NonStdMsgTests",
	6809: "EMsg_NonStdMsgUMQpipeAAPL",
	6810: "EMsg_NonStdMsgSyslog",
	6811: "EMsg_NonStdMsgLogsink",
	7000: "EMsg_UDSBase",
	7001: "EMsg_ClientUDSP2PSessionStarted",
	7002: "EMsg_ClientUDSP2PSessionEnded",
	7003: "EMsg_UDSRenderUserAuth",
	7004: "EMsg_UDSRenderUserAuthResponse",
	7005: "EMsg_ClientUDSInviteToGame",
	7006: "EMsg_UDSFindSession",
	7007: "EMsg_UDSFindSessionResponse",
	7100: "EMsg_MPASBase",
	7101: "EMsg_MPASVacBanReset",
	7200: "EMsg_KGSBase",
	7201: "EMsg_KGSAllocateKeyRange",
	7202: "EMsg_KGSAllocateKeyRangeResponse",
	7203: "EMsg_KGSGenerateKeys",
	7204: "EMsg_KGSGenerateKeysResponse",
	7205: "EMsg_KGSRemapKeys",
	7206: "EMsg_KGSRemapKeysResponse",
	7207: "EMsg_KGSGenerateGameStopWCKeys",
	7208: "EMsg_KGSGenerateGameStopWCKeysResponse",
	7300: "EMsg_UCMBase",
	7301: "EMsg_ClientUCMAddScreenshot",
	7302: "EMsg_ClientUCMAddScreenshotResponse",
	7303: "EMsg_UCMValidateObjectExists",
	7304: "EMsg_UCMValidateObjectExistsResponse",
	7307: "EMsg_UCMResetCommunityContent",
	7308: "EMsg_UCMResetCommunityContentResponse",
	7309: "EMsg_ClientUCMDeleteScreenshot",
	7310: "EMsg_ClientUCMDeleteScreenshotResponse",
	7311: "EMsg_ClientUCMPublishFile",
	7312: "EMsg_ClientUCMPublishFileResponse",
	7313: "EMsg_ClientUCMGetPublishedFileDetails",
	7314: "EMsg_ClientUCMGetPublishedFileDetailsResponse",
	7315: "EMsg_ClientUCMDeletePublishedFile",
	7316: "EMsg_ClientUCMDeletePublishedFileResponse",
	7317: "EMsg_ClientUCMEnumerateUserPublishedFiles",
	7318: "EMsg_ClientUCMEnumerateUserPublishedFilesResponse",
	7319: "EMsg_ClientUCMSubscribePublishedFile",
	7320: "EMsg_ClientUCMSubscribePublishedFileResponse",
	7321: "EMsg_ClientUCMEnumerateUserSubscribedFiles",
	7322: "EMsg_ClientUCMEnumerateUserSubscribedFilesResponse",
	7323: "EMsg_ClientUCMUnsubscribePublishedFile",
	7324: "EMsg_ClientUCMUnsubscribePublishedFileResponse",
	7325: "EMsg_ClientUCMUpdatePublishedFile",
	7326: "EMsg_ClientUCMUpdatePublishedFileResponse",
	7327: "EMsg_UCMUpdatePublishedFile",
	7328: "EMsg_UCMUpdatePublishedFileResponse",
	7329: "EMsg_UCMDeletePublishedFile",
	7330: "EMsg_UCMDeletePublishedFileResponse",
	7331: "EMsg_UCMUpdatePublishedFileStat",
	7332: "EMsg_UCMUpdatePublishedFileBan",
	7333: "EMsg_UCMUpdatePublishedFileBanResponse",
	7334: "EMsg_UCMUpdateTaggedScreenshot",
	7335: "EMsg_UCMAddTaggedScreenshot",
	7336: "EMsg_UCMRemoveTaggedScreenshot",
	7337: "EMsg_UCMReloadPublishedFile",
	7338: "EMsg_UCMReloadUserFileListCaches",
	7339: "EMsg_UCMPublishedFileReported",
	7340: "EMsg_UCMUpdatePublishedFileIncompatibleStatus",
	7341: "EMsg_UCMPublishedFilePreviewAdd",
	7342: "EMsg_UCMPublishedFilePreviewAddResponse",
	7343: "EMsg_UCMPublishedFilePreviewRemove",
	7344: "EMsg_UCMPublishedFilePreviewRemoveResponse",
	7345: "EMsg_UCMPublishedFilePreviewChangeSortOrder",
	7346: "EMsg_UCMPublishedFilePreviewChangeSortOrderResponse",
	7347: "EMsg_ClientUCMPublishedFileSubscribed",
	7348: "EMsg_ClientUCMPublishedFileUnsubscribed",
	7349: "EMsg_UCMPublishedFileSubscribed",
	7350: "EMsg_UCMPublishedFileUnsubscribed",
	7351: "EMsg_UCMPublishFile",
	7352: "EMsg_UCMPublishFileResponse",
	7353: "EMsg_UCMPublishedFileChildAdd",
	7354: "EMsg_UCMPublishedFileChildAddResponse",
	7355: "EMsg_UCMPublishedFileChildRemove",
	7356: "EMsg_UCMPublishedFileChildRemoveResponse",
	7357: "EMsg_UCMPublishedFileChildChangeSortOrder",
	7358: "EMsg_UCMPublishedFileChildChangeSortOrderResponse",
	7359: "EMsg_UCMPublishedFileParentChanged",
	7360: "EMsg_ClientUCMGetPublishedFilesForUser",
	7361: "EMsg_ClientUCMGetPublishedFilesForUserResponse",
	7362: "EMsg_UCMGetPublishedFilesForUser",
	7363: "EMsg_UCMGetPublishedFilesForUserResponse",
	7364: "EMsg_ClientUCMSetUserPublishedFileAction",
	7365: "EMsg_ClientUCMSetUserPublishedFileActionResponse",
	7366: "EMsg_ClientUCMEnumeratePublishedFilesByUserAction",
	7367: "EMsg_ClientUCMEnumeratePublishedFilesByUserActionResponse",
	7368: "EMsg_ClientUCMPublishedFileDeleted",
	7369: "EMsg_UCMGetUserSubscribedFiles",
	7370: "EMsg_UCMGetUserSubscribedFilesResponse",
	7371: "EMsg_UCMFixStatsPublishedFile",
	7372: "EMsg_UCMDeleteOldScreenshot",
	7373: "EMsg_UCMDeleteOldScreenshotResponse",
	7374: "EMsg_UCMDeleteOldVideo",
	7375: "EMsg_UCMDeleteOldVideoResponse",
	7376: "EMsg_UCMUpdateOldScreenshotPrivacy",
	7377: "EMsg_UCMUpdateOldScreenshotPrivacyResponse",
	7378: "EMsg_ClientUCMEnumerateUserSubscribedFilesWithUpdates",
	7379: "EMsg_ClientUCMEnumerateUserSubscribedFilesWithUpdatesResponse",
	7380: "EMsg_UCMPublishedFileContentUpdated",
	7381: "EMsg_UCMPublishedFileUpdated",
	7500: "EMsg_FSBase",
	7501: "EMsg_ClientRichPresenceUpload",
	7502: "EMsg_ClientRichPresenceRequest",
	7503: "EMsg_ClientRichPresenceInfo",
	7504: "EMsg_FSRichPresenceRequest",
	7505: "EMsg_FSRichPresenceResponse",
	7506: "EMsg_FSComputeFrenematrix",
	7507: "EMsg_FSComputeFrenematrixResponse",
	7508: "EMsg_FSPlayStatusNotification",
	7509: "EMsg_FSPublishPersonaStatus",
	7510: "EMsg_FSAddOrRemoveFollower",
	7511: "EMsg_FSAddOrRemoveFollowerResponse",
	7512: "EMsg_FSUpdateFollowingList",
	7513: "EMsg_FSCommentNotification",
	7514: "EMsg_FSCommentNotificationViewed",
	7515: "EMsg_ClientFSGetFollowerCount",
	7516: "EMsg_ClientFSGetFollowerCountResponse",
	7517: "EMsg_ClientFSGetIsFollowing",
	7518: "EMsg_ClientFSGetIsFollowingResponse",
	7519: "EMsg_ClientFSEnumerateFollowingList",
	7520: "EMsg_ClientFSEnumerateFollowingListResponse",
	7521: "EMsg_FSGetPendingNotificationCount",
	7522: "EMsg_FSGetPendingNotificationCountResponse",
	7523: "EMsg_ClientFSOfflineMessageNotification",
	7524: "EMsg_ClientFSRequestOfflineMessageCount",
	7525: "EMsg_ClientFSGetFriendMessageHistory",
	7526: "EMsg_ClientFSGetFriendMessageHistoryResponse",
	7527: "EMsg_ClientFSGetFriendMessageHistoryForOfflineMessages",
	7528: "EMsg_ClientFSGetFriendsSteamLevels",
	7529: "EMsg_ClientFSGetFriendsSteamLevelsResponse",
	7600: "EMsg_DRMRange2",
	7601: "EMsg_CEGVersionSetEnableDisableResponse",
	7602: "EMsg_CEGPropStatusDRMSRequest",
	7603: "EMsg_CEGPropStatusDRMSResponse",
	7604: "EMsg_CEGWhackFailureReportRequest",
	7605: "EMsg_CEGWhackFailureReportResponse",
	7606: "EMsg_DRMSFetchVersionSet",
	7607: "EMsg_DRMSFetchVersionSetResponse",
	7700: "EMsg_EconBase",
	7701: "EMsg_EconTrading_InitiateTradeRequest",
	7702: "EMsg_EconTrading_InitiateTradeProposed",
	7703: "EMsg_EconTrading_InitiateTradeResponse",
	7704: "EMsg_EconTrading_InitiateTradeResult",
	7705: "EMsg_EconTrading_StartSession",
	7706: "EMsg_EconTrading_CancelTradeRequest",
	7707: "EMsg_EconFlushInventoryCache",
	7708: "EMsg_EconFlushInventoryCacheResponse",
	7711: "EMsg_EconCDKeyProcessTransaction",
	7712: "EMsg_EconCDKeyProcessTransactionResponse",
	7713: "EMsg_EconGetErrorLogs",
	7714: "EMsg_EconGetErrorLogsResponse",
	7800: "EMsg_RMRange",
	7801: "EMsg_RMTestVerisignOTPResponse",
	7803: "EMsg_RMDeleteMemcachedKeys",
	7804: "EMsg_RMRemoteInvoke",
	7805: "EMsg_BadLoginIPList",
	7900: "EMsg_UGSBase",
	7901: "EMsg_ClientUGSGetGlobalStats",
	7902: "EMsg_ClientUGSGetGlobalStatsResponse",
	8000: "EMsg_StoreBase",
	8100: "EMsg_UMQBase",
	8101: "EMsg_UMQLogonResponse",
	8102: "EMsg_UMQLogoffRequest",
	8103: "EMsg_UMQLogoffResponse",
	8104: "EMsg_UMQSendChatMessage",
	8105: "EMsg_UMQIncomingChatMessage",
	8106: "EMsg_UMQPoll",
	8107: "EMsg_UMQPollResults",
	8108: "EMsg_UMQ2AM_ClientMsgBatch",
	8109: "EMsg_UMQEnqueueMobileSalePromotions",
	8110: "EMsg_UMQEnqueueMobileAnnouncements",
	8200: "EMsg_WorkshopBase",
	8201: "EMsg_WorkshopAcceptTOSResponse",
	8300: "EMsg_WebAPIBase",
	8301: "EMsg_WebAPIValidateOAuth2TokenResponse",
	8302: "EMsg_WebAPIInvalidateTokensForAccount",
	8303: "EMsg_WebAPIRegisterGCInterfaces",
	8304: "EMsg_WebAPIInvalidateOAuthClientCache",
	8305: "EMsg_WebAPIInvalidateOAuthTokenCache",
	8400: "EMsg_BackpackBase",
	8401: "EMsg_BackpackAddToCurrency",
	8402: "EMsg_BackpackAddToCurrencyResponse",
	8500: "EMsg_CREBase",
	8501: "EMsg_CRERankByTrend",
	8502: "EMsg_CRERankByTrendResponse",
	8503: "EMsg_CREItemVoteSummary",
	8504: "EMsg_CREItemVoteSummaryResponse",
	8505: "EMsg_CRERankByVote",
	8506: "EMsg_CRERankByVoteResponse",
	8507: "EMsg_CREUpdateUserPublishedItemVote",
	8508: "EMsg_CREUpdateUserPublishedItemVoteResponse",
	8509: "EMsg_CREGetUserPublishedItemVoteDetails",
	8510: "EMsg_CREGetUserPublishedItemVoteDetailsResponse",
	8511: "EMsg_CREEnumeratePublishedFiles",
	8512: "EMsg_CREEnumeratePublishedFilesResponse",
	8513: "EMsg_CREPublishedFileVoteAdded",
	8600: "EMsg_SecretsBase",
	8601: "EMsg_SecretsCredentialPairResponse",
	8602: "EMsg_SecretsRequestServerIdentity",
	8603: "EMsg_SecretsServerIdentityResponse",
	8604: "EMsg_SecretsUpdateServerIdentities",
	8700: "EMsg_BoxMonitorBase",
	8701: "EMsg_BoxMonitorReportResponse",
	8800: "EMsg_LogsinkBase",
	8900: "EMsg_PICSBase",
	8901: "EMsg_ClientPICSChangesSinceRequest",
	8902: "EMsg_ClientPICSChangesSinceResponse",
	8903: "EMsg_ClientPICSProductInfoRequest",
	8904: "EMsg_ClientPICSProductInfoResponse",
	8905: "EMsg_ClientPICSAccessTokenRequest",
	8906: "EMsg_ClientPICSAccessTokenResponse",
	9000: "EMsg_WorkerProcess",
	9001: "EMsg_WorkerProcessPingResponse",
	9002: "EMsg_WorkerProcessShutdown",
	9100: "EMsg_DRMWorkerProcess",
	9101: "EMsg_DRMWorkerProcessDRMAndSignResponse",
	9102: "EMsg_DRMWorkerProcessSteamworksInfoRequest",
	9103: "EMsg_DRMWorkerProcessSteamworksInfoResponse",
	9104: "EMsg_DRMWorkerProcessInstallDRMDLLRequest",
	9105: "EMsg_DRMWorkerProcessInstallDRMDLLResponse",
	9106: "EMsg_DRMWorkerProcessSecretIdStringRequest",
	9107: "EMsg_DRMWorkerProcessSecretIdStringResponse",
	9108: "EMsg_DRMWorkerProcessGetDRMGuidsFromFileRequest",
	9109: "EMsg_DRMWorkerProcessGetDRMGuidsFromFileResponse",
	9110: "EMsg_DRMWorkerProcessInstallProcessedFilesRequest",
	9111: "EMsg_DRMWorkerProcessInstallProcessedFilesResponse",
	9112: "EMsg_DRMWorkerProcessExamineBlobRequest",
	9113: "EMsg_DRMWorkerProcessExamineBlobResponse",
	9114: "EMsg_DRMWorkerProcessDescribeSecretRequest",
	9115: "EMsg_DRMWorkerProcessDescribeSecretResponse",
	9116: "EMsg_DRMWorkerProcessBackfillOriginalRequest",
	9117: "EMsg_DRMWorkerProcessBackfillOriginalResponse",
	9118: "EMsg_DRMWorkerProcessValidateDRMDLLRequest",
	9119: "EMsg_DRMWorkerProcessValidateDRMDLLResponse",
	9120: "EMsg_DRMWorkerProcessValidateFileRequest",
	9121: "EMsg_DRMWorkerProcessValidateFileResponse",
	9122: "EMsg_DRMWorkerProcessSplitAndInstallRequest",
	9123: "EMsg_DRMWorkerProcessSplitAndInstallResponse",
	9124: "EMsg_DRMWorkerProcessGetBlobRequest",
	9125: "EMsg_DRMWorkerProcessGetBlobResponse",
	9126: "EMsg_DRMWorkerProcessEvaluateCrashRequest",
	9127: "EMsg_DRMWorkerProcessEvaluateCrashResponse",
	9128: "EMsg_DRMWorkerProcessAnalyzeFileRequest",
	9129: "EMsg_DRMWorkerProcessAnalyzeFileResponse",
	9130: "EMsg_DRMWorkerProcessUnpackBlobRequest",
	9131: "EMsg_DRMWorkerProcessUnpackBlobResponse",
	9132: "EMsg_DRMWorkerProcessInstallAllRequest",
	9133: "EMsg_DRMWorkerProcessInstallAllResponse",
	9200: "EMsg_TestWorkerProcess",
	9201: "EMsg_TestWorkerProcessLoadUnloadModuleResponse",
	9202: "EMsg_TestWorkerProcessServiceModuleCallRequest",
	9203: "EMsg_TestWorkerProcessServiceModuleCallResponse",
	9330: "EMsg_ClientGetEmoticonList",
	9331: "EMsg_ClientEmoticonList",
	9400: "EMsg_ClientSharedLibraryBase",
	9403: "EMsg_ClientSharedLicensesLockStatus",
	9404: "EMsg_ClientSharedLicensesStopPlaying",
	9405: "EMsg_ClientSharedLibraryLockStatus",
	9406: "EMsg_ClientSharedLibraryStopPlaying",
	9507: "EMsg_ClientUnlockStreaming",
	9508: "EMsg_ClientUnlockStreamingResponse",
	9600: "EMsg_ClientPlayingSessionState",
	9601: "EMsg_ClientKickPlayingSession",
}

func (e EMsg) String() string {
	if s, ok := EMsg_name[e]; ok {
		return s
	}
	return "INVALID"
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

var EResult_name = map[EResult]string{
	0:  "EResult_Invalid",
	1:  "EResult_OK",
	2:  "EResult_Fail",
	3:  "EResult_NoConnection",
	5:  "EResult_InvalidPassword",
	6:  "EResult_LoggedInElsewhere",
	7:  "EResult_InvalidProtocolVer",
	8:  "EResult_InvalidParam",
	9:  "EResult_FileNotFound",
	10: "EResult_Busy",
	11: "EResult_InvalidState",
	12: "EResult_InvalidName",
	13: "EResult_InvalidEmail",
	14: "EResult_DuplicateName",
	15: "EResult_AccessDenied",
	16: "EResult_Timeout",
	17: "EResult_Banned",
	18: "EResult_AccountNotFound",
	19: "EResult_InvalidSteamID",
	20: "EResult_ServiceUnavailable",
	21: "EResult_NotLoggedOn",
	22: "EResult_Pending",
	23: "EResult_EncryptionFailure",
	24: "EResult_InsufficientPrivilege",
	25: "EResult_LimitExceeded",
	26: "EResult_Revoked",
	27: "EResult_Expired",
	28: "EResult_AlreadyRedeemed",
	29: "EResult_DuplicateRequest",
	30: "EResult_AlreadyOwned",
	31: "EResult_IPNotFound",
	32: "EResult_PersistFailed",
	33: "EResult_LockingFailed",
	34: "EResult_LogonSessionReplaced",
	35: "EResult_ConnectFailed",
	36: "EResult_HandshakeFailed",
	37: "EResult_IOFailure",
	38: "EResult_RemoteDisconnect",
	39: "EResult_ShoppingCartNotFound",
	40: "EResult_Blocked",
	41: "EResult_Ignored",
	42: "EResult_NoMatch",
	43: "EResult_AccountDisabled",
	44: "EResult_ServiceReadOnly",
	45: "EResult_AccountNotFeatured",
	46: "EResult_AdministratorOK",
	47: "EResult_ContentVersion",
	48: "EResult_TryAnotherCM",
	49: "EResult_PasswordRequiredToKickSession",
	50: "EResult_AlreadyLoggedInElsewhere",
	51: "EResult_Suspended",
	52: "EResult_Cancelled",
	53: "EResult_DataCorruption",
	54: "EResult_DiskFull",
	55: "EResult_RemoteCallFailed",
	56: "EResult_PasswordNotSet",
	57: "EResult_ExternalAccountUnlinked",
	58: "EResult_PSNTicketInvalid",
	59: "EResult_ExternalAccountAlreadyLinked",
	60: "EResult_RemoteFileConflict",
	61: "EResult_IllegalPassword",
	62: "EResult_SameAsPreviousValue",
	63: "EResult_AccountLogonDenied",
	64: "EResult_CannotUseOldPassword",
	65: "EResult_InvalidLoginAuthCode",
	66: "EResult_AccountLogonDeniedNoMailSent",
	67: "EResult_HardwareNotCapableOfIPT",
	68: "EResult_IPTInitError",
	69: "EResult_ParentalControlRestricted",
	70: "EResult_FacebookQueryError",
	71: "EResult_ExpiredLoginAuthCode",
	72: "EResult_IPLoginRestrictionFailed",
	73: "EResult_AccountLocked",
	74: "EResult_AccountLogonDeniedVerifiedEmailRequired",
	75: "EResult_NoMatchingURL",
	76: "EResult_BadResponse",
	77: "EResult_RequirePasswordReEntry",
	78: "EResult_ValueOutOfRange",
	79: "EResult_UnexpectedError",
	80: "EResult_Disabled",
	81: "EResult_InvalidCEGSubmission",
	82: "EResult_RestrictedDevice",
	83: "EResult_RegionLocked",
	84: "EResult_RateLimitExceeded",
	85: "EResult_AccountLogonDeniedNeedTwoFactorCode",
	86: "EResult_ItemOrEntryHasBeenDeleted",
}

func (e EResult) String() string {
	if s, ok := EResult_name[e]; ok {
		return s
	}
	return "INVALID"
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

var EUniverse_name = map[EUniverse]string{
	0: "EUniverse_Invalid",
	1: "EUniverse_Public",
	2: "EUniverse_Beta",
	3: "EUniverse_Internal",
	4: "EUniverse_Dev",
	5: "EUniverse_RC",
}

func (e EUniverse) String() string {
	if s, ok := EUniverse_name[e]; ok {
		return s
	}
	return "INVALID"
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

var EChatEntryType_name = map[EChatEntryType]string{
	0:  "EChatEntryType_Invalid",
	1:  "EChatEntryType_ChatMsg",
	2:  "EChatEntryType_Typing",
	3:  "EChatEntryType_InviteGame",
	4:  "EChatEntryType_Emote",
	5:  "EChatEntryType_LobbyGameStart",
	6:  "EChatEntryType_LeftConversation",
	7:  "EChatEntryType_Entered",
	8:  "EChatEntryType_WasKicked",
	9:  "EChatEntryType_WasBanned",
	10: "EChatEntryType_Disconnected",
	11: "EChatEntryType_HistoricalChat",
}

func (e EChatEntryType) String() string {
	if s, ok := EChatEntryType_name[e]; ok {
		return s
	}
	return "INVALID"
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

var EPersonaState_name = map[EPersonaState]string{
	0: "EPersonaState_Offline",
	1: "EPersonaState_Online",
	2: "EPersonaState_Busy",
	3: "EPersonaState_Away",
	4: "EPersonaState_Snooze",
	5: "EPersonaState_LookingToTrade",
	6: "EPersonaState_LookingToPlay",
	7: "EPersonaState_Max",
}

func (e EPersonaState) String() string {
	if s, ok := EPersonaState_name[e]; ok {
		return s
	}
	return "INVALID"
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

var EAccountType_name = map[EAccountType]string{
	0:  "EAccountType_Invalid",
	1:  "EAccountType_Individual",
	2:  "EAccountType_Multiseat",
	3:  "EAccountType_GameServer",
	4:  "EAccountType_AnonGameServer",
	5:  "EAccountType_Pending",
	6:  "EAccountType_ContentServer",
	7:  "EAccountType_Clan",
	8:  "EAccountType_Chat",
	9:  "EAccountType_ConsoleUser",
	10: "EAccountType_AnonUser",
	11: "EAccountType_Max",
}

func (e EAccountType) String() string {
	if s, ok := EAccountType_name[e]; ok {
		return s
	}
	return "INVALID"
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

var EFriendRelationship_name = map[EFriendRelationship]string{
	0: "EFriendRelationship_None",
	1: "EFriendRelationship_Blocked",
	2: "EFriendRelationship_PendingInvitee",
	3: "EFriendRelationship_Friend",
	4: "EFriendRelationship_RequestInitiator",
	5: "EFriendRelationship_Ignored",
	6: "EFriendRelationship_IgnoredFriend",
	7: "EFriendRelationship_SuggestedFriend",
	8: "EFriendRelationship_Max",
}

func (e EFriendRelationship) String() string {
	if s, ok := EFriendRelationship_name[e]; ok {
		return s
	}
	return "INVALID"
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

var EAccountFlags_name = map[EAccountFlags]string{
	0:         "EAccountFlags_NormalUser",
	1:         "EAccountFlags_PersonaNameSet",
	2:         "EAccountFlags_Unbannable",
	4:         "EAccountFlags_PasswordSet",
	8:         "EAccountFlags_Support",
	16:        "EAccountFlags_Admin",
	32:        "EAccountFlags_Supervisor",
	64:        "EAccountFlags_AppEditor",
	128:       "EAccountFlags_HWIDSet",
	256:       "EAccountFlags_PersonalQASet",
	512:       "EAccountFlags_VacBeta",
	1024:      "EAccountFlags_Debug",
	2048:      "EAccountFlags_Disabled",
	4096:      "EAccountFlags_LimitedUser",
	8192:      "EAccountFlags_LimitedUserForce",
	16384:     "EAccountFlags_EmailValidated",
	32768:     "EAccountFlags_MarketingTreatment",
	65536:     "EAccountFlags_OGGInviteOptOut",
	131072:    "EAccountFlags_ForcePasswordChange",
	262144:    "EAccountFlags_ForceEmailVerification",
	524288:    "EAccountFlags_LogonExtraSecurity",
	1048576:   "EAccountFlags_LogonExtraSecurityDisabled",
	2097152:   "EAccountFlags_Steam2MigrationComplete",
	4194304:   "EAccountFlags_NeedLogs",
	8388608:   "EAccountFlags_Lockdown",
	16777216:  "EAccountFlags_MasterAppEditor",
	33554432:  "EAccountFlags_BannedFromWebAPI",
	67108864:  "EAccountFlags_ClansOnlyFromFriends",
	134217728: "EAccountFlags_GlobalModerator",
}

func (e EAccountFlags) String() string {
	if s, ok := EAccountFlags_name[e]; ok {
		return s
	}
	return "INVALID"
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

var EClanPermission_name = map[EClanPermission]string{
	0:   "EClanPermission_Nobody",
	1:   "EClanPermission_Owner",
	2:   "EClanPermission_Officer",
	3:   "EClanPermission_OwnerAndOfficer",
	4:   "EClanPermission_Member",
	8:   "EClanPermission_Moderator",
	11:  "EClanPermission_OwnerOfficerModerator",
	15:  "EClanPermission_AllMembers",
	16:  "EClanPermission_OGGGameOwner",
	128: "EClanPermission_NonMember",
	132: "EClanPermission_MemberAllowed",
	140: "EClanPermission_ModeratorAllowed",
	142: "EClanPermission_OfficerAllowed",
	143: "EClanPermission_OwnerAllowed",
}

func (e EClanPermission) String() string {
	if s, ok := EClanPermission_name[e]; ok {
		return s
	}
	return "INVALID"
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

var EChatPermission_name = map[EChatPermission]string{
	1:    "EChatPermission_Close",
	2:    "EChatPermission_Invite",
	8:    "EChatPermission_Talk",
	16:   "EChatPermission_Kick",
	32:   "EChatPermission_Mute",
	64:   "EChatPermission_SetMetadata",
	128:  "EChatPermission_ChangePermissions",
	256:  "EChatPermission_Ban",
	512:  "EChatPermission_ChangeAccess",
	10:   "EChatPermission_EveryoneDefault",
	282:  "EChatPermission_MemberDefault",
	891:  "EChatPermission_OwnerDefault",
	1019: "EChatPermission_Mask",
}

func (e EChatPermission) String() string {
	if s, ok := EChatPermission_name[e]; ok {
		return s
	}
	return "INVALID"
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

var EFriendFlags_name = map[EFriendFlags]string{
	0:     "EFriendFlags_None",
	1:     "EFriendFlags_Blocked",
	2:     "EFriendFlags_FriendshipRequested",
	4:     "EFriendFlags_Immediate",
	8:     "EFriendFlags_ClanMember",
	16:    "EFriendFlags_GameServer",
	128:   "EFriendFlags_RequestingFriendship",
	256:   "EFriendFlags_RequestingInfo",
	512:   "EFriendFlags_Ignored",
	1024:  "EFriendFlags_IgnoredFriend",
	2048:  "EFriendFlags_Suggested",
	65535: "EFriendFlags_FlagAll",
}

func (e EFriendFlags) String() string {
	if s, ok := EFriendFlags_name[e]; ok {
		return s
	}
	return "INVALID"
}

type EPersonaStateFlag int32

const (
	EPersonaStateFlag_HasRichPresence       EPersonaStateFlag = 1
	EPersonaStateFlag_InJoinableGame                          = 2
	EPersonaStateFlag_OnlineUsingWeb                          = 256
	EPersonaStateFlag_OnlineUsingMobile                       = 512
	EPersonaStateFlag_OnlineUsingBigPicture                   = 1024
)

var EPersonaStateFlag_name = map[EPersonaStateFlag]string{
	1:    "EPersonaStateFlag_HasRichPresence",
	2:    "EPersonaStateFlag_InJoinableGame",
	256:  "EPersonaStateFlag_OnlineUsingWeb",
	512:  "EPersonaStateFlag_OnlineUsingMobile",
	1024: "EPersonaStateFlag_OnlineUsingBigPicture",
}

func (e EPersonaStateFlag) String() string {
	if s, ok := EPersonaStateFlag_name[e]; ok {
		return s
	}
	return "INVALID"
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

var EClientPersonaStateFlag_name = map[EClientPersonaStateFlag]string{
	1:    "EClientPersonaStateFlag_Status",
	2:    "EClientPersonaStateFlag_PlayerName",
	4:    "EClientPersonaStateFlag_QueryPort",
	8:    "EClientPersonaStateFlag_SourceID",
	16:   "EClientPersonaStateFlag_Presence",
	32:   "EClientPersonaStateFlag_Metadata",
	64:   "EClientPersonaStateFlag_LastSeen",
	128:  "EClientPersonaStateFlag_ClanInfo",
	256:  "EClientPersonaStateFlag_GameExtraInfo",
	512:  "EClientPersonaStateFlag_GameDataBlob",
	1024: "EClientPersonaStateFlag_ClanTag",
	2048: "EClientPersonaStateFlag_Facebook",
}

func (e EClientPersonaStateFlag) String() string {
	if s, ok := EClientPersonaStateFlag_name[e]; ok {
		return s
	}
	return "INVALID"
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

var EAppUsageEvent_name = map[EAppUsageEvent]string{
	1: "EAppUsageEvent_GameLaunch",
	2: "EAppUsageEvent_GameLaunchTrial",
	3: "EAppUsageEvent_Media",
	4: "EAppUsageEvent_PreloadStart",
	5: "EAppUsageEvent_PreloadFinish",
	6: "EAppUsageEvent_MarketingMessageView",
	7: "EAppUsageEvent_InGameAdViewed",
	8: "EAppUsageEvent_GameLaunchFreeWeekend",
}

func (e EAppUsageEvent) String() string {
	if s, ok := EAppUsageEvent_name[e]; ok {
		return s
	}
	return "INVALID"
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

var ELicenseFlags_name = map[ELicenseFlags]string{
	0:   "ELicenseFlags_None",
	1:   "ELicenseFlags_Renew",
	2:   "ELicenseFlags_RenewalFailed",
	4:   "ELicenseFlags_Pending",
	8:   "ELicenseFlags_Expired",
	16:  "ELicenseFlags_CancelledByUser",
	32:  "ELicenseFlags_CancelledByAdmin",
	64:  "ELicenseFlags_LowViolenceContent",
	128: "ELicenseFlags_ImportedFromSteam2",
}

func (e ELicenseFlags) String() string {
	if s, ok := ELicenseFlags_name[e]; ok {
		return s
	}
	return "INVALID"
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

var ELicenseType_name = map[ELicenseType]string{
	0: "ELicenseType_NoLicense",
	1: "ELicenseType_SinglePurchase",
	2: "ELicenseType_SinglePurchaseLimitedUse",
	3: "ELicenseType_RecurringCharge",
	4: "ELicenseType_RecurringChargeLimitedUse",
	5: "ELicenseType_RecurringChargeLimitedUseWithOverages",
	6: "ELicenseType_RecurringOption",
}

func (e ELicenseType) String() string {
	if s, ok := ELicenseType_name[e]; ok {
		return s
	}
	return "INVALID"
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

var EPaymentMethod_name = map[EPaymentMethod]string{
	0:    "EPaymentMethod_None",
	1:    "EPaymentMethod_ActivationCode",
	2:    "EPaymentMethod_CreditCard",
	3:    "EPaymentMethod_Giropay",
	4:    "EPaymentMethod_PayPal",
	5:    "EPaymentMethod_Ideal",
	6:    "EPaymentMethod_PaySafeCard",
	7:    "EPaymentMethod_Sofort",
	8:    "EPaymentMethod_GuestPass",
	9:    "EPaymentMethod_WebMoney",
	10:   "EPaymentMethod_MoneyBookers",
	11:   "EPaymentMethod_AliPay",
	12:   "EPaymentMethod_Yandex",
	13:   "EPaymentMethod_Kiosk",
	14:   "EPaymentMethod_Qiwi",
	15:   "EPaymentMethod_GameStop",
	16:   "EPaymentMethod_HardwarePromo",
	17:   "EPaymentMethod_MoPay",
	18:   "EPaymentMethod_BoletoBancario",
	19:   "EPaymentMethod_BoaCompraGold",
	20:   "EPaymentMethod_BancoDoBrasilOnline",
	21:   "EPaymentMethod_ItauOnline",
	22:   "EPaymentMethod_BradescoOnline",
	23:   "EPaymentMethod_Pagseguro",
	24:   "EPaymentMethod_VisaBrazil",
	25:   "EPaymentMethod_AmexBrazil",
	26:   "EPaymentMethod_Aura",
	27:   "EPaymentMethod_Hipercard",
	28:   "EPaymentMethod_MastercardBrazil",
	29:   "EPaymentMethod_DinersCardBrazil",
	32:   "EPaymentMethod_ClickAndBuy",
	64:   "EPaymentMethod_AutoGrant",
	128:  "EPaymentMethod_Wallet",
	129:  "EPaymentMethod_Valve",
	256:  "EPaymentMethod_OEMTicket",
	512:  "EPaymentMethod_Split",
	1024: "EPaymentMethod_Complimentary",
}

func (e EPaymentMethod) String() string {
	if s, ok := EPaymentMethod_name[e]; ok {
		return s
	}
	return "INVALID"
}

type EIntroducerRouting int32

const (
	EIntroducerRouting_FileShare     EIntroducerRouting = 0 // Deprecated
	EIntroducerRouting_P2PVoiceChat                     = 1
	EIntroducerRouting_P2PNetworking                    = 2
)

var EIntroducerRouting_name = map[EIntroducerRouting]string{
	0: "EIntroducerRouting_FileShare",
	1: "EIntroducerRouting_P2PVoiceChat",
	2: "EIntroducerRouting_P2PNetworking",
}

func (e EIntroducerRouting) String() string {
	if s, ok := EIntroducerRouting_name[e]; ok {
		return s
	}
	return "INVALID"
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

var EServerFlags_name = map[EServerFlags]string{
	0:  "EServerFlags_None",
	1:  "EServerFlags_Active",
	2:  "EServerFlags_Secure",
	4:  "EServerFlags_Dedicated",
	8:  "EServerFlags_Linux",
	16: "EServerFlags_Passworded",
	32: "EServerFlags_Private",
}

func (e EServerFlags) String() string {
	if s, ok := EServerFlags_name[e]; ok {
		return s
	}
	return "INVALID"
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

var EDenyReason_name = map[EDenyReason]string{
	1:  "EDenyReason_InvalidVersion",
	2:  "EDenyReason_Generic",
	3:  "EDenyReason_NotLoggedOn",
	4:  "EDenyReason_NoLicense",
	5:  "EDenyReason_Cheater",
	6:  "EDenyReason_LoggedInElseWhere",
	7:  "EDenyReason_UnknownText",
	8:  "EDenyReason_IncompatibleAnticheat",
	9:  "EDenyReason_MemoryCorruption",
	10: "EDenyReason_IncompatibleSoftware",
	11: "EDenyReason_SteamConnectionLost",
	12: "EDenyReason_SteamConnectionError",
	13: "EDenyReason_SteamResponseTimedOut",
	14: "EDenyReason_SteamValidationStalled",
	15: "EDenyReason_SteamOwnerLeftGuestUser",
}

func (e EDenyReason) String() string {
	if s, ok := EDenyReason_name[e]; ok {
		return s
	}
	return "INVALID"
}

type EClanRank int32

const (
	EClanRank_None      EClanRank = 0
	EClanRank_Owner               = 1
	EClanRank_Officer             = 2
	EClanRank_Member              = 3
	EClanRank_Moderator           = 4
)

var EClanRank_name = map[EClanRank]string{
	0: "EClanRank_None",
	1: "EClanRank_Owner",
	2: "EClanRank_Officer",
	3: "EClanRank_Member",
	4: "EClanRank_Moderator",
}

func (e EClanRank) String() string {
	if s, ok := EClanRank_name[e]; ok {
		return s
	}
	return "INVALID"
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

var EClanRelationship_name = map[EClanRelationship]string{
	0: "EClanRelationship_None",
	1: "EClanRelationship_Blocked",
	2: "EClanRelationship_Invited",
	3: "EClanRelationship_Member",
	4: "EClanRelationship_Kicked",
	5: "EClanRelationship_KickAcknowledged",
}

func (e EClanRelationship) String() string {
	if s, ok := EClanRelationship_name[e]; ok {
		return s
	}
	return "INVALID"
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

var EAuthSessionResponse_name = map[EAuthSessionResponse]string{
	0: "EAuthSessionResponse_OK",
	1: "EAuthSessionResponse_UserNotConnectedToSteam",
	2: "EAuthSessionResponse_NoLicenseOrExpired",
	3: "EAuthSessionResponse_VACBanned",
	4: "EAuthSessionResponse_LoggedInElseWhere",
	5: "EAuthSessionResponse_VACCheckTimedOut",
	6: "EAuthSessionResponse_AuthTicketCanceled",
	7: "EAuthSessionResponse_AuthTicketInvalidAlreadyUsed",
	8: "EAuthSessionResponse_AuthTicketInvalid",
}

func (e EAuthSessionResponse) String() string {
	if s, ok := EAuthSessionResponse_name[e]; ok {
		return s
	}
	return "INVALID"
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

var EChatRoomEnterResponse_name = map[EChatRoomEnterResponse]string{
	1:  "EChatRoomEnterResponse_Success",
	2:  "EChatRoomEnterResponse_DoesntExist",
	3:  "EChatRoomEnterResponse_NotAllowed",
	4:  "EChatRoomEnterResponse_Full",
	5:  "EChatRoomEnterResponse_Error",
	6:  "EChatRoomEnterResponse_Banned",
	7:  "EChatRoomEnterResponse_Limited",
	8:  "EChatRoomEnterResponse_ClanDisabled",
	9:  "EChatRoomEnterResponse_CommunityBan",
	10: "EChatRoomEnterResponse_MemberBlockedYou",
	11: "EChatRoomEnterResponse_YouBlockedMember",
	12: "EChatRoomEnterResponse_NoRankingDataLobby",
	13: "EChatRoomEnterResponse_NoRankingDataUser",
	14: "EChatRoomEnterResponse_RankOutOfRange",
}

func (e EChatRoomEnterResponse) String() string {
	if s, ok := EChatRoomEnterResponse_name[e]; ok {
		return s
	}
	return "INVALID"
}

type EChatRoomType int32

const (
	EChatRoomType_Friend EChatRoomType = 1
	EChatRoomType_MUC                  = 2
	EChatRoomType_Lobby                = 3
)

var EChatRoomType_name = map[EChatRoomType]string{
	1: "EChatRoomType_Friend",
	2: "EChatRoomType_MUC",
	3: "EChatRoomType_Lobby",
}

func (e EChatRoomType) String() string {
	if s, ok := EChatRoomType_name[e]; ok {
		return s
	}
	return "INVALID"
}

type EChatInfoType int32

const (
	EChatInfoType_StateChange       EChatInfoType = 1
	EChatInfoType_InfoUpdate                      = 2
	EChatInfoType_MemberLimitChange               = 3
)

var EChatInfoType_name = map[EChatInfoType]string{
	1: "EChatInfoType_StateChange",
	2: "EChatInfoType_InfoUpdate",
	3: "EChatInfoType_MemberLimitChange",
}

func (e EChatInfoType) String() string {
	if s, ok := EChatInfoType_name[e]; ok {
		return s
	}
	return "INVALID"
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

var EChatAction_name = map[EChatAction]string{
	1:  "EChatAction_InviteChat",
	2:  "EChatAction_Kick",
	3:  "EChatAction_Ban",
	4:  "EChatAction_UnBan",
	5:  "EChatAction_StartVoiceSpeak",
	6:  "EChatAction_EndVoiceSpeak",
	7:  "EChatAction_LockChat",
	8:  "EChatAction_UnlockChat",
	9:  "EChatAction_CloseChat",
	10: "EChatAction_SetJoinable",
	11: "EChatAction_SetUnjoinable",
	12: "EChatAction_SetOwner",
	13: "EChatAction_SetInvisibleToFriends",
	14: "EChatAction_SetVisibleToFriends",
	15: "EChatAction_SetModerated",
	16: "EChatAction_SetUnmoderated",
}

func (e EChatAction) String() string {
	if s, ok := EChatAction_name[e]; ok {
		return s
	}
	return "INVALID"
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

var EChatActionResult_name = map[EChatActionResult]string{
	1:  "EChatActionResult_Success",
	2:  "EChatActionResult_Error",
	3:  "EChatActionResult_NotPermitted",
	4:  "EChatActionResult_NotAllowedOnClanMember",
	5:  "EChatActionResult_NotAllowedOnBannedUser",
	6:  "EChatActionResult_NotAllowedOnChatOwner",
	7:  "EChatActionResult_NotAllowedOnSelf",
	8:  "EChatActionResult_ChatDoesntExist",
	9:  "EChatActionResult_ChatFull",
	10: "EChatActionResult_VoiceSlotsFull",
}

func (e EChatActionResult) String() string {
	if s, ok := EChatActionResult_name[e]; ok {
		return s
	}
	return "INVALID"
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

var EAppInfoSection_name = map[EAppInfoSection]string{
	0:  "EAppInfoSection_Unknown",
	1:  "EAppInfoSection_All",
	2:  "EAppInfoSection_First",
	3:  "EAppInfoSection_Extended",
	4:  "EAppInfoSection_Config",
	5:  "EAppInfoSection_Stats",
	6:  "EAppInfoSection_Install",
	7:  "EAppInfoSection_Depots",
	8:  "EAppInfoSection_VAC",
	9:  "EAppInfoSection_DRM",
	10: "EAppInfoSection_UFS",
	11: "EAppInfoSection_OGG",
	12: "EAppInfoSection_Items",
	13: "EAppInfoSection_Policies",
	14: "EAppInfoSection_SysReqs",
	15: "EAppInfoSection_Community",
	16: "EAppInfoSection_Max",
}

func (e EAppInfoSection) String() string {
	if s, ok := EAppInfoSection_name[e]; ok {
		return s
	}
	return "INVALID"
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

var EContentDownloadSourceType_name = map[EContentDownloadSourceType]string{
	0: "EContentDownloadSourceType_Invalid",
	1: "EContentDownloadSourceType_CS",
	2: "EContentDownloadSourceType_CDN",
	3: "EContentDownloadSourceType_LCS",
	4: "EContentDownloadSourceType_Proxy",
	5: "EContentDownloadSourceType_Max",
}

func (e EContentDownloadSourceType) String() string {
	if s, ok := EContentDownloadSourceType_name[e]; ok {
		return s
	}
	return "INVALID"
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

var EPlatformType_name = map[EPlatformType]string{
	0: "EPlatformType_Unknown",
	1: "EPlatformType_Win32",
	2: "EPlatformType_Win64",
	3: "EPlatformType_Linux",
	4: "EPlatformType_OSX",
	5: "EPlatformType_PS3",
	6: "EPlatformType_Max",
}

func (e EPlatformType) String() string {
	if s, ok := EPlatformType_name[e]; ok {
		return s
	}
	return "INVALID"
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

var EOSType_name = map[EOSType]string{
	-1:   "EOSType_Unknown",
	-400: "EOSType_UMQ",
	-300: "EOSType_PS3",
	-102: "EOSType_MacOSUnknown",
	-101: "EOSType_MacOS104",
	-100: "EOSType_MacOS105",
	-99:  "EOSType_MacOS1058",
	-95:  "EOSType_MacOS106",
	-94:  "EOSType_MacOS1063",
	-93:  "EOSType_MacOS1064_slgu",
	-92:  "EOSType_MacOS1067",
	-90:  "EOSType_MacOS107",
	-89:  "EOSType_MacOS108",
	-88:  "EOSType_MacOS109",
	-203: "EOSType_LinuxUnknown",
	-202: "EOSType_Linux22",
	-201: "EOSType_Linux24",
	-200: "EOSType_Linux26",
	-199: "EOSType_Linux32",
	-198: "EOSType_Linux35",
	-197: "EOSType_Linux36",
	-196: "EOSType_Linux310",
	0:    "EOSType_WinUnknown",
	1:    "EOSType_Win311",
	2:    "EOSType_Win95",
	3:    "EOSType_Win98",
	4:    "EOSType_WinME",
	5:    "EOSType_WinNT",
	6:    "EOSType_Win200",
	7:    "EOSType_WinXP",
	8:    "EOSType_Win2003",
	9:    "EOSType_WinVista",
	10:   "EOSType_Win7",
	11:   "EOSType_Win2008",
	12:   "EOSType_Win2012",
	13:   "EOSType_Windows8",
	14:   "EOSType_Windows81",
	15:   "EOSType_WinMAX",
	26:   "EOSType_Max",
}

func (e EOSType) String() string {
	if s, ok := EOSType_name[e]; ok {
		return s
	}
	return "INVALID"
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

var EServerType_name = map[EServerType]string{
	-1: "EServerType_Invalid",
	0:  "EServerType_First",
	1:  "EServerType_GM",
	2:  "EServerType_BUM",
	3:  "EServerType_AM",
	4:  "EServerType_BS",
	5:  "EServerType_VS",
	6:  "EServerType_ATS",
	7:  "EServerType_CM",
	8:  "EServerType_FBS",
	9:  "EServerType_FG",
	10: "EServerType_SS",
	11: "EServerType_DRMS",
	12: "EServerType_HubOBSOLETE",
	13: "EServerType_Console",
	14: "EServerType_ASBOBSOLETE",
	15: "EServerType_Client",
	16: "EServerType_BootstrapOBSOLETE",
	17: "EServerType_DP",
	18: "EServerType_WG",
	19: "EServerType_SM",
	21: "EServerType_UFS",
	23: "EServerType_Util",
	24: "EServerType_DSS",
	25: "EServerType_P2PRelayOBSOLETE",
	26: "EServerType_AppInformation",
	27: "EServerType_Spare",
	28: "EServerType_FTS",
	29: "EServerType_EPM",
	30: "EServerType_PS",
	31: "EServerType_IS",
	32: "EServerType_CCS",
	33: "EServerType_DFS",
	34: "EServerType_LBS",
	35: "EServerType_MDS",
	36: "EServerType_CS",
	37: "EServerType_GC",
	38: "EServerType_NS",
	39: "EServerType_OGS",
	40: "EServerType_WebAPI",
	41: "EServerType_UDS",
	42: "EServerType_MMS",
	43: "EServerType_GMS",
	44: "EServerType_KGS",
	45: "EServerType_UCM",
	46: "EServerType_RM",
	47: "EServerType_FS",
	48: "EServerType_Econ",
	49: "EServerType_Backpack",
	50: "EServerType_UGS",
	51: "EServerType_Store",
	52: "EServerType_MoneyStats",
	53: "EServerType_CRE",
	54: "EServerType_UMQ",
	55: "EServerType_Workshop",
	56: "EServerType_BRP",
	57: "EServerType_GCH",
	58: "EServerType_MPAS",
	59: "EServerType_Trade",
	60: "EServerType_Secrets",
	61: "EServerType_Logsink",
	62: "EServerType_Market",
	63: "EServerType_Quest",
	64: "EServerType_WDS",
	65: "EServerType_ACS",
	66: "EServerType_PNP",
	67: "EServerType_Max",
}

func (e EServerType) String() string {
	if s, ok := EServerType_name[e]; ok {
		return s
	}
	return "INVALID"
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

var EBillingType_name = map[EBillingType]string{
	0:  "EBillingType_NoCost",
	1:  "EBillingType_BillOnceOnly",
	2:  "EBillingType_BillMonthly",
	3:  "EBillingType_ProofOfPrepurchaseOnly",
	4:  "EBillingType_GuestPass",
	5:  "EBillingType_HardwarePromo",
	6:  "EBillingType_Gift",
	7:  "EBillingType_AutoGrant",
	8:  "EBillingType_OEMTicket",
	9:  "EBillingType_RecurringOption",
	10: "EBillingType_NumBillingTypes",
}

func (e EBillingType) String() string {
	if s, ok := EBillingType_name[e]; ok {
		return s
	}
	return "INVALID"
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

var EActivationCodeClass_name = map[EActivationCodeClass]string{
	0:          "EActivationCodeClass_WonCDKey",
	1:          "EActivationCodeClass_ValveCDKey",
	2:          "EActivationCodeClass_Doom3CDKey",
	3:          "EActivationCodeClass_DBLookup",
	4:          "EActivationCodeClass_Steam2010Key",
	5:          "EActivationCodeClass_Max",
	2147483647: "EActivationCodeClass_Test",
	4294967295: "EActivationCodeClass_Invalid",
}

func (e EActivationCodeClass) String() string {
	if s, ok := EActivationCodeClass_name[e]; ok {
		return s
	}
	return "INVALID"
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

var EChatMemberStateChange_name = map[EChatMemberStateChange]string{
	1:    "EChatMemberStateChange_Entered",
	2:    "EChatMemberStateChange_Left",
	4:    "EChatMemberStateChange_Disconnected",
	8:    "EChatMemberStateChange_Kicked",
	16:   "EChatMemberStateChange_Banned",
	4096: "EChatMemberStateChange_VoiceSpeaking",
	8192: "EChatMemberStateChange_VoiceDoneSpeaking",
}

func (e EChatMemberStateChange) String() string {
	if s, ok := EChatMemberStateChange_name[e]; ok {
		return s
	}
	return "INVALID"
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

var ERegionCode_name = map[ERegionCode]string{
	0:   "ERegionCode_USEast",
	1:   "ERegionCode_USWest",
	2:   "ERegionCode_SouthAmerica",
	3:   "ERegionCode_Europe",
	4:   "ERegionCode_Asia",
	5:   "ERegionCode_Australia",
	6:   "ERegionCode_MiddleEast",
	7:   "ERegionCode_Africa",
	255: "ERegionCode_World",
}

func (e ERegionCode) String() string {
	if s, ok := ERegionCode_name[e]; ok {
		return s
	}
	return "INVALID"
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

var ECurrencyCode_name = map[ECurrencyCode]string{
	0:  "ECurrencyCode_Invalid",
	1:  "ECurrencyCode_USD",
	2:  "ECurrencyCode_GBP",
	3:  "ECurrencyCode_EUR",
	4:  "ECurrencyCode_CHF",
	5:  "ECurrencyCode_RUB",
	6:  "ECurrencyCode_PLN",
	7:  "ECurrencyCode_BRL",
	9:  "ECurrencyCode_NOK",
	10: "ECurrencyCode_Max",
}

func (e ECurrencyCode) String() string {
	if s, ok := ECurrencyCode_name[e]; ok {
		return s
	}
	return "INVALID"
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

var EDepotFileFlag_name = map[EDepotFileFlag]string{
	1:   "EDepotFileFlag_UserConfig",
	2:   "EDepotFileFlag_VersionedUserConfig",
	4:   "EDepotFileFlag_Encrypted",
	8:   "EDepotFileFlag_ReadOnly",
	16:  "EDepotFileFlag_Hidden",
	32:  "EDepotFileFlag_Executable",
	64:  "EDepotFileFlag_Directory",
	128: "EDepotFileFlag_CustomExecutable",
	256: "EDepotFileFlag_InstallScript",
}

func (e EDepotFileFlag) String() string {
	if s, ok := EDepotFileFlag_name[e]; ok {
		return s
	}
	return "INVALID"
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

var EWorkshopEnumerationType_name = map[EWorkshopEnumerationType]string{
	0: "EWorkshopEnumerationType_RankedByVote",
	1: "EWorkshopEnumerationType_Recent",
	2: "EWorkshopEnumerationType_Trending",
	3: "EWorkshopEnumerationType_FavoriteOfFriends",
	4: "EWorkshopEnumerationType_VotedByFriends",
	5: "EWorkshopEnumerationType_ContentByFriends",
	6: "EWorkshopEnumerationType_RecentFromFollowedUsers",
}

func (e EWorkshopEnumerationType) String() string {
	if s, ok := EWorkshopEnumerationType_name[e]; ok {
		return s
	}
	return "INVALID"
}

type EPublishedFileVisibility int32

const (
	EPublishedFileVisibility_Public      EPublishedFileVisibility = 0
	EPublishedFileVisibility_FriendsOnly                          = 1
	EPublishedFileVisibility_Private                              = 2
)

var EPublishedFileVisibility_name = map[EPublishedFileVisibility]string{
	0: "EPublishedFileVisibility_Public",
	1: "EPublishedFileVisibility_FriendsOnly",
	2: "EPublishedFileVisibility_Private",
}

func (e EPublishedFileVisibility) String() string {
	if s, ok := EPublishedFileVisibility_name[e]; ok {
		return s
	}
	return "INVALID"
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

var EWorkshopFileType_name = map[EWorkshopFileType]string{
	0:  "EWorkshopFileType_First",
	1:  "EWorkshopFileType_Microtransaction",
	2:  "EWorkshopFileType_Collection",
	3:  "EWorkshopFileType_Art",
	4:  "EWorkshopFileType_Video",
	5:  "EWorkshopFileType_Screenshot",
	6:  "EWorkshopFileType_Game",
	7:  "EWorkshopFileType_Software",
	8:  "EWorkshopFileType_Concept",
	9:  "EWorkshopFileType_WebGuide",
	10: "EWorkshopFileType_IntegratedGuide",
	11: "EWorkshopFileType_Merch",
	12: "EWorkshopFileType_ControllerBinding",
	13: "EWorkshopFileType_SteamworksAccessInvite",
	14: "EWorkshopFileType_Max",
}

func (e EWorkshopFileType) String() string {
	if s, ok := EWorkshopFileType_name[e]; ok {
		return s
	}
	return "INVALID"
}

type EWorkshopFileAction int32

const (
	EWorkshopFileAction_Played    EWorkshopFileAction = 0
	EWorkshopFileAction_Completed                     = 1
)

var EWorkshopFileAction_name = map[EWorkshopFileAction]string{
	0: "EWorkshopFileAction_Played",
	1: "EWorkshopFileAction_Completed",
}

func (e EWorkshopFileAction) String() string {
	if s, ok := EWorkshopFileAction_name[e]; ok {
		return s
	}
	return "INVALID"
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

var EEconTradeResponse_name = map[EEconTradeResponse]string{
	0:  "EEconTradeResponse_Accepted",
	1:  "EEconTradeResponse_Declined",
	2:  "EEconTradeResponse_VacBannedInitiator",
	3:  "EEconTradeResponse_VacBannedTarget",
	4:  "EEconTradeResponse_TargetAlreadyTrading",
	5:  "EEconTradeResponse_Disabled",
	6:  "EEconTradeResponse_NotLoggedIn",
	7:  "EEconTradeResponse_Cancel",
	8:  "EEconTradeResponse_TooSoon",
	9:  "EEconTradeResponse_TooSoonPenalty",
	10: "EEconTradeResponse_ConnectionFailed",
	11: "EEconTradeResponse_InitiatorAlreadyTrading",
	12: "EEconTradeResponse_Error",
	13: "EEconTradeResponse_Timeout",
	14: "EEconTradeResponse_CyberCafeInitiator",
	15: "EEconTradeResponse_CyberCafeTarget",
	16: "EEconTradeResponse_SchoolLabInitiator",
	18: "EEconTradeResponse_InitiatorBlockedTarget",
	20: "EEconTradeResponse_InitiatorNeedsVerifiedEmail",
	21: "EEconTradeResponse_InitiatorNeedsSteamGuard",
	22: "EEconTradeResponse_TargetAccountCannotTrade",
	23: "EEconTradeResponse_InitiatorSteamGuardDuration",
	24: "EEconTradeResponse_InitiatorPasswordResetProbation",
	25: "EEconTradeResponse_InitiatorNewDeviceCooldown",
	50: "EEconTradeResponse_OKToDeliver",
}

func (e EEconTradeResponse) String() string {
	if s, ok := EEconTradeResponse_name[e]; ok {
		return s
	}
	return "INVALID"
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

var EMarketingMessageFlags_name = map[EMarketingMessageFlags]string{
	0:  "EMarketingMessageFlags_None",
	1:  "EMarketingMessageFlags_HighPriority",
	2:  "EMarketingMessageFlags_PlatformWindows",
	4:  "EMarketingMessageFlags_PlatformMac",
	8:  "EMarketingMessageFlags_PlatformLinux",
	14: "EMarketingMessageFlags_PlatformRestrictions",
}

func (e EMarketingMessageFlags) String() string {
	if s, ok := EMarketingMessageFlags_name[e]; ok {
		return s
	}
	return "INVALID"
}

type ENewsUpdateType int32

const (
	ENewsUpdateType_AppNews      ENewsUpdateType = 0
	ENewsUpdateType_SteamAds                     = 1
	ENewsUpdateType_SteamNews                    = 2
	ENewsUpdateType_CDDBUpdate                   = 3
	ENewsUpdateType_ClientUpdate                 = 4
)

var ENewsUpdateType_name = map[ENewsUpdateType]string{
	0: "ENewsUpdateType_AppNews",
	1: "ENewsUpdateType_SteamAds",
	2: "ENewsUpdateType_SteamNews",
	3: "ENewsUpdateType_CDDBUpdate",
	4: "ENewsUpdateType_ClientUpdate",
}

func (e ENewsUpdateType) String() string {
	if s, ok := ENewsUpdateType_name[e]; ok {
		return s
	}
	return "INVALID"
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

var ESystemIMType_name = map[ESystemIMType]string{
	0:  "ESystemIMType_RawText",
	1:  "ESystemIMType_InvalidCard",
	2:  "ESystemIMType_RecurringPurchaseFailed",
	3:  "ESystemIMType_CardWillExpire",
	4:  "ESystemIMType_SubscriptionExpired",
	5:  "ESystemIMType_GuestPassReceived",
	6:  "ESystemIMType_GuestPassGranted",
	7:  "ESystemIMType_GiftRevoked",
	8:  "ESystemIMType_SupportMessage",
	9:  "ESystemIMType_SupportMessageClearAlert",
	10: "ESystemIMType_Max",
}

func (e ESystemIMType) String() string {
	if s, ok := ESystemIMType_name[e]; ok {
		return s
	}
	return "INVALID"
}

type EChatFlags int32

const (
	EChatFlags_Locked             EChatFlags = 1
	EChatFlags_InvisibleToFriends            = 2
	EChatFlags_Moderated                     = 4
	EChatFlags_Unjoinable                    = 8
)

var EChatFlags_name = map[EChatFlags]string{
	1: "EChatFlags_Locked",
	2: "EChatFlags_InvisibleToFriends",
	4: "EChatFlags_Moderated",
	8: "EChatFlags_Unjoinable",
}

func (e EChatFlags) String() string {
	if s, ok := EChatFlags_name[e]; ok {
		return s
	}
	return "INVALID"
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

var ERemoteStoragePlatform_name = map[ERemoteStoragePlatform]string{
	0:  "ERemoteStoragePlatform_None",
	1:  "ERemoteStoragePlatform_Windows",
	2:  "ERemoteStoragePlatform_OSX",
	4:  "ERemoteStoragePlatform_PS3",
	8:  "ERemoteStoragePlatform_Reserved1",
	16: "ERemoteStoragePlatform_Reserved2",
	-1: "ERemoteStoragePlatform_All",
}

func (e ERemoteStoragePlatform) String() string {
	if s, ok := ERemoteStoragePlatform_name[e]; ok {
		return s
	}
	return "INVALID"
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

var EDRMBlobDownloadType_name = map[EDRMBlobDownloadType]string{
	0:  "EDRMBlobDownloadType_Error",
	1:  "EDRMBlobDownloadType_File",
	2:  "EDRMBlobDownloadType_Parts",
	4:  "EDRMBlobDownloadType_Compressed",
	7:  "EDRMBlobDownloadType_AllMask",
	8:  "EDRMBlobDownloadType_IsJob",
	16: "EDRMBlobDownloadType_HighPriority",
	32: "EDRMBlobDownloadType_AddTimestamp",
	64: "EDRMBlobDownloadType_LowPriority",
}

func (e EDRMBlobDownloadType) String() string {
	if s, ok := EDRMBlobDownloadType_name[e]; ok {
		return s
	}
	return "INVALID"
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

var EDRMBlobDownloadErrorDetail_name = map[EDRMBlobDownloadErrorDetail]string{
	0:      "EDRMBlobDownloadErrorDetail_None",
	1:      "EDRMBlobDownloadErrorDetail_DownloadFailed",
	2:      "EDRMBlobDownloadErrorDetail_TargetLocked",
	3:      "EDRMBlobDownloadErrorDetail_OpenZip",
	4:      "EDRMBlobDownloadErrorDetail_ReadZipDirectory",
	5:      "EDRMBlobDownloadErrorDetail_UnexpectedZipEntry",
	6:      "EDRMBlobDownloadErrorDetail_UnzipFullFile",
	7:      "EDRMBlobDownloadErrorDetail_UnknownBlobType",
	8:      "EDRMBlobDownloadErrorDetail_UnzipStrips",
	9:      "EDRMBlobDownloadErrorDetail_UnzipMergeGuid",
	10:     "EDRMBlobDownloadErrorDetail_UnzipSignature",
	11:     "EDRMBlobDownloadErrorDetail_ApplyStrips",
	12:     "EDRMBlobDownloadErrorDetail_ApplyMergeGuid",
	13:     "EDRMBlobDownloadErrorDetail_ApplySignature",
	14:     "EDRMBlobDownloadErrorDetail_AppIdMismatch",
	15:     "EDRMBlobDownloadErrorDetail_AppIdUnexpected",
	16:     "EDRMBlobDownloadErrorDetail_AppliedSignatureCorrupt",
	17:     "EDRMBlobDownloadErrorDetail_ApplyValveSignatureHeader",
	18:     "EDRMBlobDownloadErrorDetail_UnzipValveSignatureHeader",
	19:     "EDRMBlobDownloadErrorDetail_PathManipulationError",
	65536:  "EDRMBlobDownloadErrorDetail_TargetLocked_Base",
	131071: "EDRMBlobDownloadErrorDetail_TargetLocked_Max",
	131072: "EDRMBlobDownloadErrorDetail_NextBase",
}

func (e EDRMBlobDownloadErrorDetail) String() string {
	if s, ok := EDRMBlobDownloadErrorDetail_name[e]; ok {
		return s
	}
	return "INVALID"
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

var EClientStat_name = map[EClientStat]string{
	0: "EClientStat_P2PConnectionsUDP",
	1: "EClientStat_P2PConnectionsRelay",
	2: "EClientStat_P2PGameConnections",
	3: "EClientStat_P2PVoiceConnections",
	4: "EClientStat_BytesDownloaded",
	5: "EClientStat_Max",
}

func (e EClientStat) String() string {
	if s, ok := EClientStat_name[e]; ok {
		return s
	}
	return "INVALID"
}

type EClientStatAggregateMethod int32

const (
	EClientStatAggregateMethod_LatestOnly EClientStatAggregateMethod = 0
	EClientStatAggregateMethod_Sum                                   = 1
	EClientStatAggregateMethod_Event                                 = 2
	EClientStatAggregateMethod_Scalar                                = 3
)

var EClientStatAggregateMethod_name = map[EClientStatAggregateMethod]string{
	0: "EClientStatAggregateMethod_LatestOnly",
	1: "EClientStatAggregateMethod_Sum",
	2: "EClientStatAggregateMethod_Event",
	3: "EClientStatAggregateMethod_Scalar",
}

func (e EClientStatAggregateMethod) String() string {
	if s, ok := EClientStatAggregateMethod_name[e]; ok {
		return s
	}
	return "INVALID"
}

type ELeaderboardDataRequest int32

const (
	ELeaderboardDataRequest_Global           ELeaderboardDataRequest = 0
	ELeaderboardDataRequest_GlobalAroundUser                         = 1
	ELeaderboardDataRequest_Friends                                  = 2
	ELeaderboardDataRequest_Users                                    = 3
)

var ELeaderboardDataRequest_name = map[ELeaderboardDataRequest]string{
	0: "ELeaderboardDataRequest_Global",
	1: "ELeaderboardDataRequest_GlobalAroundUser",
	2: "ELeaderboardDataRequest_Friends",
	3: "ELeaderboardDataRequest_Users",
}

func (e ELeaderboardDataRequest) String() string {
	if s, ok := ELeaderboardDataRequest_name[e]; ok {
		return s
	}
	return "INVALID"
}

type ELeaderboardSortMethod int32

const (
	ELeaderboardSortMethod_None       ELeaderboardSortMethod = 0
	ELeaderboardSortMethod_Ascending                         = 1
	ELeaderboardSortMethod_Descending                        = 2
)

var ELeaderboardSortMethod_name = map[ELeaderboardSortMethod]string{
	0: "ELeaderboardSortMethod_None",
	1: "ELeaderboardSortMethod_Ascending",
	2: "ELeaderboardSortMethod_Descending",
}

func (e ELeaderboardSortMethod) String() string {
	if s, ok := ELeaderboardSortMethod_name[e]; ok {
		return s
	}
	return "INVALID"
}

type ELeaderboardUploadScoreMethod int32

const (
	ELeaderboardUploadScoreMethod_None        ELeaderboardUploadScoreMethod = 0
	ELeaderboardUploadScoreMethod_KeepBest                                  = 1
	ELeaderboardUploadScoreMethod_ForceUpdate                               = 2
)

var ELeaderboardUploadScoreMethod_name = map[ELeaderboardUploadScoreMethod]string{
	0: "ELeaderboardUploadScoreMethod_None",
	1: "ELeaderboardUploadScoreMethod_KeepBest",
	2: "ELeaderboardUploadScoreMethod_ForceUpdate",
}

func (e ELeaderboardUploadScoreMethod) String() string {
	if s, ok := ELeaderboardUploadScoreMethod_name[e]; ok {
		return s
	}
	return "INVALID"
}

type EUCMFilePrivacyState int32

const (
	EUCMFilePrivacyState_Invalid     EUCMFilePrivacyState = -1
	EUCMFilePrivacyState_Private                          = 2
	EUCMFilePrivacyState_FriendsOnly                      = 4
	EUCMFilePrivacyState_Public                           = 8
	EUCMFilePrivacyState_All                              = EUCMFilePrivacyState_Public | EUCMFilePrivacyState_FriendsOnly | EUCMFilePrivacyState_Private
)

var EUCMFilePrivacyState_name = map[EUCMFilePrivacyState]string{
	-1: "EUCMFilePrivacyState_Invalid",
	2:  "EUCMFilePrivacyState_Private",
	4:  "EUCMFilePrivacyState_FriendsOnly",
	8:  "EUCMFilePrivacyState_Public",
	14: "EUCMFilePrivacyState_All",
}

func (e EUCMFilePrivacyState) String() string {
	if s, ok := EUCMFilePrivacyState_name[e]; ok {
		return s
	}
	return "INVALID"
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

var EUdpPacketType_name = map[EUdpPacketType]string{
	0: "EUdpPacketType_Invalid",
	1: "EUdpPacketType_ChallengeReq",
	2: "EUdpPacketType_Challenge",
	3: "EUdpPacketType_Connect",
	4: "EUdpPacketType_Accept",
	5: "EUdpPacketType_Disconnect",
	6: "EUdpPacketType_Data",
	7: "EUdpPacketType_Datagram",
	8: "EUdpPacketType_Max",
}

func (e EUdpPacketType) String() string {
	if s, ok := EUdpPacketType_name[e]; ok {
		return s
	}
	return "INVALID"
}
