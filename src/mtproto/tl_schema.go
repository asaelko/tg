package mtproto

import (
	"fmt"
	"reflect"
	"strconv"
)

const (
	crc_boolFalse                               = 0xbc799737
	crc_boolTrue                                = 0x997275b5
	crc_true                                    = 0x3fedd339
	crc_error                                   = 0xc4b9f9bb
	crc_null                                    = 0x56730bcc
	crc_inputPeerEmpty                          = 0x7f3b18ea
	crc_inputPeerSelf                           = 0x7da07ec9
	crc_inputPeerChat                           = 0x179be863
	crc_inputUserEmpty                          = 0xb98886cf
	crc_inputUserSelf                           = 0xf7c1b13f
	crc_inputPhoneContact                       = 0xf392b7f4
	crc_inputFile                               = 0xf52ff27f
	crc_inputMediaEmpty                         = 0x9664f57f
	crc_inputMediaUploadedPhoto                 = 0xf7aff1c0
	crc_inputMediaPhoto                         = 0xe9bfb4f3
	crc_inputMediaGeoPoint                      = 0xf9c44144
	crc_inputMediaContact                       = 0xa6e45987
	crc_inputChatPhotoEmpty                     = 0x1ca48f57
	crc_inputChatUploadedPhoto                  = 0x94254732
	crc_inputChatPhoto                          = 0xb2e1bf08
	crc_inputGeoPointEmpty                      = 0xe4c123d6
	crc_inputGeoPoint                           = 0xf3b7acc9
	crc_inputPhotoEmpty                         = 0x1cd7bf0d
	crc_inputPhoto                              = 0xfb95c6c4
	crc_inputFileLocation                       = 0x14637196
	crc_inputPhotoCropAuto                      = 0xade6b004
	crc_inputPhotoCrop                          = 0xd9915325
	crc_inputAppEvent                           = 0x770656a8
	crc_peerUser                                = 0x9db1bc6d
	crc_peerChat                                = 0xbad0e5bb
	crc_storage_fileUnknown                     = 0xaa963b05
	crc_storage_fileJpeg                        = 0x007efe0e
	crc_storage_fileGif                         = 0xcae1aadf
	crc_storage_filePng                         = 0x0a4f63c0
	crc_storage_filePdf                         = 0xae1e508d
	crc_storage_fileMp3                         = 0x528a0677
	crc_storage_fileMov                         = 0x4b09ebbc
	crc_storage_filePartial                     = 0x40bc6f52
	crc_storage_fileMp4                         = 0xb3cea0e4
	crc_storage_fileWebp                        = 0x1081464c
	crc_fileLocationUnavailable                 = 0x7c596b46
	crc_fileLocation                            = 0x53d69076
	crc_userEmpty                               = 0x200250ba
	crc_userProfilePhotoEmpty                   = 0x4f11bae1
	crc_userProfilePhoto                        = 0xd559d8c8
	crc_userStatusEmpty                         = 0x09d05049
	crc_userStatusOnline                        = 0xedb93949
	crc_userStatusOffline                       = 0x008c703f
	crc_chatEmpty                               = 0x9ba2d800
	crc_chat                                    = 0xd91cdd54
	crc_chatForbidden                           = 0x07328bdb
	crc_chatFull                                = 0x2e02a614
	crc_chatParticipant                         = 0xc8d7493e
	crc_chatParticipantsForbidden               = 0xfc900c2b
	crc_chatParticipants                        = 0x3f460fed
	crc_chatPhotoEmpty                          = 0x37c1011c
	crc_chatPhoto                               = 0x6153276a
	crc_messageEmpty                            = 0x83e5de54
	crc_message                                 = 0xc09be45f
	crc_messageService                          = 0x9e19a1f6
	crc_messageMediaEmpty                       = 0x3ded6320
	crc_messageMediaPhoto                       = 0x3d8ce53d
	crc_messageMediaGeo                         = 0x56e0d474
	crc_messageMediaContact                     = 0x5e7d2f39
	crc_messageMediaUnsupported                 = 0x9f84f49e
	crc_messageActionEmpty                      = 0xb6aef7b0
	crc_messageActionChatCreate                 = 0xa6638b9a
	crc_messageActionChatEditTitle              = 0xb5a1ce5a
	crc_messageActionChatEditPhoto              = 0x7fcb13a8
	crc_messageActionChatDeletePhoto            = 0x95e3fbef
	crc_messageActionChatAddUser                = 0x488a7337
	crc_messageActionChatDeleteUser             = 0xb2ae9b0c
	crc_dialog                                  = 0x66ffba14
	crc_photoEmpty                              = 0x2331b22d
	crc_photo                                   = 0xcded42fe
	crc_photoSizeEmpty                          = 0x0e17e23c
	crc_photoSize                               = 0x77bfb61b
	crc_photoCachedSize                         = 0xe9a734fa
	crc_geoPointEmpty                           = 0x1117dd5f
	crc_geoPoint                                = 0x2049d70c
	crc_auth_checkedPhone                       = 0x811ea28e
	crc_auth_sentCode                           = 0x5e002502
	crc_auth_authorization                      = 0xcd050916
	crc_auth_exportedAuthorization              = 0xdf969c2d
	crc_inputNotifyPeer                         = 0xb8bc5b0c
	crc_inputNotifyUsers                        = 0x193b4417
	crc_inputNotifyChats                        = 0x4a95e84e
	crc_inputNotifyAll                          = 0xa429b886
	crc_inputPeerNotifyEventsEmpty              = 0xf03064d8
	crc_inputPeerNotifyEventsAll                = 0xe86a2c74
	crc_inputPeerNotifySettings                 = 0x38935eb2
	crc_peerNotifyEventsEmpty                   = 0xadd53cb3
	crc_peerNotifyEventsAll                     = 0x6d1ded88
	crc_peerNotifySettingsEmpty                 = 0x70a68512
	crc_peerNotifySettings                      = 0x9acda4c0
	crc_peerSettings                            = 0x818426cd
	crc_wallPaper                               = 0xccb03657
	crc_inputReportReasonSpam                   = 0x58dbcab8
	crc_inputReportReasonViolence               = 0x1e22c78d
	crc_inputReportReasonPornography            = 0x2e59d922
	crc_inputReportReasonOther                  = 0xe1746d0a
	crc_userFull                                = 0x5932fc03
	crc_contact                                 = 0xf911c994
	crc_importedContact                         = 0xd0028438
	crc_contactBlocked                          = 0x561bc879
	crc_contactStatus                           = 0xd3680c61
	crc_contacts_link                           = 0x3ace484c
	crc_contacts_contactsNotModified            = 0xb74ba9d2
	crc_contacts_contacts                       = 0x6f8b8cb2
	crc_contacts_importedContacts               = 0xad524315
	crc_contacts_blocked                        = 0x1c138d15
	crc_contacts_blockedSlice                   = 0x900802a1
	crc_messages_dialogs                        = 0x15ba6c40
	crc_messages_dialogsSlice                   = 0x71e094f3
	crc_messages_messages                       = 0x8c718e87
	crc_messages_messagesSlice                  = 0x0b446ae3
	crc_messages_chats                          = 0x64ff9fd5
	crc_messages_chatFull                       = 0xe5d7d19c
	crc_messages_affectedHistory                = 0xb45c69d1
	crc_inputMessagesFilterEmpty                = 0x57e2f66c
	crc_inputMessagesFilterPhotos               = 0x9609a51c
	crc_inputMessagesFilterVideo                = 0x9fc00e65
	crc_inputMessagesFilterPhotoVideo           = 0x56e9f0e4
	crc_inputMessagesFilterPhotoVideoDocuments  = 0xd95e73bb
	crc_inputMessagesFilterDocument             = 0x9eddf188
	crc_inputMessagesFilterUrl                  = 0x7ef0dd87
	crc_inputMessagesFilterGif                  = 0xffc86587
	crc_updateNewMessage                        = 0x1f2b0afd
	crc_updateMessageID                         = 0x4e90bfd6
	crc_updateDeleteMessages                    = 0xa20db0e5
	crc_updateUserTyping                        = 0x5c486927
	crc_updateChatUserTyping                    = 0x9a65ea1f
	crc_updateChatParticipants                  = 0x07761198
	crc_updateUserStatus                        = 0x1bfbd823
	crc_updateUserName                          = 0xa7332b73
	crc_updateUserPhoto                         = 0x95313b0c
	crc_updateContactRegistered                 = 0x2575bbb9
	crc_updateContactLink                       = 0x9d2e67c5
	crc_updateNewAuthorization                  = 0x8f06529a
	crc_updates_state                           = 0xa56c2a3e
	crc_updates_differenceEmpty                 = 0x5d75a138
	crc_updates_difference                      = 0x00f49ca0
	crc_updates_differenceSlice                 = 0xa8fb1981
	crc_updatesTooLong                          = 0xe317af7e
	crc_updateShortMessage                      = 0x914fbf11
	crc_updateShortChatMessage                  = 0x16812688
	crc_updateShort                             = 0x78d4dec1
	crc_updatesCombined                         = 0x725b04c3
	crc_updates                                 = 0x74ae4240
	crc_photos_photos                           = 0x8dca6aa5
	crc_photos_photosSlice                      = 0x15051f54
	crc_photos_photo                            = 0x20212ca8
	crc_upload_file                             = 0x096a18d5
	crc_dcOption                                = 0x05d8c6cc
	crc_config                                  = 0x9a6b2e2a
	crc_nearestDc                               = 0x8e1a1775
	crc_help_appUpdate                          = 0x8987f311
	crc_help_noAppUpdate                        = 0xc45a6536
	crc_help_inviteText                         = 0x18cb9f78
	crc_wallPaperSolid                          = 0x63117f24
	crc_updateNewEncryptedMessage               = 0x12bcbd9a
	crc_updateEncryptedChatTyping               = 0x1710f156
	crc_updateEncryption                        = 0xb4a2e88d
	crc_updateEncryptedMessagesRead             = 0x38fe25b7
	crc_encryptedChatEmpty                      = 0xab7ec0a0
	crc_encryptedChatWaiting                    = 0x3bf703dc
	crc_encryptedChatRequested                  = 0xc878527e
	crc_encryptedChat                           = 0xfa56ce36
	crc_encryptedChatDiscarded                  = 0x13d6dd27
	crc_inputEncryptedChat                      = 0xf141b5e1
	crc_encryptedFileEmpty                      = 0xc21f497e
	crc_encryptedFile                           = 0x4a70994c
	crc_inputEncryptedFileEmpty                 = 0x1837c364
	crc_inputEncryptedFileUploaded              = 0x64bd0306
	crc_inputEncryptedFile                      = 0x5a17b5e5
	crc_inputEncryptedFileLocation              = 0xf5235d55
	crc_encryptedMessage                        = 0xed18c118
	crc_encryptedMessageService                 = 0x23734b06
	crc_messages_dhConfigNotModified            = 0xc0e24635
	crc_messages_dhConfig                       = 0x2c221edd
	crc_messages_sentEncryptedMessage           = 0x560f8935
	crc_messages_sentEncryptedFile              = 0x9493ff32
	crc_inputFileBig                            = 0xfa4f0bb5
	crc_inputEncryptedFileBigUploaded           = 0x2dc173c8
	crc_updateChatParticipantAdd                = 0xea4b0e5c
	crc_updateChatParticipantDelete             = 0x6e5f8c22
	crc_updateDcOptions                         = 0x8e5e9873
	crc_inputMediaUploadedDocument              = 0x1d89306d
	crc_inputMediaUploadedThumbDocument         = 0xad613491
	crc_inputMediaDocument                      = 0x1a77f29c
	crc_messageMediaDocument                    = 0xf3e02ea8
	crc_inputDocumentEmpty                      = 0x72f0eaae
	crc_inputDocument                           = 0x18798952
	crc_inputDocumentFileLocation               = 0x430f0724
	crc_documentEmpty                           = 0x36f8c871
	crc_document                                = 0x87232bc7
	crc_help_support                            = 0x17c6b5f6
	crc_notifyPeer                              = 0x9fd40bd8
	crc_notifyUsers                             = 0xb4c83b4c
	crc_notifyChats                             = 0xc007cec3
	crc_notifyAll                               = 0x74d07c60
	crc_updateUserBlocked                       = 0x80ece81a
	crc_updateNotifySettings                    = 0xbec268ef
	crc_sendMessageTypingAction                 = 0x16bf744e
	crc_sendMessageCancelAction                 = 0xfd5ec8f5
	crc_sendMessageRecordVideoAction            = 0xa187d66f
	crc_sendMessageUploadVideoAction            = 0xe9763aec
	crc_sendMessageRecordAudioAction            = 0xd52f73f7
	crc_sendMessageUploadAudioAction            = 0xf351d7ab
	crc_sendMessageUploadPhotoAction            = 0xd1d34a26
	crc_sendMessageUploadDocumentAction         = 0xaa0cd9e4
	crc_sendMessageGeoLocationAction            = 0x176f8ba1
	crc_sendMessageChooseContactAction          = 0x628cbc6f
	crc_contacts_found                          = 0x1aa1f784
	crc_updateServiceNotification               = 0x382dd3e4
	crc_userStatusRecently                      = 0xe26f42f1
	crc_userStatusLastWeek                      = 0x07bf09fc
	crc_userStatusLastMonth                     = 0x77ebc742
	crc_updatePrivacy                           = 0xee3b272a
	crc_inputPrivacyKeyStatusTimestamp          = 0x4f96cb18
	crc_privacyKeyStatusTimestamp               = 0xbc2eab30
	crc_inputPrivacyValueAllowContacts          = 0x0d09e07b
	crc_inputPrivacyValueAllowAll               = 0x184b35ce
	crc_inputPrivacyValueAllowUsers             = 0x131cc67f
	crc_inputPrivacyValueDisallowContacts       = 0x0ba52007
	crc_inputPrivacyValueDisallowAll            = 0xd66b66c9
	crc_inputPrivacyValueDisallowUsers          = 0x90110467
	crc_privacyValueAllowContacts               = 0xfffe1bac
	crc_privacyValueAllowAll                    = 0x65427b82
	crc_privacyValueAllowUsers                  = 0x4d5bbe0c
	crc_privacyValueDisallowContacts            = 0xf888fa1a
	crc_privacyValueDisallowAll                 = 0x8b73e763
	crc_privacyValueDisallowUsers               = 0x0c7f49b7
	crc_account_privacyRules                    = 0x554abb6f
	crc_accountDaysTTL                          = 0xb8d0afdf
	crc_updateUserPhone                         = 0x12b9417b
	crc_documentAttributeImageSize              = 0x6c37c15c
	crc_documentAttributeAnimated               = 0x11b58939
	crc_documentAttributeSticker                = 0x3a556302
	crc_documentAttributeVideo                  = 0x5910cccb
	crc_documentAttributeAudio                  = 0x9852f9c6
	crc_documentAttributeFilename               = 0x15590068
	crc_messages_stickersNotModified            = 0xf1749a22
	crc_messages_stickers                       = 0x8a8ecd32
	crc_stickerPack                             = 0x12b299d4
	crc_messages_allStickersNotModified         = 0xe86602c3
	crc_messages_allStickers                    = 0xedfd405f
	crc_disabledFeature                         = 0xae636f24
	crc_updateReadHistoryInbox                  = 0x9961fd5c
	crc_updateReadHistoryOutbox                 = 0x2f2f21bf
	crc_messages_affectedMessages               = 0x84d19185
	crc_contactLinkUnknown                      = 0x5f4f9247
	crc_contactLinkNone                         = 0xfeedd3ad
	crc_contactLinkHasPhone                     = 0x268f3f59
	crc_contactLinkContact                      = 0xd502c2d0
	crc_updateWebPage                           = 0x7f891213
	crc_webPageEmpty                            = 0xeb1477e8
	crc_webPagePending                          = 0xc586da1c
	crc_webPage                                 = 0xca820ed7
	crc_messageMediaWebPage                     = 0xa32dd600
	crc_authorization                           = 0x7bf2e6f6
	crc_account_authorizations                  = 0x1250abde
	crc_account_noPassword                      = 0x96dabc18
	crc_account_password                        = 0x7c18141c
	crc_account_passwordSettings                = 0xb7b72ab3
	crc_account_passwordInputSettings           = 0x86916deb
	crc_auth_passwordRecovery                   = 0x137948a5
	crc_inputMediaVenue                         = 0x2827a81a
	crc_messageMediaVenue                       = 0x7912b71f
	crc_receivedNotifyMessage                   = 0xa384b779
	crc_chatInviteEmpty                         = 0x69df3769
	crc_chatInviteExported                      = 0xfc2e05bc
	crc_chatInviteAlready                       = 0x5a686d7c
	crc_chatInvite                              = 0xdb74f558
	crc_messageActionChatJoinedByLink           = 0xf89cf5e8
	crc_updateReadMessagesContents              = 0x68c13933
	crc_inputStickerSetEmpty                    = 0xffb62b95
	crc_inputStickerSetID                       = 0x9de7a269
	crc_inputStickerSetShortName                = 0x861cc8a0
	crc_stickerSet                              = 0xcd303b41
	crc_messages_stickerSet                     = 0xb60a24a6
	crc_user                                    = 0xd10d979a
	crc_botCommand                              = 0xc27ac8c7
	crc_botInfo                                 = 0x98e81d3a
	crc_keyboardButton                          = 0xa2fa4880
	crc_keyboardButtonRow                       = 0x77608b83
	crc_replyKeyboardHide                       = 0xa03e5b85
	crc_replyKeyboardForceReply                 = 0xf4108aa0
	crc_replyKeyboardMarkup                     = 0x3502758c
	crc_inputPeerUser                           = 0x7b8e7de6
	crc_inputUser                               = 0xd8292816
	crc_help_appChangelogEmpty                  = 0xaf7e0394
	crc_help_appChangelog                       = 0x4668e6bd
	crc_messageEntityUnknown                    = 0xbb92ba95
	crc_messageEntityMention                    = 0xfa04579d
	crc_messageEntityHashtag                    = 0x6f635b0d
	crc_messageEntityBotCommand                 = 0x6cef8ac7
	crc_messageEntityUrl                        = 0x6ed02538
	crc_messageEntityEmail                      = 0x64e475c2
	crc_messageEntityBold                       = 0xbd610bc9
	crc_messageEntityItalic                     = 0x826f8b60
	crc_messageEntityCode                       = 0x28a20571
	crc_messageEntityPre                        = 0x73924be0
	crc_messageEntityTextUrl                    = 0x76a6d327
	crc_updateShortSentMessage                  = 0x11f1331c
	crc_inputChannelEmpty                       = 0xee8c1e86
	crc_inputChannel                            = 0xafeb712e
	crc_peerChannel                             = 0xbddde532
	crc_inputPeerChannel                        = 0x20adaef8
	crc_channel                                 = 0xa14dca52
	crc_channelForbidden                        = 0x8537784f
	crc_contacts_resolvedPeer                   = 0x7f077ad9
	crc_channelFull                             = 0xc3d5512f
	crc_messageRange                            = 0x0ae30253
	crc_messages_channelMessages                = 0x99262e37
	crc_messageActionChannelCreate              = 0x95d2ac92
	crc_updateChannelTooLong                    = 0xeb0467fb
	crc_updateChannel                           = 0xb6d45656
	crc_updateNewChannelMessage                 = 0x62ba04d9
	crc_updateReadChannelInbox                  = 0x4214f37f
	crc_updateDeleteChannelMessages             = 0xc37521c9
	crc_updateChannelMessageViews               = 0x98a12b4b
	crc_updates_channelDifferenceEmpty          = 0x3e11affb
	crc_updates_channelDifferenceTooLong        = 0x410dee07
	crc_updates_channelDifference               = 0x2064674e
	crc_channelMessagesFilterEmpty              = 0x94d42ee7
	crc_channelMessagesFilter                   = 0xcd77d957
	crc_channelParticipant                      = 0x15ebac1d
	crc_channelParticipantSelf                  = 0xa3289a6d
	crc_channelParticipantModerator             = 0x91057fef
	crc_channelParticipantEditor                = 0x98192d61
	crc_channelParticipantKicked                = 0x8cc5e69a
	crc_channelParticipantCreator               = 0xe3e2e1f9
	crc_channelParticipantsRecent               = 0xde3f3c79
	crc_channelParticipantsAdmins               = 0xb4608969
	crc_channelParticipantsKicked               = 0x3c37bb7a
	crc_channelRoleEmpty                        = 0xb285a0c6
	crc_channelRoleModerator                    = 0x9618d975
	crc_channelRoleEditor                       = 0x820bfe8c
	crc_channels_channelParticipants            = 0xf56ee2a8
	crc_channels_channelParticipant             = 0xd0d9b163
	crc_chatParticipantCreator                  = 0xda13538a
	crc_chatParticipantAdmin                    = 0xe2d6e436
	crc_updateChatAdmins                        = 0x6e947941
	crc_updateChatParticipantAdmin              = 0xb6901959
	crc_messageActionChatMigrateTo              = 0x51bdb021
	crc_messageActionChannelMigrateFrom         = 0xb055eaee
	crc_channelParticipantsBots                 = 0xb0d1865b
	crc_help_termsOfService                     = 0xf1ee3e90
	crc_updateNewStickerSet                     = 0x688a30aa
	crc_updateStickerSetsOrder                  = 0xf0dfb451
	crc_updateStickerSets                       = 0x43ae3dec
	crc_foundGif                                = 0x162ecc1f
	crc_foundGifCached                          = 0x9c750409
	crc_inputMediaGifExternal                   = 0x4843b0fd
	crc_messages_foundGifs                      = 0x450a1c0a
	crc_messages_savedGifsNotModified           = 0xe8025ca2
	crc_messages_savedGifs                      = 0x2e0709a5
	crc_updateSavedGifs                         = 0x9375341e
	crc_inputBotInlineMessageMediaAuto          = 0x292fed13
	crc_inputBotInlineMessageText               = 0x3dcd7a87
	crc_inputBotInlineResult                    = 0x2cbbe15a
	crc_botInlineMessageMediaAuto               = 0x0a74b15b
	crc_botInlineMessageText                    = 0x8c7f65e2
	crc_botInlineResult                         = 0x9bebaeb9
	crc_messages_botResults                     = 0x256709a6
	crc_updateBotInlineQuery                    = 0x54826690
	crc_updateBotInlineSend                     = 0x0e48f964
	crc_inputMessagesFilterVoice                = 0x50f5c392
	crc_inputMessagesFilterMusic                = 0x3751b49e
	crc_inputPrivacyKeyChatInvite               = 0xbdfb0426
	crc_privacyKeyChatInvite                    = 0x500e6dfa
	crc_exportedMessageLink                     = 0x1f486803
	crc_messageFwdHeader                        = 0xc786ddcb
	crc_updateEditChannelMessage                = 0x1b3f4df7
	crc_updateChannelPinnedMessage              = 0x98592475
	crc_messageActionPinMessage                 = 0x94bd38ed
	crc_auth_codeTypeSms                        = 0x72a3158c
	crc_auth_codeTypeCall                       = 0x741cd3e3
	crc_auth_codeTypeFlashCall                  = 0x226ccefb
	crc_auth_sentCodeTypeApp                    = 0x3dbb5986
	crc_auth_sentCodeTypeSms                    = 0xc000bba2
	crc_auth_sentCodeTypeCall                   = 0x5353e5a7
	crc_auth_sentCodeTypeFlashCall              = 0xab03c6d9
	crc_keyboardButtonUrl                       = 0x258aff05
	crc_keyboardButtonCallback                  = 0x683a5e46
	crc_keyboardButtonRequestPhone              = 0xb16a6c29
	crc_keyboardButtonRequestGeoLocation        = 0xfc796b3f
	crc_keyboardButtonSwitchInline              = 0xea1b7a14
	crc_replyInlineMarkup                       = 0x48a30254
	crc_messages_botCallbackAnswer              = 0xb10df1fb
	crc_updateBotCallbackQuery                  = 0xa68c688c
	crc_messages_messageEditData                = 0x26b5dde6
	crc_updateEditMessage                       = 0xe40370a3
	crc_inputBotInlineMessageMediaGeo           = 0xf4a59de1
	crc_inputBotInlineMessageMediaVenue         = 0xaaafadc8
	crc_inputBotInlineMessageMediaContact       = 0x2daf01a7
	crc_botInlineMessageMediaGeo                = 0x3a8fd8b8
	crc_botInlineMessageMediaVenue              = 0x4366232e
	crc_botInlineMessageMediaContact            = 0x35edb4d4
	crc_inputBotInlineResultPhoto               = 0xa8d864a7
	crc_inputBotInlineResultDocument            = 0xfff8fdc4
	crc_botInlineMediaResult                    = 0x17db940b
	crc_inputBotInlineMessageID                 = 0x890c3d89
	crc_updateInlineBotCallbackQuery            = 0x2cbd95af
	crc_inlineBotSwitchPM                       = 0x3c20629f
	crc_messages_peerDialogs                    = 0x3371c354
	crc_topPeer                                 = 0xedcdc05b
	crc_topPeerCategoryBotsPM                   = 0xab661b5b
	crc_topPeerCategoryBotsInline               = 0x148677e2
	crc_topPeerCategoryCorrespondents           = 0x0637b7ed
	crc_topPeerCategoryGroups                   = 0xbd17a14a
	crc_topPeerCategoryChannels                 = 0x161d9628
	crc_topPeerCategoryPeers                    = 0xfb834291
	crc_contacts_topPeersNotModified            = 0xde266ef5
	crc_contacts_topPeers                       = 0x70b772a8
	crc_messageEntityMentionName                = 0x352dca58
	crc_inputMessageEntityMentionName           = 0x208e68c9
	crc_inputMessagesFilterChatPhotos           = 0x3a20ecb8
	crc_updateReadChannelOutbox                 = 0x25d6c9c7
	crc_updateDraftMessage                      = 0xee2bb969
	crc_draftMessageEmpty                       = 0xba4baec5
	crc_draftMessage                            = 0xfd8e711f
	crc_messageActionHistoryClear               = 0x9fbab604
	crc_messages_featuredStickersNotModified    = 0x04ede3cf
	crc_messages_featuredStickers               = 0xf89d88e5
	crc_updateReadFeaturedStickers              = 0x571d2742
	crc_messages_recentStickersNotModified      = 0x0b17f890
	crc_messages_recentStickers                 = 0x5ce20970
	crc_updateRecentStickers                    = 0x9a422c20
	crc_messages_archivedStickers               = 0x4fcba9c8
	crc_messages_stickerSetInstallResultSuccess = 0x38641628
	crc_messages_stickerSetInstallResultArchive = 0x35e410a8
	crc_stickerSetCovered                       = 0x6410a5d2
	crc_invokeAfterMsg                          = 0xcb9f372d
	crc_invokeAfterMsgs                         = 0x3dc4b4f0
	crc_auth_checkPhone                         = 0x6fe51dfb
	crc_auth_sendCode                           = 0x86aef0ec
	crc_auth_signUp                             = 0x1b067634
	crc_auth_signIn                             = 0xbcd51581
	crc_auth_logOut                             = 0x5717da40
	crc_auth_resetAuthorizations                = 0x9fab0d1a
	crc_auth_sendInvites                        = 0x771c1d97
	crc_auth_exportAuthorization                = 0xe5bfffcd
	crc_auth_importAuthorization                = 0xe3ef9613
	crc_auth_bindTempAuthKey                    = 0xcdd42a05
	crc_account_registerDevice                  = 0x637ea878
	crc_account_unregisterDevice                = 0x65c55b40
	crc_account_updateNotifySettings            = 0x84be5b93
	crc_account_getNotifySettings               = 0x12b3ad31
	crc_account_resetNotifySettings             = 0xdb7e1747
	crc_account_updateProfile                   = 0x78515775
	crc_account_updateStatus                    = 0x6628562c
	crc_account_getWallPapers                   = 0xc04cfac2
	crc_account_reportPeer                      = 0xae189d5f
	crc_users_getUsers                          = 0x0d91a548
	crc_users_getFullUser                       = 0xca30a5b1
	crc_contacts_getStatuses                    = 0xc4a353ee
	crc_contacts_getContacts                    = 0x22c6aa08
	crc_contacts_importContacts                 = 0xda30b32d
	crc_contacts_deleteContact                  = 0x8e953744
	crc_contacts_deleteContacts                 = 0x59ab389e
	crc_contacts_block                          = 0x332b49fc
	crc_contacts_unblock                        = 0xe54100bd
	crc_contacts_getBlocked                     = 0xf57c350f
	crc_contacts_exportCard                     = 0x84e53737
	crc_contacts_importCard                     = 0x4fe196fe
	crc_messages_getMessages                    = 0x4222fa74
	crc_messages_getDialogs                     = 0x6b47f94d
	crc_messages_getHistory                     = 0xafa92846
	crc_messages_search                         = 0xd4569248
	crc_messages_readHistory                    = 0x0e306d3a
	crc_messages_deleteHistory                  = 0x1c015b09
	crc_messages_deleteMessages                 = 0xa5f18925
	crc_messages_receivedMessages               = 0x05a954c0
	crc_messages_setTyping                      = 0xa3825e50
	crc_messages_sendMessage                    = 0xfa88427a
	crc_messages_sendMedia                      = 0xc8f16791
	crc_messages_forwardMessages                = 0x708e0195
	crc_messages_reportSpam                     = 0xcf1592db
	crc_messages_hideReportSpam                 = 0xa8f1709b
	crc_messages_getPeerSettings                = 0x3672e09c
	crc_messages_getChats                       = 0x3c6aa187
	crc_messages_getFullChat                    = 0x3b831c66
	crc_messages_editChatTitle                  = 0xdc452855
	crc_messages_editChatPhoto                  = 0xca4c79d8
	crc_messages_addChatUser                    = 0xf9a0aa09
	crc_messages_deleteChatUser                 = 0xe0611f16
	crc_messages_createChat                     = 0x09cb126e
	crc_updates_getState                        = 0xedd4882a
	crc_updates_getDifference                   = 0x0a041495
	crc_photos_updateProfilePhoto               = 0xeef579a0
	crc_photos_uploadProfilePhoto               = 0xd50f9c88
	crc_photos_deletePhotos                     = 0x87cf7f2f
	crc_upload_saveFilePart                     = 0xb304a621
	crc_upload_getFile                          = 0xe3a6cfb5
	crc_help_getConfig                          = 0xc4f9186b
	crc_help_getNearestDc                       = 0x1fb33026
	crc_help_getAppUpdate                       = 0xae2de196
	crc_help_saveAppLog                         = 0x6f02f748
	crc_help_getInviteText                      = 0x4d392343
	crc_photos_getUserPhotos                    = 0x91cd32a8
	crc_messages_forwardMessage                 = 0x33963bf9
	crc_messages_sendBroadcast                  = 0xbf73f4da
	crc_messages_getDhConfig                    = 0x26cf8950
	crc_messages_requestEncryption              = 0xf64daf43
	crc_messages_acceptEncryption               = 0x3dbc0415
	crc_messages_discardEncryption              = 0xedd923c5
	crc_messages_setEncryptedTyping             = 0x791451ed
	crc_messages_readEncryptedHistory           = 0x7f4b690a
	crc_messages_sendEncrypted                  = 0xa9776773
	crc_messages_sendEncryptedFile              = 0x9a901b66
	crc_messages_sendEncryptedService           = 0x32d439a4
	crc_messages_receivedQueue                  = 0x55a5bb66
	crc_upload_saveBigFilePart                  = 0xde7b673d
	crc_initConnection                          = 0x69796de9
	crc_help_getSupport                         = 0x9cdf08cd
	crc_messages_readMessageContents            = 0x36a73f77
	crc_account_checkUsername                   = 0x2714d86c
	crc_account_updateUsername                  = 0x3e0bdd7c
	crc_contacts_search                         = 0x11f812d8
	crc_account_getPrivacy                      = 0xdadbc950
	crc_account_setPrivacy                      = 0xc9f81ce8
	crc_account_deleteAccount                   = 0x418d4e0b
	crc_account_getAccountTTL                   = 0x08fc711d
	crc_account_setAccountTTL                   = 0x2442485e
	crc_invokeWithLayer                         = 0xda9b0d0d
	crc_contacts_resolveUsername                = 0xf93ccba3
	crc_account_sendChangePhoneCode             = 0x08e57deb
	crc_account_changePhone                     = 0x70c32edb
	crc_messages_getStickers                    = 0xae22e045
	crc_messages_getAllStickers                 = 0x1c9618b1
	crc_account_updateDeviceLocked              = 0x38df3532
	crc_auth_importBotAuthorization             = 0x67a3ff2c
	crc_messages_getWebPagePreview              = 0x25223e24
	crc_account_getAuthorizations               = 0xe320c158
	crc_account_resetAuthorization              = 0xdf77f3bc
	crc_account_getPassword                     = 0x548a30f5
	crc_account_getPasswordSettings             = 0xbc8d11bb
	crc_account_updatePasswordSettings          = 0xfa7c4b86
	crc_auth_checkPassword                      = 0x0a63011e
	crc_auth_requestPasswordRecovery            = 0xd897bc66
	crc_auth_recoverPassword                    = 0x4ea56e92
	crc_invokeWithoutUpdates                    = 0xbf9459b7
	crc_messages_exportChatInvite               = 0x7d885289
	crc_messages_checkChatInvite                = 0x3eadb1bb
	crc_messages_importChatInvite               = 0x6c50051c
	crc_messages_getStickerSet                  = 0x2619a90e
	crc_messages_installStickerSet              = 0xc78fe460
	crc_messages_uninstallStickerSet            = 0xf96e55de
	crc_messages_startBot                       = 0xe6df7378
	crc_help_getAppChangelog                    = 0xb921197a
	crc_messages_getMessagesViews               = 0xc4c8a55d
	crc_channels_readHistory                    = 0xcc104937
	crc_channels_deleteMessages                 = 0x84c1fd4e
	crc_channels_deleteUserHistory              = 0xd10dd71b
	crc_channels_reportSpam                     = 0xfe087810
	crc_channels_getMessages                    = 0x93d7b347
	crc_channels_getParticipants                = 0x24d98f92
	crc_channels_getParticipant                 = 0x546dd7a6
	crc_channels_getChannels                    = 0x0a7f6bbb
	crc_channels_getFullChannel                 = 0x08736a09
	crc_channels_createChannel                  = 0xf4893d7f
	crc_channels_editAbout                      = 0x13e27f1e
	crc_channels_editAdmin                      = 0xeb7611d0
	crc_channels_editTitle                      = 0x566decd0
	crc_channels_editPhoto                      = 0xf12e57c9
	crc_channels_checkUsername                  = 0x10e6bd2c
	crc_channels_updateUsername                 = 0x3514b3de
	crc_channels_joinChannel                    = 0x24b524c5
	crc_channels_leaveChannel                   = 0xf836aa95
	crc_channels_inviteToChannel                = 0x199f3a6c
	crc_channels_kickFromChannel                = 0xa672de14
	crc_channels_exportInvite                   = 0xc7560885
	crc_channels_deleteChannel                  = 0xc0111fe3
	crc_updates_getChannelDifference            = 0xbb32d7c0
	crc_messages_toggleChatAdmins               = 0xec8bd9e1
	crc_messages_editChatAdmin                  = 0xa9e69f2e
	crc_messages_migrateChat                    = 0x15a3b8e3
	crc_messages_searchGlobal                   = 0x9e3cacb0
	crc_help_getTermsOfService                  = 0x350170f3
	crc_messages_reorderStickerSets             = 0x9fcfbc30
	crc_messages_getDocumentByHash              = 0x338e2464
	crc_messages_searchGifs                     = 0xbf9a776b
	crc_messages_getSavedGifs                   = 0x83bf3d52
	crc_messages_saveGif                        = 0x327a30cb
	crc_messages_getInlineBotResults            = 0x514e999d
	crc_messages_setInlineBotResults            = 0xeb5ea206
	crc_messages_sendInlineBotResult            = 0xb16e06fe
	crc_channels_toggleInvites                  = 0x49609307
	crc_channels_exportMessageLink              = 0xc846d22d
	crc_channels_toggleSignatures               = 0x1f69b606
	crc_channels_updatePinnedMessage            = 0xa72ded52
	crc_auth_resendCode                         = 0x3ef1a9bf
	crc_auth_cancelCode                         = 0x1f040578
	crc_messages_getMessageEditData             = 0xfda68d36
	crc_messages_editMessage                    = 0xce91e4ca
	crc_messages_editInlineBotMessage           = 0x130c2c85
	crc_messages_getBotCallbackAnswer           = 0xa6e94f04
	crc_messages_setBotCallbackAnswer           = 0xc927d44b
	crc_contacts_getTopPeers                    = 0xd4982db5
	crc_contacts_resetTopPeerRating             = 0x1ae373ac
	crc_messages_getPeerDialogs                 = 0x2d9776b9
	crc_messages_saveDraft                      = 0xbc39e14b
	crc_messages_getAllDrafts                   = 0x6a3f8d65
	crc_messages_getFeaturedStickers            = 0x2dacca4f
	crc_messages_readFeaturedStickers           = 0x00e21cbb
	crc_messages_getRecentStickers              = 0x99197c2c
	crc_messages_saveRecentSticker              = 0x348e39bf
	crc_messages_clearRecentStickers            = 0xab02e5d2
	crc_messages_getUnusedStickers              = 0x4309d65b
	crc_messages_getArchivedStickers            = 0x906e241f
	crc_account_sendConfirmPhoneCode            = 0x1516d7bd
	crc_account_confirmPhone                    = 0x5f2178c3
)

type TL_boolFalse struct {
}

type TL_boolTrue struct {
}

type TL_true struct {
}

type TL_error struct {
	code int32
	text string
}

type TL_null struct {
}

type TL_inputPeerEmpty struct {
}

type TL_inputPeerSelf struct {
}

type TL_inputPeerChat struct {
	chat_id int32
}

type TL_inputUserEmpty struct {
}

type TL_inputUserSelf struct {
}

type TL_inputPhoneContact struct {
	client_id  int64
	phone      string
	first_name string
	last_name  string
}

type TL_inputFile struct {
	id           int64
	parts        int32
	name         string
	md5_checksum string
}

type TL_inputMediaEmpty struct {
}

type TL_inputMediaUploadedPhoto struct {
	file    TL // InputFile
	caption string
}

type TL_inputMediaPhoto struct {
	id      TL // InputPhoto
	caption string
}

type TL_inputMediaGeoPoint struct {
	geo_point TL // InputGeoPoint
}

type TL_inputMediaContact struct {
	phone_number string
	first_name   string
	last_name    string
}

type TL_inputChatPhotoEmpty struct {
}

type TL_inputChatUploadedPhoto struct {
	file TL // InputFile
	crop TL // InputPhotoCrop
}

type TL_inputChatPhoto struct {
	id   TL // InputPhoto
	crop TL // InputPhotoCrop
}

type TL_inputGeoPointEmpty struct {
}

type TL_inputGeoPoint struct {
	lat  float64
	long float64
}

type TL_inputPhotoEmpty struct {
}

type TL_inputPhoto struct {
	id          int64
	access_hash int64
}

type TL_inputFileLocation struct {
	volume_id int64
	local_id  int32
	secret    int64
}

type TL_inputPhotoCropAuto struct {
}

type TL_inputPhotoCrop struct {
	crop_left  float64
	crop_top   float64
	crop_width float64
}

type TL_inputAppEvent struct {
	time  float64
	_type string
	peer  int64
	data  string
}

type TL_peerUser struct {
	user_id int32
}

type TL_peerChat struct {
	chat_id int32
}

type TL_storage_fileUnknown struct {
}

type TL_storage_fileJpeg struct {
}

type TL_storage_fileGif struct {
}

type TL_storage_filePng struct {
}

type TL_storage_filePdf struct {
}

type TL_storage_fileMp3 struct {
}

type TL_storage_fileMov struct {
}

type TL_storage_filePartial struct {
}

type TL_storage_fileMp4 struct {
}

type TL_storage_fileWebp struct {
}

type TL_fileLocationUnavailable struct {
	volume_id int64
	local_id  int32
	secret    int64
}

type TL_fileLocation struct {
	dc_id     int32
	volume_id int64
	local_id  int32
	secret    int64
}

type TL_userEmpty struct {
	id int32
}

type TL_userProfilePhotoEmpty struct {
}

type TL_userProfilePhoto struct {
	photo_id    int64
	photo_small TL // FileLocation
	photo_big   TL // FileLocation
}

type TL_userStatusEmpty struct {
}

type TL_userStatusOnline struct {
	expires int32
}

type TL_userStatusOffline struct {
	was_online int32
}

type TL_chatEmpty struct {
	id int32
}

type TL_chat struct {
	flags              int32
	creator            bool `flag_byte:"0"`
	kicked             bool `flag_byte:"1"`
	left               bool `flag_byte:"2"`
	admins_enabled     bool `flag_byte:"3"`
	admin              bool `flag_byte:"4"`
	deactivated        bool `flag_byte:"5"`
	id                 int32
	title              string
	photo              TL // ChatPhoto
	participants_count int32
	date               int32
	version            int32
	migrated_to        TL `flag_byte:"6"` // InputChannel
}

type TL_chatForbidden struct {
	id    int32
	title string
}

type TL_chatFull struct {
	id              int32
	participants    TL   // ChatParticipants
	chat_photo      TL   // Photo
	notify_settings TL   // PeerNotifySettings
	exported_invite TL   // ExportedChatInvite
	bot_info        []TL // BotInfo
}

type TL_chatParticipant struct {
	user_id    int32
	inviter_id int32
	date       int32
}

type TL_chatParticipantsForbidden struct {
	flags            int32
	chat_id          int32
	self_participant TL `flag_byte:"0"` // ChatParticipant
}

type TL_chatParticipants struct {
	chat_id      int32
	participants []TL // ChatParticipant
	version      int32
}

type TL_chatPhotoEmpty struct {
}

type TL_chatPhoto struct {
	photo_small TL // FileLocation
	photo_big   TL // FileLocation
}

type TL_messageEmpty struct {
	id int32
}

type TL_message struct {
	flags           int32
	out             bool `flag_byte:"1"`
	mentioned       bool `flag_byte:"4"`
	media_unread    bool `flag_byte:"5"`
	silent          bool `flag_byte:"13"`
	post            bool `flag_byte:"14"`
	id              int32
	from_id         int32 `flag_byte:"8"`
	to_id           TL    // Peer
	fwd_from        TL    `flag_byte:"2"` // MessageFwdHeader
	via_bot_id      int32 `flag_byte:"11"`
	reply_to_msg_id int32 `flag_byte:"3"`
	date            int32
	message         string
	media           TL    `flag_byte:"9"` // MessageMedia
	reply_markup    TL    `flag_byte:"6"` // ReplyMarkup
	entities        []TL  `flag_byte:"7"` // MessageEntity
	views           int32 `flag_byte:"10"`
	edit_date       int32 `flag_byte:"15"`
}

type TL_messageService struct {
	flags           int32
	out             bool `flag_byte:"1"`
	mentioned       bool `flag_byte:"4"`
	media_unread    bool `flag_byte:"5"`
	silent          bool `flag_byte:"13"`
	post            bool `flag_byte:"14"`
	id              int32
	from_id         int32 `flag_byte:"8"`
	to_id           TL    // Peer
	reply_to_msg_id int32 `flag_byte:"3"`
	date            int32
	action          TL // MessageAction
}

type TL_messageMediaEmpty struct {
}

type TL_messageMediaPhoto struct {
	photo   TL // Photo
	caption string
}

type TL_messageMediaGeo struct {
	geo TL // GeoPoint
}

type TL_messageMediaContact struct {
	phone_number string
	first_name   string
	last_name    string
	user_id      int32
}

type TL_messageMediaUnsupported struct {
}

type TL_messageActionEmpty struct {
}

type TL_messageActionChatCreate struct {
	title string
	users []int32
}

type TL_messageActionChatEditTitle struct {
	title string
}

type TL_messageActionChatEditPhoto struct {
	photo TL // Photo
}

type TL_messageActionChatDeletePhoto struct {
}

type TL_messageActionChatAddUser struct {
	users []int32
}

type TL_messageActionChatDeleteUser struct {
	user_id int32
}

type TL_dialog struct {
	flags              int32
	peer               TL // Peer
	top_message        int32
	read_inbox_max_id  int32
	read_outbox_max_id int32
	unread_count       int32
	notify_settings    TL    // PeerNotifySettings
	pts                int32 `flag_byte:"0"`
	draft              TL    `flag_byte:"1"` // DraftMessage
}

type TL_photoEmpty struct {
	id int64
}

type TL_photo struct {
	id          int64
	access_hash int64
	date        int32
	sizes       []TL // PhotoSize
}

type TL_photoSizeEmpty struct {
	_type string
}

type TL_photoSize struct {
	_type    string
	location TL // FileLocation
	w        int32
	h        int32
	size     int32
}

type TL_photoCachedSize struct {
	_type    string
	location TL // FileLocation
	w        int32
	h        int32
	bytes    []byte
}

type TL_geoPointEmpty struct {
}

type TL_geoPoint struct {
	long float64
	lat  float64
}

type TL_auth_checkedPhone struct {
	phone_registered TL // Bool
}

type TL_auth_sentCode struct {
	flags            int32
	phone_registered bool `flag_byte:"0"`
	_type            TL   // auth_SentCodeType
	phone_code_hash  string
	next_type        TL    `flag_byte:"1"` // auth_CodeType
	timeout          int32 `flag_byte:"2"`
}

type TL_auth_authorization struct {
	flags        int32
	tmp_sessions int32 `flag_byte:"0"`
	user         TL    // User
}

type TL_auth_exportedAuthorization struct {
	id    int32
	bytes []byte
}

type TL_inputNotifyPeer struct {
	peer TL // InputPeer
}

type TL_inputNotifyUsers struct {
}

type TL_inputNotifyChats struct {
}

type TL_inputNotifyAll struct {
}

type TL_inputPeerNotifyEventsEmpty struct {
}

type TL_inputPeerNotifyEventsAll struct {
}

type TL_inputPeerNotifySettings struct {
	flags         int32
	show_previews bool `flag_byte:"0"`
	silent        bool `flag_byte:"1"`
	mute_until    int32
	sound         string
}

type TL_peerNotifyEventsEmpty struct {
}

type TL_peerNotifyEventsAll struct {
}

type TL_peerNotifySettingsEmpty struct {
}

type TL_peerNotifySettings struct {
	flags         int32
	show_previews bool `flag_byte:"0"`
	silent        bool `flag_byte:"1"`
	mute_until    int32
	sound         string
}

type TL_peerSettings struct {
	flags       int32
	report_spam bool `flag_byte:"0"`
}

type TL_wallPaper struct {
	id    int32
	title string
	sizes []TL // PhotoSize
	color int32
}

type TL_inputReportReasonSpam struct {
}

type TL_inputReportReasonViolence struct {
}

type TL_inputReportReasonPornography struct {
}

type TL_inputReportReasonOther struct {
	text string
}

type TL_userFull struct {
	flags           int32
	blocked         bool   `flag_byte:"0"`
	user            TL     // User
	about           string `flag_byte:"1"`
	link            TL     // contacts_Link
	profile_photo   TL     `flag_byte:"2"` // Photo
	notify_settings TL     // PeerNotifySettings
	bot_info        TL     `flag_byte:"3"` // BotInfo
}

type TL_contact struct {
	user_id int32
	mutual  TL // Bool
}

type TL_importedContact struct {
	user_id   int32
	client_id int64
}

type TL_contactBlocked struct {
	user_id int32
	date    int32
}

type TL_contactStatus struct {
	user_id int32
	status  TL // UserStatus
}

type TL_contacts_link struct {
	my_link      TL // ContactLink
	foreign_link TL // ContactLink
	user         TL // User
}

type TL_contacts_contactsNotModified struct {
}

type TL_contacts_contacts struct {
	contacts []TL // Contact
	users    []TL // User
}

type TL_contacts_importedContacts struct {
	imported       []TL // ImportedContact
	retry_contacts []int64
	users          []TL // User
}

type TL_contacts_blocked struct {
	blocked []TL // ContactBlocked
	users   []TL // User
}

type TL_contacts_blockedSlice struct {
	count   int32
	blocked []TL // ContactBlocked
	users   []TL // User
}

type TL_messages_dialogs struct {
	dialogs  []TL // Dialog
	messages []TL // Message
	chats    []TL // Chat
	users    []TL // User
}

type TL_messages_dialogsSlice struct {
	count    int32
	dialogs  []TL // Dialog
	messages []TL // Message
	chats    []TL // Chat
	users    []TL // User
}

type TL_messages_messages struct {
	messages []TL // Message
	chats    []TL // Chat
	users    []TL // User
}

type TL_messages_messagesSlice struct {
	count    int32
	messages []TL // Message
	chats    []TL // Chat
	users    []TL // User
}

type TL_messages_chats struct {
	chats []TL // Chat
}

type TL_messages_chatFull struct {
	full_chat TL   // ChatFull
	chats     []TL // Chat
	users     []TL // User
}

type TL_messages_affectedHistory struct {
	pts       int32
	pts_count int32
	offset    int32
}

type TL_inputMessagesFilterEmpty struct {
}

type TL_inputMessagesFilterPhotos struct {
}

type TL_inputMessagesFilterVideo struct {
}

type TL_inputMessagesFilterPhotoVideo struct {
}

type TL_inputMessagesFilterPhotoVideoDocuments struct {
}

type TL_inputMessagesFilterDocument struct {
}

type TL_inputMessagesFilterUrl struct {
}

type TL_inputMessagesFilterGif struct {
}

type TL_updateNewMessage struct {
	message   TL // Message
	pts       int32
	pts_count int32
}

type TL_updateMessageID struct {
	id        int32
	random_id int64
}

type TL_updateDeleteMessages struct {
	messages  []int32
	pts       int32
	pts_count int32
}

type TL_updateUserTyping struct {
	user_id int32
	action  TL // SendMessageAction
}

type TL_updateChatUserTyping struct {
	chat_id int32
	user_id int32
	action  TL // SendMessageAction
}

type TL_updateChatParticipants struct {
	participants TL // ChatParticipants
}

type TL_updateUserStatus struct {
	user_id int32
	status  TL // UserStatus
}

type TL_updateUserName struct {
	user_id    int32
	first_name string
	last_name  string
	username   string
}

type TL_updateUserPhoto struct {
	user_id  int32
	date     int32
	photo    TL // UserProfilePhoto
	previous TL // Bool
}

type TL_updateContactRegistered struct {
	user_id int32
	date    int32
}

type TL_updateContactLink struct {
	user_id      int32
	my_link      TL // ContactLink
	foreign_link TL // ContactLink
}

type TL_updateNewAuthorization struct {
	auth_key_id int64
	date        int32
	device      string
	location    string
}

type TL_updates_state struct {
	pts          int32
	qts          int32
	date         int32
	seq          int32
	unread_count int32
}

type TL_updates_differenceEmpty struct {
	date int32
	seq  int32
}

type TL_updates_difference struct {
	new_messages           []TL // Message
	new_encrypted_messages []TL // EncryptedMessage
	other_updates          []TL // Update
	chats                  []TL // Chat
	users                  []TL // User
	state                  TL   // updates_State
}

type TL_updates_differenceSlice struct {
	new_messages           []TL // Message
	new_encrypted_messages []TL // EncryptedMessage
	other_updates          []TL // Update
	chats                  []TL // Chat
	users                  []TL // User
	intermediate_state     TL   // updates_State
}

type TL_updatesTooLong struct {
}

type TL_updateShortMessage struct {
	flags           int32
	out             bool `flag_byte:"1"`
	mentioned       bool `flag_byte:"4"`
	media_unread    bool `flag_byte:"5"`
	silent          bool `flag_byte:"13"`
	id              int32
	user_id         int32
	message         string
	pts             int32
	pts_count       int32
	date            int32
	fwd_from        TL    `flag_byte:"2"` // MessageFwdHeader
	via_bot_id      int32 `flag_byte:"11"`
	reply_to_msg_id int32 `flag_byte:"3"`
	entities        []TL  `flag_byte:"7"` // MessageEntity
}

type TL_updateShortChatMessage struct {
	flags           int32
	out             bool `flag_byte:"1"`
	mentioned       bool `flag_byte:"4"`
	media_unread    bool `flag_byte:"5"`
	silent          bool `flag_byte:"13"`
	id              int32
	from_id         int32
	chat_id         int32
	message         string
	pts             int32
	pts_count       int32
	date            int32
	fwd_from        TL    `flag_byte:"2"` // MessageFwdHeader
	via_bot_id      int32 `flag_byte:"11"`
	reply_to_msg_id int32 `flag_byte:"3"`
	entities        []TL  `flag_byte:"7"` // MessageEntity
}

type TL_updateShort struct {
	update TL // Update
	date   int32
}

type TL_updatesCombined struct {
	updates   []TL // Update
	users     []TL // User
	chats     []TL // Chat
	date      int32
	seq_start int32
	seq       int32
}

type TL_updates struct {
	updates []TL // Update
	users   []TL // User
	chats   []TL // Chat
	date    int32
	seq     int32
}

type TL_photos_photos struct {
	photos []TL // Photo
	users  []TL // User
}

type TL_photos_photosSlice struct {
	count  int32
	photos []TL // Photo
	users  []TL // User
}

type TL_photos_photo struct {
	photo TL   // Photo
	users []TL // User
}

type TL_upload_file struct {
	_type TL // storage_FileType
	mtime int32
	bytes []byte
}

type TL_dcOption struct {
	flags      int32
	ipv6       bool `flag_byte:"0"`
	media_only bool `flag_byte:"1"`
	tcpo_only  bool `flag_byte:"2"`
	id         int32
	ip_address string
	port       int32
}

type TL_config struct {
	flags                   int32
	date                    int32
	expires                 int32
	test_mode               TL // Bool
	this_dc                 int32
	dc_options              []TL // DcOption
	chat_size_max           int32
	megagroup_size_max      int32
	forwarded_count_max     int32
	online_update_period_ms int32
	offline_blur_timeout_ms int32
	offline_idle_timeout_ms int32
	online_cloud_timeout_ms int32
	notify_cloud_delay_ms   int32
	notify_default_delay_ms int32
	chat_big_size           int32
	push_chat_period_ms     int32
	push_chat_limit         int32
	saved_gifs_limit        int32
	edit_time_limit         int32
	rating_e_decay          int32
	stickers_recent_limit   int32
	tmp_sessions            int32 `flag_byte:"0"`
	disabled_features       []TL  // DisabledFeature
}

type TL_nearestDc struct {
	country    string
	this_dc    int32
	nearest_dc int32
}

type TL_help_appUpdate struct {
	id       int32
	critical TL // Bool
	url      string
	text     string
}

type TL_help_noAppUpdate struct {
}

type TL_help_inviteText struct {
	message string
}

type TL_wallPaperSolid struct {
	id       int32
	title    string
	bg_color int32
	color    int32
}

type TL_updateNewEncryptedMessage struct {
	message TL // EncryptedMessage
	qts     int32
}

type TL_updateEncryptedChatTyping struct {
	chat_id int32
}

type TL_updateEncryption struct {
	chat TL // EncryptedChat
	date int32
}

type TL_updateEncryptedMessagesRead struct {
	chat_id  int32
	max_date int32
	date     int32
}

type TL_encryptedChatEmpty struct {
	id int32
}

type TL_encryptedChatWaiting struct {
	id             int32
	access_hash    int64
	date           int32
	admin_id       int32
	participant_id int32
}

type TL_encryptedChatRequested struct {
	id             int32
	access_hash    int64
	date           int32
	admin_id       int32
	participant_id int32
	g_a            []byte
}

type TL_encryptedChat struct {
	id              int32
	access_hash     int64
	date            int32
	admin_id        int32
	participant_id  int32
	g_a_or_b        []byte
	key_fingerprint int64
}

type TL_encryptedChatDiscarded struct {
	id int32
}

type TL_inputEncryptedChat struct {
	chat_id     int32
	access_hash int64
}

type TL_encryptedFileEmpty struct {
}

type TL_encryptedFile struct {
	id              int64
	access_hash     int64
	size            int32
	dc_id           int32
	key_fingerprint int32
}

type TL_inputEncryptedFileEmpty struct {
}

type TL_inputEncryptedFileUploaded struct {
	id              int64
	parts           int32
	md5_checksum    string
	key_fingerprint int32
}

type TL_inputEncryptedFile struct {
	id          int64
	access_hash int64
}

type TL_inputEncryptedFileLocation struct {
	id          int64
	access_hash int64
}

type TL_encryptedMessage struct {
	random_id int64
	chat_id   int32
	date      int32
	bytes     []byte
	file      TL // EncryptedFile
}

type TL_encryptedMessageService struct {
	random_id int64
	chat_id   int32
	date      int32
	bytes     []byte
}

type TL_messages_dhConfigNotModified struct {
	random []byte
}

type TL_messages_dhConfig struct {
	g       int32
	p       []byte
	version int32
	random  []byte
}

type TL_messages_sentEncryptedMessage struct {
	date int32
}

type TL_messages_sentEncryptedFile struct {
	date int32
	file TL // EncryptedFile
}

type TL_inputFileBig struct {
	id    int64
	parts int32
	name  string
}

type TL_inputEncryptedFileBigUploaded struct {
	id              int64
	parts           int32
	key_fingerprint int32
}

type TL_updateChatParticipantAdd struct {
	chat_id    int32
	user_id    int32
	inviter_id int32
	date       int32
	version    int32
}

type TL_updateChatParticipantDelete struct {
	chat_id int32
	user_id int32
	version int32
}

type TL_updateDcOptions struct {
	dc_options []TL // DcOption
}

type TL_inputMediaUploadedDocument struct {
	file       TL // InputFile
	mime_type  string
	attributes []TL // DocumentAttribute
	caption    string
}

type TL_inputMediaUploadedThumbDocument struct {
	file       TL // InputFile
	thumb      TL // InputFile
	mime_type  string
	attributes []TL // DocumentAttribute
	caption    string
}

type TL_inputMediaDocument struct {
	id      TL // InputDocument
	caption string
}

type TL_messageMediaDocument struct {
	document TL // Document
	caption  string
}

type TL_inputDocumentEmpty struct {
}

type TL_inputDocument struct {
	id          int64
	access_hash int64
}

type TL_inputDocumentFileLocation struct {
	id          int64
	access_hash int64
	version     int32
}

type TL_documentEmpty struct {
	id int64
}

type TL_document struct {
	id          int64
	access_hash int64
	date        int32
	mime_type   string
	size        int32
	thumb       TL // PhotoSize
	dc_id       int32
	version     int32
	attributes  []TL // DocumentAttribute
}

type TL_help_support struct {
	phone_number string
	user         TL // User
}

type TL_notifyPeer struct {
	peer TL // Peer
}

type TL_notifyUsers struct {
}

type TL_notifyChats struct {
}

type TL_notifyAll struct {
}

type TL_updateUserBlocked struct {
	user_id int32
	blocked TL // Bool
}

type TL_updateNotifySettings struct {
	peer            TL // NotifyPeer
	notify_settings TL // PeerNotifySettings
}

type TL_sendMessageTypingAction struct {
}

type TL_sendMessageCancelAction struct {
}

type TL_sendMessageRecordVideoAction struct {
}

type TL_sendMessageUploadVideoAction struct {
	progress int32
}

type TL_sendMessageRecordAudioAction struct {
}

type TL_sendMessageUploadAudioAction struct {
	progress int32
}

type TL_sendMessageUploadPhotoAction struct {
	progress int32
}

type TL_sendMessageUploadDocumentAction struct {
	progress int32
}

type TL_sendMessageGeoLocationAction struct {
}

type TL_sendMessageChooseContactAction struct {
}

type TL_contacts_found struct {
	results []TL // Peer
	chats   []TL // Chat
	users   []TL // User
}

type TL_updateServiceNotification struct {
	_type   string
	message string
	media   TL // MessageMedia
	popup   TL // Bool
}

type TL_userStatusRecently struct {
}

type TL_userStatusLastWeek struct {
}

type TL_userStatusLastMonth struct {
}

type TL_updatePrivacy struct {
	key   TL   // PrivacyKey
	rules []TL // PrivacyRule
}

type TL_inputPrivacyKeyStatusTimestamp struct {
}

type TL_privacyKeyStatusTimestamp struct {
}

type TL_inputPrivacyValueAllowContacts struct {
}

type TL_inputPrivacyValueAllowAll struct {
}

type TL_inputPrivacyValueAllowUsers struct {
	users []TL // InputUser
}

type TL_inputPrivacyValueDisallowContacts struct {
}

type TL_inputPrivacyValueDisallowAll struct {
}

type TL_inputPrivacyValueDisallowUsers struct {
	users []TL // InputUser
}

type TL_privacyValueAllowContacts struct {
}

type TL_privacyValueAllowAll struct {
}

type TL_privacyValueAllowUsers struct {
	users []int32
}

type TL_privacyValueDisallowContacts struct {
}

type TL_privacyValueDisallowAll struct {
}

type TL_privacyValueDisallowUsers struct {
	users []int32
}

type TL_account_privacyRules struct {
	rules []TL // PrivacyRule
	users []TL // User
}

type TL_accountDaysTTL struct {
	days int32
}

type TL_updateUserPhone struct {
	user_id int32
	phone   string
}

type TL_documentAttributeImageSize struct {
	w int32
	h int32
}

type TL_documentAttributeAnimated struct {
}

type TL_documentAttributeSticker struct {
	alt        string
	stickerset TL // InputStickerSet
}

type TL_documentAttributeVideo struct {
	duration int32
	w        int32
	h        int32
}

type TL_documentAttributeAudio struct {
	flags     int32
	voice     bool `flag_byte:"10"`
	duration  int32
	title     string `flag_byte:"0"`
	performer string `flag_byte:"1"`
	waveform  []byte `flag_byte:"2"`
}

type TL_documentAttributeFilename struct {
	file_name string
}

type TL_messages_stickersNotModified struct {
}

type TL_messages_stickers struct {
	hash     string
	stickers []TL // Document
}

type TL_stickerPack struct {
	emoticon  string
	documents []int64
}

type TL_messages_allStickersNotModified struct {
}

type TL_messages_allStickers struct {
	hash int32
	sets []TL // StickerSet
}

type TL_disabledFeature struct {
	feature     string
	description string
}

type TL_updateReadHistoryInbox struct {
	peer      TL // Peer
	max_id    int32
	pts       int32
	pts_count int32
}

type TL_updateReadHistoryOutbox struct {
	peer      TL // Peer
	max_id    int32
	pts       int32
	pts_count int32
}

type TL_messages_affectedMessages struct {
	pts       int32
	pts_count int32
}

type TL_contactLinkUnknown struct {
}

type TL_contactLinkNone struct {
}

type TL_contactLinkHasPhone struct {
}

type TL_contactLinkContact struct {
}

type TL_updateWebPage struct {
	webpage   TL // WebPage
	pts       int32
	pts_count int32
}

type TL_webPageEmpty struct {
	id int64
}

type TL_webPagePending struct {
	id   int64
	date int32
}

type TL_webPage struct {
	flags        int32
	id           int64
	url          string
	display_url  string
	_type        string `flag_byte:"0"`
	site_name    string `flag_byte:"1"`
	title        string `flag_byte:"2"`
	description  string `flag_byte:"3"`
	photo        TL     `flag_byte:"4"` // Photo
	embed_url    string `flag_byte:"5"`
	embed_type   string `flag_byte:"5"`
	embed_width  int32  `flag_byte:"6"`
	embed_height int32  `flag_byte:"6"`
	duration     int32  `flag_byte:"7"`
	author       string `flag_byte:"8"`
	document     TL     `flag_byte:"9"` // Document
}

type TL_messageMediaWebPage struct {
	webpage TL // WebPage
}

type TL_authorization struct {
	hash           int64
	flags          int32
	device_model   string
	platform       string
	system_version string
	api_id         int32
	app_name       string
	app_version    string
	date_created   int32
	date_active    int32
	ip             string
	country        string
	region         string
}

type TL_account_authorizations struct {
	authorizations []TL // Authorization
}

type TL_account_noPassword struct {
	new_salt                  []byte
	email_unconfirmed_pattern string
}

type TL_account_password struct {
	current_salt              []byte
	new_salt                  []byte
	hint                      string
	has_recovery              TL // Bool
	email_unconfirmed_pattern string
}

type TL_account_passwordSettings struct {
	email string
}

type TL_account_passwordInputSettings struct {
	flags             int32
	new_salt          []byte `flag_byte:"0"`
	new_password_hash []byte `flag_byte:"0"`
	hint              string `flag_byte:"0"`
	email             string `flag_byte:"1"`
}

type TL_auth_passwordRecovery struct {
	email_pattern string
}

type TL_inputMediaVenue struct {
	geo_point TL // InputGeoPoint
	title     string
	address   string
	provider  string
	venue_id  string
}

type TL_messageMediaVenue struct {
	geo      TL // GeoPoint
	title    string
	address  string
	provider string
	venue_id string
}

type TL_receivedNotifyMessage struct {
	id    int32
	flags int32
}

type TL_chatInviteEmpty struct {
}

type TL_chatInviteExported struct {
	link string
}

type TL_chatInviteAlready struct {
	chat TL // Chat
}

type TL_chatInvite struct {
	flags              int32
	channel            bool `flag_byte:"0"`
	broadcast          bool `flag_byte:"1"`
	public             bool `flag_byte:"2"`
	megagroup          bool `flag_byte:"3"`
	title              string
	photo              TL // ChatPhoto
	participants_count int32
	participants       []TL `flag_byte:"4"` // User
}

type TL_messageActionChatJoinedByLink struct {
	inviter_id int32
}

type TL_updateReadMessagesContents struct {
	messages  []int32
	pts       int32
	pts_count int32
}

type TL_inputStickerSetEmpty struct {
}

type TL_inputStickerSetID struct {
	id          int64
	access_hash int64
}

type TL_inputStickerSetShortName struct {
	short_name string
}

type TL_stickerSet struct {
	flags       int32
	installed   bool `flag_byte:"0"`
	archived    bool `flag_byte:"1"`
	official    bool `flag_byte:"2"`
	id          int64
	access_hash int64
	title       string
	short_name  string
	count       int32
	hash        int32
}

type TL_messages_stickerSet struct {
	set       TL   // StickerSet
	packs     []TL // StickerPack
	documents []TL // Document
}

type TL_user struct {
	flags                  int32
	self                   bool `flag_byte:"10"`
	contact                bool `flag_byte:"11"`
	mutual_contact         bool `flag_byte:"12"`
	deleted                bool `flag_byte:"13"`
	bot                    bool `flag_byte:"14"`
	bot_chat_history       bool `flag_byte:"15"`
	bot_nochats            bool `flag_byte:"16"`
	verified               bool `flag_byte:"17"`
	restricted             bool `flag_byte:"18"`
	min                    bool `flag_byte:"20"`
	bot_inline_geo         bool `flag_byte:"21"`
	id                     int32
	access_hash            int64  `flag_byte:"0"`
	first_name             string `flag_byte:"1"`
	last_name              string `flag_byte:"2"`
	username               string `flag_byte:"3"`
	phone                  string `flag_byte:"4"`
	photo                  TL     `flag_byte:"5"` // UserProfilePhoto
	status                 TL     `flag_byte:"6"` // UserStatus
	bot_info_version       int32  `flag_byte:"14"`
	restriction_reason     string `flag_byte:"18"`
	bot_inline_placeholder string `flag_byte:"19"`
}

type TL_botCommand struct {
	command     string
	description string
}

type TL_botInfo struct {
	user_id     int32
	description string
	commands    []TL // BotCommand
}

type TL_keyboardButton struct {
	text string
}

type TL_keyboardButtonRow struct {
	buttons []TL // KeyboardButton
}

type TL_replyKeyboardHide struct {
	flags     int32
	selective bool `flag_byte:"2"`
}

type TL_replyKeyboardForceReply struct {
	flags      int32
	single_use bool `flag_byte:"1"`
	selective  bool `flag_byte:"2"`
}

type TL_replyKeyboardMarkup struct {
	flags      int32
	resize     bool `flag_byte:"0"`
	single_use bool `flag_byte:"1"`
	selective  bool `flag_byte:"2"`
	rows       []TL // KeyboardButtonRow
}

type TL_inputPeerUser struct {
	user_id     int32
	access_hash int64
}

type TL_inputUser struct {
	user_id     int32
	access_hash int64
}

type TL_help_appChangelogEmpty struct {
}

type TL_help_appChangelog struct {
	text string
}

type TL_messageEntityUnknown struct {
	offset int32
	length int32
}

type TL_messageEntityMention struct {
	offset int32
	length int32
}

type TL_messageEntityHashtag struct {
	offset int32
	length int32
}

type TL_messageEntityBotCommand struct {
	offset int32
	length int32
}

type TL_messageEntityUrl struct {
	offset int32
	length int32
}

type TL_messageEntityEmail struct {
	offset int32
	length int32
}

type TL_messageEntityBold struct {
	offset int32
	length int32
}

type TL_messageEntityItalic struct {
	offset int32
	length int32
}

type TL_messageEntityCode struct {
	offset int32
	length int32
}

type TL_messageEntityPre struct {
	offset   int32
	length   int32
	language string
}

type TL_messageEntityTextUrl struct {
	offset int32
	length int32
	url    string
}

type TL_updateShortSentMessage struct {
	flags     int32
	out       bool `flag_byte:"1"`
	id        int32
	pts       int32
	pts_count int32
	date      int32
	media     TL   `flag_byte:"9"` // MessageMedia
	entities  []TL `flag_byte:"7"` // MessageEntity
}

type TL_inputChannelEmpty struct {
}

type TL_inputChannel struct {
	channel_id  int32
	access_hash int64
}

type TL_peerChannel struct {
	channel_id int32
}

type TL_inputPeerChannel struct {
	channel_id  int32
	access_hash int64
}

type TL_channel struct {
	flags              int32
	creator            bool `flag_byte:"0"`
	kicked             bool `flag_byte:"1"`
	left               bool `flag_byte:"2"`
	editor             bool `flag_byte:"3"`
	moderator          bool `flag_byte:"4"`
	broadcast          bool `flag_byte:"5"`
	verified           bool `flag_byte:"7"`
	megagroup          bool `flag_byte:"8"`
	restricted         bool `flag_byte:"9"`
	democracy          bool `flag_byte:"10"`
	signatures         bool `flag_byte:"11"`
	min                bool `flag_byte:"12"`
	id                 int32
	access_hash        int64 `flag_byte:"13"`
	title              string
	username           string `flag_byte:"6"`
	photo              TL     // ChatPhoto
	date               int32
	version            int32
	restriction_reason string `flag_byte:"9"`
}

type TL_channelForbidden struct {
	flags       int32
	broadcast   bool `flag_byte:"5"`
	megagroup   bool `flag_byte:"8"`
	id          int32
	access_hash int64
	title       string
}

type TL_contacts_resolvedPeer struct {
	peer  TL   // Peer
	chats []TL // Chat
	users []TL // User
}

type TL_channelFull struct {
	flags                 int32
	can_view_participants bool `flag_byte:"3"`
	can_set_username      bool `flag_byte:"6"`
	id                    int32
	about                 string
	participants_count    int32 `flag_byte:"0"`
	admins_count          int32 `flag_byte:"1"`
	kicked_count          int32 `flag_byte:"2"`
	read_inbox_max_id     int32
	read_outbox_max_id    int32
	unread_count          int32
	chat_photo            TL    // Photo
	notify_settings       TL    // PeerNotifySettings
	exported_invite       TL    // ExportedChatInvite
	bot_info              []TL  // BotInfo
	migrated_from_chat_id int32 `flag_byte:"4"`
	migrated_from_max_id  int32 `flag_byte:"4"`
	pinned_msg_id         int32 `flag_byte:"5"`
}

type TL_messageRange struct {
	min_id int32
	max_id int32
}

type TL_messages_channelMessages struct {
	flags    int32
	pts      int32
	count    int32
	messages []TL // Message
	chats    []TL // Chat
	users    []TL // User
}

type TL_messageActionChannelCreate struct {
	title string
}

type TL_updateChannelTooLong struct {
	flags      int32
	channel_id int32
	pts        int32 `flag_byte:"0"`
}

type TL_updateChannel struct {
	channel_id int32
}

type TL_updateNewChannelMessage struct {
	message   TL // Message
	pts       int32
	pts_count int32
}

type TL_updateReadChannelInbox struct {
	channel_id int32
	max_id     int32
}

type TL_updateDeleteChannelMessages struct {
	channel_id int32
	messages   []int32
	pts        int32
	pts_count  int32
}

type TL_updateChannelMessageViews struct {
	channel_id int32
	id         int32
	views      int32
}

type TL_updates_channelDifferenceEmpty struct {
	flags   int32
	final   bool `flag_byte:"0"`
	pts     int32
	timeout int32 `flag_byte:"1"`
}

type TL_updates_channelDifferenceTooLong struct {
	flags              int32
	final              bool `flag_byte:"0"`
	pts                int32
	timeout            int32 `flag_byte:"1"`
	top_message        int32
	read_inbox_max_id  int32
	read_outbox_max_id int32
	unread_count       int32
	messages           []TL // Message
	chats              []TL // Chat
	users              []TL // User
}

type TL_updates_channelDifference struct {
	flags         int32
	final         bool `flag_byte:"0"`
	pts           int32
	timeout       int32 `flag_byte:"1"`
	new_messages  []TL  // Message
	other_updates []TL  // Update
	chats         []TL  // Chat
	users         []TL  // User
}

type TL_channelMessagesFilterEmpty struct {
}

type TL_channelMessagesFilter struct {
	flags                int32
	exclude_new_messages bool `flag_byte:"1"`
	ranges               []TL // MessageRange
}

type TL_channelParticipant struct {
	user_id int32
	date    int32
}

type TL_channelParticipantSelf struct {
	user_id    int32
	inviter_id int32
	date       int32
}

type TL_channelParticipantModerator struct {
	user_id    int32
	inviter_id int32
	date       int32
}

type TL_channelParticipantEditor struct {
	user_id    int32
	inviter_id int32
	date       int32
}

type TL_channelParticipantKicked struct {
	user_id   int32
	kicked_by int32
	date      int32
}

type TL_channelParticipantCreator struct {
	user_id int32
}

type TL_channelParticipantsRecent struct {
}

type TL_channelParticipantsAdmins struct {
}

type TL_channelParticipantsKicked struct {
}

type TL_channelRoleEmpty struct {
}

type TL_channelRoleModerator struct {
}

type TL_channelRoleEditor struct {
}

type TL_channels_channelParticipants struct {
	count        int32
	participants []TL // ChannelParticipant
	users        []TL // User
}

type TL_channels_channelParticipant struct {
	participant TL   // ChannelParticipant
	users       []TL // User
}

type TL_chatParticipantCreator struct {
	user_id int32
}

type TL_chatParticipantAdmin struct {
	user_id    int32
	inviter_id int32
	date       int32
}

type TL_updateChatAdmins struct {
	chat_id int32
	enabled TL // Bool
	version int32
}

type TL_updateChatParticipantAdmin struct {
	chat_id  int32
	user_id  int32
	is_admin TL // Bool
	version  int32
}

type TL_messageActionChatMigrateTo struct {
	channel_id int32
}

type TL_messageActionChannelMigrateFrom struct {
	title   string
	chat_id int32
}

type TL_channelParticipantsBots struct {
}

type TL_help_termsOfService struct {
	text string
}

type TL_updateNewStickerSet struct {
	stickerset TL // messages_StickerSet
}

type TL_updateStickerSetsOrder struct {
	order []int64
}

type TL_updateStickerSets struct {
}

type TL_foundGif struct {
	url          string
	thumb_url    string
	content_url  string
	content_type string
	w            int32
	h            int32
}

type TL_foundGifCached struct {
	url      string
	photo    TL // Photo
	document TL // Document
}

type TL_inputMediaGifExternal struct {
	url string
	q   string
}

type TL_messages_foundGifs struct {
	next_offset int32
	results     []TL // FoundGif
}

type TL_messages_savedGifsNotModified struct {
}

type TL_messages_savedGifs struct {
	hash int32
	gifs []TL // Document
}

type TL_updateSavedGifs struct {
}

type TL_inputBotInlineMessageMediaAuto struct {
	flags        int32
	caption      string
	reply_markup TL `flag_byte:"2"` // ReplyMarkup
}

type TL_inputBotInlineMessageText struct {
	flags        int32
	no_webpage   bool `flag_byte:"0"`
	message      string
	entities     []TL `flag_byte:"1"` // MessageEntity
	reply_markup TL   `flag_byte:"2"` // ReplyMarkup
}

type TL_inputBotInlineResult struct {
	flags        int32
	id           string
	_type        string
	title        string `flag_byte:"1"`
	description  string `flag_byte:"2"`
	url          string `flag_byte:"3"`
	thumb_url    string `flag_byte:"4"`
	content_url  string `flag_byte:"5"`
	content_type string `flag_byte:"5"`
	w            int32  `flag_byte:"6"`
	h            int32  `flag_byte:"6"`
	duration     int32  `flag_byte:"7"`
	send_message TL     // InputBotInlineMessage
}

type TL_botInlineMessageMediaAuto struct {
	flags        int32
	caption      string
	reply_markup TL `flag_byte:"2"` // ReplyMarkup
}

type TL_botInlineMessageText struct {
	flags        int32
	no_webpage   bool `flag_byte:"0"`
	message      string
	entities     []TL `flag_byte:"1"` // MessageEntity
	reply_markup TL   `flag_byte:"2"` // ReplyMarkup
}

type TL_botInlineResult struct {
	flags        int32
	id           string
	_type        string
	title        string `flag_byte:"1"`
	description  string `flag_byte:"2"`
	url          string `flag_byte:"3"`
	thumb_url    string `flag_byte:"4"`
	content_url  string `flag_byte:"5"`
	content_type string `flag_byte:"5"`
	w            int32  `flag_byte:"6"`
	h            int32  `flag_byte:"6"`
	duration     int32  `flag_byte:"7"`
	send_message TL     // BotInlineMessage
}

type TL_messages_botResults struct {
	flags       int32
	gallery     bool `flag_byte:"0"`
	query_id    int64
	next_offset string `flag_byte:"1"`
	switch_pm   TL     `flag_byte:"2"` // InlineBotSwitchPM
	results     []TL   // BotInlineResult
}

type TL_updateBotInlineQuery struct {
	flags    int32
	query_id int64
	user_id  int32
	query    string
	geo      TL `flag_byte:"0"` // GeoPoint
	offset   string
}

type TL_updateBotInlineSend struct {
	flags   int32
	user_id int32
	query   string
	geo     TL `flag_byte:"0"` // GeoPoint
	id      string
	msg_id  TL `flag_byte:"1"` // InputBotInlineMessageID
}

type TL_inputMessagesFilterVoice struct {
}

type TL_inputMessagesFilterMusic struct {
}

type TL_inputPrivacyKeyChatInvite struct {
}

type TL_privacyKeyChatInvite struct {
}

type TL_exportedMessageLink struct {
	link string
}

type TL_messageFwdHeader struct {
	flags        int32
	from_id      int32 `flag_byte:"0"`
	date         int32
	channel_id   int32 `flag_byte:"1"`
	channel_post int32 `flag_byte:"2"`
}

type TL_updateEditChannelMessage struct {
	message   TL // Message
	pts       int32
	pts_count int32
}

type TL_updateChannelPinnedMessage struct {
	channel_id int32
	id         int32
}

type TL_messageActionPinMessage struct {
}

type TL_auth_codeTypeSms struct {
}

type TL_auth_codeTypeCall struct {
}

type TL_auth_codeTypeFlashCall struct {
}

type TL_auth_sentCodeTypeApp struct {
	length int32
}

type TL_auth_sentCodeTypeSms struct {
	length int32
}

type TL_auth_sentCodeTypeCall struct {
	length int32
}

type TL_auth_sentCodeTypeFlashCall struct {
	pattern string
}

type TL_keyboardButtonUrl struct {
	text string
	url  string
}

type TL_keyboardButtonCallback struct {
	text string
	data []byte
}

type TL_keyboardButtonRequestPhone struct {
	text string
}

type TL_keyboardButtonRequestGeoLocation struct {
	text string
}

type TL_keyboardButtonSwitchInline struct {
	text  string
	query string
}

type TL_replyInlineMarkup struct {
	rows []TL // KeyboardButtonRow
}

type TL_messages_botCallbackAnswer struct {
	flags   int32
	alert   bool   `flag_byte:"1"`
	message string `flag_byte:"0"`
	url     string `flag_byte:"2"`
}

type TL_updateBotCallbackQuery struct {
	query_id int64
	user_id  int32
	peer     TL // Peer
	msg_id   int32
	data     []byte
}

type TL_messages_messageEditData struct {
	flags   int32
	caption bool `flag_byte:"0"`
}

type TL_updateEditMessage struct {
	message   TL // Message
	pts       int32
	pts_count int32
}

type TL_inputBotInlineMessageMediaGeo struct {
	flags        int32
	geo_point    TL // InputGeoPoint
	reply_markup TL `flag_byte:"2"` // ReplyMarkup
}

type TL_inputBotInlineMessageMediaVenue struct {
	flags        int32
	geo_point    TL // InputGeoPoint
	title        string
	address      string
	provider     string
	venue_id     string
	reply_markup TL `flag_byte:"2"` // ReplyMarkup
}

type TL_inputBotInlineMessageMediaContact struct {
	flags        int32
	phone_number string
	first_name   string
	last_name    string
	reply_markup TL `flag_byte:"2"` // ReplyMarkup
}

type TL_botInlineMessageMediaGeo struct {
	flags        int32
	geo          TL // GeoPoint
	reply_markup TL `flag_byte:"2"` // ReplyMarkup
}

type TL_botInlineMessageMediaVenue struct {
	flags        int32
	geo          TL // GeoPoint
	title        string
	address      string
	provider     string
	venue_id     string
	reply_markup TL `flag_byte:"2"` // ReplyMarkup
}

type TL_botInlineMessageMediaContact struct {
	flags        int32
	phone_number string
	first_name   string
	last_name    string
	reply_markup TL `flag_byte:"2"` // ReplyMarkup
}

type TL_inputBotInlineResultPhoto struct {
	id           string
	_type        string
	photo        TL // InputPhoto
	send_message TL // InputBotInlineMessage
}

type TL_inputBotInlineResultDocument struct {
	flags        int32
	id           string
	_type        string
	title        string `flag_byte:"1"`
	description  string `flag_byte:"2"`
	document     TL     // InputDocument
	send_message TL     // InputBotInlineMessage
}

type TL_botInlineMediaResult struct {
	flags        int32
	id           string
	_type        string
	photo        TL     `flag_byte:"0"` // Photo
	document     TL     `flag_byte:"1"` // Document
	title        string `flag_byte:"2"`
	description  string `flag_byte:"3"`
	send_message TL     // BotInlineMessage
}

type TL_inputBotInlineMessageID struct {
	dc_id       int32
	id          int64
	access_hash int64
}

type TL_updateInlineBotCallbackQuery struct {
	query_id int64
	user_id  int32
	msg_id   TL // InputBotInlineMessageID
	data     []byte
}

type TL_inlineBotSwitchPM struct {
	text        string
	start_param string
}

type TL_messages_peerDialogs struct {
	dialogs  []TL // Dialog
	messages []TL // Message
	chats    []TL // Chat
	users    []TL // User
	state    TL   // updates_State
}

type TL_topPeer struct {
	peer   TL // Peer
	rating float64
}

type TL_topPeerCategoryBotsPM struct {
}

type TL_topPeerCategoryBotsInline struct {
}

type TL_topPeerCategoryCorrespondents struct {
}

type TL_topPeerCategoryGroups struct {
}

type TL_topPeerCategoryChannels struct {
}

type TL_topPeerCategoryPeers struct {
	category TL // TopPeerCategory
	count    int32
	peers    []TL // TopPeer
}

type TL_contacts_topPeersNotModified struct {
}

type TL_contacts_topPeers struct {
	categories []TL // TopPeerCategoryPeers
	chats      []TL // Chat
	users      []TL // User
}

type TL_messageEntityMentionName struct {
	offset  int32
	length  int32
	user_id int32
}

type TL_inputMessageEntityMentionName struct {
	offset  int32
	length  int32
	user_id TL // InputUser
}

type TL_inputMessagesFilterChatPhotos struct {
}

type TL_updateReadChannelOutbox struct {
	channel_id int32
	max_id     int32
}

type TL_updateDraftMessage struct {
	peer  TL // Peer
	draft TL // DraftMessage
}

type TL_draftMessageEmpty struct {
}

type TL_draftMessage struct {
	flags           int32
	no_webpage      bool  `flag_byte:"1"`
	reply_to_msg_id int32 `flag_byte:"0"`
	message         string
	entities        []TL `flag_byte:"3"` // MessageEntity
	date            int32
}

type TL_messageActionHistoryClear struct {
}

type TL_messages_featuredStickersNotModified struct {
}

type TL_messages_featuredStickers struct {
	hash   int32
	sets   []TL // StickerSetCovered
	unread []int64
}

type TL_updateReadFeaturedStickers struct {
}

type TL_messages_recentStickersNotModified struct {
}

type TL_messages_recentStickers struct {
	hash     int32
	stickers []TL // Document
}

type TL_updateRecentStickers struct {
}

type TL_messages_archivedStickers struct {
	count int32
	sets  []TL // StickerSetCovered
}

type TL_messages_stickerSetInstallResultSuccess struct {
}

type TL_messages_stickerSetInstallResultArchive struct {
	sets []TL // StickerSetCovered
}

type TL_stickerSetCovered struct {
	set   TL // StickerSet
	cover TL // Document
}

type TL_invokeAfterMsg struct {
	msg_id int64
	query  TL
}

type TL_invokeAfterMsgs struct {
	msg_ids []int64
	query   TL
}

type TL_auth_checkPhone struct {
	phone_number string
}

type TL_auth_sendCode struct {
	flags           int32
	allow_flashcall bool `flag_byte:"0"`
	phone_number    string
	current_number  TL `flag_byte:"0"` // Bool
	api_id          int32
	api_hash        string
}

type TL_auth_signUp struct {
	phone_number    string
	phone_code_hash string
	phone_code      string
	first_name      string
	last_name       string
}

type TL_auth_signIn struct {
	phone_number    string
	phone_code_hash string
	phone_code      string
}

type TL_auth_logOut struct {
}

type TL_auth_resetAuthorizations struct {
}

type TL_auth_sendInvites struct {
	phone_numbers []string
	message       string
}

type TL_auth_exportAuthorization struct {
	dc_id int32
}

type TL_auth_importAuthorization struct {
	id    int32
	bytes []byte
}

type TL_auth_bindTempAuthKey struct {
	perm_auth_key_id  int64
	nonce             int64
	expires_at        int32
	encrypted_message []byte
}

type TL_account_registerDevice struct {
	token_type int32
	token      string
}

type TL_account_unregisterDevice struct {
	token_type int32
	token      string
}

type TL_account_updateNotifySettings struct {
	peer     TL // InputNotifyPeer
	settings TL // InputPeerNotifySettings
}

type TL_account_getNotifySettings struct {
	peer TL // InputNotifyPeer
}

type TL_account_resetNotifySettings struct {
}

type TL_account_updateProfile struct {
	flags      int32
	first_name string `flag_byte:"0"`
	last_name  string `flag_byte:"1"`
	about      string `flag_byte:"2"`
}

type TL_account_updateStatus struct {
	offline TL // Bool
}

type TL_account_getWallPapers struct {
}

type TL_account_reportPeer struct {
	peer   TL // InputPeer
	reason TL // ReportReason
}

type TL_users_getUsers struct {
	id []TL // InputUser
}

type TL_users_getFullUser struct {
	id TL // InputUser
}

type TL_contacts_getStatuses struct {
}

type TL_contacts_getContacts struct {
	hash string
}

type TL_contacts_importContacts struct {
	contacts []TL // InputContact
	replace  TL   // Bool
}

type TL_contacts_deleteContact struct {
	id TL // InputUser
}

type TL_contacts_deleteContacts struct {
	id []TL // InputUser
}

type TL_contacts_block struct {
	id TL // InputUser
}

type TL_contacts_unblock struct {
	id TL // InputUser
}

type TL_contacts_getBlocked struct {
	offset int32
	limit  int32
}

type TL_contacts_exportCard struct {
}

type TL_contacts_importCard struct {
	export_card []int32
}

type TL_messages_getMessages struct {
	id []int32
}

type TL_messages_getDialogs struct {
	offset_date int32
	offset_id   int32
	offset_peer TL // InputPeer
	limit       int32
}

type TL_messages_getHistory struct {
	peer        TL // InputPeer
	offset_id   int32
	offset_date int32
	add_offset  int32
	limit       int32
	max_id      int32
	min_id      int32
}

type TL_messages_search struct {
	flags    int32
	peer     TL // InputPeer
	q        string
	filter   TL // MessagesFilter
	min_date int32
	max_date int32
	offset   int32
	max_id   int32
	limit    int32
}

type TL_messages_readHistory struct {
	peer   TL // InputPeer
	max_id int32
}

type TL_messages_deleteHistory struct {
	flags      int32
	just_clear bool `flag_byte:"0"`
	peer       TL   // InputPeer
	max_id     int32
}

type TL_messages_deleteMessages struct {
	id []int32
}

type TL_messages_receivedMessages struct {
	max_id int32
}

type TL_messages_setTyping struct {
	peer   TL // InputPeer
	action TL // SendMessageAction
}

type TL_messages_sendMessage struct {
	flags           int32
	no_webpage      bool  `flag_byte:"1"`
	silent          bool  `flag_byte:"5"`
	background      bool  `flag_byte:"6"`
	clear_draft     bool  `flag_byte:"7"`
	peer            TL    // InputPeer
	reply_to_msg_id int32 `flag_byte:"0"`
	message         string
	random_id       int64
	reply_markup    TL   `flag_byte:"2"` // ReplyMarkup
	entities        []TL `flag_byte:"3"` // MessageEntity
}

type TL_messages_sendMedia struct {
	flags           int32
	silent          bool  `flag_byte:"5"`
	background      bool  `flag_byte:"6"`
	clear_draft     bool  `flag_byte:"7"`
	peer            TL    // InputPeer
	reply_to_msg_id int32 `flag_byte:"0"`
	media           TL    // InputMedia
	random_id       int64
	reply_markup    TL `flag_byte:"2"` // ReplyMarkup
}

type TL_messages_forwardMessages struct {
	flags      int32
	silent     bool `flag_byte:"5"`
	background bool `flag_byte:"6"`
	from_peer  TL   // InputPeer
	id         []int32
	random_id  []int64
	to_peer    TL // InputPeer
}

type TL_messages_reportSpam struct {
	peer TL // InputPeer
}

type TL_messages_hideReportSpam struct {
	peer TL // InputPeer
}

type TL_messages_getPeerSettings struct {
	peer TL // InputPeer
}

type TL_messages_getChats struct {
	id []int32
}

type TL_messages_getFullChat struct {
	chat_id int32
}

type TL_messages_editChatTitle struct {
	chat_id int32
	title   string
}

type TL_messages_editChatPhoto struct {
	chat_id int32
	photo   TL // InputChatPhoto
}

type TL_messages_addChatUser struct {
	chat_id   int32
	user_id   TL // InputUser
	fwd_limit int32
}

type TL_messages_deleteChatUser struct {
	chat_id int32
	user_id TL // InputUser
}

type TL_messages_createChat struct {
	users []TL // InputUser
	title string
}

type TL_updates_getState struct {
}

type TL_updates_getDifference struct {
	pts  int32
	date int32
	qts  int32
}

type TL_photos_updateProfilePhoto struct {
	id   TL // InputPhoto
	crop TL // InputPhotoCrop
}

type TL_photos_uploadProfilePhoto struct {
	file      TL // InputFile
	caption   string
	geo_point TL // InputGeoPoint
	crop      TL // InputPhotoCrop
}

type TL_photos_deletePhotos struct {
	id []TL // InputPhoto
}

type TL_upload_saveFilePart struct {
	file_id   int64
	file_part int32
	bytes     []byte
}

type TL_upload_getFile struct {
	location TL // InputFileLocation
	offset   int32
	limit    int32
}

type TL_help_getConfig struct {
}

type TL_help_getNearestDc struct {
}

type TL_help_getAppUpdate struct {
}

type TL_help_saveAppLog struct {
	events []TL // InputAppEvent
}

type TL_help_getInviteText struct {
}

type TL_photos_getUserPhotos struct {
	user_id TL // InputUser
	offset  int32
	max_id  int64
	limit   int32
}

type TL_messages_forwardMessage struct {
	peer      TL // InputPeer
	id        int32
	random_id int64
}

type TL_messages_sendBroadcast struct {
	contacts  []TL // InputUser
	random_id []int64
	message   string
	media     TL // InputMedia
}

type TL_messages_getDhConfig struct {
	version       int32
	random_length int32
}

type TL_messages_requestEncryption struct {
	user_id   TL // InputUser
	random_id int32
	g_a       []byte
}

type TL_messages_acceptEncryption struct {
	peer            TL // InputEncryptedChat
	g_b             []byte
	key_fingerprint int64
}

type TL_messages_discardEncryption struct {
	chat_id int32
}

type TL_messages_setEncryptedTyping struct {
	peer   TL // InputEncryptedChat
	typing TL // Bool
}

type TL_messages_readEncryptedHistory struct {
	peer     TL // InputEncryptedChat
	max_date int32
}

type TL_messages_sendEncrypted struct {
	peer      TL // InputEncryptedChat
	random_id int64
	data      []byte
}

type TL_messages_sendEncryptedFile struct {
	peer      TL // InputEncryptedChat
	random_id int64
	data      []byte
	file      TL // InputEncryptedFile
}

type TL_messages_sendEncryptedService struct {
	peer      TL // InputEncryptedChat
	random_id int64
	data      []byte
}

type TL_messages_receivedQueue struct {
	max_qts int32
}

type TL_upload_saveBigFilePart struct {
	file_id          int64
	file_part        int32
	file_total_parts int32
	bytes            []byte
}

type TL_initConnection struct {
	api_id         int32
	device_model   string
	system_version string
	app_version    string
	lang_code      string
	query          TL
}

type TL_help_getSupport struct {
}

type TL_messages_readMessageContents struct {
	id []int32
}

type TL_account_checkUsername struct {
	username string
}

type TL_account_updateUsername struct {
	username string
}

type TL_contacts_search struct {
	q     string
	limit int32
}

type TL_account_getPrivacy struct {
	key TL // InputPrivacyKey
}

type TL_account_setPrivacy struct {
	key   TL   // InputPrivacyKey
	rules []TL // InputPrivacyRule
}

type TL_account_deleteAccount struct {
	reason string
}

type TL_account_getAccountTTL struct {
}

type TL_account_setAccountTTL struct {
	ttl TL // AccountDaysTTL
}

type TL_invokeWithLayer struct {
	layer int32
	query TL
}

type TL_contacts_resolveUsername struct {
	username string
}

type TL_account_sendChangePhoneCode struct {
	flags           int32
	allow_flashcall bool `flag_byte:"0"`
	phone_number    string
	current_number  TL `flag_byte:"0"` // Bool
}

type TL_account_changePhone struct {
	phone_number    string
	phone_code_hash string
	phone_code      string
}

type TL_messages_getStickers struct {
	emoticon string
	hash     string
}

type TL_messages_getAllStickers struct {
	hash int32
}

type TL_account_updateDeviceLocked struct {
	period int32
}

type TL_auth_importBotAuthorization struct {
	flags          int32
	api_id         int32
	api_hash       string
	bot_auth_token string
}

type TL_messages_getWebPagePreview struct {
	message string
}

type TL_account_getAuthorizations struct {
}

type TL_account_resetAuthorization struct {
	hash int64
}

type TL_account_getPassword struct {
}

type TL_account_getPasswordSettings struct {
	current_password_hash []byte
}

type TL_account_updatePasswordSettings struct {
	current_password_hash []byte
	new_settings          TL // account_PasswordInputSettings
}

type TL_auth_checkPassword struct {
	password_hash []byte
}

type TL_auth_requestPasswordRecovery struct {
}

type TL_auth_recoverPassword struct {
	code string
}

type TL_invokeWithoutUpdates struct {
	query TL
}

type TL_messages_exportChatInvite struct {
	chat_id int32
}

type TL_messages_checkChatInvite struct {
	hash string
}

type TL_messages_importChatInvite struct {
	hash string
}

type TL_messages_getStickerSet struct {
	stickerset TL // InputStickerSet
}

type TL_messages_installStickerSet struct {
	stickerset TL // InputStickerSet
	archived   TL // Bool
}

type TL_messages_uninstallStickerSet struct {
	stickerset TL // InputStickerSet
}

type TL_messages_startBot struct {
	bot         TL // InputUser
	peer        TL // InputPeer
	random_id   int64
	start_param string
}

type TL_help_getAppChangelog struct {
}

type TL_messages_getMessagesViews struct {
	peer      TL // InputPeer
	id        []int32
	increment TL // Bool
}

type TL_channels_readHistory struct {
	channel TL // InputChannel
	max_id  int32
}

type TL_channels_deleteMessages struct {
	channel TL // InputChannel
	id      []int32
}

type TL_channels_deleteUserHistory struct {
	channel TL // InputChannel
	user_id TL // InputUser
}

type TL_channels_reportSpam struct {
	channel TL // InputChannel
	user_id TL // InputUser
	id      []int32
}

type TL_channels_getMessages struct {
	channel TL // InputChannel
	id      []int32
}

type TL_channels_getParticipants struct {
	channel TL // InputChannel
	filter  TL // ChannelParticipantsFilter
	offset  int32
	limit   int32
}

type TL_channels_getParticipant struct {
	channel TL // InputChannel
	user_id TL // InputUser
}

type TL_channels_getChannels struct {
	id []TL // InputChannel
}

type TL_channels_getFullChannel struct {
	channel TL // InputChannel
}

type TL_channels_createChannel struct {
	flags     int32
	broadcast bool `flag_byte:"0"`
	megagroup bool `flag_byte:"1"`
	title     string
	about     string
}

type TL_channels_editAbout struct {
	channel TL // InputChannel
	about   string
}

type TL_channels_editAdmin struct {
	channel TL // InputChannel
	user_id TL // InputUser
	role    TL // ChannelParticipantRole
}

type TL_channels_editTitle struct {
	channel TL // InputChannel
	title   string
}

type TL_channels_editPhoto struct {
	channel TL // InputChannel
	photo   TL // InputChatPhoto
}

type TL_channels_checkUsername struct {
	channel  TL // InputChannel
	username string
}

type TL_channels_updateUsername struct {
	channel  TL // InputChannel
	username string
}

type TL_channels_joinChannel struct {
	channel TL // InputChannel
}

type TL_channels_leaveChannel struct {
	channel TL // InputChannel
}

type TL_channels_inviteToChannel struct {
	channel TL   // InputChannel
	users   []TL // InputUser
}

type TL_channels_kickFromChannel struct {
	channel TL // InputChannel
	user_id TL // InputUser
	kicked  TL // Bool
}

type TL_channels_exportInvite struct {
	channel TL // InputChannel
}

type TL_channels_deleteChannel struct {
	channel TL // InputChannel
}

type TL_updates_getChannelDifference struct {
	channel TL // InputChannel
	filter  TL // ChannelMessagesFilter
	pts     int32
	limit   int32
}

type TL_messages_toggleChatAdmins struct {
	chat_id int32
	enabled TL // Bool
}

type TL_messages_editChatAdmin struct {
	chat_id  int32
	user_id  TL // InputUser
	is_admin TL // Bool
}

type TL_messages_migrateChat struct {
	chat_id int32
}

type TL_messages_searchGlobal struct {
	q           string
	offset_date int32
	offset_peer TL // InputPeer
	offset_id   int32
	limit       int32
}

type TL_help_getTermsOfService struct {
}

type TL_messages_reorderStickerSets struct {
	order []int64
}

type TL_messages_getDocumentByHash struct {
	sha256    []byte
	size      int32
	mime_type string
}

type TL_messages_searchGifs struct {
	q      string
	offset int32
}

type TL_messages_getSavedGifs struct {
	hash int32
}

type TL_messages_saveGif struct {
	id     TL // InputDocument
	unsave TL // Bool
}

type TL_messages_getInlineBotResults struct {
	flags     int32
	bot       TL // InputUser
	peer      TL // InputPeer
	geo_point TL `flag_byte:"0"` // InputGeoPoint
	query     string
	offset    string
}

type TL_messages_setInlineBotResults struct {
	flags       int32
	gallery     bool `flag_byte:"0"`
	private     bool `flag_byte:"1"`
	query_id    int64
	results     []TL // InputBotInlineResult
	cache_time  int32
	next_offset string `flag_byte:"2"`
	switch_pm   TL     `flag_byte:"3"` // InlineBotSwitchPM
}

type TL_messages_sendInlineBotResult struct {
	flags           int32
	silent          bool  `flag_byte:"5"`
	background      bool  `flag_byte:"6"`
	clear_draft     bool  `flag_byte:"7"`
	peer            TL    // InputPeer
	reply_to_msg_id int32 `flag_byte:"0"`
	random_id       int64
	query_id        int64
	id              string
}

type TL_channels_toggleInvites struct {
	channel TL // InputChannel
	enabled TL // Bool
}

type TL_channels_exportMessageLink struct {
	channel TL // InputChannel
	id      int32
}

type TL_channels_toggleSignatures struct {
	channel TL // InputChannel
	enabled TL // Bool
}

type TL_channels_updatePinnedMessage struct {
	flags   int32
	silent  bool `flag_byte:"0"`
	channel TL   // InputChannel
	id      int32
}

type TL_auth_resendCode struct {
	phone_number    string
	phone_code_hash string
}

type TL_auth_cancelCode struct {
	phone_number    string
	phone_code_hash string
}

type TL_messages_getMessageEditData struct {
	peer TL // InputPeer
	id   int32
}

type TL_messages_editMessage struct {
	flags        int32
	no_webpage   bool `flag_byte:"1"`
	peer         TL   // InputPeer
	id           int32
	message      string `flag_byte:"11"`
	reply_markup TL     `flag_byte:"2"` // ReplyMarkup
	entities     []TL   `flag_byte:"3"` // MessageEntity
}

type TL_messages_editInlineBotMessage struct {
	flags        int32
	no_webpage   bool   `flag_byte:"1"`
	id           TL     // InputBotInlineMessageID
	message      string `flag_byte:"11"`
	reply_markup TL     `flag_byte:"2"` // ReplyMarkup
	entities     []TL   `flag_byte:"3"` // MessageEntity
}

type TL_messages_getBotCallbackAnswer struct {
	peer   TL // InputPeer
	msg_id int32
	data   []byte
}

type TL_messages_setBotCallbackAnswer struct {
	flags    int32
	alert    bool `flag_byte:"1"`
	query_id int64
	message  string `flag_byte:"0"`
	url      string `flag_byte:"2"`
}

type TL_contacts_getTopPeers struct {
	flags          int32
	correspondents bool `flag_byte:"0"`
	bots_pm        bool `flag_byte:"1"`
	bots_inline    bool `flag_byte:"2"`
	groups         bool `flag_byte:"10"`
	channels       bool `flag_byte:"15"`
	offset         int32
	limit          int32
	hash           int32
}

type TL_contacts_resetTopPeerRating struct {
	category TL // TopPeerCategory
	peer     TL // InputPeer
}

type TL_messages_getPeerDialogs struct {
	peers []TL // InputPeer
}

type TL_messages_saveDraft struct {
	flags           int32
	no_webpage      bool  `flag_byte:"1"`
	reply_to_msg_id int32 `flag_byte:"0"`
	peer            TL    // InputPeer
	message         string
	entities        []TL `flag_byte:"3"` // MessageEntity
}

type TL_messages_getAllDrafts struct {
}

type TL_messages_getFeaturedStickers struct {
	hash int32
}

type TL_messages_readFeaturedStickers struct {
}

type TL_messages_getRecentStickers struct {
	hash int32
}

type TL_messages_saveRecentSticker struct {
	id     TL // InputDocument
	unsave TL // Bool
}

type TL_messages_clearRecentStickers struct {
}

type TL_messages_getUnusedStickers struct {
	limit int32
}

type TL_messages_getArchivedStickers struct {
	offset_id int64
	limit     int32
}

type TL_account_sendConfirmPhoneCode struct {
	flags           int32
	allow_flashcall bool `flag_byte:"0"`
	hash            string
	current_number  TL `flag_byte:"0"` // Bool
}

type TL_account_confirmPhone struct {
	phone_code_hash string
	phone_code      string
}

func (e TL_boolFalse) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_boolFalse)
	return x.buf
}

func (e TL_boolTrue) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_boolTrue)
	return x.buf
}

func (e TL_true) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_true)
	return x.buf
}

func (e TL_error) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_error)
	x.Int(e.code)
	x.String(e.text)
	return x.buf
}

func (e TL_null) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_null)
	return x.buf
}

func (e TL_inputPeerEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPeerEmpty)
	return x.buf
}

func (e TL_inputPeerSelf) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPeerSelf)
	return x.buf
}

func (e TL_inputPeerChat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPeerChat)
	x.Int(e.chat_id)
	return x.buf
}

func (e TL_inputUserEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputUserEmpty)
	return x.buf
}

func (e TL_inputUserSelf) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputUserSelf)
	return x.buf
}

func (e TL_inputPhoneContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPhoneContact)
	x.Long(e.client_id)
	x.String(e.phone)
	x.String(e.first_name)
	x.String(e.last_name)
	return x.buf
}

func (e TL_inputFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputFile)
	x.Long(e.id)
	x.Int(e.parts)
	x.String(e.name)
	x.String(e.md5_checksum)
	return x.buf
}

func (e TL_inputMediaEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaEmpty)
	return x.buf
}

func (e TL_inputMediaUploadedPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaUploadedPhoto)
	x.Bytes(e.file.encode())
	x.String(e.caption)
	return x.buf
}

func (e TL_inputMediaPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaPhoto)
	x.Bytes(e.id.encode())
	x.String(e.caption)
	return x.buf
}

func (e TL_inputMediaGeoPoint) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaGeoPoint)
	x.Bytes(e.geo_point.encode())
	return x.buf
}

func (e TL_inputMediaContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaContact)
	x.String(e.phone_number)
	x.String(e.first_name)
	x.String(e.last_name)
	return x.buf
}

func (e TL_inputChatPhotoEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputChatPhotoEmpty)
	return x.buf
}

func (e TL_inputChatUploadedPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputChatUploadedPhoto)
	x.Bytes(e.file.encode())
	x.Bytes(e.crop.encode())
	return x.buf
}

func (e TL_inputChatPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputChatPhoto)
	x.Bytes(e.id.encode())
	x.Bytes(e.crop.encode())
	return x.buf
}

func (e TL_inputGeoPointEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputGeoPointEmpty)
	return x.buf
}

func (e TL_inputGeoPoint) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputGeoPoint)
	x.Double(e.lat)
	x.Double(e.long)
	return x.buf
}

func (e TL_inputPhotoEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPhotoEmpty)
	return x.buf
}

func (e TL_inputPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPhoto)
	x.Long(e.id)
	x.Long(e.access_hash)
	return x.buf
}

func (e TL_inputFileLocation) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputFileLocation)
	x.Long(e.volume_id)
	x.Int(e.local_id)
	x.Long(e.secret)
	return x.buf
}

func (e TL_inputPhotoCropAuto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPhotoCropAuto)
	return x.buf
}

func (e TL_inputPhotoCrop) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPhotoCrop)
	x.Double(e.crop_left)
	x.Double(e.crop_top)
	x.Double(e.crop_width)
	return x.buf
}

func (e TL_inputAppEvent) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputAppEvent)
	x.Double(e.time)
	x.String(e._type)
	x.Long(e.peer)
	x.String(e.data)
	return x.buf
}

func (e TL_peerUser) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_peerUser)
	x.Int(e.user_id)
	return x.buf
}

func (e TL_peerChat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_peerChat)
	x.Int(e.chat_id)
	return x.buf
}

func (e TL_storage_fileUnknown) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_storage_fileUnknown)
	return x.buf
}

func (e TL_storage_fileJpeg) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_storage_fileJpeg)
	return x.buf
}

func (e TL_storage_fileGif) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_storage_fileGif)
	return x.buf
}

func (e TL_storage_filePng) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_storage_filePng)
	return x.buf
}

func (e TL_storage_filePdf) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_storage_filePdf)
	return x.buf
}

func (e TL_storage_fileMp3) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_storage_fileMp3)
	return x.buf
}

func (e TL_storage_fileMov) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_storage_fileMov)
	return x.buf
}

func (e TL_storage_filePartial) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_storage_filePartial)
	return x.buf
}

func (e TL_storage_fileMp4) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_storage_fileMp4)
	return x.buf
}

func (e TL_storage_fileWebp) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_storage_fileWebp)
	return x.buf
}

func (e TL_fileLocationUnavailable) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_fileLocationUnavailable)
	x.Long(e.volume_id)
	x.Int(e.local_id)
	x.Long(e.secret)
	return x.buf
}

func (e TL_fileLocation) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_fileLocation)
	x.Int(e.dc_id)
	x.Long(e.volume_id)
	x.Int(e.local_id)
	x.Long(e.secret)
	return x.buf
}

func (e TL_userEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_userEmpty)
	x.Int(e.id)
	return x.buf
}

func (e TL_userProfilePhotoEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_userProfilePhotoEmpty)
	return x.buf
}

func (e TL_userProfilePhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_userProfilePhoto)
	x.Long(e.photo_id)
	x.Bytes(e.photo_small.encode())
	x.Bytes(e.photo_big.encode())
	return x.buf
}

func (e TL_userStatusEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_userStatusEmpty)
	return x.buf
}

func (e TL_userStatusOnline) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_userStatusOnline)
	x.Int(e.expires)
	return x.buf
}

func (e TL_userStatusOffline) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_userStatusOffline)
	x.Int(e.was_online)
	return x.buf
}

func (e TL_chatEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatEmpty)
	x.Int(e.id)
	return x.buf
}

func (e TL_chat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chat)
	var flags int32 = 0
	if e.creator {
		flags = flags | 1<<0
	}
	if e.kicked {
		flags = flags | 1<<1
	}
	if e.left {
		flags = flags | 1<<2
	}
	if e.admins_enabled {
		flags = flags | 1<<3
	}
	if e.admin {
		flags = flags | 1<<4
	}
	if e.deactivated {
		flags = flags | 1<<5
	}
	s := TL_chat{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.Int(e.id)
	x.String(e.title)
	x.Bytes(e.photo.encode())
	x.Int(e.participants_count)
	x.Int(e.date)
	x.Int(e.version)
	x.Bytes(e.migrated_to.encode())
	return x.buf
}

func (e TL_chatForbidden) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatForbidden)
	x.Int(e.id)
	x.String(e.title)
	return x.buf
}

func (e TL_chatFull) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatFull)
	x.Int(e.id)
	x.Bytes(e.participants.encode())
	x.Bytes(e.chat_photo.encode())
	x.Bytes(e.notify_settings.encode())
	x.Bytes(e.exported_invite.encode())
	x.Vector(e.bot_info)
	return x.buf
}

func (e TL_chatParticipant) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatParticipant)
	x.Int(e.user_id)
	x.Int(e.inviter_id)
	x.Int(e.date)
	return x.buf
}

func (e TL_chatParticipantsForbidden) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatParticipantsForbidden)
	var flags int32 = 0
	s := TL_chatParticipantsForbidden{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.Int(e.chat_id)
	x.Bytes(e.self_participant.encode())
	return x.buf
}

func (e TL_chatParticipants) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatParticipants)
	x.Int(e.chat_id)
	x.Vector(e.participants)
	x.Int(e.version)
	return x.buf
}

func (e TL_chatPhotoEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatPhotoEmpty)
	return x.buf
}

func (e TL_chatPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatPhoto)
	x.Bytes(e.photo_small.encode())
	x.Bytes(e.photo_big.encode())
	return x.buf
}

func (e TL_messageEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEmpty)
	x.Int(e.id)
	return x.buf
}

func (e TL_message) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_message)
	var flags int32 = 0
	if e.out {
		flags = flags | 1<<1
	}
	if e.mentioned {
		flags = flags | 1<<4
	}
	if e.media_unread {
		flags = flags | 1<<5
	}
	if e.silent {
		flags = flags | 1<<13
	}
	if e.post {
		flags = flags | 1<<14
	}
	s := TL_message{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.Int(e.id)
	x.Int(e.from_id)
	x.Bytes(e.to_id.encode())
	x.Bytes(e.fwd_from.encode())
	x.Int(e.via_bot_id)
	x.Int(e.reply_to_msg_id)
	x.Int(e.date)
	x.String(e.message)
	x.Bytes(e.media.encode())
	x.Bytes(e.reply_markup.encode())
	x.Vector(e.entities)
	x.Int(e.views)
	x.Int(e.edit_date)
	return x.buf
}

func (e TL_messageService) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageService)
	var flags int32 = 0
	if e.out {
		flags = flags | 1<<1
	}
	if e.mentioned {
		flags = flags | 1<<4
	}
	if e.media_unread {
		flags = flags | 1<<5
	}
	if e.silent {
		flags = flags | 1<<13
	}
	if e.post {
		flags = flags | 1<<14
	}
	s := TL_messageService{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.Int(e.id)
	x.Int(e.from_id)
	x.Bytes(e.to_id.encode())
	x.Int(e.reply_to_msg_id)
	x.Int(e.date)
	x.Bytes(e.action.encode())
	return x.buf
}

func (e TL_messageMediaEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageMediaEmpty)
	return x.buf
}

func (e TL_messageMediaPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageMediaPhoto)
	x.Bytes(e.photo.encode())
	x.String(e.caption)
	return x.buf
}

func (e TL_messageMediaGeo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageMediaGeo)
	x.Bytes(e.geo.encode())
	return x.buf
}

func (e TL_messageMediaContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageMediaContact)
	x.String(e.phone_number)
	x.String(e.first_name)
	x.String(e.last_name)
	x.Int(e.user_id)
	return x.buf
}

func (e TL_messageMediaUnsupported) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageMediaUnsupported)
	return x.buf
}

func (e TL_messageActionEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionEmpty)
	return x.buf
}

func (e TL_messageActionChatCreate) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionChatCreate)
	x.String(e.title)
	x.VectorInt(e.users)
	return x.buf
}

func (e TL_messageActionChatEditTitle) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionChatEditTitle)
	x.String(e.title)
	return x.buf
}

func (e TL_messageActionChatEditPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionChatEditPhoto)
	x.Bytes(e.photo.encode())
	return x.buf
}

func (e TL_messageActionChatDeletePhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionChatDeletePhoto)
	return x.buf
}

func (e TL_messageActionChatAddUser) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionChatAddUser)
	x.VectorInt(e.users)
	return x.buf
}

func (e TL_messageActionChatDeleteUser) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionChatDeleteUser)
	x.Int(e.user_id)
	return x.buf
}

func (e TL_dialog) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_dialog)
	var flags int32 = 0
	s := TL_dialog{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.Bytes(e.peer.encode())
	x.Int(e.top_message)
	x.Int(e.read_inbox_max_id)
	x.Int(e.read_outbox_max_id)
	x.Int(e.unread_count)
	x.Bytes(e.notify_settings.encode())
	x.Int(e.pts)
	x.Bytes(e.draft.encode())
	return x.buf
}

func (e TL_photoEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photoEmpty)
	x.Long(e.id)
	return x.buf
}

func (e TL_photo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photo)
	x.Long(e.id)
	x.Long(e.access_hash)
	x.Int(e.date)
	x.Vector(e.sizes)
	return x.buf
}

func (e TL_photoSizeEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photoSizeEmpty)
	x.String(e._type)
	return x.buf
}

func (e TL_photoSize) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photoSize)
	x.String(e._type)
	x.Bytes(e.location.encode())
	x.Int(e.w)
	x.Int(e.h)
	x.Int(e.size)
	return x.buf
}

func (e TL_photoCachedSize) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photoCachedSize)
	x.String(e._type)
	x.Bytes(e.location.encode())
	x.Int(e.w)
	x.Int(e.h)
	x.StringBytes(e.bytes)
	return x.buf
}

func (e TL_geoPointEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_geoPointEmpty)
	return x.buf
}

func (e TL_geoPoint) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_geoPoint)
	x.Double(e.long)
	x.Double(e.lat)
	return x.buf
}

func (e TL_auth_checkedPhone) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_checkedPhone)
	x.Bytes(e.phone_registered.encode())
	return x.buf
}

func (e TL_auth_sentCode) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_sentCode)
	var flags int32 = 0
	if e.phone_registered {
		flags = flags | 1<<0
	}
	s := TL_auth_sentCode{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.Bytes(e._type.encode())
	x.String(e.phone_code_hash)
	x.Bytes(e.next_type.encode())
	x.Int(e.timeout)
	return x.buf
}

func (e TL_auth_authorization) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_authorization)
	var flags int32 = 0
	s := TL_auth_authorization{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.Int(e.tmp_sessions)
	x.Bytes(e.user.encode())
	return x.buf
}

func (e TL_auth_exportedAuthorization) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_exportedAuthorization)
	x.Int(e.id)
	x.StringBytes(e.bytes)
	return x.buf
}

func (e TL_inputNotifyPeer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputNotifyPeer)
	x.Bytes(e.peer.encode())
	return x.buf
}

func (e TL_inputNotifyUsers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputNotifyUsers)
	return x.buf
}

func (e TL_inputNotifyChats) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputNotifyChats)
	return x.buf
}

func (e TL_inputNotifyAll) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputNotifyAll)
	return x.buf
}

func (e TL_inputPeerNotifyEventsEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPeerNotifyEventsEmpty)
	return x.buf
}

func (e TL_inputPeerNotifyEventsAll) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPeerNotifyEventsAll)
	return x.buf
}

func (e TL_inputPeerNotifySettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPeerNotifySettings)
	var flags int32 = 0
	if e.show_previews {
		flags = flags | 1<<0
	}
	if e.silent {
		flags = flags | 1<<1
	}
	s := TL_inputPeerNotifySettings{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.Int(e.mute_until)
	x.String(e.sound)
	return x.buf
}

func (e TL_peerNotifyEventsEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_peerNotifyEventsEmpty)
	return x.buf
}

func (e TL_peerNotifyEventsAll) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_peerNotifyEventsAll)
	return x.buf
}

func (e TL_peerNotifySettingsEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_peerNotifySettingsEmpty)
	return x.buf
}

func (e TL_peerNotifySettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_peerNotifySettings)
	var flags int32 = 0
	if e.show_previews {
		flags = flags | 1<<0
	}
	if e.silent {
		flags = flags | 1<<1
	}
	s := TL_peerNotifySettings{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.Int(e.mute_until)
	x.String(e.sound)
	return x.buf
}

func (e TL_peerSettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_peerSettings)
	var flags int32 = 0
	if e.report_spam {
		flags = flags | 1<<0
	}
	s := TL_peerSettings{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	return x.buf
}

func (e TL_wallPaper) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_wallPaper)
	x.Int(e.id)
	x.String(e.title)
	x.Vector(e.sizes)
	x.Int(e.color)
	return x.buf
}

func (e TL_inputReportReasonSpam) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputReportReasonSpam)
	return x.buf
}

func (e TL_inputReportReasonViolence) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputReportReasonViolence)
	return x.buf
}

func (e TL_inputReportReasonPornography) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputReportReasonPornography)
	return x.buf
}

func (e TL_inputReportReasonOther) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputReportReasonOther)
	x.String(e.text)
	return x.buf
}

func (e TL_userFull) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_userFull)
	var flags int32 = 0
	if e.blocked {
		flags = flags | 1<<0
	}
	s := TL_userFull{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.Bytes(e.user.encode())
	x.String(e.about)
	x.Bytes(e.link.encode())
	x.Bytes(e.profile_photo.encode())
	x.Bytes(e.notify_settings.encode())
	x.Bytes(e.bot_info.encode())
	return x.buf
}

func (e TL_contact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contact)
	x.Int(e.user_id)
	x.Bytes(e.mutual.encode())
	return x.buf
}

func (e TL_importedContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_importedContact)
	x.Int(e.user_id)
	x.Long(e.client_id)
	return x.buf
}

func (e TL_contactBlocked) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contactBlocked)
	x.Int(e.user_id)
	x.Int(e.date)
	return x.buf
}

func (e TL_contactStatus) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contactStatus)
	x.Int(e.user_id)
	x.Bytes(e.status.encode())
	return x.buf
}

func (e TL_contacts_link) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_link)
	x.Bytes(e.my_link.encode())
	x.Bytes(e.foreign_link.encode())
	x.Bytes(e.user.encode())
	return x.buf
}

func (e TL_contacts_contactsNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_contactsNotModified)
	return x.buf
}

func (e TL_contacts_contacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_contacts)
	x.Vector(e.contacts)
	x.Vector(e.users)
	return x.buf
}

func (e TL_contacts_importedContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_importedContacts)
	x.Vector(e.imported)
	x.VectorLong(e.retry_contacts)
	x.Vector(e.users)
	return x.buf
}

func (e TL_contacts_blocked) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_blocked)
	x.Vector(e.blocked)
	x.Vector(e.users)
	return x.buf
}

func (e TL_contacts_blockedSlice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_blockedSlice)
	x.Int(e.count)
	x.Vector(e.blocked)
	x.Vector(e.users)
	return x.buf
}

func (e TL_messages_dialogs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_dialogs)
	x.Vector(e.dialogs)
	x.Vector(e.messages)
	x.Vector(e.chats)
	x.Vector(e.users)
	return x.buf
}

func (e TL_messages_dialogsSlice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_dialogsSlice)
	x.Int(e.count)
	x.Vector(e.dialogs)
	x.Vector(e.messages)
	x.Vector(e.chats)
	x.Vector(e.users)
	return x.buf
}

func (e TL_messages_messages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_messages)
	x.Vector(e.messages)
	x.Vector(e.chats)
	x.Vector(e.users)
	return x.buf
}

func (e TL_messages_messagesSlice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_messagesSlice)
	x.Int(e.count)
	x.Vector(e.messages)
	x.Vector(e.chats)
	x.Vector(e.users)
	return x.buf
}

func (e TL_messages_chats) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_chats)
	x.Vector(e.chats)
	return x.buf
}

func (e TL_messages_chatFull) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_chatFull)
	x.Bytes(e.full_chat.encode())
	x.Vector(e.chats)
	x.Vector(e.users)
	return x.buf
}

func (e TL_messages_affectedHistory) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_affectedHistory)
	x.Int(e.pts)
	x.Int(e.pts_count)
	x.Int(e.offset)
	return x.buf
}

func (e TL_inputMessagesFilterEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterEmpty)
	return x.buf
}

func (e TL_inputMessagesFilterPhotos) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterPhotos)
	return x.buf
}

func (e TL_inputMessagesFilterVideo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterVideo)
	return x.buf
}

func (e TL_inputMessagesFilterPhotoVideo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterPhotoVideo)
	return x.buf
}

func (e TL_inputMessagesFilterPhotoVideoDocuments) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterPhotoVideoDocuments)
	return x.buf
}

func (e TL_inputMessagesFilterDocument) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterDocument)
	return x.buf
}

func (e TL_inputMessagesFilterUrl) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterUrl)
	return x.buf
}

func (e TL_inputMessagesFilterGif) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterGif)
	return x.buf
}

func (e TL_updateNewMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateNewMessage)
	x.Bytes(e.message.encode())
	x.Int(e.pts)
	x.Int(e.pts_count)
	return x.buf
}

func (e TL_updateMessageID) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateMessageID)
	x.Int(e.id)
	x.Long(e.random_id)
	return x.buf
}

func (e TL_updateDeleteMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateDeleteMessages)
	x.VectorInt(e.messages)
	x.Int(e.pts)
	x.Int(e.pts_count)
	return x.buf
}

func (e TL_updateUserTyping) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateUserTyping)
	x.Int(e.user_id)
	x.Bytes(e.action.encode())
	return x.buf
}

func (e TL_updateChatUserTyping) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateChatUserTyping)
	x.Int(e.chat_id)
	x.Int(e.user_id)
	x.Bytes(e.action.encode())
	return x.buf
}

func (e TL_updateChatParticipants) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateChatParticipants)
	x.Bytes(e.participants.encode())
	return x.buf
}

func (e TL_updateUserStatus) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateUserStatus)
	x.Int(e.user_id)
	x.Bytes(e.status.encode())
	return x.buf
}

func (e TL_updateUserName) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateUserName)
	x.Int(e.user_id)
	x.String(e.first_name)
	x.String(e.last_name)
	x.String(e.username)
	return x.buf
}

func (e TL_updateUserPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateUserPhoto)
	x.Int(e.user_id)
	x.Int(e.date)
	x.Bytes(e.photo.encode())
	x.Bytes(e.previous.encode())
	return x.buf
}

func (e TL_updateContactRegistered) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateContactRegistered)
	x.Int(e.user_id)
	x.Int(e.date)
	return x.buf
}

func (e TL_updateContactLink) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateContactLink)
	x.Int(e.user_id)
	x.Bytes(e.my_link.encode())
	x.Bytes(e.foreign_link.encode())
	return x.buf
}

func (e TL_updateNewAuthorization) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateNewAuthorization)
	x.Long(e.auth_key_id)
	x.Int(e.date)
	x.String(e.device)
	x.String(e.location)
	return x.buf
}

func (e TL_updates_state) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updates_state)
	x.Int(e.pts)
	x.Int(e.qts)
	x.Int(e.date)
	x.Int(e.seq)
	x.Int(e.unread_count)
	return x.buf
}

func (e TL_updates_differenceEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updates_differenceEmpty)
	x.Int(e.date)
	x.Int(e.seq)
	return x.buf
}

func (e TL_updates_difference) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updates_difference)
	x.Vector(e.new_messages)
	x.Vector(e.new_encrypted_messages)
	x.Vector(e.other_updates)
	x.Vector(e.chats)
	x.Vector(e.users)
	x.Bytes(e.state.encode())
	return x.buf
}

func (e TL_updates_differenceSlice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updates_differenceSlice)
	x.Vector(e.new_messages)
	x.Vector(e.new_encrypted_messages)
	x.Vector(e.other_updates)
	x.Vector(e.chats)
	x.Vector(e.users)
	x.Bytes(e.intermediate_state.encode())
	return x.buf
}

func (e TL_updatesTooLong) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updatesTooLong)
	return x.buf
}

func (e TL_updateShortMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateShortMessage)
	var flags int32 = 0
	if e.out {
		flags = flags | 1<<1
	}
	if e.mentioned {
		flags = flags | 1<<4
	}
	if e.media_unread {
		flags = flags | 1<<5
	}
	if e.silent {
		flags = flags | 1<<13
	}
	s := TL_updateShortMessage{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.Int(e.id)
	x.Int(e.user_id)
	x.String(e.message)
	x.Int(e.pts)
	x.Int(e.pts_count)
	x.Int(e.date)
	x.Bytes(e.fwd_from.encode())
	x.Int(e.via_bot_id)
	x.Int(e.reply_to_msg_id)
	x.Vector(e.entities)
	return x.buf
}

func (e TL_updateShortChatMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateShortChatMessage)
	var flags int32 = 0
	if e.out {
		flags = flags | 1<<1
	}
	if e.mentioned {
		flags = flags | 1<<4
	}
	if e.media_unread {
		flags = flags | 1<<5
	}
	if e.silent {
		flags = flags | 1<<13
	}
	s := TL_updateShortChatMessage{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.Int(e.id)
	x.Int(e.from_id)
	x.Int(e.chat_id)
	x.String(e.message)
	x.Int(e.pts)
	x.Int(e.pts_count)
	x.Int(e.date)
	x.Bytes(e.fwd_from.encode())
	x.Int(e.via_bot_id)
	x.Int(e.reply_to_msg_id)
	x.Vector(e.entities)
	return x.buf
}

func (e TL_updateShort) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateShort)
	x.Bytes(e.update.encode())
	x.Int(e.date)
	return x.buf
}

func (e TL_updatesCombined) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updatesCombined)
	x.Vector(e.updates)
	x.Vector(e.users)
	x.Vector(e.chats)
	x.Int(e.date)
	x.Int(e.seq_start)
	x.Int(e.seq)
	return x.buf
}

func (e TL_updates) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updates)
	x.Vector(e.updates)
	x.Vector(e.users)
	x.Vector(e.chats)
	x.Int(e.date)
	x.Int(e.seq)
	return x.buf
}

func (e TL_photos_photos) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photos_photos)
	x.Vector(e.photos)
	x.Vector(e.users)
	return x.buf
}

func (e TL_photos_photosSlice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photos_photosSlice)
	x.Int(e.count)
	x.Vector(e.photos)
	x.Vector(e.users)
	return x.buf
}

func (e TL_photos_photo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photos_photo)
	x.Bytes(e.photo.encode())
	x.Vector(e.users)
	return x.buf
}

func (e TL_upload_file) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_upload_file)
	x.Bytes(e._type.encode())
	x.Int(e.mtime)
	x.StringBytes(e.bytes)
	return x.buf
}

func (e TL_dcOption) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_dcOption)
	var flags int32 = 0
	if e.ipv6 {
		flags = flags | 1<<0
	}
	if e.media_only {
		flags = flags | 1<<1
	}
	if e.tcpo_only {
		flags = flags | 1<<2
	}
	s := TL_dcOption{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.Int(e.id)
	x.String(e.ip_address)
	x.Int(e.port)
	return x.buf
}

func (e TL_config) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_config)
	var flags int32 = 0
	s := TL_config{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.Int(e.date)
	x.Int(e.expires)
	x.Bytes(e.test_mode.encode())
	x.Int(e.this_dc)
	x.Vector(e.dc_options)
	x.Int(e.chat_size_max)
	x.Int(e.megagroup_size_max)
	x.Int(e.forwarded_count_max)
	x.Int(e.online_update_period_ms)
	x.Int(e.offline_blur_timeout_ms)
	x.Int(e.offline_idle_timeout_ms)
	x.Int(e.online_cloud_timeout_ms)
	x.Int(e.notify_cloud_delay_ms)
	x.Int(e.notify_default_delay_ms)
	x.Int(e.chat_big_size)
	x.Int(e.push_chat_period_ms)
	x.Int(e.push_chat_limit)
	x.Int(e.saved_gifs_limit)
	x.Int(e.edit_time_limit)
	x.Int(e.rating_e_decay)
	x.Int(e.stickers_recent_limit)
	x.Int(e.tmp_sessions)
	x.Vector(e.disabled_features)
	return x.buf
}

func (e TL_nearestDc) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_nearestDc)
	x.String(e.country)
	x.Int(e.this_dc)
	x.Int(e.nearest_dc)
	return x.buf
}

func (e TL_help_appUpdate) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_appUpdate)
	x.Int(e.id)
	x.Bytes(e.critical.encode())
	x.String(e.url)
	x.String(e.text)
	return x.buf
}

func (e TL_help_noAppUpdate) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_noAppUpdate)
	return x.buf
}

func (e TL_help_inviteText) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_inviteText)
	x.String(e.message)
	return x.buf
}

func (e TL_wallPaperSolid) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_wallPaperSolid)
	x.Int(e.id)
	x.String(e.title)
	x.Int(e.bg_color)
	x.Int(e.color)
	return x.buf
}

func (e TL_updateNewEncryptedMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateNewEncryptedMessage)
	x.Bytes(e.message.encode())
	x.Int(e.qts)
	return x.buf
}

func (e TL_updateEncryptedChatTyping) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateEncryptedChatTyping)
	x.Int(e.chat_id)
	return x.buf
}

func (e TL_updateEncryption) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateEncryption)
	x.Bytes(e.chat.encode())
	x.Int(e.date)
	return x.buf
}

func (e TL_updateEncryptedMessagesRead) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateEncryptedMessagesRead)
	x.Int(e.chat_id)
	x.Int(e.max_date)
	x.Int(e.date)
	return x.buf
}

func (e TL_encryptedChatEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_encryptedChatEmpty)
	x.Int(e.id)
	return x.buf
}

func (e TL_encryptedChatWaiting) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_encryptedChatWaiting)
	x.Int(e.id)
	x.Long(e.access_hash)
	x.Int(e.date)
	x.Int(e.admin_id)
	x.Int(e.participant_id)
	return x.buf
}

func (e TL_encryptedChatRequested) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_encryptedChatRequested)
	x.Int(e.id)
	x.Long(e.access_hash)
	x.Int(e.date)
	x.Int(e.admin_id)
	x.Int(e.participant_id)
	x.StringBytes(e.g_a)
	return x.buf
}

func (e TL_encryptedChat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_encryptedChat)
	x.Int(e.id)
	x.Long(e.access_hash)
	x.Int(e.date)
	x.Int(e.admin_id)
	x.Int(e.participant_id)
	x.StringBytes(e.g_a_or_b)
	x.Long(e.key_fingerprint)
	return x.buf
}

func (e TL_encryptedChatDiscarded) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_encryptedChatDiscarded)
	x.Int(e.id)
	return x.buf
}

func (e TL_inputEncryptedChat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputEncryptedChat)
	x.Int(e.chat_id)
	x.Long(e.access_hash)
	return x.buf
}

func (e TL_encryptedFileEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_encryptedFileEmpty)
	return x.buf
}

func (e TL_encryptedFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_encryptedFile)
	x.Long(e.id)
	x.Long(e.access_hash)
	x.Int(e.size)
	x.Int(e.dc_id)
	x.Int(e.key_fingerprint)
	return x.buf
}

func (e TL_inputEncryptedFileEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputEncryptedFileEmpty)
	return x.buf
}

func (e TL_inputEncryptedFileUploaded) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputEncryptedFileUploaded)
	x.Long(e.id)
	x.Int(e.parts)
	x.String(e.md5_checksum)
	x.Int(e.key_fingerprint)
	return x.buf
}

func (e TL_inputEncryptedFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputEncryptedFile)
	x.Long(e.id)
	x.Long(e.access_hash)
	return x.buf
}

func (e TL_inputEncryptedFileLocation) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputEncryptedFileLocation)
	x.Long(e.id)
	x.Long(e.access_hash)
	return x.buf
}

func (e TL_encryptedMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_encryptedMessage)
	x.Long(e.random_id)
	x.Int(e.chat_id)
	x.Int(e.date)
	x.StringBytes(e.bytes)
	x.Bytes(e.file.encode())
	return x.buf
}

func (e TL_encryptedMessageService) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_encryptedMessageService)
	x.Long(e.random_id)
	x.Int(e.chat_id)
	x.Int(e.date)
	x.StringBytes(e.bytes)
	return x.buf
}

func (e TL_messages_dhConfigNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_dhConfigNotModified)
	x.StringBytes(e.random)
	return x.buf
}

func (e TL_messages_dhConfig) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_dhConfig)
	x.Int(e.g)
	x.StringBytes(e.p)
	x.Int(e.version)
	x.StringBytes(e.random)
	return x.buf
}

func (e TL_messages_sentEncryptedMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_sentEncryptedMessage)
	x.Int(e.date)
	return x.buf
}

func (e TL_messages_sentEncryptedFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_sentEncryptedFile)
	x.Int(e.date)
	x.Bytes(e.file.encode())
	return x.buf
}

func (e TL_inputFileBig) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputFileBig)
	x.Long(e.id)
	x.Int(e.parts)
	x.String(e.name)
	return x.buf
}

func (e TL_inputEncryptedFileBigUploaded) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputEncryptedFileBigUploaded)
	x.Long(e.id)
	x.Int(e.parts)
	x.Int(e.key_fingerprint)
	return x.buf
}

func (e TL_updateChatParticipantAdd) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateChatParticipantAdd)
	x.Int(e.chat_id)
	x.Int(e.user_id)
	x.Int(e.inviter_id)
	x.Int(e.date)
	x.Int(e.version)
	return x.buf
}

func (e TL_updateChatParticipantDelete) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateChatParticipantDelete)
	x.Int(e.chat_id)
	x.Int(e.user_id)
	x.Int(e.version)
	return x.buf
}

func (e TL_updateDcOptions) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateDcOptions)
	x.Vector(e.dc_options)
	return x.buf
}

func (e TL_inputMediaUploadedDocument) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaUploadedDocument)
	x.Bytes(e.file.encode())
	x.String(e.mime_type)
	x.Vector(e.attributes)
	x.String(e.caption)
	return x.buf
}

func (e TL_inputMediaUploadedThumbDocument) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaUploadedThumbDocument)
	x.Bytes(e.file.encode())
	x.Bytes(e.thumb.encode())
	x.String(e.mime_type)
	x.Vector(e.attributes)
	x.String(e.caption)
	return x.buf
}

func (e TL_inputMediaDocument) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaDocument)
	x.Bytes(e.id.encode())
	x.String(e.caption)
	return x.buf
}

func (e TL_messageMediaDocument) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageMediaDocument)
	x.Bytes(e.document.encode())
	x.String(e.caption)
	return x.buf
}

func (e TL_inputDocumentEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputDocumentEmpty)
	return x.buf
}

func (e TL_inputDocument) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputDocument)
	x.Long(e.id)
	x.Long(e.access_hash)
	return x.buf
}

func (e TL_inputDocumentFileLocation) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputDocumentFileLocation)
	x.Long(e.id)
	x.Long(e.access_hash)
	x.Int(e.version)
	return x.buf
}

func (e TL_documentEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_documentEmpty)
	x.Long(e.id)
	return x.buf
}

func (e TL_document) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_document)
	x.Long(e.id)
	x.Long(e.access_hash)
	x.Int(e.date)
	x.String(e.mime_type)
	x.Int(e.size)
	x.Bytes(e.thumb.encode())
	x.Int(e.dc_id)
	x.Int(e.version)
	x.Vector(e.attributes)
	return x.buf
}

func (e TL_help_support) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_support)
	x.String(e.phone_number)
	x.Bytes(e.user.encode())
	return x.buf
}

func (e TL_notifyPeer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_notifyPeer)
	x.Bytes(e.peer.encode())
	return x.buf
}

func (e TL_notifyUsers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_notifyUsers)
	return x.buf
}

func (e TL_notifyChats) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_notifyChats)
	return x.buf
}

func (e TL_notifyAll) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_notifyAll)
	return x.buf
}

func (e TL_updateUserBlocked) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateUserBlocked)
	x.Int(e.user_id)
	x.Bytes(e.blocked.encode())
	return x.buf
}

func (e TL_updateNotifySettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateNotifySettings)
	x.Bytes(e.peer.encode())
	x.Bytes(e.notify_settings.encode())
	return x.buf
}

func (e TL_sendMessageTypingAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageTypingAction)
	return x.buf
}

func (e TL_sendMessageCancelAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageCancelAction)
	return x.buf
}

func (e TL_sendMessageRecordVideoAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageRecordVideoAction)
	return x.buf
}

func (e TL_sendMessageUploadVideoAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageUploadVideoAction)
	x.Int(e.progress)
	return x.buf
}

func (e TL_sendMessageRecordAudioAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageRecordAudioAction)
	return x.buf
}

func (e TL_sendMessageUploadAudioAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageUploadAudioAction)
	x.Int(e.progress)
	return x.buf
}

func (e TL_sendMessageUploadPhotoAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageUploadPhotoAction)
	x.Int(e.progress)
	return x.buf
}

func (e TL_sendMessageUploadDocumentAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageUploadDocumentAction)
	x.Int(e.progress)
	return x.buf
}

func (e TL_sendMessageGeoLocationAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageGeoLocationAction)
	return x.buf
}

func (e TL_sendMessageChooseContactAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageChooseContactAction)
	return x.buf
}

func (e TL_contacts_found) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_found)
	x.Vector(e.results)
	x.Vector(e.chats)
	x.Vector(e.users)
	return x.buf
}

func (e TL_updateServiceNotification) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateServiceNotification)
	x.String(e._type)
	x.String(e.message)
	x.Bytes(e.media.encode())
	x.Bytes(e.popup.encode())
	return x.buf
}

func (e TL_userStatusRecently) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_userStatusRecently)
	return x.buf
}

func (e TL_userStatusLastWeek) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_userStatusLastWeek)
	return x.buf
}

func (e TL_userStatusLastMonth) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_userStatusLastMonth)
	return x.buf
}

func (e TL_updatePrivacy) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updatePrivacy)
	x.Bytes(e.key.encode())
	x.Vector(e.rules)
	return x.buf
}

func (e TL_inputPrivacyKeyStatusTimestamp) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPrivacyKeyStatusTimestamp)
	return x.buf
}

func (e TL_privacyKeyStatusTimestamp) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_privacyKeyStatusTimestamp)
	return x.buf
}

func (e TL_inputPrivacyValueAllowContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPrivacyValueAllowContacts)
	return x.buf
}

func (e TL_inputPrivacyValueAllowAll) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPrivacyValueAllowAll)
	return x.buf
}

func (e TL_inputPrivacyValueAllowUsers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPrivacyValueAllowUsers)
	x.Vector(e.users)
	return x.buf
}

func (e TL_inputPrivacyValueDisallowContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPrivacyValueDisallowContacts)
	return x.buf
}

func (e TL_inputPrivacyValueDisallowAll) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPrivacyValueDisallowAll)
	return x.buf
}

func (e TL_inputPrivacyValueDisallowUsers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPrivacyValueDisallowUsers)
	x.Vector(e.users)
	return x.buf
}

func (e TL_privacyValueAllowContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_privacyValueAllowContacts)
	return x.buf
}

func (e TL_privacyValueAllowAll) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_privacyValueAllowAll)
	return x.buf
}

func (e TL_privacyValueAllowUsers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_privacyValueAllowUsers)
	x.VectorInt(e.users)
	return x.buf
}

func (e TL_privacyValueDisallowContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_privacyValueDisallowContacts)
	return x.buf
}

func (e TL_privacyValueDisallowAll) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_privacyValueDisallowAll)
	return x.buf
}

func (e TL_privacyValueDisallowUsers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_privacyValueDisallowUsers)
	x.VectorInt(e.users)
	return x.buf
}

func (e TL_account_privacyRules) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_privacyRules)
	x.Vector(e.rules)
	x.Vector(e.users)
	return x.buf
}

func (e TL_accountDaysTTL) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_accountDaysTTL)
	x.Int(e.days)
	return x.buf
}

func (e TL_updateUserPhone) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateUserPhone)
	x.Int(e.user_id)
	x.String(e.phone)
	return x.buf
}

func (e TL_documentAttributeImageSize) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_documentAttributeImageSize)
	x.Int(e.w)
	x.Int(e.h)
	return x.buf
}

func (e TL_documentAttributeAnimated) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_documentAttributeAnimated)
	return x.buf
}

func (e TL_documentAttributeSticker) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_documentAttributeSticker)
	x.String(e.alt)
	x.Bytes(e.stickerset.encode())
	return x.buf
}

func (e TL_documentAttributeVideo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_documentAttributeVideo)
	x.Int(e.duration)
	x.Int(e.w)
	x.Int(e.h)
	return x.buf
}

func (e TL_documentAttributeAudio) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_documentAttributeAudio)
	var flags int32 = 0
	if e.voice {
		flags = flags | 1<<10
	}
	s := TL_documentAttributeAudio{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.Int(e.duration)
	x.String(e.title)
	x.String(e.performer)
	x.StringBytes(e.waveform)
	return x.buf
}

func (e TL_documentAttributeFilename) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_documentAttributeFilename)
	x.String(e.file_name)
	return x.buf
}

func (e TL_messages_stickersNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_stickersNotModified)
	return x.buf
}

func (e TL_messages_stickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_stickers)
	x.String(e.hash)
	x.Vector(e.stickers)
	return x.buf
}

func (e TL_stickerPack) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_stickerPack)
	x.String(e.emoticon)
	x.VectorLong(e.documents)
	return x.buf
}

func (e TL_messages_allStickersNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_allStickersNotModified)
	return x.buf
}

func (e TL_messages_allStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_allStickers)
	x.Int(e.hash)
	x.Vector(e.sets)
	return x.buf
}

func (e TL_disabledFeature) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_disabledFeature)
	x.String(e.feature)
	x.String(e.description)
	return x.buf
}

func (e TL_updateReadHistoryInbox) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateReadHistoryInbox)
	x.Bytes(e.peer.encode())
	x.Int(e.max_id)
	x.Int(e.pts)
	x.Int(e.pts_count)
	return x.buf
}

func (e TL_updateReadHistoryOutbox) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateReadHistoryOutbox)
	x.Bytes(e.peer.encode())
	x.Int(e.max_id)
	x.Int(e.pts)
	x.Int(e.pts_count)
	return x.buf
}

func (e TL_messages_affectedMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_affectedMessages)
	x.Int(e.pts)
	x.Int(e.pts_count)
	return x.buf
}

func (e TL_contactLinkUnknown) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contactLinkUnknown)
	return x.buf
}

func (e TL_contactLinkNone) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contactLinkNone)
	return x.buf
}

func (e TL_contactLinkHasPhone) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contactLinkHasPhone)
	return x.buf
}

func (e TL_contactLinkContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contactLinkContact)
	return x.buf
}

func (e TL_updateWebPage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateWebPage)
	x.Bytes(e.webpage.encode())
	x.Int(e.pts)
	x.Int(e.pts_count)
	return x.buf
}

func (e TL_webPageEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_webPageEmpty)
	x.Long(e.id)
	return x.buf
}

func (e TL_webPagePending) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_webPagePending)
	x.Long(e.id)
	x.Int(e.date)
	return x.buf
}

func (e TL_webPage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_webPage)
	var flags int32 = 0
	s := TL_webPage{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.Long(e.id)
	x.String(e.url)
	x.String(e.display_url)
	x.String(e._type)
	x.String(e.site_name)
	x.String(e.title)
	x.String(e.description)
	x.Bytes(e.photo.encode())
	x.String(e.embed_url)
	x.String(e.embed_type)
	x.Int(e.embed_width)
	x.Int(e.embed_height)
	x.Int(e.duration)
	x.String(e.author)
	x.Bytes(e.document.encode())
	return x.buf
}

func (e TL_messageMediaWebPage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageMediaWebPage)
	x.Bytes(e.webpage.encode())
	return x.buf
}

func (e TL_authorization) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_authorization)
	x.Long(e.hash)
	x.Int(e.flags)
	x.String(e.device_model)
	x.String(e.platform)
	x.String(e.system_version)
	x.Int(e.api_id)
	x.String(e.app_name)
	x.String(e.app_version)
	x.Int(e.date_created)
	x.Int(e.date_active)
	x.String(e.ip)
	x.String(e.country)
	x.String(e.region)
	return x.buf
}

func (e TL_account_authorizations) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_authorizations)
	x.Vector(e.authorizations)
	return x.buf
}

func (e TL_account_noPassword) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_noPassword)
	x.StringBytes(e.new_salt)
	x.String(e.email_unconfirmed_pattern)
	return x.buf
}

func (e TL_account_password) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_password)
	x.StringBytes(e.current_salt)
	x.StringBytes(e.new_salt)
	x.String(e.hint)
	x.Bytes(e.has_recovery.encode())
	x.String(e.email_unconfirmed_pattern)
	return x.buf
}

func (e TL_account_passwordSettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_passwordSettings)
	x.String(e.email)
	return x.buf
}

func (e TL_account_passwordInputSettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_passwordInputSettings)
	var flags int32 = 0
	s := TL_account_passwordInputSettings{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.StringBytes(e.new_salt)
	x.StringBytes(e.new_password_hash)
	x.String(e.hint)
	x.String(e.email)
	return x.buf
}

func (e TL_auth_passwordRecovery) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_passwordRecovery)
	x.String(e.email_pattern)
	return x.buf
}

func (e TL_inputMediaVenue) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaVenue)
	x.Bytes(e.geo_point.encode())
	x.String(e.title)
	x.String(e.address)
	x.String(e.provider)
	x.String(e.venue_id)
	return x.buf
}

func (e TL_messageMediaVenue) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageMediaVenue)
	x.Bytes(e.geo.encode())
	x.String(e.title)
	x.String(e.address)
	x.String(e.provider)
	x.String(e.venue_id)
	return x.buf
}

func (e TL_receivedNotifyMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_receivedNotifyMessage)
	x.Int(e.id)
	x.Int(e.flags)
	return x.buf
}

func (e TL_chatInviteEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatInviteEmpty)
	return x.buf
}

func (e TL_chatInviteExported) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatInviteExported)
	x.String(e.link)
	return x.buf
}

func (e TL_chatInviteAlready) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatInviteAlready)
	x.Bytes(e.chat.encode())
	return x.buf
}

func (e TL_chatInvite) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatInvite)
	var flags int32 = 0
	if e.channel {
		flags = flags | 1<<0
	}
	if e.broadcast {
		flags = flags | 1<<1
	}
	if e.public {
		flags = flags | 1<<2
	}
	if e.megagroup {
		flags = flags | 1<<3
	}
	s := TL_chatInvite{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.String(e.title)
	x.Bytes(e.photo.encode())
	x.Int(e.participants_count)
	x.Vector(e.participants)
	return x.buf
}

func (e TL_messageActionChatJoinedByLink) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionChatJoinedByLink)
	x.Int(e.inviter_id)
	return x.buf
}

func (e TL_updateReadMessagesContents) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateReadMessagesContents)
	x.VectorInt(e.messages)
	x.Int(e.pts)
	x.Int(e.pts_count)
	return x.buf
}

func (e TL_inputStickerSetEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputStickerSetEmpty)
	return x.buf
}

func (e TL_inputStickerSetID) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputStickerSetID)
	x.Long(e.id)
	x.Long(e.access_hash)
	return x.buf
}

func (e TL_inputStickerSetShortName) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputStickerSetShortName)
	x.String(e.short_name)
	return x.buf
}

func (e TL_stickerSet) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_stickerSet)
	var flags int32 = 0
	if e.installed {
		flags = flags | 1<<0
	}
	if e.archived {
		flags = flags | 1<<1
	}
	if e.official {
		flags = flags | 1<<2
	}
	s := TL_stickerSet{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.Long(e.id)
	x.Long(e.access_hash)
	x.String(e.title)
	x.String(e.short_name)
	x.Int(e.count)
	x.Int(e.hash)
	return x.buf
}

func (e TL_messages_stickerSet) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_stickerSet)
	x.Bytes(e.set.encode())
	x.Vector(e.packs)
	x.Vector(e.documents)
	return x.buf
}

func (e TL_user) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_user)
	var flags int32 = 0
	if e.self {
		flags = flags | 1<<10
	}
	if e.contact {
		flags = flags | 1<<11
	}
	if e.mutual_contact {
		flags = flags | 1<<12
	}
	if e.deleted {
		flags = flags | 1<<13
	}
	if e.bot {
		flags = flags | 1<<14
	}
	if e.bot_chat_history {
		flags = flags | 1<<15
	}
	if e.bot_nochats {
		flags = flags | 1<<16
	}
	if e.verified {
		flags = flags | 1<<17
	}
	if e.restricted {
		flags = flags | 1<<18
	}
	if e.min {
		flags = flags | 1<<20
	}
	if e.bot_inline_geo {
		flags = flags | 1<<21
	}
	s := TL_user{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.Int(e.id)
	x.Long(e.access_hash)
	x.String(e.first_name)
	x.String(e.last_name)
	x.String(e.username)
	x.String(e.phone)
	x.Bytes(e.photo.encode())
	x.Bytes(e.status.encode())
	x.Int(e.bot_info_version)
	x.String(e.restriction_reason)
	x.String(e.bot_inline_placeholder)
	return x.buf
}

func (e TL_botCommand) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_botCommand)
	x.String(e.command)
	x.String(e.description)
	return x.buf
}

func (e TL_botInfo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_botInfo)
	x.Int(e.user_id)
	x.String(e.description)
	x.Vector(e.commands)
	return x.buf
}

func (e TL_keyboardButton) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_keyboardButton)
	x.String(e.text)
	return x.buf
}

func (e TL_keyboardButtonRow) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_keyboardButtonRow)
	x.Vector(e.buttons)
	return x.buf
}

func (e TL_replyKeyboardHide) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_replyKeyboardHide)
	var flags int32 = 0
	if e.selective {
		flags = flags | 1<<2
	}
	s := TL_replyKeyboardHide{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	return x.buf
}

func (e TL_replyKeyboardForceReply) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_replyKeyboardForceReply)
	var flags int32 = 0
	if e.single_use {
		flags = flags | 1<<1
	}
	if e.selective {
		flags = flags | 1<<2
	}
	s := TL_replyKeyboardForceReply{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	return x.buf
}

func (e TL_replyKeyboardMarkup) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_replyKeyboardMarkup)
	var flags int32 = 0
	if e.resize {
		flags = flags | 1<<0
	}
	if e.single_use {
		flags = flags | 1<<1
	}
	if e.selective {
		flags = flags | 1<<2
	}
	s := TL_replyKeyboardMarkup{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.Vector(e.rows)
	return x.buf
}

func (e TL_inputPeerUser) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPeerUser)
	x.Int(e.user_id)
	x.Long(e.access_hash)
	return x.buf
}

func (e TL_inputUser) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputUser)
	x.Int(e.user_id)
	x.Long(e.access_hash)
	return x.buf
}

func (e TL_help_appChangelogEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_appChangelogEmpty)
	return x.buf
}

func (e TL_help_appChangelog) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_appChangelog)
	x.String(e.text)
	return x.buf
}

func (e TL_messageEntityUnknown) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEntityUnknown)
	x.Int(e.offset)
	x.Int(e.length)
	return x.buf
}

func (e TL_messageEntityMention) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEntityMention)
	x.Int(e.offset)
	x.Int(e.length)
	return x.buf
}

func (e TL_messageEntityHashtag) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEntityHashtag)
	x.Int(e.offset)
	x.Int(e.length)
	return x.buf
}

func (e TL_messageEntityBotCommand) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEntityBotCommand)
	x.Int(e.offset)
	x.Int(e.length)
	return x.buf
}

func (e TL_messageEntityUrl) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEntityUrl)
	x.Int(e.offset)
	x.Int(e.length)
	return x.buf
}

func (e TL_messageEntityEmail) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEntityEmail)
	x.Int(e.offset)
	x.Int(e.length)
	return x.buf
}

func (e TL_messageEntityBold) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEntityBold)
	x.Int(e.offset)
	x.Int(e.length)
	return x.buf
}

func (e TL_messageEntityItalic) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEntityItalic)
	x.Int(e.offset)
	x.Int(e.length)
	return x.buf
}

func (e TL_messageEntityCode) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEntityCode)
	x.Int(e.offset)
	x.Int(e.length)
	return x.buf
}

func (e TL_messageEntityPre) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEntityPre)
	x.Int(e.offset)
	x.Int(e.length)
	x.String(e.language)
	return x.buf
}

func (e TL_messageEntityTextUrl) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEntityTextUrl)
	x.Int(e.offset)
	x.Int(e.length)
	x.String(e.url)
	return x.buf
}

func (e TL_updateShortSentMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateShortSentMessage)
	var flags int32 = 0
	if e.out {
		flags = flags | 1<<1
	}
	s := TL_updateShortSentMessage{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.Int(e.id)
	x.Int(e.pts)
	x.Int(e.pts_count)
	x.Int(e.date)
	x.Bytes(e.media.encode())
	x.Vector(e.entities)
	return x.buf
}

func (e TL_inputChannelEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputChannelEmpty)
	return x.buf
}

func (e TL_inputChannel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputChannel)
	x.Int(e.channel_id)
	x.Long(e.access_hash)
	return x.buf
}

func (e TL_peerChannel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_peerChannel)
	x.Int(e.channel_id)
	return x.buf
}

func (e TL_inputPeerChannel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPeerChannel)
	x.Int(e.channel_id)
	x.Long(e.access_hash)
	return x.buf
}

func (e TL_channel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channel)
	var flags int32 = 0
	if e.creator {
		flags = flags | 1<<0
	}
	if e.kicked {
		flags = flags | 1<<1
	}
	if e.left {
		flags = flags | 1<<2
	}
	if e.editor {
		flags = flags | 1<<3
	}
	if e.moderator {
		flags = flags | 1<<4
	}
	if e.broadcast {
		flags = flags | 1<<5
	}
	if e.verified {
		flags = flags | 1<<7
	}
	if e.megagroup {
		flags = flags | 1<<8
	}
	if e.restricted {
		flags = flags | 1<<9
	}
	if e.democracy {
		flags = flags | 1<<10
	}
	if e.signatures {
		flags = flags | 1<<11
	}
	if e.min {
		flags = flags | 1<<12
	}
	s := TL_channel{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.Int(e.id)
	x.Long(e.access_hash)
	x.String(e.title)
	x.String(e.username)
	x.Bytes(e.photo.encode())
	x.Int(e.date)
	x.Int(e.version)
	x.String(e.restriction_reason)
	return x.buf
}

func (e TL_channelForbidden) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelForbidden)
	var flags int32 = 0
	if e.broadcast {
		flags = flags | 1<<5
	}
	if e.megagroup {
		flags = flags | 1<<8
	}
	s := TL_channelForbidden{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.Int(e.id)
	x.Long(e.access_hash)
	x.String(e.title)
	return x.buf
}

func (e TL_contacts_resolvedPeer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_resolvedPeer)
	x.Bytes(e.peer.encode())
	x.Vector(e.chats)
	x.Vector(e.users)
	return x.buf
}

func (e TL_channelFull) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelFull)
	var flags int32 = 0
	if e.can_view_participants {
		flags = flags | 1<<3
	}
	if e.can_set_username {
		flags = flags | 1<<6
	}
	s := TL_channelFull{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.Int(e.id)
	x.String(e.about)
	x.Int(e.participants_count)
	x.Int(e.admins_count)
	x.Int(e.kicked_count)
	x.Int(e.read_inbox_max_id)
	x.Int(e.read_outbox_max_id)
	x.Int(e.unread_count)
	x.Bytes(e.chat_photo.encode())
	x.Bytes(e.notify_settings.encode())
	x.Bytes(e.exported_invite.encode())
	x.Vector(e.bot_info)
	x.Int(e.migrated_from_chat_id)
	x.Int(e.migrated_from_max_id)
	x.Int(e.pinned_msg_id)
	return x.buf
}

func (e TL_messageRange) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageRange)
	x.Int(e.min_id)
	x.Int(e.max_id)
	return x.buf
}

func (e TL_messages_channelMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_channelMessages)
	var flags int32 = 0
	s := TL_messages_channelMessages{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.Int(e.pts)
	x.Int(e.count)
	x.Vector(e.messages)
	x.Vector(e.chats)
	x.Vector(e.users)
	return x.buf
}

func (e TL_messageActionChannelCreate) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionChannelCreate)
	x.String(e.title)
	return x.buf
}

func (e TL_updateChannelTooLong) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateChannelTooLong)
	var flags int32 = 0
	s := TL_updateChannelTooLong{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.Int(e.channel_id)
	x.Int(e.pts)
	return x.buf
}

func (e TL_updateChannel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateChannel)
	x.Int(e.channel_id)
	return x.buf
}

func (e TL_updateNewChannelMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateNewChannelMessage)
	x.Bytes(e.message.encode())
	x.Int(e.pts)
	x.Int(e.pts_count)
	return x.buf
}

func (e TL_updateReadChannelInbox) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateReadChannelInbox)
	x.Int(e.channel_id)
	x.Int(e.max_id)
	return x.buf
}

func (e TL_updateDeleteChannelMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateDeleteChannelMessages)
	x.Int(e.channel_id)
	x.VectorInt(e.messages)
	x.Int(e.pts)
	x.Int(e.pts_count)
	return x.buf
}

func (e TL_updateChannelMessageViews) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateChannelMessageViews)
	x.Int(e.channel_id)
	x.Int(e.id)
	x.Int(e.views)
	return x.buf
}

func (e TL_updates_channelDifferenceEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updates_channelDifferenceEmpty)
	var flags int32 = 0
	if e.final {
		flags = flags | 1<<0
	}
	s := TL_updates_channelDifferenceEmpty{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.Int(e.pts)
	x.Int(e.timeout)
	return x.buf
}

func (e TL_updates_channelDifferenceTooLong) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updates_channelDifferenceTooLong)
	var flags int32 = 0
	if e.final {
		flags = flags | 1<<0
	}
	s := TL_updates_channelDifferenceTooLong{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.Int(e.pts)
	x.Int(e.timeout)
	x.Int(e.top_message)
	x.Int(e.read_inbox_max_id)
	x.Int(e.read_outbox_max_id)
	x.Int(e.unread_count)
	x.Vector(e.messages)
	x.Vector(e.chats)
	x.Vector(e.users)
	return x.buf
}

func (e TL_updates_channelDifference) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updates_channelDifference)
	var flags int32 = 0
	if e.final {
		flags = flags | 1<<0
	}
	s := TL_updates_channelDifference{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.Int(e.pts)
	x.Int(e.timeout)
	x.Vector(e.new_messages)
	x.Vector(e.other_updates)
	x.Vector(e.chats)
	x.Vector(e.users)
	return x.buf
}

func (e TL_channelMessagesFilterEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelMessagesFilterEmpty)
	return x.buf
}

func (e TL_channelMessagesFilter) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelMessagesFilter)
	var flags int32 = 0
	if e.exclude_new_messages {
		flags = flags | 1<<1
	}
	s := TL_channelMessagesFilter{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.Vector(e.ranges)
	return x.buf
}

func (e TL_channelParticipant) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelParticipant)
	x.Int(e.user_id)
	x.Int(e.date)
	return x.buf
}

func (e TL_channelParticipantSelf) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelParticipantSelf)
	x.Int(e.user_id)
	x.Int(e.inviter_id)
	x.Int(e.date)
	return x.buf
}

func (e TL_channelParticipantModerator) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelParticipantModerator)
	x.Int(e.user_id)
	x.Int(e.inviter_id)
	x.Int(e.date)
	return x.buf
}

func (e TL_channelParticipantEditor) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelParticipantEditor)
	x.Int(e.user_id)
	x.Int(e.inviter_id)
	x.Int(e.date)
	return x.buf
}

func (e TL_channelParticipantKicked) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelParticipantKicked)
	x.Int(e.user_id)
	x.Int(e.kicked_by)
	x.Int(e.date)
	return x.buf
}

func (e TL_channelParticipantCreator) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelParticipantCreator)
	x.Int(e.user_id)
	return x.buf
}

func (e TL_channelParticipantsRecent) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelParticipantsRecent)
	return x.buf
}

func (e TL_channelParticipantsAdmins) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelParticipantsAdmins)
	return x.buf
}

func (e TL_channelParticipantsKicked) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelParticipantsKicked)
	return x.buf
}

func (e TL_channelRoleEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelRoleEmpty)
	return x.buf
}

func (e TL_channelRoleModerator) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelRoleModerator)
	return x.buf
}

func (e TL_channelRoleEditor) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelRoleEditor)
	return x.buf
}

func (e TL_channels_channelParticipants) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_channelParticipants)
	x.Int(e.count)
	x.Vector(e.participants)
	x.Vector(e.users)
	return x.buf
}

func (e TL_channels_channelParticipant) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_channelParticipant)
	x.Bytes(e.participant.encode())
	x.Vector(e.users)
	return x.buf
}

func (e TL_chatParticipantCreator) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatParticipantCreator)
	x.Int(e.user_id)
	return x.buf
}

func (e TL_chatParticipantAdmin) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatParticipantAdmin)
	x.Int(e.user_id)
	x.Int(e.inviter_id)
	x.Int(e.date)
	return x.buf
}

func (e TL_updateChatAdmins) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateChatAdmins)
	x.Int(e.chat_id)
	x.Bytes(e.enabled.encode())
	x.Int(e.version)
	return x.buf
}

func (e TL_updateChatParticipantAdmin) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateChatParticipantAdmin)
	x.Int(e.chat_id)
	x.Int(e.user_id)
	x.Bytes(e.is_admin.encode())
	x.Int(e.version)
	return x.buf
}

func (e TL_messageActionChatMigrateTo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionChatMigrateTo)
	x.Int(e.channel_id)
	return x.buf
}

func (e TL_messageActionChannelMigrateFrom) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionChannelMigrateFrom)
	x.String(e.title)
	x.Int(e.chat_id)
	return x.buf
}

func (e TL_channelParticipantsBots) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channelParticipantsBots)
	return x.buf
}

func (e TL_help_termsOfService) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_termsOfService)
	x.String(e.text)
	return x.buf
}

func (e TL_updateNewStickerSet) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateNewStickerSet)
	x.Bytes(e.stickerset.encode())
	return x.buf
}

func (e TL_updateStickerSetsOrder) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateStickerSetsOrder)
	x.VectorLong(e.order)
	return x.buf
}

func (e TL_updateStickerSets) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateStickerSets)
	return x.buf
}

func (e TL_foundGif) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_foundGif)
	x.String(e.url)
	x.String(e.thumb_url)
	x.String(e.content_url)
	x.String(e.content_type)
	x.Int(e.w)
	x.Int(e.h)
	return x.buf
}

func (e TL_foundGifCached) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_foundGifCached)
	x.String(e.url)
	x.Bytes(e.photo.encode())
	x.Bytes(e.document.encode())
	return x.buf
}

func (e TL_inputMediaGifExternal) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaGifExternal)
	x.String(e.url)
	x.String(e.q)
	return x.buf
}

func (e TL_messages_foundGifs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_foundGifs)
	x.Int(e.next_offset)
	x.Vector(e.results)
	return x.buf
}

func (e TL_messages_savedGifsNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_savedGifsNotModified)
	return x.buf
}

func (e TL_messages_savedGifs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_savedGifs)
	x.Int(e.hash)
	x.Vector(e.gifs)
	return x.buf
}

func (e TL_updateSavedGifs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateSavedGifs)
	return x.buf
}

func (e TL_inputBotInlineMessageMediaAuto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputBotInlineMessageMediaAuto)
	var flags int32 = 0
	s := TL_inputBotInlineMessageMediaAuto{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.String(e.caption)
	x.Bytes(e.reply_markup.encode())
	return x.buf
}

func (e TL_inputBotInlineMessageText) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputBotInlineMessageText)
	var flags int32 = 0
	if e.no_webpage {
		flags = flags | 1<<0
	}
	s := TL_inputBotInlineMessageText{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.String(e.message)
	x.Vector(e.entities)
	x.Bytes(e.reply_markup.encode())
	return x.buf
}

func (e TL_inputBotInlineResult) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputBotInlineResult)
	var flags int32 = 0
	s := TL_inputBotInlineResult{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.String(e.id)
	x.String(e._type)
	x.String(e.title)
	x.String(e.description)
	x.String(e.url)
	x.String(e.thumb_url)
	x.String(e.content_url)
	x.String(e.content_type)
	x.Int(e.w)
	x.Int(e.h)
	x.Int(e.duration)
	x.Bytes(e.send_message.encode())
	return x.buf
}

func (e TL_botInlineMessageMediaAuto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_botInlineMessageMediaAuto)
	var flags int32 = 0
	s := TL_botInlineMessageMediaAuto{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.String(e.caption)
	x.Bytes(e.reply_markup.encode())
	return x.buf
}

func (e TL_botInlineMessageText) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_botInlineMessageText)
	var flags int32 = 0
	if e.no_webpage {
		flags = flags | 1<<0
	}
	s := TL_botInlineMessageText{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.String(e.message)
	x.Vector(e.entities)
	x.Bytes(e.reply_markup.encode())
	return x.buf
}

func (e TL_botInlineResult) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_botInlineResult)
	var flags int32 = 0
	s := TL_botInlineResult{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.String(e.id)
	x.String(e._type)
	x.String(e.title)
	x.String(e.description)
	x.String(e.url)
	x.String(e.thumb_url)
	x.String(e.content_url)
	x.String(e.content_type)
	x.Int(e.w)
	x.Int(e.h)
	x.Int(e.duration)
	x.Bytes(e.send_message.encode())
	return x.buf
}

func (e TL_messages_botResults) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_botResults)
	var flags int32 = 0
	if e.gallery {
		flags = flags | 1<<0
	}
	s := TL_messages_botResults{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.Long(e.query_id)
	x.String(e.next_offset)
	x.Bytes(e.switch_pm.encode())
	x.Vector(e.results)
	return x.buf
}

func (e TL_updateBotInlineQuery) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateBotInlineQuery)
	var flags int32 = 0
	s := TL_updateBotInlineQuery{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.Long(e.query_id)
	x.Int(e.user_id)
	x.String(e.query)
	x.Bytes(e.geo.encode())
	x.String(e.offset)
	return x.buf
}

func (e TL_updateBotInlineSend) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateBotInlineSend)
	var flags int32 = 0
	s := TL_updateBotInlineSend{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.Int(e.user_id)
	x.String(e.query)
	x.Bytes(e.geo.encode())
	x.String(e.id)
	x.Bytes(e.msg_id.encode())
	return x.buf
}

func (e TL_inputMessagesFilterVoice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterVoice)
	return x.buf
}

func (e TL_inputMessagesFilterMusic) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterMusic)
	return x.buf
}

func (e TL_inputPrivacyKeyChatInvite) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPrivacyKeyChatInvite)
	return x.buf
}

func (e TL_privacyKeyChatInvite) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_privacyKeyChatInvite)
	return x.buf
}

func (e TL_exportedMessageLink) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_exportedMessageLink)
	x.String(e.link)
	return x.buf
}

func (e TL_messageFwdHeader) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageFwdHeader)
	var flags int32 = 0
	s := TL_messageFwdHeader{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.Int(e.from_id)
	x.Int(e.date)
	x.Int(e.channel_id)
	x.Int(e.channel_post)
	return x.buf
}

func (e TL_updateEditChannelMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateEditChannelMessage)
	x.Bytes(e.message.encode())
	x.Int(e.pts)
	x.Int(e.pts_count)
	return x.buf
}

func (e TL_updateChannelPinnedMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateChannelPinnedMessage)
	x.Int(e.channel_id)
	x.Int(e.id)
	return x.buf
}

func (e TL_messageActionPinMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionPinMessage)
	return x.buf
}

func (e TL_auth_codeTypeSms) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_codeTypeSms)
	return x.buf
}

func (e TL_auth_codeTypeCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_codeTypeCall)
	return x.buf
}

func (e TL_auth_codeTypeFlashCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_codeTypeFlashCall)
	return x.buf
}

func (e TL_auth_sentCodeTypeApp) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_sentCodeTypeApp)
	x.Int(e.length)
	return x.buf
}

func (e TL_auth_sentCodeTypeSms) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_sentCodeTypeSms)
	x.Int(e.length)
	return x.buf
}

func (e TL_auth_sentCodeTypeCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_sentCodeTypeCall)
	x.Int(e.length)
	return x.buf
}

func (e TL_auth_sentCodeTypeFlashCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_sentCodeTypeFlashCall)
	x.String(e.pattern)
	return x.buf
}

func (e TL_keyboardButtonUrl) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_keyboardButtonUrl)
	x.String(e.text)
	x.String(e.url)
	return x.buf
}

func (e TL_keyboardButtonCallback) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_keyboardButtonCallback)
	x.String(e.text)
	x.StringBytes(e.data)
	return x.buf
}

func (e TL_keyboardButtonRequestPhone) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_keyboardButtonRequestPhone)
	x.String(e.text)
	return x.buf
}

func (e TL_keyboardButtonRequestGeoLocation) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_keyboardButtonRequestGeoLocation)
	x.String(e.text)
	return x.buf
}

func (e TL_keyboardButtonSwitchInline) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_keyboardButtonSwitchInline)
	x.String(e.text)
	x.String(e.query)
	return x.buf
}

func (e TL_replyInlineMarkup) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_replyInlineMarkup)
	x.Vector(e.rows)
	return x.buf
}

func (e TL_messages_botCallbackAnswer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_botCallbackAnswer)
	var flags int32 = 0
	if e.alert {
		flags = flags | 1<<1
	}
	s := TL_messages_botCallbackAnswer{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.String(e.message)
	x.String(e.url)
	return x.buf
}

func (e TL_updateBotCallbackQuery) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateBotCallbackQuery)
	x.Long(e.query_id)
	x.Int(e.user_id)
	x.Bytes(e.peer.encode())
	x.Int(e.msg_id)
	x.StringBytes(e.data)
	return x.buf
}

func (e TL_messages_messageEditData) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_messageEditData)
	var flags int32 = 0
	if e.caption {
		flags = flags | 1<<0
	}
	s := TL_messages_messageEditData{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	return x.buf
}

func (e TL_updateEditMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateEditMessage)
	x.Bytes(e.message.encode())
	x.Int(e.pts)
	x.Int(e.pts_count)
	return x.buf
}

func (e TL_inputBotInlineMessageMediaGeo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputBotInlineMessageMediaGeo)
	var flags int32 = 0
	s := TL_inputBotInlineMessageMediaGeo{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.Bytes(e.geo_point.encode())
	x.Bytes(e.reply_markup.encode())
	return x.buf
}

func (e TL_inputBotInlineMessageMediaVenue) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputBotInlineMessageMediaVenue)
	var flags int32 = 0
	s := TL_inputBotInlineMessageMediaVenue{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.Bytes(e.geo_point.encode())
	x.String(e.title)
	x.String(e.address)
	x.String(e.provider)
	x.String(e.venue_id)
	x.Bytes(e.reply_markup.encode())
	return x.buf
}

func (e TL_inputBotInlineMessageMediaContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputBotInlineMessageMediaContact)
	var flags int32 = 0
	s := TL_inputBotInlineMessageMediaContact{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.String(e.phone_number)
	x.String(e.first_name)
	x.String(e.last_name)
	x.Bytes(e.reply_markup.encode())
	return x.buf
}

func (e TL_botInlineMessageMediaGeo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_botInlineMessageMediaGeo)
	var flags int32 = 0
	s := TL_botInlineMessageMediaGeo{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.Bytes(e.geo.encode())
	x.Bytes(e.reply_markup.encode())
	return x.buf
}

func (e TL_botInlineMessageMediaVenue) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_botInlineMessageMediaVenue)
	var flags int32 = 0
	s := TL_botInlineMessageMediaVenue{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.Bytes(e.geo.encode())
	x.String(e.title)
	x.String(e.address)
	x.String(e.provider)
	x.String(e.venue_id)
	x.Bytes(e.reply_markup.encode())
	return x.buf
}

func (e TL_botInlineMessageMediaContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_botInlineMessageMediaContact)
	var flags int32 = 0
	s := TL_botInlineMessageMediaContact{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.String(e.phone_number)
	x.String(e.first_name)
	x.String(e.last_name)
	x.Bytes(e.reply_markup.encode())
	return x.buf
}

func (e TL_inputBotInlineResultPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputBotInlineResultPhoto)
	x.String(e.id)
	x.String(e._type)
	x.Bytes(e.photo.encode())
	x.Bytes(e.send_message.encode())
	return x.buf
}

func (e TL_inputBotInlineResultDocument) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputBotInlineResultDocument)
	var flags int32 = 0
	s := TL_inputBotInlineResultDocument{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.String(e.id)
	x.String(e._type)
	x.String(e.title)
	x.String(e.description)
	x.Bytes(e.document.encode())
	x.Bytes(e.send_message.encode())
	return x.buf
}

func (e TL_botInlineMediaResult) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_botInlineMediaResult)
	var flags int32 = 0
	s := TL_botInlineMediaResult{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.String(e.id)
	x.String(e._type)
	x.Bytes(e.photo.encode())
	x.Bytes(e.document.encode())
	x.String(e.title)
	x.String(e.description)
	x.Bytes(e.send_message.encode())
	return x.buf
}

func (e TL_inputBotInlineMessageID) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputBotInlineMessageID)
	x.Int(e.dc_id)
	x.Long(e.id)
	x.Long(e.access_hash)
	return x.buf
}

func (e TL_updateInlineBotCallbackQuery) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateInlineBotCallbackQuery)
	x.Long(e.query_id)
	x.Int(e.user_id)
	x.Bytes(e.msg_id.encode())
	x.StringBytes(e.data)
	return x.buf
}

func (e TL_inlineBotSwitchPM) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inlineBotSwitchPM)
	x.String(e.text)
	x.String(e.start_param)
	return x.buf
}

func (e TL_messages_peerDialogs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_peerDialogs)
	x.Vector(e.dialogs)
	x.Vector(e.messages)
	x.Vector(e.chats)
	x.Vector(e.users)
	x.Bytes(e.state.encode())
	return x.buf
}

func (e TL_topPeer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_topPeer)
	x.Bytes(e.peer.encode())
	x.Double(e.rating)
	return x.buf
}

func (e TL_topPeerCategoryBotsPM) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_topPeerCategoryBotsPM)
	return x.buf
}

func (e TL_topPeerCategoryBotsInline) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_topPeerCategoryBotsInline)
	return x.buf
}

func (e TL_topPeerCategoryCorrespondents) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_topPeerCategoryCorrespondents)
	return x.buf
}

func (e TL_topPeerCategoryGroups) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_topPeerCategoryGroups)
	return x.buf
}

func (e TL_topPeerCategoryChannels) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_topPeerCategoryChannels)
	return x.buf
}

func (e TL_topPeerCategoryPeers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_topPeerCategoryPeers)
	x.Bytes(e.category.encode())
	x.Int(e.count)
	x.Vector(e.peers)
	return x.buf
}

func (e TL_contacts_topPeersNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_topPeersNotModified)
	return x.buf
}

func (e TL_contacts_topPeers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_topPeers)
	x.Vector(e.categories)
	x.Vector(e.chats)
	x.Vector(e.users)
	return x.buf
}

func (e TL_messageEntityMentionName) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEntityMentionName)
	x.Int(e.offset)
	x.Int(e.length)
	x.Int(e.user_id)
	return x.buf
}

func (e TL_inputMessageEntityMentionName) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessageEntityMentionName)
	x.Int(e.offset)
	x.Int(e.length)
	x.Bytes(e.user_id.encode())
	return x.buf
}

func (e TL_inputMessagesFilterChatPhotos) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterChatPhotos)
	return x.buf
}

func (e TL_updateReadChannelOutbox) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateReadChannelOutbox)
	x.Int(e.channel_id)
	x.Int(e.max_id)
	return x.buf
}

func (e TL_updateDraftMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateDraftMessage)
	x.Bytes(e.peer.encode())
	x.Bytes(e.draft.encode())
	return x.buf
}

func (e TL_draftMessageEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_draftMessageEmpty)
	return x.buf
}

func (e TL_draftMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_draftMessage)
	var flags int32 = 0
	if e.no_webpage {
		flags = flags | 1<<1
	}
	s := TL_draftMessage{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.Int(e.reply_to_msg_id)
	x.String(e.message)
	x.Vector(e.entities)
	x.Int(e.date)
	return x.buf
}

func (e TL_messageActionHistoryClear) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionHistoryClear)
	return x.buf
}

func (e TL_messages_featuredStickersNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_featuredStickersNotModified)
	return x.buf
}

func (e TL_messages_featuredStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_featuredStickers)
	x.Int(e.hash)
	x.Vector(e.sets)
	x.VectorLong(e.unread)
	return x.buf
}

func (e TL_updateReadFeaturedStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateReadFeaturedStickers)
	return x.buf
}

func (e TL_messages_recentStickersNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_recentStickersNotModified)
	return x.buf
}

func (e TL_messages_recentStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_recentStickers)
	x.Int(e.hash)
	x.Vector(e.stickers)
	return x.buf
}

func (e TL_updateRecentStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateRecentStickers)
	return x.buf
}

func (e TL_messages_archivedStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_archivedStickers)
	x.Int(e.count)
	x.Vector(e.sets)
	return x.buf
}

func (e TL_messages_stickerSetInstallResultSuccess) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_stickerSetInstallResultSuccess)
	return x.buf
}

func (e TL_messages_stickerSetInstallResultArchive) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_stickerSetInstallResultArchive)
	x.Vector(e.sets)
	return x.buf
}

func (e TL_stickerSetCovered) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_stickerSetCovered)
	x.Bytes(e.set.encode())
	x.Bytes(e.cover.encode())
	return x.buf
}

func (e TL_invokeAfterMsg) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_invokeAfterMsg)
	x.Long(e.msg_id)
	x.Bytes(e.query.encode())
	return x.buf
}

func (e TL_invokeAfterMsgs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_invokeAfterMsgs)
	x.VectorLong(e.msg_ids)
	x.Bytes(e.query.encode())
	return x.buf
}

func (e TL_auth_checkPhone) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_checkPhone)
	x.String(e.phone_number)
	return x.buf
}

func (e TL_auth_sendCode) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_sendCode)
	var flags int32 = 0
	if e.allow_flashcall {
		flags = flags | 1<<0
	}
	s := TL_auth_sendCode{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.String(e.phone_number)
	x.Bytes(e.current_number.encode())
	x.Int(e.api_id)
	x.String(e.api_hash)
	return x.buf
}

func (e TL_auth_signUp) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_signUp)
	x.String(e.phone_number)
	x.String(e.phone_code_hash)
	x.String(e.phone_code)
	x.String(e.first_name)
	x.String(e.last_name)
	return x.buf
}

func (e TL_auth_signIn) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_signIn)
	x.String(e.phone_number)
	x.String(e.phone_code_hash)
	x.String(e.phone_code)
	return x.buf
}

func (e TL_auth_logOut) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_logOut)
	return x.buf
}

func (e TL_auth_resetAuthorizations) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_resetAuthorizations)
	return x.buf
}

func (e TL_auth_sendInvites) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_sendInvites)
	x.VectorString(e.phone_numbers)
	x.String(e.message)
	return x.buf
}

func (e TL_auth_exportAuthorization) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_exportAuthorization)
	x.Int(e.dc_id)
	return x.buf
}

func (e TL_auth_importAuthorization) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_importAuthorization)
	x.Int(e.id)
	x.StringBytes(e.bytes)
	return x.buf
}

func (e TL_auth_bindTempAuthKey) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_bindTempAuthKey)
	x.Long(e.perm_auth_key_id)
	x.Long(e.nonce)
	x.Int(e.expires_at)
	x.StringBytes(e.encrypted_message)
	return x.buf
}

func (e TL_account_registerDevice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_registerDevice)
	x.Int(e.token_type)
	x.String(e.token)
	return x.buf
}

func (e TL_account_unregisterDevice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_unregisterDevice)
	x.Int(e.token_type)
	x.String(e.token)
	return x.buf
}

func (e TL_account_updateNotifySettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_updateNotifySettings)
	x.Bytes(e.peer.encode())
	x.Bytes(e.settings.encode())
	return x.buf
}

func (e TL_account_getNotifySettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_getNotifySettings)
	x.Bytes(e.peer.encode())
	return x.buf
}

func (e TL_account_resetNotifySettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_resetNotifySettings)
	return x.buf
}

func (e TL_account_updateProfile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_updateProfile)
	var flags int32 = 0
	s := TL_account_updateProfile{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.String(e.first_name)
	x.String(e.last_name)
	x.String(e.about)
	return x.buf
}

func (e TL_account_updateStatus) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_updateStatus)
	x.Bytes(e.offline.encode())
	return x.buf
}

func (e TL_account_getWallPapers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_getWallPapers)
	return x.buf
}

func (e TL_account_reportPeer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_reportPeer)
	x.Bytes(e.peer.encode())
	x.Bytes(e.reason.encode())
	return x.buf
}

func (e TL_users_getUsers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_users_getUsers)
	x.Vector(e.id)
	return x.buf
}

func (e TL_users_getFullUser) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_users_getFullUser)
	x.Bytes(e.id.encode())
	return x.buf
}

func (e TL_contacts_getStatuses) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_getStatuses)
	return x.buf
}

func (e TL_contacts_getContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_getContacts)
	x.String(e.hash)
	return x.buf
}

func (e TL_contacts_importContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_importContacts)
	x.Vector(e.contacts)
	x.Bytes(e.replace.encode())
	return x.buf
}

func (e TL_contacts_deleteContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_deleteContact)
	x.Bytes(e.id.encode())
	return x.buf
}

func (e TL_contacts_deleteContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_deleteContacts)
	x.Vector(e.id)
	return x.buf
}

func (e TL_contacts_block) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_block)
	x.Bytes(e.id.encode())
	return x.buf
}

func (e TL_contacts_unblock) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_unblock)
	x.Bytes(e.id.encode())
	return x.buf
}

func (e TL_contacts_getBlocked) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_getBlocked)
	x.Int(e.offset)
	x.Int(e.limit)
	return x.buf
}

func (e TL_contacts_exportCard) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_exportCard)
	return x.buf
}

func (e TL_contacts_importCard) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_importCard)
	x.VectorInt(e.export_card)
	return x.buf
}

func (e TL_messages_getMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getMessages)
	x.VectorInt(e.id)
	return x.buf
}

func (e TL_messages_getDialogs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getDialogs)
	x.Int(e.offset_date)
	x.Int(e.offset_id)
	x.Bytes(e.offset_peer.encode())
	x.Int(e.limit)
	return x.buf
}

func (e TL_messages_getHistory) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getHistory)
	x.Bytes(e.peer.encode())
	x.Int(e.offset_id)
	x.Int(e.offset_date)
	x.Int(e.add_offset)
	x.Int(e.limit)
	x.Int(e.max_id)
	x.Int(e.min_id)
	return x.buf
}

func (e TL_messages_search) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_search)
	var flags int32 = 0
	s := TL_messages_search{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.Bytes(e.peer.encode())
	x.String(e.q)
	x.Bytes(e.filter.encode())
	x.Int(e.min_date)
	x.Int(e.max_date)
	x.Int(e.offset)
	x.Int(e.max_id)
	x.Int(e.limit)
	return x.buf
}

func (e TL_messages_readHistory) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_readHistory)
	x.Bytes(e.peer.encode())
	x.Int(e.max_id)
	return x.buf
}

func (e TL_messages_deleteHistory) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_deleteHistory)
	var flags int32 = 0
	if e.just_clear {
		flags = flags | 1<<0
	}
	s := TL_messages_deleteHistory{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.Bytes(e.peer.encode())
	x.Int(e.max_id)
	return x.buf
}

func (e TL_messages_deleteMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_deleteMessages)
	x.VectorInt(e.id)
	return x.buf
}

func (e TL_messages_receivedMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_receivedMessages)
	x.Int(e.max_id)
	return x.buf
}

func (e TL_messages_setTyping) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_setTyping)
	x.Bytes(e.peer.encode())
	x.Bytes(e.action.encode())
	return x.buf
}

func (e TL_messages_sendMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_sendMessage)
	var flags int32 = 0
	if e.no_webpage {
		flags = flags | 1<<1
	}
	if e.silent {
		flags = flags | 1<<5
	}
	if e.background {
		flags = flags | 1<<6
	}
	if e.clear_draft {
		flags = flags | 1<<7
	}
	s := TL_messages_sendMessage{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.Bytes(e.peer.encode())
	x.Int(e.reply_to_msg_id)
	x.String(e.message)
	x.Long(e.random_id)
	x.Bytes(e.reply_markup.encode())
	x.Vector(e.entities)
	return x.buf
}

func (e TL_messages_sendMedia) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_sendMedia)
	var flags int32 = 0
	if e.silent {
		flags = flags | 1<<5
	}
	if e.background {
		flags = flags | 1<<6
	}
	if e.clear_draft {
		flags = flags | 1<<7
	}
	s := TL_messages_sendMedia{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.Bytes(e.peer.encode())
	x.Int(e.reply_to_msg_id)
	x.Bytes(e.media.encode())
	x.Long(e.random_id)
	x.Bytes(e.reply_markup.encode())
	return x.buf
}

func (e TL_messages_forwardMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_forwardMessages)
	var flags int32 = 0
	if e.silent {
		flags = flags | 1<<5
	}
	if e.background {
		flags = flags | 1<<6
	}
	s := TL_messages_forwardMessages{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.Bytes(e.from_peer.encode())
	x.VectorInt(e.id)
	x.VectorLong(e.random_id)
	x.Bytes(e.to_peer.encode())
	return x.buf
}

func (e TL_messages_reportSpam) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_reportSpam)
	x.Bytes(e.peer.encode())
	return x.buf
}

func (e TL_messages_hideReportSpam) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_hideReportSpam)
	x.Bytes(e.peer.encode())
	return x.buf
}

func (e TL_messages_getPeerSettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getPeerSettings)
	x.Bytes(e.peer.encode())
	return x.buf
}

func (e TL_messages_getChats) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getChats)
	x.VectorInt(e.id)
	return x.buf
}

func (e TL_messages_getFullChat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getFullChat)
	x.Int(e.chat_id)
	return x.buf
}

func (e TL_messages_editChatTitle) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_editChatTitle)
	x.Int(e.chat_id)
	x.String(e.title)
	return x.buf
}

func (e TL_messages_editChatPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_editChatPhoto)
	x.Int(e.chat_id)
	x.Bytes(e.photo.encode())
	return x.buf
}

func (e TL_messages_addChatUser) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_addChatUser)
	x.Int(e.chat_id)
	x.Bytes(e.user_id.encode())
	x.Int(e.fwd_limit)
	return x.buf
}

func (e TL_messages_deleteChatUser) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_deleteChatUser)
	x.Int(e.chat_id)
	x.Bytes(e.user_id.encode())
	return x.buf
}

func (e TL_messages_createChat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_createChat)
	x.Vector(e.users)
	x.String(e.title)
	return x.buf
}

func (e TL_updates_getState) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updates_getState)
	return x.buf
}

func (e TL_updates_getDifference) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updates_getDifference)
	x.Int(e.pts)
	x.Int(e.date)
	x.Int(e.qts)
	return x.buf
}

func (e TL_photos_updateProfilePhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photos_updateProfilePhoto)
	x.Bytes(e.id.encode())
	x.Bytes(e.crop.encode())
	return x.buf
}

func (e TL_photos_uploadProfilePhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photos_uploadProfilePhoto)
	x.Bytes(e.file.encode())
	x.String(e.caption)
	x.Bytes(e.geo_point.encode())
	x.Bytes(e.crop.encode())
	return x.buf
}

func (e TL_photos_deletePhotos) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photos_deletePhotos)
	x.Vector(e.id)
	return x.buf
}

func (e TL_upload_saveFilePart) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_upload_saveFilePart)
	x.Long(e.file_id)
	x.Int(e.file_part)
	x.StringBytes(e.bytes)
	return x.buf
}

func (e TL_upload_getFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_upload_getFile)
	x.Bytes(e.location.encode())
	x.Int(e.offset)
	x.Int(e.limit)
	return x.buf
}

func (e TL_help_getConfig) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_getConfig)
	return x.buf
}

func (e TL_help_getNearestDc) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_getNearestDc)
	return x.buf
}

func (e TL_help_getAppUpdate) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_getAppUpdate)
	return x.buf
}

func (e TL_help_saveAppLog) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_saveAppLog)
	x.Vector(e.events)
	return x.buf
}

func (e TL_help_getInviteText) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_getInviteText)
	return x.buf
}

func (e TL_photos_getUserPhotos) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photos_getUserPhotos)
	x.Bytes(e.user_id.encode())
	x.Int(e.offset)
	x.Long(e.max_id)
	x.Int(e.limit)
	return x.buf
}

func (e TL_messages_forwardMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_forwardMessage)
	x.Bytes(e.peer.encode())
	x.Int(e.id)
	x.Long(e.random_id)
	return x.buf
}

func (e TL_messages_sendBroadcast) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_sendBroadcast)
	x.Vector(e.contacts)
	x.VectorLong(e.random_id)
	x.String(e.message)
	x.Bytes(e.media.encode())
	return x.buf
}

func (e TL_messages_getDhConfig) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getDhConfig)
	x.Int(e.version)
	x.Int(e.random_length)
	return x.buf
}

func (e TL_messages_requestEncryption) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_requestEncryption)
	x.Bytes(e.user_id.encode())
	x.Int(e.random_id)
	x.StringBytes(e.g_a)
	return x.buf
}

func (e TL_messages_acceptEncryption) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_acceptEncryption)
	x.Bytes(e.peer.encode())
	x.StringBytes(e.g_b)
	x.Long(e.key_fingerprint)
	return x.buf
}

func (e TL_messages_discardEncryption) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_discardEncryption)
	x.Int(e.chat_id)
	return x.buf
}

func (e TL_messages_setEncryptedTyping) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_setEncryptedTyping)
	x.Bytes(e.peer.encode())
	x.Bytes(e.typing.encode())
	return x.buf
}

func (e TL_messages_readEncryptedHistory) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_readEncryptedHistory)
	x.Bytes(e.peer.encode())
	x.Int(e.max_date)
	return x.buf
}

func (e TL_messages_sendEncrypted) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_sendEncrypted)
	x.Bytes(e.peer.encode())
	x.Long(e.random_id)
	x.StringBytes(e.data)
	return x.buf
}

func (e TL_messages_sendEncryptedFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_sendEncryptedFile)
	x.Bytes(e.peer.encode())
	x.Long(e.random_id)
	x.StringBytes(e.data)
	x.Bytes(e.file.encode())
	return x.buf
}

func (e TL_messages_sendEncryptedService) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_sendEncryptedService)
	x.Bytes(e.peer.encode())
	x.Long(e.random_id)
	x.StringBytes(e.data)
	return x.buf
}

func (e TL_messages_receivedQueue) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_receivedQueue)
	x.Int(e.max_qts)
	return x.buf
}

func (e TL_upload_saveBigFilePart) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_upload_saveBigFilePart)
	x.Long(e.file_id)
	x.Int(e.file_part)
	x.Int(e.file_total_parts)
	x.StringBytes(e.bytes)
	return x.buf
}

func (e TL_initConnection) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_initConnection)
	x.Int(e.api_id)
	x.String(e.device_model)
	x.String(e.system_version)
	x.String(e.app_version)
	x.String(e.lang_code)
	x.Bytes(e.query.encode())
	return x.buf
}

func (e TL_help_getSupport) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_getSupport)
	return x.buf
}

func (e TL_messages_readMessageContents) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_readMessageContents)
	x.VectorInt(e.id)
	return x.buf
}

func (e TL_account_checkUsername) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_checkUsername)
	x.String(e.username)
	return x.buf
}

func (e TL_account_updateUsername) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_updateUsername)
	x.String(e.username)
	return x.buf
}

func (e TL_contacts_search) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_search)
	x.String(e.q)
	x.Int(e.limit)
	return x.buf
}

func (e TL_account_getPrivacy) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_getPrivacy)
	x.Bytes(e.key.encode())
	return x.buf
}

func (e TL_account_setPrivacy) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_setPrivacy)
	x.Bytes(e.key.encode())
	x.Vector(e.rules)
	return x.buf
}

func (e TL_account_deleteAccount) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_deleteAccount)
	x.String(e.reason)
	return x.buf
}

func (e TL_account_getAccountTTL) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_getAccountTTL)
	return x.buf
}

func (e TL_account_setAccountTTL) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_setAccountTTL)
	x.Bytes(e.ttl.encode())
	return x.buf
}

func (e TL_invokeWithLayer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_invokeWithLayer)
	x.Int(e.layer)
	x.Bytes(e.query.encode())
	return x.buf
}

func (e TL_contacts_resolveUsername) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_resolveUsername)
	x.String(e.username)
	return x.buf
}

func (e TL_account_sendChangePhoneCode) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_sendChangePhoneCode)
	var flags int32 = 0
	if e.allow_flashcall {
		flags = flags | 1<<0
	}
	s := TL_account_sendChangePhoneCode{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.String(e.phone_number)
	x.Bytes(e.current_number.encode())
	return x.buf
}

func (e TL_account_changePhone) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_changePhone)
	x.String(e.phone_number)
	x.String(e.phone_code_hash)
	x.String(e.phone_code)
	return x.buf
}

func (e TL_messages_getStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getStickers)
	x.String(e.emoticon)
	x.String(e.hash)
	return x.buf
}

func (e TL_messages_getAllStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getAllStickers)
	x.Int(e.hash)
	return x.buf
}

func (e TL_account_updateDeviceLocked) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_updateDeviceLocked)
	x.Int(e.period)
	return x.buf
}

func (e TL_auth_importBotAuthorization) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_importBotAuthorization)
	x.Int(e.flags)
	x.Int(e.api_id)
	x.String(e.api_hash)
	x.String(e.bot_auth_token)
	return x.buf
}

func (e TL_messages_getWebPagePreview) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getWebPagePreview)
	x.String(e.message)
	return x.buf
}

func (e TL_account_getAuthorizations) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_getAuthorizations)
	return x.buf
}

func (e TL_account_resetAuthorization) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_resetAuthorization)
	x.Long(e.hash)
	return x.buf
}

func (e TL_account_getPassword) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_getPassword)
	return x.buf
}

func (e TL_account_getPasswordSettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_getPasswordSettings)
	x.StringBytes(e.current_password_hash)
	return x.buf
}

func (e TL_account_updatePasswordSettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_updatePasswordSettings)
	x.StringBytes(e.current_password_hash)
	x.Bytes(e.new_settings.encode())
	return x.buf
}

func (e TL_auth_checkPassword) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_checkPassword)
	x.StringBytes(e.password_hash)
	return x.buf
}

func (e TL_auth_requestPasswordRecovery) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_requestPasswordRecovery)
	return x.buf
}

func (e TL_auth_recoverPassword) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_recoverPassword)
	x.String(e.code)
	return x.buf
}

func (e TL_invokeWithoutUpdates) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_invokeWithoutUpdates)
	x.Bytes(e.query.encode())
	return x.buf
}

func (e TL_messages_exportChatInvite) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_exportChatInvite)
	x.Int(e.chat_id)
	return x.buf
}

func (e TL_messages_checkChatInvite) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_checkChatInvite)
	x.String(e.hash)
	return x.buf
}

func (e TL_messages_importChatInvite) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_importChatInvite)
	x.String(e.hash)
	return x.buf
}

func (e TL_messages_getStickerSet) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getStickerSet)
	x.Bytes(e.stickerset.encode())
	return x.buf
}

func (e TL_messages_installStickerSet) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_installStickerSet)
	x.Bytes(e.stickerset.encode())
	x.Bytes(e.archived.encode())
	return x.buf
}

func (e TL_messages_uninstallStickerSet) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_uninstallStickerSet)
	x.Bytes(e.stickerset.encode())
	return x.buf
}

func (e TL_messages_startBot) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_startBot)
	x.Bytes(e.bot.encode())
	x.Bytes(e.peer.encode())
	x.Long(e.random_id)
	x.String(e.start_param)
	return x.buf
}

func (e TL_help_getAppChangelog) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_getAppChangelog)
	return x.buf
}

func (e TL_messages_getMessagesViews) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getMessagesViews)
	x.Bytes(e.peer.encode())
	x.VectorInt(e.id)
	x.Bytes(e.increment.encode())
	return x.buf
}

func (e TL_channels_readHistory) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_readHistory)
	x.Bytes(e.channel.encode())
	x.Int(e.max_id)
	return x.buf
}

func (e TL_channels_deleteMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_deleteMessages)
	x.Bytes(e.channel.encode())
	x.VectorInt(e.id)
	return x.buf
}

func (e TL_channels_deleteUserHistory) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_deleteUserHistory)
	x.Bytes(e.channel.encode())
	x.Bytes(e.user_id.encode())
	return x.buf
}

func (e TL_channels_reportSpam) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_reportSpam)
	x.Bytes(e.channel.encode())
	x.Bytes(e.user_id.encode())
	x.VectorInt(e.id)
	return x.buf
}

func (e TL_channels_getMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_getMessages)
	x.Bytes(e.channel.encode())
	x.VectorInt(e.id)
	return x.buf
}

func (e TL_channels_getParticipants) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_getParticipants)
	x.Bytes(e.channel.encode())
	x.Bytes(e.filter.encode())
	x.Int(e.offset)
	x.Int(e.limit)
	return x.buf
}

func (e TL_channels_getParticipant) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_getParticipant)
	x.Bytes(e.channel.encode())
	x.Bytes(e.user_id.encode())
	return x.buf
}

func (e TL_channels_getChannels) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_getChannels)
	x.Vector(e.id)
	return x.buf
}

func (e TL_channels_getFullChannel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_getFullChannel)
	x.Bytes(e.channel.encode())
	return x.buf
}

func (e TL_channels_createChannel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_createChannel)
	var flags int32 = 0
	if e.broadcast {
		flags = flags | 1<<0
	}
	if e.megagroup {
		flags = flags | 1<<1
	}
	s := TL_channels_createChannel{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.String(e.title)
	x.String(e.about)
	return x.buf
}

func (e TL_channels_editAbout) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_editAbout)
	x.Bytes(e.channel.encode())
	x.String(e.about)
	return x.buf
}

func (e TL_channels_editAdmin) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_editAdmin)
	x.Bytes(e.channel.encode())
	x.Bytes(e.user_id.encode())
	x.Bytes(e.role.encode())
	return x.buf
}

func (e TL_channels_editTitle) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_editTitle)
	x.Bytes(e.channel.encode())
	x.String(e.title)
	return x.buf
}

func (e TL_channels_editPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_editPhoto)
	x.Bytes(e.channel.encode())
	x.Bytes(e.photo.encode())
	return x.buf
}

func (e TL_channels_checkUsername) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_checkUsername)
	x.Bytes(e.channel.encode())
	x.String(e.username)
	return x.buf
}

func (e TL_channels_updateUsername) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_updateUsername)
	x.Bytes(e.channel.encode())
	x.String(e.username)
	return x.buf
}

func (e TL_channels_joinChannel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_joinChannel)
	x.Bytes(e.channel.encode())
	return x.buf
}

func (e TL_channels_leaveChannel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_leaveChannel)
	x.Bytes(e.channel.encode())
	return x.buf
}

func (e TL_channels_inviteToChannel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_inviteToChannel)
	x.Bytes(e.channel.encode())
	x.Vector(e.users)
	return x.buf
}

func (e TL_channels_kickFromChannel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_kickFromChannel)
	x.Bytes(e.channel.encode())
	x.Bytes(e.user_id.encode())
	x.Bytes(e.kicked.encode())
	return x.buf
}

func (e TL_channels_exportInvite) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_exportInvite)
	x.Bytes(e.channel.encode())
	return x.buf
}

func (e TL_channels_deleteChannel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_deleteChannel)
	x.Bytes(e.channel.encode())
	return x.buf
}

func (e TL_updates_getChannelDifference) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updates_getChannelDifference)
	x.Bytes(e.channel.encode())
	x.Bytes(e.filter.encode())
	x.Int(e.pts)
	x.Int(e.limit)
	return x.buf
}

func (e TL_messages_toggleChatAdmins) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_toggleChatAdmins)
	x.Int(e.chat_id)
	x.Bytes(e.enabled.encode())
	return x.buf
}

func (e TL_messages_editChatAdmin) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_editChatAdmin)
	x.Int(e.chat_id)
	x.Bytes(e.user_id.encode())
	x.Bytes(e.is_admin.encode())
	return x.buf
}

func (e TL_messages_migrateChat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_migrateChat)
	x.Int(e.chat_id)
	return x.buf
}

func (e TL_messages_searchGlobal) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_searchGlobal)
	x.String(e.q)
	x.Int(e.offset_date)
	x.Bytes(e.offset_peer.encode())
	x.Int(e.offset_id)
	x.Int(e.limit)
	return x.buf
}

func (e TL_help_getTermsOfService) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_getTermsOfService)
	return x.buf
}

func (e TL_messages_reorderStickerSets) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_reorderStickerSets)
	x.VectorLong(e.order)
	return x.buf
}

func (e TL_messages_getDocumentByHash) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getDocumentByHash)
	x.StringBytes(e.sha256)
	x.Int(e.size)
	x.String(e.mime_type)
	return x.buf
}

func (e TL_messages_searchGifs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_searchGifs)
	x.String(e.q)
	x.Int(e.offset)
	return x.buf
}

func (e TL_messages_getSavedGifs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getSavedGifs)
	x.Int(e.hash)
	return x.buf
}

func (e TL_messages_saveGif) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_saveGif)
	x.Bytes(e.id.encode())
	x.Bytes(e.unsave.encode())
	return x.buf
}

func (e TL_messages_getInlineBotResults) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getInlineBotResults)
	var flags int32 = 0
	s := TL_messages_getInlineBotResults{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.Bytes(e.bot.encode())
	x.Bytes(e.peer.encode())
	x.Bytes(e.geo_point.encode())
	x.String(e.query)
	x.String(e.offset)
	return x.buf
}

func (e TL_messages_setInlineBotResults) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_setInlineBotResults)
	var flags int32 = 0
	if e.gallery {
		flags = flags | 1<<0
	}
	if e.private {
		flags = flags | 1<<1
	}
	s := TL_messages_setInlineBotResults{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.Long(e.query_id)
	x.Vector(e.results)
	x.Int(e.cache_time)
	x.String(e.next_offset)
	x.Bytes(e.switch_pm.encode())
	return x.buf
}

func (e TL_messages_sendInlineBotResult) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_sendInlineBotResult)
	var flags int32 = 0
	if e.silent {
		flags = flags | 1<<5
	}
	if e.background {
		flags = flags | 1<<6
	}
	if e.clear_draft {
		flags = flags | 1<<7
	}
	s := TL_messages_sendInlineBotResult{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.Bytes(e.peer.encode())
	x.Int(e.reply_to_msg_id)
	x.Long(e.random_id)
	x.Long(e.query_id)
	x.String(e.id)
	return x.buf
}

func (e TL_channels_toggleInvites) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_toggleInvites)
	x.Bytes(e.channel.encode())
	x.Bytes(e.enabled.encode())
	return x.buf
}

func (e TL_channels_exportMessageLink) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_exportMessageLink)
	x.Bytes(e.channel.encode())
	x.Int(e.id)
	return x.buf
}

func (e TL_channels_toggleSignatures) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_toggleSignatures)
	x.Bytes(e.channel.encode())
	x.Bytes(e.enabled.encode())
	return x.buf
}

func (e TL_channels_updatePinnedMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_channels_updatePinnedMessage)
	var flags int32 = 0
	if e.silent {
		flags = flags | 1<<0
	}
	s := TL_channels_updatePinnedMessage{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.Bytes(e.channel.encode())
	x.Int(e.id)
	return x.buf
}

func (e TL_auth_resendCode) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_resendCode)
	x.String(e.phone_number)
	x.String(e.phone_code_hash)
	return x.buf
}

func (e TL_auth_cancelCode) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_cancelCode)
	x.String(e.phone_number)
	x.String(e.phone_code_hash)
	return x.buf
}

func (e TL_messages_getMessageEditData) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getMessageEditData)
	x.Bytes(e.peer.encode())
	x.Int(e.id)
	return x.buf
}

func (e TL_messages_editMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_editMessage)
	var flags int32 = 0
	if e.no_webpage {
		flags = flags | 1<<1
	}
	s := TL_messages_editMessage{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.Bytes(e.peer.encode())
	x.Int(e.id)
	x.String(e.message)
	x.Bytes(e.reply_markup.encode())
	x.Vector(e.entities)
	return x.buf
}

func (e TL_messages_editInlineBotMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_editInlineBotMessage)
	var flags int32 = 0
	if e.no_webpage {
		flags = flags | 1<<1
	}
	s := TL_messages_editInlineBotMessage{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.Bytes(e.id.encode())
	x.String(e.message)
	x.Bytes(e.reply_markup.encode())
	x.Vector(e.entities)
	return x.buf
}

func (e TL_messages_getBotCallbackAnswer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getBotCallbackAnswer)
	x.Bytes(e.peer.encode())
	x.Int(e.msg_id)
	x.StringBytes(e.data)
	return x.buf
}

func (e TL_messages_setBotCallbackAnswer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_setBotCallbackAnswer)
	var flags int32 = 0
	if e.alert {
		flags = flags | 1<<1
	}
	s := TL_messages_setBotCallbackAnswer{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.Long(e.query_id)
	x.String(e.message)
	x.String(e.url)
	return x.buf
}

func (e TL_contacts_getTopPeers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_getTopPeers)
	var flags int32 = 0
	if e.correspondents {
		flags = flags | 1<<0
	}
	if e.bots_pm {
		flags = flags | 1<<1
	}
	if e.bots_inline {
		flags = flags | 1<<2
	}
	if e.groups {
		flags = flags | 1<<10
	}
	if e.channels {
		flags = flags | 1<<15
	}
	s := TL_contacts_getTopPeers{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.Int(e.offset)
	x.Int(e.limit)
	x.Int(e.hash)
	return x.buf
}

func (e TL_contacts_resetTopPeerRating) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_resetTopPeerRating)
	x.Bytes(e.category.encode())
	x.Bytes(e.peer.encode())
	return x.buf
}

func (e TL_messages_getPeerDialogs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getPeerDialogs)
	x.Vector(e.peers)
	return x.buf
}

func (e TL_messages_saveDraft) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_saveDraft)
	var flags int32 = 0
	if e.no_webpage {
		flags = flags | 1<<1
	}
	s := TL_messages_saveDraft{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.Int(e.reply_to_msg_id)
	x.Bytes(e.peer.encode())
	x.String(e.message)
	x.Vector(e.entities)
	return x.buf
}

func (e TL_messages_getAllDrafts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getAllDrafts)
	return x.buf
}

func (e TL_messages_getFeaturedStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getFeaturedStickers)
	x.Int(e.hash)
	return x.buf
}

func (e TL_messages_readFeaturedStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_readFeaturedStickers)
	return x.buf
}

func (e TL_messages_getRecentStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getRecentStickers)
	x.Int(e.hash)
	return x.buf
}

func (e TL_messages_saveRecentSticker) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_saveRecentSticker)
	x.Bytes(e.id.encode())
	x.Bytes(e.unsave.encode())
	return x.buf
}

func (e TL_messages_clearRecentStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_clearRecentStickers)
	return x.buf
}

func (e TL_messages_getUnusedStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getUnusedStickers)
	x.Int(e.limit)
	return x.buf
}

func (e TL_messages_getArchivedStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getArchivedStickers)
	x.Long(e.offset_id)
	x.Int(e.limit)
	return x.buf
}

func (e TL_account_sendConfirmPhoneCode) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_sendConfirmPhoneCode)
	var flags int32 = 0
	if e.allow_flashcall {
		flags = flags | 1<<0
	}
	s := TL_account_sendConfirmPhoneCode{}
	sRef := reflect.TypeOf(s)

	for i := 0; i < sRef.NumField(); i++ {
		field := sRef.Field(i)
		if flag_byte, ok := field.Tag.Lookup("flag_byte"); ok {
			b, _ := strconv.ParseUint(flag_byte, 10, 64)
			flags = flags | 1<<b
		}
	}
	x.Int(flags)
	x.String(e.hash)
	x.Bytes(e.current_number.encode())
	return x.buf
}

func (e TL_account_confirmPhone) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_confirmPhone)
	x.String(e.phone_code_hash)
	x.String(e.phone_code)
	return x.buf
}

func (m *DecodeBuf) ObjectGenerated(constructor uint32) (r TL) {
	switch constructor {
	case crc_boolFalse:
		r = TL_boolFalse{}

	case crc_boolTrue:
		r = TL_boolTrue{}

	case crc_true:
		r = TL_true{}

	case crc_error:
		r = TL_error{
			m.Int(),
			m.String(),
		}

	case crc_null:
		r = TL_null{}

	case crc_inputPeerEmpty:
		r = TL_inputPeerEmpty{}

	case crc_inputPeerSelf:
		r = TL_inputPeerSelf{}

	case crc_inputPeerChat:
		r = TL_inputPeerChat{
			m.Int(),
		}

	case crc_inputUserEmpty:
		r = TL_inputUserEmpty{}

	case crc_inputUserSelf:
		r = TL_inputUserSelf{}

	case crc_inputPhoneContact:
		r = TL_inputPhoneContact{
			m.Long(),
			m.String(),
			m.String(),
			m.String(),
		}

	case crc_inputFile:
		r = TL_inputFile{
			m.Long(),
			m.Int(),
			m.String(),
			m.String(),
		}

	case crc_inputMediaEmpty:
		r = TL_inputMediaEmpty{}

	case crc_inputMediaUploadedPhoto:
		r = TL_inputMediaUploadedPhoto{
			m.Object(),
			m.String(),
		}

	case crc_inputMediaPhoto:
		r = TL_inputMediaPhoto{
			m.Object(),
			m.String(),
		}

	case crc_inputMediaGeoPoint:
		r = TL_inputMediaGeoPoint{
			m.Object(),
		}

	case crc_inputMediaContact:
		r = TL_inputMediaContact{
			m.String(),
			m.String(),
			m.String(),
		}

	case crc_inputChatPhotoEmpty:
		r = TL_inputChatPhotoEmpty{}

	case crc_inputChatUploadedPhoto:
		r = TL_inputChatUploadedPhoto{
			m.Object(),
			m.Object(),
		}

	case crc_inputChatPhoto:
		r = TL_inputChatPhoto{
			m.Object(),
			m.Object(),
		}

	case crc_inputGeoPointEmpty:
		r = TL_inputGeoPointEmpty{}

	case crc_inputGeoPoint:
		r = TL_inputGeoPoint{
			m.Double(),
			m.Double(),
		}

	case crc_inputPhotoEmpty:
		r = TL_inputPhotoEmpty{}

	case crc_inputPhoto:
		r = TL_inputPhoto{
			m.Long(),
			m.Long(),
		}

	case crc_inputFileLocation:
		r = TL_inputFileLocation{
			m.Long(),
			m.Int(),
			m.Long(),
		}

	case crc_inputPhotoCropAuto:
		r = TL_inputPhotoCropAuto{}

	case crc_inputPhotoCrop:
		r = TL_inputPhotoCrop{
			m.Double(),
			m.Double(),
			m.Double(),
		}

	case crc_inputAppEvent:
		r = TL_inputAppEvent{
			m.Double(),
			m.String(),
			m.Long(),
			m.String(),
		}

	case crc_peerUser:
		r = TL_peerUser{
			m.Int(),
		}

	case crc_peerChat:
		r = TL_peerChat{
			m.Int(),
		}

	case crc_storage_fileUnknown:
		r = TL_storage_fileUnknown{}

	case crc_storage_fileJpeg:
		r = TL_storage_fileJpeg{}

	case crc_storage_fileGif:
		r = TL_storage_fileGif{}

	case crc_storage_filePng:
		r = TL_storage_filePng{}

	case crc_storage_filePdf:
		r = TL_storage_filePdf{}

	case crc_storage_fileMp3:
		r = TL_storage_fileMp3{}

	case crc_storage_fileMov:
		r = TL_storage_fileMov{}

	case crc_storage_filePartial:
		r = TL_storage_filePartial{}

	case crc_storage_fileMp4:
		r = TL_storage_fileMp4{}

	case crc_storage_fileWebp:
		r = TL_storage_fileWebp{}

	case crc_fileLocationUnavailable:
		r = TL_fileLocationUnavailable{
			m.Long(),
			m.Int(),
			m.Long(),
		}

	case crc_fileLocation:
		r = TL_fileLocation{
			m.Int(),
			m.Long(),
			m.Int(),
			m.Long(),
		}

	case crc_userEmpty:
		r = TL_userEmpty{
			m.Int(),
		}

	case crc_userProfilePhotoEmpty:
		r = TL_userProfilePhotoEmpty{}

	case crc_userProfilePhoto:
		r = TL_userProfilePhoto{
			m.Long(),
			m.Object(),
			m.Object(),
		}

	case crc_userStatusEmpty:
		r = TL_userStatusEmpty{}

	case crc_userStatusOnline:
		r = TL_userStatusOnline{
			m.Int(),
		}

	case crc_userStatusOffline:
		r = TL_userStatusOffline{
			m.Int(),
		}

	case crc_chatEmpty:
		r = TL_chatEmpty{
			m.Int(),
		}

	case crc_chat:
		rInit := TL_chat{}
		rInit.flags = m.Int()
		if rInit.flags&(1<<0) == (1 << 0) {
			rInit.creator = true
		}
		if rInit.flags&(1<<1) == (1 << 1) {
			rInit.kicked = true
		}
		if rInit.flags&(1<<2) == (1 << 2) {
			rInit.left = true
		}
		if rInit.flags&(1<<3) == (1 << 3) {
			rInit.admins_enabled = true
		}
		if rInit.flags&(1<<4) == (1 << 4) {
			rInit.admin = true
		}
		if rInit.flags&(1<<5) == (1 << 5) {
			rInit.deactivated = true
		}
		rInit.id = m.Int()
		rInit.title = m.String()
		rInit.photo = m.Object()
		rInit.participants_count = m.Int()
		rInit.date = m.Int()
		rInit.version = m.Int()
		if rInit.flags&(1<<6) == (1 << 6) {
			rInit.migrated_to = m.Object()
		}
		r = rInit

	case crc_chatForbidden:
		r = TL_chatForbidden{
			m.Int(),
			m.String(),
		}

	case crc_chatFull:
		r = TL_chatFull{
			m.Int(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Vector(),
		}

	case crc_chatParticipant:
		r = TL_chatParticipant{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_chatParticipantsForbidden:
		rInit := TL_chatParticipantsForbidden{}
		rInit.flags = m.Int()
		rInit.chat_id = m.Int()
		if rInit.flags&(1<<0) == (1 << 0) {
			rInit.self_participant = m.Object()
		}
		r = rInit

	case crc_chatParticipants:
		r = TL_chatParticipants{
			m.Int(),
			m.Vector(),
			m.Int(),
		}

	case crc_chatPhotoEmpty:
		r = TL_chatPhotoEmpty{}

	case crc_chatPhoto:
		r = TL_chatPhoto{
			m.Object(),
			m.Object(),
		}

	case crc_messageEmpty:
		r = TL_messageEmpty{
			m.Int(),
		}

	case crc_message:
		rInit := TL_message{}
		rInit.flags = m.Int()
		if rInit.flags&(1<<1) == (1 << 1) {
			rInit.out = true
		}
		if rInit.flags&(1<<4) == (1 << 4) {
			rInit.mentioned = true
		}
		if rInit.flags&(1<<5) == (1 << 5) {
			rInit.media_unread = true
		}
		if rInit.flags&(1<<13) == (1 << 13) {
			rInit.silent = true
		}
		if rInit.flags&(1<<14) == (1 << 14) {
			rInit.post = true
		}
		rInit.id = m.Int()
		if rInit.flags&(1<<8) == (1 << 8) {
			rInit.from_id = m.Int()
		}
		rInit.to_id = m.Object()
		if rInit.flags&(1<<2) == (1 << 2) {
			rInit.fwd_from = m.Object()
		}
		if rInit.flags&(1<<11) == (1 << 11) {
			rInit.via_bot_id = m.Int()
		}
		if rInit.flags&(1<<3) == (1 << 3) {
			rInit.reply_to_msg_id = m.Int()
		}
		rInit.date = m.Int()
		rInit.message = m.String()
		if rInit.flags&(1<<9) == (1 << 9) {
			rInit.media = m.Object()
		}
		if rInit.flags&(1<<6) == (1 << 6) {
			rInit.reply_markup = m.Object()
		}
		if rInit.flags&(1<<7) == (1 << 7) {
			rInit.entities = m.Vector()
		}
		if rInit.flags&(1<<10) == (1 << 10) {
			rInit.views = m.Int()
		}
		if rInit.flags&(1<<15) == (1 << 15) {
			rInit.edit_date = m.Int()
		}
		r = rInit

	case crc_messageService:
		rInit := TL_messageService{}
		rInit.flags = m.Int()
		if rInit.flags&(1<<1) == (1 << 1) {
			rInit.out = true
		}
		if rInit.flags&(1<<4) == (1 << 4) {
			rInit.mentioned = true
		}
		if rInit.flags&(1<<5) == (1 << 5) {
			rInit.media_unread = true
		}
		if rInit.flags&(1<<13) == (1 << 13) {
			rInit.silent = true
		}
		if rInit.flags&(1<<14) == (1 << 14) {
			rInit.post = true
		}
		rInit.id = m.Int()
		if rInit.flags&(1<<8) == (1 << 8) {
			rInit.from_id = m.Int()
		}
		rInit.to_id = m.Object()
		if rInit.flags&(1<<3) == (1 << 3) {
			rInit.reply_to_msg_id = m.Int()
		}
		rInit.date = m.Int()
		rInit.action = m.Object()
		r = rInit

	case crc_messageMediaEmpty:
		r = TL_messageMediaEmpty{}

	case crc_messageMediaPhoto:
		r = TL_messageMediaPhoto{
			m.Object(),
			m.String(),
		}

	case crc_messageMediaGeo:
		r = TL_messageMediaGeo{
			m.Object(),
		}

	case crc_messageMediaContact:
		r = TL_messageMediaContact{
			m.String(),
			m.String(),
			m.String(),
			m.Int(),
		}

	case crc_messageMediaUnsupported:
		r = TL_messageMediaUnsupported{}

	case crc_messageActionEmpty:
		r = TL_messageActionEmpty{}

	case crc_messageActionChatCreate:
		r = TL_messageActionChatCreate{
			m.String(),
			m.VectorInt(),
		}

	case crc_messageActionChatEditTitle:
		r = TL_messageActionChatEditTitle{
			m.String(),
		}

	case crc_messageActionChatEditPhoto:
		r = TL_messageActionChatEditPhoto{
			m.Object(),
		}

	case crc_messageActionChatDeletePhoto:
		r = TL_messageActionChatDeletePhoto{}

	case crc_messageActionChatAddUser:
		r = TL_messageActionChatAddUser{
			m.VectorInt(),
		}

	case crc_messageActionChatDeleteUser:
		r = TL_messageActionChatDeleteUser{
			m.Int(),
		}

	case crc_dialog:
		rInit := TL_dialog{}
		rInit.flags = m.Int()
		rInit.peer = m.Object()
		rInit.top_message = m.Int()
		rInit.read_inbox_max_id = m.Int()
		rInit.read_outbox_max_id = m.Int()
		rInit.unread_count = m.Int()
		rInit.notify_settings = m.Object()
		if rInit.flags&(1<<0) == (1 << 0) {
			rInit.pts = m.Int()
		}
		if rInit.flags&(1<<1) == (1 << 1) {
			rInit.draft = m.Object()
		}
		r = rInit

	case crc_photoEmpty:
		r = TL_photoEmpty{
			m.Long(),
		}

	case crc_photo:
		r = TL_photo{
			m.Long(),
			m.Long(),
			m.Int(),
			m.Vector(),
		}

	case crc_photoSizeEmpty:
		r = TL_photoSizeEmpty{
			m.String(),
		}

	case crc_photoSize:
		r = TL_photoSize{
			m.String(),
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_photoCachedSize:
		r = TL_photoCachedSize{
			m.String(),
			m.Object(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
		}

	case crc_geoPointEmpty:
		r = TL_geoPointEmpty{}

	case crc_geoPoint:
		r = TL_geoPoint{
			m.Double(),
			m.Double(),
		}

	case crc_auth_checkedPhone:
		r = TL_auth_checkedPhone{
			m.Object(),
		}

	case crc_auth_sentCode:
		rInit := TL_auth_sentCode{}
		rInit.flags = m.Int()
		if rInit.flags&(1<<0) == (1 << 0) {
			rInit.phone_registered = true
		}
		rInit._type = m.Object()
		rInit.phone_code_hash = m.String()
		if rInit.flags&(1<<1) == (1 << 1) {
			rInit.next_type = m.Object()
		}
		if rInit.flags&(1<<2) == (1 << 2) {
			rInit.timeout = m.Int()
		}
		r = rInit

	case crc_auth_authorization:
		rInit := TL_auth_authorization{}
		rInit.flags = m.Int()
		if rInit.flags&(1<<0) == (1 << 0) {
			rInit.tmp_sessions = m.Int()
		}
		rInit.user = m.Object()
		r = rInit

	case crc_auth_exportedAuthorization:
		r = TL_auth_exportedAuthorization{
			m.Int(),
			m.StringBytes(),
		}

	case crc_inputNotifyPeer:
		r = TL_inputNotifyPeer{
			m.Object(),
		}

	case crc_inputNotifyUsers:
		r = TL_inputNotifyUsers{}

	case crc_inputNotifyChats:
		r = TL_inputNotifyChats{}

	case crc_inputNotifyAll:
		r = TL_inputNotifyAll{}

	case crc_inputPeerNotifyEventsEmpty:
		r = TL_inputPeerNotifyEventsEmpty{}

	case crc_inputPeerNotifyEventsAll:
		r = TL_inputPeerNotifyEventsAll{}

	case crc_inputPeerNotifySettings:
		rInit := TL_inputPeerNotifySettings{}
		rInit.flags = m.Int()
		if rInit.flags&(1<<0) == (1 << 0) {
			rInit.show_previews = true
		}
		if rInit.flags&(1<<1) == (1 << 1) {
			rInit.silent = true
		}
		rInit.mute_until = m.Int()
		rInit.sound = m.String()
		r = rInit

	case crc_peerNotifyEventsEmpty:
		r = TL_peerNotifyEventsEmpty{}

	case crc_peerNotifyEventsAll:
		r = TL_peerNotifyEventsAll{}

	case crc_peerNotifySettingsEmpty:
		r = TL_peerNotifySettingsEmpty{}

	case crc_peerNotifySettings:
		rInit := TL_peerNotifySettings{}
		rInit.flags = m.Int()
		if rInit.flags&(1<<0) == (1 << 0) {
			rInit.show_previews = true
		}
		if rInit.flags&(1<<1) == (1 << 1) {
			rInit.silent = true
		}
		rInit.mute_until = m.Int()
		rInit.sound = m.String()
		r = rInit

	case crc_peerSettings:
		rInit := TL_peerSettings{}
		rInit.flags = m.Int()
		if rInit.flags&(1<<0) == (1 << 0) {
			rInit.report_spam = true
		}
		r = rInit

	case crc_wallPaper:
		r = TL_wallPaper{
			m.Int(),
			m.String(),
			m.Vector(),
			m.Int(),
		}

	case crc_inputReportReasonSpam:
		r = TL_inputReportReasonSpam{}

	case crc_inputReportReasonViolence:
		r = TL_inputReportReasonViolence{}

	case crc_inputReportReasonPornography:
		r = TL_inputReportReasonPornography{}

	case crc_inputReportReasonOther:
		r = TL_inputReportReasonOther{
			m.String(),
		}

	case crc_userFull:
		rInit := TL_userFull{}
		rInit.flags = m.Int()
		if rInit.flags&(1<<0) == (1 << 0) {
			rInit.blocked = true
		}
		rInit.user = m.Object()
		if rInit.flags&(1<<1) == (1 << 1) {
			rInit.about = m.String()
		}
		rInit.link = m.Object()
		if rInit.flags&(1<<2) == (1 << 2) {
			rInit.profile_photo = m.Object()
		}
		rInit.notify_settings = m.Object()
		if rInit.flags&(1<<3) == (1 << 3) {
			rInit.bot_info = m.Object()
		}
		r = rInit

	case crc_contact:
		r = TL_contact{
			m.Int(),
			m.Object(),
		}

	case crc_importedContact:
		r = TL_importedContact{
			m.Int(),
			m.Long(),
		}

	case crc_contactBlocked:
		r = TL_contactBlocked{
			m.Int(),
			m.Int(),
		}

	case crc_contactStatus:
		r = TL_contactStatus{
			m.Int(),
			m.Object(),
		}

	case crc_contacts_link:
		r = TL_contacts_link{
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_contacts_contactsNotModified:
		r = TL_contacts_contactsNotModified{}

	case crc_contacts_contacts:
		r = TL_contacts_contacts{
			m.Vector(),
			m.Vector(),
		}

	case crc_contacts_importedContacts:
		r = TL_contacts_importedContacts{
			m.Vector(),
			m.VectorLong(),
			m.Vector(),
		}

	case crc_contacts_blocked:
		r = TL_contacts_blocked{
			m.Vector(),
			m.Vector(),
		}

	case crc_contacts_blockedSlice:
		r = TL_contacts_blockedSlice{
			m.Int(),
			m.Vector(),
			m.Vector(),
		}

	case crc_messages_dialogs:
		r = TL_messages_dialogs{
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case crc_messages_dialogsSlice:
		r = TL_messages_dialogsSlice{
			m.Int(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case crc_messages_messages:
		r = TL_messages_messages{
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case crc_messages_messagesSlice:
		r = TL_messages_messagesSlice{
			m.Int(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case crc_messages_chats:
		r = TL_messages_chats{
			m.Vector(),
		}

	case crc_messages_chatFull:
		r = TL_messages_chatFull{
			m.Object(),
			m.Vector(),
			m.Vector(),
		}

	case crc_messages_affectedHistory:
		r = TL_messages_affectedHistory{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_inputMessagesFilterEmpty:
		r = TL_inputMessagesFilterEmpty{}

	case crc_inputMessagesFilterPhotos:
		r = TL_inputMessagesFilterPhotos{}

	case crc_inputMessagesFilterVideo:
		r = TL_inputMessagesFilterVideo{}

	case crc_inputMessagesFilterPhotoVideo:
		r = TL_inputMessagesFilterPhotoVideo{}

	case crc_inputMessagesFilterPhotoVideoDocuments:
		r = TL_inputMessagesFilterPhotoVideoDocuments{}

	case crc_inputMessagesFilterDocument:
		r = TL_inputMessagesFilterDocument{}

	case crc_inputMessagesFilterUrl:
		r = TL_inputMessagesFilterUrl{}

	case crc_inputMessagesFilterGif:
		r = TL_inputMessagesFilterGif{}

	case crc_updateNewMessage:
		r = TL_updateNewMessage{
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crc_updateMessageID:
		r = TL_updateMessageID{
			m.Int(),
			m.Long(),
		}

	case crc_updateDeleteMessages:
		r = TL_updateDeleteMessages{
			m.VectorInt(),
			m.Int(),
			m.Int(),
		}

	case crc_updateUserTyping:
		r = TL_updateUserTyping{
			m.Int(),
			m.Object(),
		}

	case crc_updateChatUserTyping:
		r = TL_updateChatUserTyping{
			m.Int(),
			m.Int(),
			m.Object(),
		}

	case crc_updateChatParticipants:
		r = TL_updateChatParticipants{
			m.Object(),
		}

	case crc_updateUserStatus:
		r = TL_updateUserStatus{
			m.Int(),
			m.Object(),
		}

	case crc_updateUserName:
		r = TL_updateUserName{
			m.Int(),
			m.String(),
			m.String(),
			m.String(),
		}

	case crc_updateUserPhoto:
		r = TL_updateUserPhoto{
			m.Int(),
			m.Int(),
			m.Object(),
			m.Object(),
		}

	case crc_updateContactRegistered:
		r = TL_updateContactRegistered{
			m.Int(),
			m.Int(),
		}

	case crc_updateContactLink:
		r = TL_updateContactLink{
			m.Int(),
			m.Object(),
			m.Object(),
		}

	case crc_updateNewAuthorization:
		r = TL_updateNewAuthorization{
			m.Long(),
			m.Int(),
			m.String(),
			m.String(),
		}

	case crc_updates_state:
		r = TL_updates_state{
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_updates_differenceEmpty:
		r = TL_updates_differenceEmpty{
			m.Int(),
			m.Int(),
		}

	case crc_updates_difference:
		r = TL_updates_difference{
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Object(),
		}

	case crc_updates_differenceSlice:
		r = TL_updates_differenceSlice{
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Object(),
		}

	case crc_updatesTooLong:
		r = TL_updatesTooLong{}

	case crc_updateShortMessage:
		rInit := TL_updateShortMessage{}
		rInit.flags = m.Int()
		if rInit.flags&(1<<1) == (1 << 1) {
			rInit.out = true
		}
		if rInit.flags&(1<<4) == (1 << 4) {
			rInit.mentioned = true
		}
		if rInit.flags&(1<<5) == (1 << 5) {
			rInit.media_unread = true
		}
		if rInit.flags&(1<<13) == (1 << 13) {
			rInit.silent = true
		}
		rInit.id = m.Int()
		rInit.user_id = m.Int()
		rInit.message = m.String()
		rInit.pts = m.Int()
		rInit.pts_count = m.Int()
		rInit.date = m.Int()
		if rInit.flags&(1<<2) == (1 << 2) {
			rInit.fwd_from = m.Object()
		}
		if rInit.flags&(1<<11) == (1 << 11) {
			rInit.via_bot_id = m.Int()
		}
		if rInit.flags&(1<<3) == (1 << 3) {
			rInit.reply_to_msg_id = m.Int()
		}
		if rInit.flags&(1<<7) == (1 << 7) {
			rInit.entities = m.Vector()
		}
		r = rInit

	case crc_updateShortChatMessage:
		rInit := TL_updateShortChatMessage{}
		rInit.flags = m.Int()
		if rInit.flags&(1<<1) == (1 << 1) {
			rInit.out = true
		}
		if rInit.flags&(1<<4) == (1 << 4) {
			rInit.mentioned = true
		}
		if rInit.flags&(1<<5) == (1 << 5) {
			rInit.media_unread = true
		}
		if rInit.flags&(1<<13) == (1 << 13) {
			rInit.silent = true
		}
		rInit.id = m.Int()
		rInit.from_id = m.Int()
		rInit.chat_id = m.Int()
		rInit.message = m.String()
		rInit.pts = m.Int()
		rInit.pts_count = m.Int()
		rInit.date = m.Int()
		if rInit.flags&(1<<2) == (1 << 2) {
			rInit.fwd_from = m.Object()
		}
		if rInit.flags&(1<<11) == (1 << 11) {
			rInit.via_bot_id = m.Int()
		}
		if rInit.flags&(1<<3) == (1 << 3) {
			rInit.reply_to_msg_id = m.Int()
		}
		if rInit.flags&(1<<7) == (1 << 7) {
			rInit.entities = m.Vector()
		}
		r = rInit

	case crc_updateShort:
		r = TL_updateShort{
			m.Object(),
			m.Int(),
		}

	case crc_updatesCombined:
		r = TL_updatesCombined{
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_updates:
		r = TL_updates{
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Int(),
			m.Int(),
		}

	case crc_photos_photos:
		r = TL_photos_photos{
			m.Vector(),
			m.Vector(),
		}

	case crc_photos_photosSlice:
		r = TL_photos_photosSlice{
			m.Int(),
			m.Vector(),
			m.Vector(),
		}

	case crc_photos_photo:
		r = TL_photos_photo{
			m.Object(),
			m.Vector(),
		}

	case crc_upload_file:
		r = TL_upload_file{
			m.Object(),
			m.Int(),
			m.StringBytes(),
		}

	case crc_dcOption:
		rInit := TL_dcOption{}
		rInit.flags = m.Int()
		if rInit.flags&(1<<0) == (1 << 0) {
			rInit.ipv6 = true
		}
		if rInit.flags&(1<<1) == (1 << 1) {
			rInit.media_only = true
		}
		if rInit.flags&(1<<2) == (1 << 2) {
			rInit.tcpo_only = true
		}
		rInit.id = m.Int()
		rInit.ip_address = m.String()
		rInit.port = m.Int()
		r = rInit

	case crc_config:
		rInit := TL_config{}
		rInit.flags = m.Int()
		rInit.date = m.Int()
		rInit.expires = m.Int()
		rInit.test_mode = m.Object()
		rInit.this_dc = m.Int()
		rInit.dc_options = m.Vector()
		rInit.chat_size_max = m.Int()
		rInit.megagroup_size_max = m.Int()
		rInit.forwarded_count_max = m.Int()
		rInit.online_update_period_ms = m.Int()
		rInit.offline_blur_timeout_ms = m.Int()
		rInit.offline_idle_timeout_ms = m.Int()
		rInit.online_cloud_timeout_ms = m.Int()
		rInit.notify_cloud_delay_ms = m.Int()
		rInit.notify_default_delay_ms = m.Int()
		rInit.chat_big_size = m.Int()
		rInit.push_chat_period_ms = m.Int()
		rInit.push_chat_limit = m.Int()
		rInit.saved_gifs_limit = m.Int()
		rInit.edit_time_limit = m.Int()
		rInit.rating_e_decay = m.Int()
		rInit.stickers_recent_limit = m.Int()
		if rInit.flags&(1<<0) == (1 << 0) {
			rInit.tmp_sessions = m.Int()
		}
		rInit.disabled_features = m.Vector()
		r = rInit

	case crc_nearestDc:
		r = TL_nearestDc{
			m.String(),
			m.Int(),
			m.Int(),
		}

	case crc_help_appUpdate:
		r = TL_help_appUpdate{
			m.Int(),
			m.Object(),
			m.String(),
			m.String(),
		}

	case crc_help_noAppUpdate:
		r = TL_help_noAppUpdate{}

	case crc_help_inviteText:
		r = TL_help_inviteText{
			m.String(),
		}

	case crc_wallPaperSolid:
		r = TL_wallPaperSolid{
			m.Int(),
			m.String(),
			m.Int(),
			m.Int(),
		}

	case crc_updateNewEncryptedMessage:
		r = TL_updateNewEncryptedMessage{
			m.Object(),
			m.Int(),
		}

	case crc_updateEncryptedChatTyping:
		r = TL_updateEncryptedChatTyping{
			m.Int(),
		}

	case crc_updateEncryption:
		r = TL_updateEncryption{
			m.Object(),
			m.Int(),
		}

	case crc_updateEncryptedMessagesRead:
		r = TL_updateEncryptedMessagesRead{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_encryptedChatEmpty:
		r = TL_encryptedChatEmpty{
			m.Int(),
		}

	case crc_encryptedChatWaiting:
		r = TL_encryptedChatWaiting{
			m.Int(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_encryptedChatRequested:
		r = TL_encryptedChatRequested{
			m.Int(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
		}

	case crc_encryptedChat:
		r = TL_encryptedChat{
			m.Int(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
			m.Long(),
		}

	case crc_encryptedChatDiscarded:
		r = TL_encryptedChatDiscarded{
			m.Int(),
		}

	case crc_inputEncryptedChat:
		r = TL_inputEncryptedChat{
			m.Int(),
			m.Long(),
		}

	case crc_encryptedFileEmpty:
		r = TL_encryptedFileEmpty{}

	case crc_encryptedFile:
		r = TL_encryptedFile{
			m.Long(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_inputEncryptedFileEmpty:
		r = TL_inputEncryptedFileEmpty{}

	case crc_inputEncryptedFileUploaded:
		r = TL_inputEncryptedFileUploaded{
			m.Long(),
			m.Int(),
			m.String(),
			m.Int(),
		}

	case crc_inputEncryptedFile:
		r = TL_inputEncryptedFile{
			m.Long(),
			m.Long(),
		}

	case crc_inputEncryptedFileLocation:
		r = TL_inputEncryptedFileLocation{
			m.Long(),
			m.Long(),
		}

	case crc_encryptedMessage:
		r = TL_encryptedMessage{
			m.Long(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
			m.Object(),
		}

	case crc_encryptedMessageService:
		r = TL_encryptedMessageService{
			m.Long(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
		}

	case crc_messages_dhConfigNotModified:
		r = TL_messages_dhConfigNotModified{
			m.StringBytes(),
		}

	case crc_messages_dhConfig:
		r = TL_messages_dhConfig{
			m.Int(),
			m.StringBytes(),
			m.Int(),
			m.StringBytes(),
		}

	case crc_messages_sentEncryptedMessage:
		r = TL_messages_sentEncryptedMessage{
			m.Int(),
		}

	case crc_messages_sentEncryptedFile:
		r = TL_messages_sentEncryptedFile{
			m.Int(),
			m.Object(),
		}

	case crc_inputFileBig:
		r = TL_inputFileBig{
			m.Long(),
			m.Int(),
			m.String(),
		}

	case crc_inputEncryptedFileBigUploaded:
		r = TL_inputEncryptedFileBigUploaded{
			m.Long(),
			m.Int(),
			m.Int(),
		}

	case crc_updateChatParticipantAdd:
		r = TL_updateChatParticipantAdd{
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_updateChatParticipantDelete:
		r = TL_updateChatParticipantDelete{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_updateDcOptions:
		r = TL_updateDcOptions{
			m.Vector(),
		}

	case crc_inputMediaUploadedDocument:
		r = TL_inputMediaUploadedDocument{
			m.Object(),
			m.String(),
			m.Vector(),
			m.String(),
		}

	case crc_inputMediaUploadedThumbDocument:
		r = TL_inputMediaUploadedThumbDocument{
			m.Object(),
			m.Object(),
			m.String(),
			m.Vector(),
			m.String(),
		}

	case crc_inputMediaDocument:
		r = TL_inputMediaDocument{
			m.Object(),
			m.String(),
		}

	case crc_messageMediaDocument:
		r = TL_messageMediaDocument{
			m.Object(),
			m.String(),
		}

	case crc_inputDocumentEmpty:
		r = TL_inputDocumentEmpty{}

	case crc_inputDocument:
		r = TL_inputDocument{
			m.Long(),
			m.Long(),
		}

	case crc_inputDocumentFileLocation:
		r = TL_inputDocumentFileLocation{
			m.Long(),
			m.Long(),
			m.Int(),
		}

	case crc_documentEmpty:
		r = TL_documentEmpty{
			m.Long(),
		}

	case crc_document:
		r = TL_document{
			m.Long(),
			m.Long(),
			m.Int(),
			m.String(),
			m.Int(),
			m.Object(),
			m.Int(),
			m.Int(),
			m.Vector(),
		}

	case crc_help_support:
		r = TL_help_support{
			m.String(),
			m.Object(),
		}

	case crc_notifyPeer:
		r = TL_notifyPeer{
			m.Object(),
		}

	case crc_notifyUsers:
		r = TL_notifyUsers{}

	case crc_notifyChats:
		r = TL_notifyChats{}

	case crc_notifyAll:
		r = TL_notifyAll{}

	case crc_updateUserBlocked:
		r = TL_updateUserBlocked{
			m.Int(),
			m.Object(),
		}

	case crc_updateNotifySettings:
		r = TL_updateNotifySettings{
			m.Object(),
			m.Object(),
		}

	case crc_sendMessageTypingAction:
		r = TL_sendMessageTypingAction{}

	case crc_sendMessageCancelAction:
		r = TL_sendMessageCancelAction{}

	case crc_sendMessageRecordVideoAction:
		r = TL_sendMessageRecordVideoAction{}

	case crc_sendMessageUploadVideoAction:
		r = TL_sendMessageUploadVideoAction{
			m.Int(),
		}

	case crc_sendMessageRecordAudioAction:
		r = TL_sendMessageRecordAudioAction{}

	case crc_sendMessageUploadAudioAction:
		r = TL_sendMessageUploadAudioAction{
			m.Int(),
		}

	case crc_sendMessageUploadPhotoAction:
		r = TL_sendMessageUploadPhotoAction{
			m.Int(),
		}

	case crc_sendMessageUploadDocumentAction:
		r = TL_sendMessageUploadDocumentAction{
			m.Int(),
		}

	case crc_sendMessageGeoLocationAction:
		r = TL_sendMessageGeoLocationAction{}

	case crc_sendMessageChooseContactAction:
		r = TL_sendMessageChooseContactAction{}

	case crc_contacts_found:
		r = TL_contacts_found{
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case crc_updateServiceNotification:
		r = TL_updateServiceNotification{
			m.String(),
			m.String(),
			m.Object(),
			m.Object(),
		}

	case crc_userStatusRecently:
		r = TL_userStatusRecently{}

	case crc_userStatusLastWeek:
		r = TL_userStatusLastWeek{}

	case crc_userStatusLastMonth:
		r = TL_userStatusLastMonth{}

	case crc_updatePrivacy:
		r = TL_updatePrivacy{
			m.Object(),
			m.Vector(),
		}

	case crc_inputPrivacyKeyStatusTimestamp:
		r = TL_inputPrivacyKeyStatusTimestamp{}

	case crc_privacyKeyStatusTimestamp:
		r = TL_privacyKeyStatusTimestamp{}

	case crc_inputPrivacyValueAllowContacts:
		r = TL_inputPrivacyValueAllowContacts{}

	case crc_inputPrivacyValueAllowAll:
		r = TL_inputPrivacyValueAllowAll{}

	case crc_inputPrivacyValueAllowUsers:
		r = TL_inputPrivacyValueAllowUsers{
			m.Vector(),
		}

	case crc_inputPrivacyValueDisallowContacts:
		r = TL_inputPrivacyValueDisallowContacts{}

	case crc_inputPrivacyValueDisallowAll:
		r = TL_inputPrivacyValueDisallowAll{}

	case crc_inputPrivacyValueDisallowUsers:
		r = TL_inputPrivacyValueDisallowUsers{
			m.Vector(),
		}

	case crc_privacyValueAllowContacts:
		r = TL_privacyValueAllowContacts{}

	case crc_privacyValueAllowAll:
		r = TL_privacyValueAllowAll{}

	case crc_privacyValueAllowUsers:
		r = TL_privacyValueAllowUsers{
			m.VectorInt(),
		}

	case crc_privacyValueDisallowContacts:
		r = TL_privacyValueDisallowContacts{}

	case crc_privacyValueDisallowAll:
		r = TL_privacyValueDisallowAll{}

	case crc_privacyValueDisallowUsers:
		r = TL_privacyValueDisallowUsers{
			m.VectorInt(),
		}

	case crc_account_privacyRules:
		r = TL_account_privacyRules{
			m.Vector(),
			m.Vector(),
		}

	case crc_accountDaysTTL:
		r = TL_accountDaysTTL{
			m.Int(),
		}

	case crc_updateUserPhone:
		r = TL_updateUserPhone{
			m.Int(),
			m.String(),
		}

	case crc_documentAttributeImageSize:
		r = TL_documentAttributeImageSize{
			m.Int(),
			m.Int(),
		}

	case crc_documentAttributeAnimated:
		r = TL_documentAttributeAnimated{}

	case crc_documentAttributeSticker:
		r = TL_documentAttributeSticker{
			m.String(),
			m.Object(),
		}

	case crc_documentAttributeVideo:
		r = TL_documentAttributeVideo{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_documentAttributeAudio:
		rInit := TL_documentAttributeAudio{}
		rInit.flags = m.Int()
		if rInit.flags&(1<<10) == (1 << 10) {
			rInit.voice = true
		}
		rInit.duration = m.Int()
		if rInit.flags&(1<<0) == (1 << 0) {
			rInit.title = m.String()
		}
		if rInit.flags&(1<<1) == (1 << 1) {
			rInit.performer = m.String()
		}
		if rInit.flags&(1<<2) == (1 << 2) {
			rInit.waveform = m.StringBytes()
		}
		r = rInit

	case crc_documentAttributeFilename:
		r = TL_documentAttributeFilename{
			m.String(),
		}

	case crc_messages_stickersNotModified:
		r = TL_messages_stickersNotModified{}

	case crc_messages_stickers:
		r = TL_messages_stickers{
			m.String(),
			m.Vector(),
		}

	case crc_stickerPack:
		r = TL_stickerPack{
			m.String(),
			m.VectorLong(),
		}

	case crc_messages_allStickersNotModified:
		r = TL_messages_allStickersNotModified{}

	case crc_messages_allStickers:
		r = TL_messages_allStickers{
			m.Int(),
			m.Vector(),
		}

	case crc_disabledFeature:
		r = TL_disabledFeature{
			m.String(),
			m.String(),
		}

	case crc_updateReadHistoryInbox:
		r = TL_updateReadHistoryInbox{
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_updateReadHistoryOutbox:
		r = TL_updateReadHistoryOutbox{
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_messages_affectedMessages:
		r = TL_messages_affectedMessages{
			m.Int(),
			m.Int(),
		}

	case crc_contactLinkUnknown:
		r = TL_contactLinkUnknown{}

	case crc_contactLinkNone:
		r = TL_contactLinkNone{}

	case crc_contactLinkHasPhone:
		r = TL_contactLinkHasPhone{}

	case crc_contactLinkContact:
		r = TL_contactLinkContact{}

	case crc_updateWebPage:
		r = TL_updateWebPage{
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crc_webPageEmpty:
		r = TL_webPageEmpty{
			m.Long(),
		}

	case crc_webPagePending:
		r = TL_webPagePending{
			m.Long(),
			m.Int(),
		}

	case crc_webPage:
		rInit := TL_webPage{}
		rInit.flags = m.Int()
		rInit.id = m.Long()
		rInit.url = m.String()
		rInit.display_url = m.String()
		if rInit.flags&(1<<0) == (1 << 0) {
			rInit._type = m.String()
		}
		if rInit.flags&(1<<1) == (1 << 1) {
			rInit.site_name = m.String()
		}
		if rInit.flags&(1<<2) == (1 << 2) {
			rInit.title = m.String()
		}
		if rInit.flags&(1<<3) == (1 << 3) {
			rInit.description = m.String()
		}
		if rInit.flags&(1<<4) == (1 << 4) {
			rInit.photo = m.Object()
		}
		if rInit.flags&(1<<5) == (1 << 5) {
			rInit.embed_url = m.String()
		}
		if rInit.flags&(1<<5) == (1 << 5) {
			rInit.embed_type = m.String()
		}
		if rInit.flags&(1<<6) == (1 << 6) {
			rInit.embed_width = m.Int()
		}
		if rInit.flags&(1<<6) == (1 << 6) {
			rInit.embed_height = m.Int()
		}
		if rInit.flags&(1<<7) == (1 << 7) {
			rInit.duration = m.Int()
		}
		if rInit.flags&(1<<8) == (1 << 8) {
			rInit.author = m.String()
		}
		if rInit.flags&(1<<9) == (1 << 9) {
			rInit.document = m.Object()
		}
		r = rInit

	case crc_messageMediaWebPage:
		r = TL_messageMediaWebPage{
			m.Object(),
		}

	case crc_authorization:
		r = TL_authorization{
			m.Long(),
			m.Int(),
			m.String(),
			m.String(),
			m.String(),
			m.Int(),
			m.String(),
			m.String(),
			m.Int(),
			m.Int(),
			m.String(),
			m.String(),
			m.String(),
		}

	case crc_account_authorizations:
		r = TL_account_authorizations{
			m.Vector(),
		}

	case crc_account_noPassword:
		r = TL_account_noPassword{
			m.StringBytes(),
			m.String(),
		}

	case crc_account_password:
		r = TL_account_password{
			m.StringBytes(),
			m.StringBytes(),
			m.String(),
			m.Object(),
			m.String(),
		}

	case crc_account_passwordSettings:
		r = TL_account_passwordSettings{
			m.String(),
		}

	case crc_account_passwordInputSettings:
		rInit := TL_account_passwordInputSettings{}
		rInit.flags = m.Int()
		if rInit.flags&(1<<0) == (1 << 0) {
			rInit.new_salt = m.StringBytes()
		}
		if rInit.flags&(1<<0) == (1 << 0) {
			rInit.new_password_hash = m.StringBytes()
		}
		if rInit.flags&(1<<0) == (1 << 0) {
			rInit.hint = m.String()
		}
		if rInit.flags&(1<<1) == (1 << 1) {
			rInit.email = m.String()
		}
		r = rInit

	case crc_auth_passwordRecovery:
		r = TL_auth_passwordRecovery{
			m.String(),
		}

	case crc_inputMediaVenue:
		r = TL_inputMediaVenue{
			m.Object(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
		}

	case crc_messageMediaVenue:
		r = TL_messageMediaVenue{
			m.Object(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
		}

	case crc_receivedNotifyMessage:
		r = TL_receivedNotifyMessage{
			m.Int(),
			m.Int(),
		}

	case crc_chatInviteEmpty:
		r = TL_chatInviteEmpty{}

	case crc_chatInviteExported:
		r = TL_chatInviteExported{
			m.String(),
		}

	case crc_chatInviteAlready:
		r = TL_chatInviteAlready{
			m.Object(),
		}

	case crc_chatInvite:
		rInit := TL_chatInvite{}
		rInit.flags = m.Int()
		if rInit.flags&(1<<0) == (1 << 0) {
			rInit.channel = true
		}
		if rInit.flags&(1<<1) == (1 << 1) {
			rInit.broadcast = true
		}
		if rInit.flags&(1<<2) == (1 << 2) {
			rInit.public = true
		}
		if rInit.flags&(1<<3) == (1 << 3) {
			rInit.megagroup = true
		}
		rInit.title = m.String()
		rInit.photo = m.Object()
		rInit.participants_count = m.Int()
		if rInit.flags&(1<<4) == (1 << 4) {
			rInit.participants = m.Vector()
		}
		r = rInit

	case crc_messageActionChatJoinedByLink:
		r = TL_messageActionChatJoinedByLink{
			m.Int(),
		}

	case crc_updateReadMessagesContents:
		r = TL_updateReadMessagesContents{
			m.VectorInt(),
			m.Int(),
			m.Int(),
		}

	case crc_inputStickerSetEmpty:
		r = TL_inputStickerSetEmpty{}

	case crc_inputStickerSetID:
		r = TL_inputStickerSetID{
			m.Long(),
			m.Long(),
		}

	case crc_inputStickerSetShortName:
		r = TL_inputStickerSetShortName{
			m.String(),
		}

	case crc_stickerSet:
		rInit := TL_stickerSet{}
		rInit.flags = m.Int()
		if rInit.flags&(1<<0) == (1 << 0) {
			rInit.installed = true
		}
		if rInit.flags&(1<<1) == (1 << 1) {
			rInit.archived = true
		}
		if rInit.flags&(1<<2) == (1 << 2) {
			rInit.official = true
		}
		rInit.id = m.Long()
		rInit.access_hash = m.Long()
		rInit.title = m.String()
		rInit.short_name = m.String()
		rInit.count = m.Int()
		rInit.hash = m.Int()
		r = rInit

	case crc_messages_stickerSet:
		r = TL_messages_stickerSet{
			m.Object(),
			m.Vector(),
			m.Vector(),
		}

	case crc_user:
		rInit := TL_user{}
		rInit.flags = m.Int()
		if rInit.flags&(1<<10) == (1 << 10) {
			rInit.self = true
		}
		if rInit.flags&(1<<11) == (1 << 11) {
			rInit.contact = true
		}
		if rInit.flags&(1<<12) == (1 << 12) {
			rInit.mutual_contact = true
		}
		if rInit.flags&(1<<13) == (1 << 13) {
			rInit.deleted = true
		}
		if rInit.flags&(1<<14) == (1 << 14) {
			rInit.bot = true
		}
		if rInit.flags&(1<<15) == (1 << 15) {
			rInit.bot_chat_history = true
		}
		if rInit.flags&(1<<16) == (1 << 16) {
			rInit.bot_nochats = true
		}
		if rInit.flags&(1<<17) == (1 << 17) {
			rInit.verified = true
		}
		if rInit.flags&(1<<18) == (1 << 18) {
			rInit.restricted = true
		}
		if rInit.flags&(1<<20) == (1 << 20) {
			rInit.min = true
		}
		if rInit.flags&(1<<21) == (1 << 21) {
			rInit.bot_inline_geo = true
		}
		rInit.id = m.Int()
		if rInit.flags&(1<<0) == (1 << 0) {
			rInit.access_hash = m.Long()
		}
		if rInit.flags&(1<<1) == (1 << 1) {
			rInit.first_name = m.String()
		}
		if rInit.flags&(1<<2) == (1 << 2) {
			rInit.last_name = m.String()
		}
		if rInit.flags&(1<<3) == (1 << 3) {
			rInit.username = m.String()
		}
		if rInit.flags&(1<<4) == (1 << 4) {
			rInit.phone = m.String()
		}
		if rInit.flags&(1<<5) == (1 << 5) {
			rInit.photo = m.Object()
		}
		if rInit.flags&(1<<6) == (1 << 6) {
			rInit.status = m.Object()
		}
		if rInit.flags&(1<<14) == (1 << 14) {
			rInit.bot_info_version = m.Int()
		}
		if rInit.flags&(1<<18) == (1 << 18) {
			rInit.restriction_reason = m.String()
		}
		if rInit.flags&(1<<19) == (1 << 19) {
			rInit.bot_inline_placeholder = m.String()
		}
		r = rInit

	case crc_botCommand:
		r = TL_botCommand{
			m.String(),
			m.String(),
		}

	case crc_botInfo:
		r = TL_botInfo{
			m.Int(),
			m.String(),
			m.Vector(),
		}

	case crc_keyboardButton:
		r = TL_keyboardButton{
			m.String(),
		}

	case crc_keyboardButtonRow:
		r = TL_keyboardButtonRow{
			m.Vector(),
		}

	case crc_replyKeyboardHide:
		rInit := TL_replyKeyboardHide{}
		rInit.flags = m.Int()
		if rInit.flags&(1<<2) == (1 << 2) {
			rInit.selective = true
		}
		r = rInit

	case crc_replyKeyboardForceReply:
		rInit := TL_replyKeyboardForceReply{}
		rInit.flags = m.Int()
		if rInit.flags&(1<<1) == (1 << 1) {
			rInit.single_use = true
		}
		if rInit.flags&(1<<2) == (1 << 2) {
			rInit.selective = true
		}
		r = rInit

	case crc_replyKeyboardMarkup:
		rInit := TL_replyKeyboardMarkup{}
		rInit.flags = m.Int()
		if rInit.flags&(1<<0) == (1 << 0) {
			rInit.resize = true
		}
		if rInit.flags&(1<<1) == (1 << 1) {
			rInit.single_use = true
		}
		if rInit.flags&(1<<2) == (1 << 2) {
			rInit.selective = true
		}
		rInit.rows = m.Vector()
		r = rInit

	case crc_inputPeerUser:
		r = TL_inputPeerUser{
			m.Int(),
			m.Long(),
		}

	case crc_inputUser:
		r = TL_inputUser{
			m.Int(),
			m.Long(),
		}

	case crc_help_appChangelogEmpty:
		r = TL_help_appChangelogEmpty{}

	case crc_help_appChangelog:
		r = TL_help_appChangelog{
			m.String(),
		}

	case crc_messageEntityUnknown:
		r = TL_messageEntityUnknown{
			m.Int(),
			m.Int(),
		}

	case crc_messageEntityMention:
		r = TL_messageEntityMention{
			m.Int(),
			m.Int(),
		}

	case crc_messageEntityHashtag:
		r = TL_messageEntityHashtag{
			m.Int(),
			m.Int(),
		}

	case crc_messageEntityBotCommand:
		r = TL_messageEntityBotCommand{
			m.Int(),
			m.Int(),
		}

	case crc_messageEntityUrl:
		r = TL_messageEntityUrl{
			m.Int(),
			m.Int(),
		}

	case crc_messageEntityEmail:
		r = TL_messageEntityEmail{
			m.Int(),
			m.Int(),
		}

	case crc_messageEntityBold:
		r = TL_messageEntityBold{
			m.Int(),
			m.Int(),
		}

	case crc_messageEntityItalic:
		r = TL_messageEntityItalic{
			m.Int(),
			m.Int(),
		}

	case crc_messageEntityCode:
		r = TL_messageEntityCode{
			m.Int(),
			m.Int(),
		}

	case crc_messageEntityPre:
		r = TL_messageEntityPre{
			m.Int(),
			m.Int(),
			m.String(),
		}

	case crc_messageEntityTextUrl:
		r = TL_messageEntityTextUrl{
			m.Int(),
			m.Int(),
			m.String(),
		}

	case crc_updateShortSentMessage:
		rInit := TL_updateShortSentMessage{}
		rInit.flags = m.Int()
		if rInit.flags&(1<<1) == (1 << 1) {
			rInit.out = true
		}
		rInit.id = m.Int()
		rInit.pts = m.Int()
		rInit.pts_count = m.Int()
		rInit.date = m.Int()
		if rInit.flags&(1<<9) == (1 << 9) {
			rInit.media = m.Object()
		}
		if rInit.flags&(1<<7) == (1 << 7) {
			rInit.entities = m.Vector()
		}
		r = rInit

	case crc_inputChannelEmpty:
		r = TL_inputChannelEmpty{}

	case crc_inputChannel:
		r = TL_inputChannel{
			m.Int(),
			m.Long(),
		}

	case crc_peerChannel:
		r = TL_peerChannel{
			m.Int(),
		}

	case crc_inputPeerChannel:
		r = TL_inputPeerChannel{
			m.Int(),
			m.Long(),
		}

	case crc_channel:
		rInit := TL_channel{}
		rInit.flags = m.Int()
		if rInit.flags&(1<<0) == (1 << 0) {
			rInit.creator = true
		}
		if rInit.flags&(1<<1) == (1 << 1) {
			rInit.kicked = true
		}
		if rInit.flags&(1<<2) == (1 << 2) {
			rInit.left = true
		}
		if rInit.flags&(1<<3) == (1 << 3) {
			rInit.editor = true
		}
		if rInit.flags&(1<<4) == (1 << 4) {
			rInit.moderator = true
		}
		if rInit.flags&(1<<5) == (1 << 5) {
			rInit.broadcast = true
		}
		if rInit.flags&(1<<7) == (1 << 7) {
			rInit.verified = true
		}
		if rInit.flags&(1<<8) == (1 << 8) {
			rInit.megagroup = true
		}
		if rInit.flags&(1<<9) == (1 << 9) {
			rInit.restricted = true
		}
		if rInit.flags&(1<<10) == (1 << 10) {
			rInit.democracy = true
		}
		if rInit.flags&(1<<11) == (1 << 11) {
			rInit.signatures = true
		}
		if rInit.flags&(1<<12) == (1 << 12) {
			rInit.min = true
		}
		rInit.id = m.Int()
		if rInit.flags&(1<<13) == (1 << 13) {
			rInit.access_hash = m.Long()
		}
		rInit.title = m.String()
		if rInit.flags&(1<<6) == (1 << 6) {
			rInit.username = m.String()
		}
		rInit.photo = m.Object()
		rInit.date = m.Int()
		rInit.version = m.Int()
		if rInit.flags&(1<<9) == (1 << 9) {
			rInit.restriction_reason = m.String()
		}
		r = rInit

	case crc_channelForbidden:
		rInit := TL_channelForbidden{}
		rInit.flags = m.Int()
		if rInit.flags&(1<<5) == (1 << 5) {
			rInit.broadcast = true
		}
		if rInit.flags&(1<<8) == (1 << 8) {
			rInit.megagroup = true
		}
		rInit.id = m.Int()
		rInit.access_hash = m.Long()
		rInit.title = m.String()
		r = rInit

	case crc_contacts_resolvedPeer:
		r = TL_contacts_resolvedPeer{
			m.Object(),
			m.Vector(),
			m.Vector(),
		}

	case crc_channelFull:
		rInit := TL_channelFull{}
		rInit.flags = m.Int()
		if rInit.flags&(1<<3) == (1 << 3) {
			rInit.can_view_participants = true
		}
		if rInit.flags&(1<<6) == (1 << 6) {
			rInit.can_set_username = true
		}
		rInit.id = m.Int()
		rInit.about = m.String()
		if rInit.flags&(1<<0) == (1 << 0) {
			rInit.participants_count = m.Int()
		}
		if rInit.flags&(1<<1) == (1 << 1) {
			rInit.admins_count = m.Int()
		}
		if rInit.flags&(1<<2) == (1 << 2) {
			rInit.kicked_count = m.Int()
		}
		rInit.read_inbox_max_id = m.Int()
		rInit.read_outbox_max_id = m.Int()
		rInit.unread_count = m.Int()
		rInit.chat_photo = m.Object()
		rInit.notify_settings = m.Object()
		rInit.exported_invite = m.Object()
		rInit.bot_info = m.Vector()
		if rInit.flags&(1<<4) == (1 << 4) {
			rInit.migrated_from_chat_id = m.Int()
		}
		if rInit.flags&(1<<4) == (1 << 4) {
			rInit.migrated_from_max_id = m.Int()
		}
		if rInit.flags&(1<<5) == (1 << 5) {
			rInit.pinned_msg_id = m.Int()
		}
		r = rInit

	case crc_messageRange:
		r = TL_messageRange{
			m.Int(),
			m.Int(),
		}

	case crc_messages_channelMessages:
		rInit := TL_messages_channelMessages{}
		rInit.flags = m.Int()
		rInit.pts = m.Int()
		rInit.count = m.Int()
		rInit.messages = m.Vector()
		rInit.chats = m.Vector()
		rInit.users = m.Vector()
		r = rInit

	case crc_messageActionChannelCreate:
		r = TL_messageActionChannelCreate{
			m.String(),
		}

	case crc_updateChannelTooLong:
		rInit := TL_updateChannelTooLong{}
		rInit.flags = m.Int()
		rInit.channel_id = m.Int()
		if rInit.flags&(1<<0) == (1 << 0) {
			rInit.pts = m.Int()
		}
		r = rInit

	case crc_updateChannel:
		r = TL_updateChannel{
			m.Int(),
		}

	case crc_updateNewChannelMessage:
		r = TL_updateNewChannelMessage{
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crc_updateReadChannelInbox:
		r = TL_updateReadChannelInbox{
			m.Int(),
			m.Int(),
		}

	case crc_updateDeleteChannelMessages:
		r = TL_updateDeleteChannelMessages{
			m.Int(),
			m.VectorInt(),
			m.Int(),
			m.Int(),
		}

	case crc_updateChannelMessageViews:
		r = TL_updateChannelMessageViews{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_updates_channelDifferenceEmpty:
		rInit := TL_updates_channelDifferenceEmpty{}
		rInit.flags = m.Int()
		if rInit.flags&(1<<0) == (1 << 0) {
			rInit.final = true
		}
		rInit.pts = m.Int()
		if rInit.flags&(1<<1) == (1 << 1) {
			rInit.timeout = m.Int()
		}
		r = rInit

	case crc_updates_channelDifferenceTooLong:
		rInit := TL_updates_channelDifferenceTooLong{}
		rInit.flags = m.Int()
		if rInit.flags&(1<<0) == (1 << 0) {
			rInit.final = true
		}
		rInit.pts = m.Int()
		if rInit.flags&(1<<1) == (1 << 1) {
			rInit.timeout = m.Int()
		}
		rInit.top_message = m.Int()
		rInit.read_inbox_max_id = m.Int()
		rInit.read_outbox_max_id = m.Int()
		rInit.unread_count = m.Int()
		rInit.messages = m.Vector()
		rInit.chats = m.Vector()
		rInit.users = m.Vector()
		r = rInit

	case crc_updates_channelDifference:
		rInit := TL_updates_channelDifference{}
		rInit.flags = m.Int()
		if rInit.flags&(1<<0) == (1 << 0) {
			rInit.final = true
		}
		rInit.pts = m.Int()
		if rInit.flags&(1<<1) == (1 << 1) {
			rInit.timeout = m.Int()
		}
		rInit.new_messages = m.Vector()
		rInit.other_updates = m.Vector()
		rInit.chats = m.Vector()
		rInit.users = m.Vector()
		r = rInit

	case crc_channelMessagesFilterEmpty:
		r = TL_channelMessagesFilterEmpty{}

	case crc_channelMessagesFilter:
		rInit := TL_channelMessagesFilter{}
		rInit.flags = m.Int()
		if rInit.flags&(1<<1) == (1 << 1) {
			rInit.exclude_new_messages = true
		}
		rInit.ranges = m.Vector()
		r = rInit

	case crc_channelParticipant:
		r = TL_channelParticipant{
			m.Int(),
			m.Int(),
		}

	case crc_channelParticipantSelf:
		r = TL_channelParticipantSelf{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_channelParticipantModerator:
		r = TL_channelParticipantModerator{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_channelParticipantEditor:
		r = TL_channelParticipantEditor{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_channelParticipantKicked:
		r = TL_channelParticipantKicked{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_channelParticipantCreator:
		r = TL_channelParticipantCreator{
			m.Int(),
		}

	case crc_channelParticipantsRecent:
		r = TL_channelParticipantsRecent{}

	case crc_channelParticipantsAdmins:
		r = TL_channelParticipantsAdmins{}

	case crc_channelParticipantsKicked:
		r = TL_channelParticipantsKicked{}

	case crc_channelRoleEmpty:
		r = TL_channelRoleEmpty{}

	case crc_channelRoleModerator:
		r = TL_channelRoleModerator{}

	case crc_channelRoleEditor:
		r = TL_channelRoleEditor{}

	case crc_channels_channelParticipants:
		r = TL_channels_channelParticipants{
			m.Int(),
			m.Vector(),
			m.Vector(),
		}

	case crc_channels_channelParticipant:
		r = TL_channels_channelParticipant{
			m.Object(),
			m.Vector(),
		}

	case crc_chatParticipantCreator:
		r = TL_chatParticipantCreator{
			m.Int(),
		}

	case crc_chatParticipantAdmin:
		r = TL_chatParticipantAdmin{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_updateChatAdmins:
		r = TL_updateChatAdmins{
			m.Int(),
			m.Object(),
			m.Int(),
		}

	case crc_updateChatParticipantAdmin:
		r = TL_updateChatParticipantAdmin{
			m.Int(),
			m.Int(),
			m.Object(),
			m.Int(),
		}

	case crc_messageActionChatMigrateTo:
		r = TL_messageActionChatMigrateTo{
			m.Int(),
		}

	case crc_messageActionChannelMigrateFrom:
		r = TL_messageActionChannelMigrateFrom{
			m.String(),
			m.Int(),
		}

	case crc_channelParticipantsBots:
		r = TL_channelParticipantsBots{}

	case crc_help_termsOfService:
		r = TL_help_termsOfService{
			m.String(),
		}

	case crc_updateNewStickerSet:
		r = TL_updateNewStickerSet{
			m.Object(),
		}

	case crc_updateStickerSetsOrder:
		r = TL_updateStickerSetsOrder{
			m.VectorLong(),
		}

	case crc_updateStickerSets:
		r = TL_updateStickerSets{}

	case crc_foundGif:
		r = TL_foundGif{
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.Int(),
			m.Int(),
		}

	case crc_foundGifCached:
		r = TL_foundGifCached{
			m.String(),
			m.Object(),
			m.Object(),
		}

	case crc_inputMediaGifExternal:
		r = TL_inputMediaGifExternal{
			m.String(),
			m.String(),
		}

	case crc_messages_foundGifs:
		r = TL_messages_foundGifs{
			m.Int(),
			m.Vector(),
		}

	case crc_messages_savedGifsNotModified:
		r = TL_messages_savedGifsNotModified{}

	case crc_messages_savedGifs:
		r = TL_messages_savedGifs{
			m.Int(),
			m.Vector(),
		}

	case crc_updateSavedGifs:
		r = TL_updateSavedGifs{}

	case crc_inputBotInlineMessageMediaAuto:
		rInit := TL_inputBotInlineMessageMediaAuto{}
		rInit.flags = m.Int()
		rInit.caption = m.String()
		if rInit.flags&(1<<2) == (1 << 2) {
			rInit.reply_markup = m.Object()
		}
		r = rInit

	case crc_inputBotInlineMessageText:
		rInit := TL_inputBotInlineMessageText{}
		rInit.flags = m.Int()
		if rInit.flags&(1<<0) == (1 << 0) {
			rInit.no_webpage = true
		}
		rInit.message = m.String()
		if rInit.flags&(1<<1) == (1 << 1) {
			rInit.entities = m.Vector()
		}
		if rInit.flags&(1<<2) == (1 << 2) {
			rInit.reply_markup = m.Object()
		}
		r = rInit

	case crc_inputBotInlineResult:
		rInit := TL_inputBotInlineResult{}
		rInit.flags = m.Int()
		rInit.id = m.String()
		rInit._type = m.String()
		if rInit.flags&(1<<1) == (1 << 1) {
			rInit.title = m.String()
		}
		if rInit.flags&(1<<2) == (1 << 2) {
			rInit.description = m.String()
		}
		if rInit.flags&(1<<3) == (1 << 3) {
			rInit.url = m.String()
		}
		if rInit.flags&(1<<4) == (1 << 4) {
			rInit.thumb_url = m.String()
		}
		if rInit.flags&(1<<5) == (1 << 5) {
			rInit.content_url = m.String()
		}
		if rInit.flags&(1<<5) == (1 << 5) {
			rInit.content_type = m.String()
		}
		if rInit.flags&(1<<6) == (1 << 6) {
			rInit.w = m.Int()
		}
		if rInit.flags&(1<<6) == (1 << 6) {
			rInit.h = m.Int()
		}
		if rInit.flags&(1<<7) == (1 << 7) {
			rInit.duration = m.Int()
		}
		rInit.send_message = m.Object()
		r = rInit

	case crc_botInlineMessageMediaAuto:
		rInit := TL_botInlineMessageMediaAuto{}
		rInit.flags = m.Int()
		rInit.caption = m.String()
		if rInit.flags&(1<<2) == (1 << 2) {
			rInit.reply_markup = m.Object()
		}
		r = rInit

	case crc_botInlineMessageText:
		rInit := TL_botInlineMessageText{}
		rInit.flags = m.Int()
		if rInit.flags&(1<<0) == (1 << 0) {
			rInit.no_webpage = true
		}
		rInit.message = m.String()
		if rInit.flags&(1<<1) == (1 << 1) {
			rInit.entities = m.Vector()
		}
		if rInit.flags&(1<<2) == (1 << 2) {
			rInit.reply_markup = m.Object()
		}
		r = rInit

	case crc_botInlineResult:
		rInit := TL_botInlineResult{}
		rInit.flags = m.Int()
		rInit.id = m.String()
		rInit._type = m.String()
		if rInit.flags&(1<<1) == (1 << 1) {
			rInit.title = m.String()
		}
		if rInit.flags&(1<<2) == (1 << 2) {
			rInit.description = m.String()
		}
		if rInit.flags&(1<<3) == (1 << 3) {
			rInit.url = m.String()
		}
		if rInit.flags&(1<<4) == (1 << 4) {
			rInit.thumb_url = m.String()
		}
		if rInit.flags&(1<<5) == (1 << 5) {
			rInit.content_url = m.String()
		}
		if rInit.flags&(1<<5) == (1 << 5) {
			rInit.content_type = m.String()
		}
		if rInit.flags&(1<<6) == (1 << 6) {
			rInit.w = m.Int()
		}
		if rInit.flags&(1<<6) == (1 << 6) {
			rInit.h = m.Int()
		}
		if rInit.flags&(1<<7) == (1 << 7) {
			rInit.duration = m.Int()
		}
		rInit.send_message = m.Object()
		r = rInit

	case crc_messages_botResults:
		rInit := TL_messages_botResults{}
		rInit.flags = m.Int()
		if rInit.flags&(1<<0) == (1 << 0) {
			rInit.gallery = true
		}
		rInit.query_id = m.Long()
		if rInit.flags&(1<<1) == (1 << 1) {
			rInit.next_offset = m.String()
		}
		if rInit.flags&(1<<2) == (1 << 2) {
			rInit.switch_pm = m.Object()
		}
		rInit.results = m.Vector()
		r = rInit

	case crc_updateBotInlineQuery:
		rInit := TL_updateBotInlineQuery{}
		rInit.flags = m.Int()
		rInit.query_id = m.Long()
		rInit.user_id = m.Int()
		rInit.query = m.String()
		if rInit.flags&(1<<0) == (1 << 0) {
			rInit.geo = m.Object()
		}
		rInit.offset = m.String()
		r = rInit

	case crc_updateBotInlineSend:
		rInit := TL_updateBotInlineSend{}
		rInit.flags = m.Int()
		rInit.user_id = m.Int()
		rInit.query = m.String()
		if rInit.flags&(1<<0) == (1 << 0) {
			rInit.geo = m.Object()
		}
		rInit.id = m.String()
		if rInit.flags&(1<<1) == (1 << 1) {
			rInit.msg_id = m.Object()
		}
		r = rInit

	case crc_inputMessagesFilterVoice:
		r = TL_inputMessagesFilterVoice{}

	case crc_inputMessagesFilterMusic:
		r = TL_inputMessagesFilterMusic{}

	case crc_inputPrivacyKeyChatInvite:
		r = TL_inputPrivacyKeyChatInvite{}

	case crc_privacyKeyChatInvite:
		r = TL_privacyKeyChatInvite{}

	case crc_exportedMessageLink:
		r = TL_exportedMessageLink{
			m.String(),
		}

	case crc_messageFwdHeader:
		rInit := TL_messageFwdHeader{}
		rInit.flags = m.Int()
		if rInit.flags&(1<<0) == (1 << 0) {
			rInit.from_id = m.Int()
		}
		rInit.date = m.Int()
		if rInit.flags&(1<<1) == (1 << 1) {
			rInit.channel_id = m.Int()
		}
		if rInit.flags&(1<<2) == (1 << 2) {
			rInit.channel_post = m.Int()
		}
		r = rInit

	case crc_updateEditChannelMessage:
		r = TL_updateEditChannelMessage{
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crc_updateChannelPinnedMessage:
		r = TL_updateChannelPinnedMessage{
			m.Int(),
			m.Int(),
		}

	case crc_messageActionPinMessage:
		r = TL_messageActionPinMessage{}

	case crc_auth_codeTypeSms:
		r = TL_auth_codeTypeSms{}

	case crc_auth_codeTypeCall:
		r = TL_auth_codeTypeCall{}

	case crc_auth_codeTypeFlashCall:
		r = TL_auth_codeTypeFlashCall{}

	case crc_auth_sentCodeTypeApp:
		r = TL_auth_sentCodeTypeApp{
			m.Int(),
		}

	case crc_auth_sentCodeTypeSms:
		r = TL_auth_sentCodeTypeSms{
			m.Int(),
		}

	case crc_auth_sentCodeTypeCall:
		r = TL_auth_sentCodeTypeCall{
			m.Int(),
		}

	case crc_auth_sentCodeTypeFlashCall:
		r = TL_auth_sentCodeTypeFlashCall{
			m.String(),
		}

	case crc_keyboardButtonUrl:
		r = TL_keyboardButtonUrl{
			m.String(),
			m.String(),
		}

	case crc_keyboardButtonCallback:
		r = TL_keyboardButtonCallback{
			m.String(),
			m.StringBytes(),
		}

	case crc_keyboardButtonRequestPhone:
		r = TL_keyboardButtonRequestPhone{
			m.String(),
		}

	case crc_keyboardButtonRequestGeoLocation:
		r = TL_keyboardButtonRequestGeoLocation{
			m.String(),
		}

	case crc_keyboardButtonSwitchInline:
		r = TL_keyboardButtonSwitchInline{
			m.String(),
			m.String(),
		}

	case crc_replyInlineMarkup:
		r = TL_replyInlineMarkup{
			m.Vector(),
		}

	case crc_messages_botCallbackAnswer:
		rInit := TL_messages_botCallbackAnswer{}
		rInit.flags = m.Int()
		if rInit.flags&(1<<1) == (1 << 1) {
			rInit.alert = true
		}
		if rInit.flags&(1<<0) == (1 << 0) {
			rInit.message = m.String()
		}
		if rInit.flags&(1<<2) == (1 << 2) {
			rInit.url = m.String()
		}
		r = rInit

	case crc_updateBotCallbackQuery:
		r = TL_updateBotCallbackQuery{
			m.Long(),
			m.Int(),
			m.Object(),
			m.Int(),
			m.StringBytes(),
		}

	case crc_messages_messageEditData:
		rInit := TL_messages_messageEditData{}
		rInit.flags = m.Int()
		if rInit.flags&(1<<0) == (1 << 0) {
			rInit.caption = true
		}
		r = rInit

	case crc_updateEditMessage:
		r = TL_updateEditMessage{
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crc_inputBotInlineMessageMediaGeo:
		rInit := TL_inputBotInlineMessageMediaGeo{}
		rInit.flags = m.Int()
		rInit.geo_point = m.Object()
		if rInit.flags&(1<<2) == (1 << 2) {
			rInit.reply_markup = m.Object()
		}
		r = rInit

	case crc_inputBotInlineMessageMediaVenue:
		rInit := TL_inputBotInlineMessageMediaVenue{}
		rInit.flags = m.Int()
		rInit.geo_point = m.Object()
		rInit.title = m.String()
		rInit.address = m.String()
		rInit.provider = m.String()
		rInit.venue_id = m.String()
		if rInit.flags&(1<<2) == (1 << 2) {
			rInit.reply_markup = m.Object()
		}
		r = rInit

	case crc_inputBotInlineMessageMediaContact:
		rInit := TL_inputBotInlineMessageMediaContact{}
		rInit.flags = m.Int()
		rInit.phone_number = m.String()
		rInit.first_name = m.String()
		rInit.last_name = m.String()
		if rInit.flags&(1<<2) == (1 << 2) {
			rInit.reply_markup = m.Object()
		}
		r = rInit

	case crc_botInlineMessageMediaGeo:
		rInit := TL_botInlineMessageMediaGeo{}
		rInit.flags = m.Int()
		rInit.geo = m.Object()
		if rInit.flags&(1<<2) == (1 << 2) {
			rInit.reply_markup = m.Object()
		}
		r = rInit

	case crc_botInlineMessageMediaVenue:
		rInit := TL_botInlineMessageMediaVenue{}
		rInit.flags = m.Int()
		rInit.geo = m.Object()
		rInit.title = m.String()
		rInit.address = m.String()
		rInit.provider = m.String()
		rInit.venue_id = m.String()
		if rInit.flags&(1<<2) == (1 << 2) {
			rInit.reply_markup = m.Object()
		}
		r = rInit

	case crc_botInlineMessageMediaContact:
		rInit := TL_botInlineMessageMediaContact{}
		rInit.flags = m.Int()
		rInit.phone_number = m.String()
		rInit.first_name = m.String()
		rInit.last_name = m.String()
		if rInit.flags&(1<<2) == (1 << 2) {
			rInit.reply_markup = m.Object()
		}
		r = rInit

	case crc_inputBotInlineResultPhoto:
		r = TL_inputBotInlineResultPhoto{
			m.String(),
			m.String(),
			m.Object(),
			m.Object(),
		}

	case crc_inputBotInlineResultDocument:
		rInit := TL_inputBotInlineResultDocument{}
		rInit.flags = m.Int()
		rInit.id = m.String()
		rInit._type = m.String()
		if rInit.flags&(1<<1) == (1 << 1) {
			rInit.title = m.String()
		}
		if rInit.flags&(1<<2) == (1 << 2) {
			rInit.description = m.String()
		}
		rInit.document = m.Object()
		rInit.send_message = m.Object()
		r = rInit

	case crc_botInlineMediaResult:
		rInit := TL_botInlineMediaResult{}
		rInit.flags = m.Int()
		rInit.id = m.String()
		rInit._type = m.String()
		if rInit.flags&(1<<0) == (1 << 0) {
			rInit.photo = m.Object()
		}
		if rInit.flags&(1<<1) == (1 << 1) {
			rInit.document = m.Object()
		}
		if rInit.flags&(1<<2) == (1 << 2) {
			rInit.title = m.String()
		}
		if rInit.flags&(1<<3) == (1 << 3) {
			rInit.description = m.String()
		}
		rInit.send_message = m.Object()
		r = rInit

	case crc_inputBotInlineMessageID:
		r = TL_inputBotInlineMessageID{
			m.Int(),
			m.Long(),
			m.Long(),
		}

	case crc_updateInlineBotCallbackQuery:
		r = TL_updateInlineBotCallbackQuery{
			m.Long(),
			m.Int(),
			m.Object(),
			m.StringBytes(),
		}

	case crc_inlineBotSwitchPM:
		r = TL_inlineBotSwitchPM{
			m.String(),
			m.String(),
		}

	case crc_messages_peerDialogs:
		r = TL_messages_peerDialogs{
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Object(),
		}

	case crc_topPeer:
		r = TL_topPeer{
			m.Object(),
			m.Double(),
		}

	case crc_topPeerCategoryBotsPM:
		r = TL_topPeerCategoryBotsPM{}

	case crc_topPeerCategoryBotsInline:
		r = TL_topPeerCategoryBotsInline{}

	case crc_topPeerCategoryCorrespondents:
		r = TL_topPeerCategoryCorrespondents{}

	case crc_topPeerCategoryGroups:
		r = TL_topPeerCategoryGroups{}

	case crc_topPeerCategoryChannels:
		r = TL_topPeerCategoryChannels{}

	case crc_topPeerCategoryPeers:
		r = TL_topPeerCategoryPeers{
			m.Object(),
			m.Int(),
			m.Vector(),
		}

	case crc_contacts_topPeersNotModified:
		r = TL_contacts_topPeersNotModified{}

	case crc_contacts_topPeers:
		r = TL_contacts_topPeers{
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case crc_messageEntityMentionName:
		r = TL_messageEntityMentionName{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_inputMessageEntityMentionName:
		r = TL_inputMessageEntityMentionName{
			m.Int(),
			m.Int(),
			m.Object(),
		}

	case crc_inputMessagesFilterChatPhotos:
		r = TL_inputMessagesFilterChatPhotos{}

	case crc_updateReadChannelOutbox:
		r = TL_updateReadChannelOutbox{
			m.Int(),
			m.Int(),
		}

	case crc_updateDraftMessage:
		r = TL_updateDraftMessage{
			m.Object(),
			m.Object(),
		}

	case crc_draftMessageEmpty:
		r = TL_draftMessageEmpty{}

	case crc_draftMessage:
		rInit := TL_draftMessage{}
		rInit.flags = m.Int()
		if rInit.flags&(1<<1) == (1 << 1) {
			rInit.no_webpage = true
		}
		if rInit.flags&(1<<0) == (1 << 0) {
			rInit.reply_to_msg_id = m.Int()
		}
		rInit.message = m.String()
		if rInit.flags&(1<<3) == (1 << 3) {
			rInit.entities = m.Vector()
		}
		rInit.date = m.Int()
		r = rInit

	case crc_messageActionHistoryClear:
		r = TL_messageActionHistoryClear{}

	case crc_messages_featuredStickersNotModified:
		r = TL_messages_featuredStickersNotModified{}

	case crc_messages_featuredStickers:
		r = TL_messages_featuredStickers{
			m.Int(),
			m.Vector(),
			m.VectorLong(),
		}

	case crc_updateReadFeaturedStickers:
		r = TL_updateReadFeaturedStickers{}

	case crc_messages_recentStickersNotModified:
		r = TL_messages_recentStickersNotModified{}

	case crc_messages_recentStickers:
		r = TL_messages_recentStickers{
			m.Int(),
			m.Vector(),
		}

	case crc_updateRecentStickers:
		r = TL_updateRecentStickers{}

	case crc_messages_archivedStickers:
		r = TL_messages_archivedStickers{
			m.Int(),
			m.Vector(),
		}

	case crc_messages_stickerSetInstallResultSuccess:
		r = TL_messages_stickerSetInstallResultSuccess{}

	case crc_messages_stickerSetInstallResultArchive:
		r = TL_messages_stickerSetInstallResultArchive{
			m.Vector(),
		}

	case crc_stickerSetCovered:
		r = TL_stickerSetCovered{
			m.Object(),
			m.Object(),
		}

	case crc_invokeAfterMsg:
		r = TL_invokeAfterMsg{
			m.Long(),
			m.Object(),
		}

	case crc_invokeAfterMsgs:
		r = TL_invokeAfterMsgs{
			m.VectorLong(),
			m.Object(),
		}

	case crc_auth_checkPhone:
		r = TL_auth_checkPhone{
			m.String(),
		}

	case crc_auth_sendCode:
		rInit := TL_auth_sendCode{}
		rInit.flags = m.Int()
		if rInit.flags&(1<<0) == (1 << 0) {
			rInit.allow_flashcall = true
		}
		rInit.phone_number = m.String()
		if rInit.flags&(1<<0) == (1 << 0) {
			rInit.current_number = m.Object()
		}
		rInit.api_id = m.Int()
		rInit.api_hash = m.String()
		r = rInit

	case crc_auth_signUp:
		r = TL_auth_signUp{
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
		}

	case crc_auth_signIn:
		r = TL_auth_signIn{
			m.String(),
			m.String(),
			m.String(),
		}

	case crc_auth_logOut:
		r = TL_auth_logOut{}

	case crc_auth_resetAuthorizations:
		r = TL_auth_resetAuthorizations{}

	case crc_auth_sendInvites:
		r = TL_auth_sendInvites{
			m.VectorString(),
			m.String(),
		}

	case crc_auth_exportAuthorization:
		r = TL_auth_exportAuthorization{
			m.Int(),
		}

	case crc_auth_importAuthorization:
		r = TL_auth_importAuthorization{
			m.Int(),
			m.StringBytes(),
		}

	case crc_auth_bindTempAuthKey:
		r = TL_auth_bindTempAuthKey{
			m.Long(),
			m.Long(),
			m.Int(),
			m.StringBytes(),
		}

	case crc_account_registerDevice:
		r = TL_account_registerDevice{
			m.Int(),
			m.String(),
		}

	case crc_account_unregisterDevice:
		r = TL_account_unregisterDevice{
			m.Int(),
			m.String(),
		}

	case crc_account_updateNotifySettings:
		r = TL_account_updateNotifySettings{
			m.Object(),
			m.Object(),
		}

	case crc_account_getNotifySettings:
		r = TL_account_getNotifySettings{
			m.Object(),
		}

	case crc_account_resetNotifySettings:
		r = TL_account_resetNotifySettings{}

	case crc_account_updateProfile:
		rInit := TL_account_updateProfile{}
		rInit.flags = m.Int()
		if rInit.flags&(1<<0) == (1 << 0) {
			rInit.first_name = m.String()
		}
		if rInit.flags&(1<<1) == (1 << 1) {
			rInit.last_name = m.String()
		}
		if rInit.flags&(1<<2) == (1 << 2) {
			rInit.about = m.String()
		}
		r = rInit

	case crc_account_updateStatus:
		r = TL_account_updateStatus{
			m.Object(),
		}

	case crc_account_getWallPapers:
		r = TL_account_getWallPapers{}

	case crc_account_reportPeer:
		r = TL_account_reportPeer{
			m.Object(),
			m.Object(),
		}

	case crc_users_getUsers:
		r = TL_users_getUsers{
			m.Vector(),
		}

	case crc_users_getFullUser:
		r = TL_users_getFullUser{
			m.Object(),
		}

	case crc_contacts_getStatuses:
		r = TL_contacts_getStatuses{}

	case crc_contacts_getContacts:
		r = TL_contacts_getContacts{
			m.String(),
		}

	case crc_contacts_importContacts:
		r = TL_contacts_importContacts{
			m.Vector(),
			m.Object(),
		}

	case crc_contacts_deleteContact:
		r = TL_contacts_deleteContact{
			m.Object(),
		}

	case crc_contacts_deleteContacts:
		r = TL_contacts_deleteContacts{
			m.Vector(),
		}

	case crc_contacts_block:
		r = TL_contacts_block{
			m.Object(),
		}

	case crc_contacts_unblock:
		r = TL_contacts_unblock{
			m.Object(),
		}

	case crc_contacts_getBlocked:
		r = TL_contacts_getBlocked{
			m.Int(),
			m.Int(),
		}

	case crc_contacts_exportCard:
		r = TL_contacts_exportCard{}

	case crc_contacts_importCard:
		r = TL_contacts_importCard{
			m.VectorInt(),
		}

	case crc_messages_getMessages:
		r = TL_messages_getMessages{
			m.VectorInt(),
		}

	case crc_messages_getDialogs:
		r = TL_messages_getDialogs{
			m.Int(),
			m.Int(),
			m.Object(),
			m.Int(),
		}

	case crc_messages_getHistory:
		r = TL_messages_getHistory{
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_messages_search:
		rInit := TL_messages_search{}
		rInit.flags = m.Int()
		rInit.peer = m.Object()
		rInit.q = m.String()
		rInit.filter = m.Object()
		rInit.min_date = m.Int()
		rInit.max_date = m.Int()
		rInit.offset = m.Int()
		rInit.max_id = m.Int()
		rInit.limit = m.Int()
		r = rInit

	case crc_messages_readHistory:
		r = TL_messages_readHistory{
			m.Object(),
			m.Int(),
		}

	case crc_messages_deleteHistory:
		rInit := TL_messages_deleteHistory{}
		rInit.flags = m.Int()
		if rInit.flags&(1<<0) == (1 << 0) {
			rInit.just_clear = true
		}
		rInit.peer = m.Object()
		rInit.max_id = m.Int()
		r = rInit

	case crc_messages_deleteMessages:
		r = TL_messages_deleteMessages{
			m.VectorInt(),
		}

	case crc_messages_receivedMessages:
		r = TL_messages_receivedMessages{
			m.Int(),
		}

	case crc_messages_setTyping:
		r = TL_messages_setTyping{
			m.Object(),
			m.Object(),
		}

	case crc_messages_sendMessage:
		rInit := TL_messages_sendMessage{}
		rInit.flags = m.Int()
		if rInit.flags&(1<<1) == (1 << 1) {
			rInit.no_webpage = true
		}
		if rInit.flags&(1<<5) == (1 << 5) {
			rInit.silent = true
		}
		if rInit.flags&(1<<6) == (1 << 6) {
			rInit.background = true
		}
		if rInit.flags&(1<<7) == (1 << 7) {
			rInit.clear_draft = true
		}
		rInit.peer = m.Object()
		if rInit.flags&(1<<0) == (1 << 0) {
			rInit.reply_to_msg_id = m.Int()
		}
		rInit.message = m.String()
		rInit.random_id = m.Long()
		if rInit.flags&(1<<2) == (1 << 2) {
			rInit.reply_markup = m.Object()
		}
		if rInit.flags&(1<<3) == (1 << 3) {
			rInit.entities = m.Vector()
		}
		r = rInit

	case crc_messages_sendMedia:
		rInit := TL_messages_sendMedia{}
		rInit.flags = m.Int()
		if rInit.flags&(1<<5) == (1 << 5) {
			rInit.silent = true
		}
		if rInit.flags&(1<<6) == (1 << 6) {
			rInit.background = true
		}
		if rInit.flags&(1<<7) == (1 << 7) {
			rInit.clear_draft = true
		}
		rInit.peer = m.Object()
		if rInit.flags&(1<<0) == (1 << 0) {
			rInit.reply_to_msg_id = m.Int()
		}
		rInit.media = m.Object()
		rInit.random_id = m.Long()
		if rInit.flags&(1<<2) == (1 << 2) {
			rInit.reply_markup = m.Object()
		}
		r = rInit

	case crc_messages_forwardMessages:
		rInit := TL_messages_forwardMessages{}
		rInit.flags = m.Int()
		if rInit.flags&(1<<5) == (1 << 5) {
			rInit.silent = true
		}
		if rInit.flags&(1<<6) == (1 << 6) {
			rInit.background = true
		}
		rInit.from_peer = m.Object()
		rInit.id = m.VectorInt()
		rInit.random_id = m.VectorLong()
		rInit.to_peer = m.Object()
		r = rInit

	case crc_messages_reportSpam:
		r = TL_messages_reportSpam{
			m.Object(),
		}

	case crc_messages_hideReportSpam:
		r = TL_messages_hideReportSpam{
			m.Object(),
		}

	case crc_messages_getPeerSettings:
		r = TL_messages_getPeerSettings{
			m.Object(),
		}

	case crc_messages_getChats:
		r = TL_messages_getChats{
			m.VectorInt(),
		}

	case crc_messages_getFullChat:
		r = TL_messages_getFullChat{
			m.Int(),
		}

	case crc_messages_editChatTitle:
		r = TL_messages_editChatTitle{
			m.Int(),
			m.String(),
		}

	case crc_messages_editChatPhoto:
		r = TL_messages_editChatPhoto{
			m.Int(),
			m.Object(),
		}

	case crc_messages_addChatUser:
		r = TL_messages_addChatUser{
			m.Int(),
			m.Object(),
			m.Int(),
		}

	case crc_messages_deleteChatUser:
		r = TL_messages_deleteChatUser{
			m.Int(),
			m.Object(),
		}

	case crc_messages_createChat:
		r = TL_messages_createChat{
			m.Vector(),
			m.String(),
		}

	case crc_updates_getState:
		r = TL_updates_getState{}

	case crc_updates_getDifference:
		r = TL_updates_getDifference{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_photos_updateProfilePhoto:
		r = TL_photos_updateProfilePhoto{
			m.Object(),
			m.Object(),
		}

	case crc_photos_uploadProfilePhoto:
		r = TL_photos_uploadProfilePhoto{
			m.Object(),
			m.String(),
			m.Object(),
			m.Object(),
		}

	case crc_photos_deletePhotos:
		r = TL_photos_deletePhotos{
			m.Vector(),
		}

	case crc_upload_saveFilePart:
		r = TL_upload_saveFilePart{
			m.Long(),
			m.Int(),
			m.StringBytes(),
		}

	case crc_upload_getFile:
		r = TL_upload_getFile{
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crc_help_getConfig:
		r = TL_help_getConfig{}

	case crc_help_getNearestDc:
		r = TL_help_getNearestDc{}

	case crc_help_getAppUpdate:
		r = TL_help_getAppUpdate{}

	case crc_help_saveAppLog:
		r = TL_help_saveAppLog{
			m.Vector(),
		}

	case crc_help_getInviteText:
		r = TL_help_getInviteText{}

	case crc_photos_getUserPhotos:
		r = TL_photos_getUserPhotos{
			m.Object(),
			m.Int(),
			m.Long(),
			m.Int(),
		}

	case crc_messages_forwardMessage:
		r = TL_messages_forwardMessage{
			m.Object(),
			m.Int(),
			m.Long(),
		}

	case crc_messages_sendBroadcast:
		r = TL_messages_sendBroadcast{
			m.Vector(),
			m.VectorLong(),
			m.String(),
			m.Object(),
		}

	case crc_messages_getDhConfig:
		r = TL_messages_getDhConfig{
			m.Int(),
			m.Int(),
		}

	case crc_messages_requestEncryption:
		r = TL_messages_requestEncryption{
			m.Object(),
			m.Int(),
			m.StringBytes(),
		}

	case crc_messages_acceptEncryption:
		r = TL_messages_acceptEncryption{
			m.Object(),
			m.StringBytes(),
			m.Long(),
		}

	case crc_messages_discardEncryption:
		r = TL_messages_discardEncryption{
			m.Int(),
		}

	case crc_messages_setEncryptedTyping:
		r = TL_messages_setEncryptedTyping{
			m.Object(),
			m.Object(),
		}

	case crc_messages_readEncryptedHistory:
		r = TL_messages_readEncryptedHistory{
			m.Object(),
			m.Int(),
		}

	case crc_messages_sendEncrypted:
		r = TL_messages_sendEncrypted{
			m.Object(),
			m.Long(),
			m.StringBytes(),
		}

	case crc_messages_sendEncryptedFile:
		r = TL_messages_sendEncryptedFile{
			m.Object(),
			m.Long(),
			m.StringBytes(),
			m.Object(),
		}

	case crc_messages_sendEncryptedService:
		r = TL_messages_sendEncryptedService{
			m.Object(),
			m.Long(),
			m.StringBytes(),
		}

	case crc_messages_receivedQueue:
		r = TL_messages_receivedQueue{
			m.Int(),
		}

	case crc_upload_saveBigFilePart:
		r = TL_upload_saveBigFilePart{
			m.Long(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
		}

	case crc_initConnection:
		r = TL_initConnection{
			m.Int(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.Object(),
		}

	case crc_help_getSupport:
		r = TL_help_getSupport{}

	case crc_messages_readMessageContents:
		r = TL_messages_readMessageContents{
			m.VectorInt(),
		}

	case crc_account_checkUsername:
		r = TL_account_checkUsername{
			m.String(),
		}

	case crc_account_updateUsername:
		r = TL_account_updateUsername{
			m.String(),
		}

	case crc_contacts_search:
		r = TL_contacts_search{
			m.String(),
			m.Int(),
		}

	case crc_account_getPrivacy:
		r = TL_account_getPrivacy{
			m.Object(),
		}

	case crc_account_setPrivacy:
		r = TL_account_setPrivacy{
			m.Object(),
			m.Vector(),
		}

	case crc_account_deleteAccount:
		r = TL_account_deleteAccount{
			m.String(),
		}

	case crc_account_getAccountTTL:
		r = TL_account_getAccountTTL{}

	case crc_account_setAccountTTL:
		r = TL_account_setAccountTTL{
			m.Object(),
		}

	case crc_invokeWithLayer:
		r = TL_invokeWithLayer{
			m.Int(),
			m.Object(),
		}

	case crc_contacts_resolveUsername:
		r = TL_contacts_resolveUsername{
			m.String(),
		}

	case crc_account_sendChangePhoneCode:
		rInit := TL_account_sendChangePhoneCode{}
		rInit.flags = m.Int()
		if rInit.flags&(1<<0) == (1 << 0) {
			rInit.allow_flashcall = true
		}
		rInit.phone_number = m.String()
		if rInit.flags&(1<<0) == (1 << 0) {
			rInit.current_number = m.Object()
		}
		r = rInit

	case crc_account_changePhone:
		r = TL_account_changePhone{
			m.String(),
			m.String(),
			m.String(),
		}

	case crc_messages_getStickers:
		r = TL_messages_getStickers{
			m.String(),
			m.String(),
		}

	case crc_messages_getAllStickers:
		r = TL_messages_getAllStickers{
			m.Int(),
		}

	case crc_account_updateDeviceLocked:
		r = TL_account_updateDeviceLocked{
			m.Int(),
		}

	case crc_auth_importBotAuthorization:
		r = TL_auth_importBotAuthorization{
			m.Int(),
			m.Int(),
			m.String(),
			m.String(),
		}

	case crc_messages_getWebPagePreview:
		r = TL_messages_getWebPagePreview{
			m.String(),
		}

	case crc_account_getAuthorizations:
		r = TL_account_getAuthorizations{}

	case crc_account_resetAuthorization:
		r = TL_account_resetAuthorization{
			m.Long(),
		}

	case crc_account_getPassword:
		r = TL_account_getPassword{}

	case crc_account_getPasswordSettings:
		r = TL_account_getPasswordSettings{
			m.StringBytes(),
		}

	case crc_account_updatePasswordSettings:
		r = TL_account_updatePasswordSettings{
			m.StringBytes(),
			m.Object(),
		}

	case crc_auth_checkPassword:
		r = TL_auth_checkPassword{
			m.StringBytes(),
		}

	case crc_auth_requestPasswordRecovery:
		r = TL_auth_requestPasswordRecovery{}

	case crc_auth_recoverPassword:
		r = TL_auth_recoverPassword{
			m.String(),
		}

	case crc_invokeWithoutUpdates:
		r = TL_invokeWithoutUpdates{
			m.Object(),
		}

	case crc_messages_exportChatInvite:
		r = TL_messages_exportChatInvite{
			m.Int(),
		}

	case crc_messages_checkChatInvite:
		r = TL_messages_checkChatInvite{
			m.String(),
		}

	case crc_messages_importChatInvite:
		r = TL_messages_importChatInvite{
			m.String(),
		}

	case crc_messages_getStickerSet:
		r = TL_messages_getStickerSet{
			m.Object(),
		}

	case crc_messages_installStickerSet:
		r = TL_messages_installStickerSet{
			m.Object(),
			m.Object(),
		}

	case crc_messages_uninstallStickerSet:
		r = TL_messages_uninstallStickerSet{
			m.Object(),
		}

	case crc_messages_startBot:
		r = TL_messages_startBot{
			m.Object(),
			m.Object(),
			m.Long(),
			m.String(),
		}

	case crc_help_getAppChangelog:
		r = TL_help_getAppChangelog{}

	case crc_messages_getMessagesViews:
		r = TL_messages_getMessagesViews{
			m.Object(),
			m.VectorInt(),
			m.Object(),
		}

	case crc_channels_readHistory:
		r = TL_channels_readHistory{
			m.Object(),
			m.Int(),
		}

	case crc_channels_deleteMessages:
		r = TL_channels_deleteMessages{
			m.Object(),
			m.VectorInt(),
		}

	case crc_channels_deleteUserHistory:
		r = TL_channels_deleteUserHistory{
			m.Object(),
			m.Object(),
		}

	case crc_channels_reportSpam:
		r = TL_channels_reportSpam{
			m.Object(),
			m.Object(),
			m.VectorInt(),
		}

	case crc_channels_getMessages:
		r = TL_channels_getMessages{
			m.Object(),
			m.VectorInt(),
		}

	case crc_channels_getParticipants:
		r = TL_channels_getParticipants{
			m.Object(),
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crc_channels_getParticipant:
		r = TL_channels_getParticipant{
			m.Object(),
			m.Object(),
		}

	case crc_channels_getChannels:
		r = TL_channels_getChannels{
			m.Vector(),
		}

	case crc_channels_getFullChannel:
		r = TL_channels_getFullChannel{
			m.Object(),
		}

	case crc_channels_createChannel:
		rInit := TL_channels_createChannel{}
		rInit.flags = m.Int()
		if rInit.flags&(1<<0) == (1 << 0) {
			rInit.broadcast = true
		}
		if rInit.flags&(1<<1) == (1 << 1) {
			rInit.megagroup = true
		}
		rInit.title = m.String()
		rInit.about = m.String()
		r = rInit

	case crc_channels_editAbout:
		r = TL_channels_editAbout{
			m.Object(),
			m.String(),
		}

	case crc_channels_editAdmin:
		r = TL_channels_editAdmin{
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_channels_editTitle:
		r = TL_channels_editTitle{
			m.Object(),
			m.String(),
		}

	case crc_channels_editPhoto:
		r = TL_channels_editPhoto{
			m.Object(),
			m.Object(),
		}

	case crc_channels_checkUsername:
		r = TL_channels_checkUsername{
			m.Object(),
			m.String(),
		}

	case crc_channels_updateUsername:
		r = TL_channels_updateUsername{
			m.Object(),
			m.String(),
		}

	case crc_channels_joinChannel:
		r = TL_channels_joinChannel{
			m.Object(),
		}

	case crc_channels_leaveChannel:
		r = TL_channels_leaveChannel{
			m.Object(),
		}

	case crc_channels_inviteToChannel:
		r = TL_channels_inviteToChannel{
			m.Object(),
			m.Vector(),
		}

	case crc_channels_kickFromChannel:
		r = TL_channels_kickFromChannel{
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_channels_exportInvite:
		r = TL_channels_exportInvite{
			m.Object(),
		}

	case crc_channels_deleteChannel:
		r = TL_channels_deleteChannel{
			m.Object(),
		}

	case crc_updates_getChannelDifference:
		r = TL_updates_getChannelDifference{
			m.Object(),
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crc_messages_toggleChatAdmins:
		r = TL_messages_toggleChatAdmins{
			m.Int(),
			m.Object(),
		}

	case crc_messages_editChatAdmin:
		r = TL_messages_editChatAdmin{
			m.Int(),
			m.Object(),
			m.Object(),
		}

	case crc_messages_migrateChat:
		r = TL_messages_migrateChat{
			m.Int(),
		}

	case crc_messages_searchGlobal:
		r = TL_messages_searchGlobal{
			m.String(),
			m.Int(),
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crc_help_getTermsOfService:
		r = TL_help_getTermsOfService{}

	case crc_messages_reorderStickerSets:
		r = TL_messages_reorderStickerSets{
			m.VectorLong(),
		}

	case crc_messages_getDocumentByHash:
		r = TL_messages_getDocumentByHash{
			m.StringBytes(),
			m.Int(),
			m.String(),
		}

	case crc_messages_searchGifs:
		r = TL_messages_searchGifs{
			m.String(),
			m.Int(),
		}

	case crc_messages_getSavedGifs:
		r = TL_messages_getSavedGifs{
			m.Int(),
		}

	case crc_messages_saveGif:
		r = TL_messages_saveGif{
			m.Object(),
			m.Object(),
		}

	case crc_messages_getInlineBotResults:
		rInit := TL_messages_getInlineBotResults{}
		rInit.flags = m.Int()
		rInit.bot = m.Object()
		rInit.peer = m.Object()
		if rInit.flags&(1<<0) == (1 << 0) {
			rInit.geo_point = m.Object()
		}
		rInit.query = m.String()
		rInit.offset = m.String()
		r = rInit

	case crc_messages_setInlineBotResults:
		rInit := TL_messages_setInlineBotResults{}
		rInit.flags = m.Int()
		if rInit.flags&(1<<0) == (1 << 0) {
			rInit.gallery = true
		}
		if rInit.flags&(1<<1) == (1 << 1) {
			rInit.private = true
		}
		rInit.query_id = m.Long()
		rInit.results = m.Vector()
		rInit.cache_time = m.Int()
		if rInit.flags&(1<<2) == (1 << 2) {
			rInit.next_offset = m.String()
		}
		if rInit.flags&(1<<3) == (1 << 3) {
			rInit.switch_pm = m.Object()
		}
		r = rInit

	case crc_messages_sendInlineBotResult:
		rInit := TL_messages_sendInlineBotResult{}
		rInit.flags = m.Int()
		if rInit.flags&(1<<5) == (1 << 5) {
			rInit.silent = true
		}
		if rInit.flags&(1<<6) == (1 << 6) {
			rInit.background = true
		}
		if rInit.flags&(1<<7) == (1 << 7) {
			rInit.clear_draft = true
		}
		rInit.peer = m.Object()
		if rInit.flags&(1<<0) == (1 << 0) {
			rInit.reply_to_msg_id = m.Int()
		}
		rInit.random_id = m.Long()
		rInit.query_id = m.Long()
		rInit.id = m.String()
		r = rInit

	case crc_channels_toggleInvites:
		r = TL_channels_toggleInvites{
			m.Object(),
			m.Object(),
		}

	case crc_channels_exportMessageLink:
		r = TL_channels_exportMessageLink{
			m.Object(),
			m.Int(),
		}

	case crc_channels_toggleSignatures:
		r = TL_channels_toggleSignatures{
			m.Object(),
			m.Object(),
		}

	case crc_channels_updatePinnedMessage:
		rInit := TL_channels_updatePinnedMessage{}
		rInit.flags = m.Int()
		if rInit.flags&(1<<0) == (1 << 0) {
			rInit.silent = true
		}
		rInit.channel = m.Object()
		rInit.id = m.Int()
		r = rInit

	case crc_auth_resendCode:
		r = TL_auth_resendCode{
			m.String(),
			m.String(),
		}

	case crc_auth_cancelCode:
		r = TL_auth_cancelCode{
			m.String(),
			m.String(),
		}

	case crc_messages_getMessageEditData:
		r = TL_messages_getMessageEditData{
			m.Object(),
			m.Int(),
		}

	case crc_messages_editMessage:
		rInit := TL_messages_editMessage{}
		rInit.flags = m.Int()
		if rInit.flags&(1<<1) == (1 << 1) {
			rInit.no_webpage = true
		}
		rInit.peer = m.Object()
		rInit.id = m.Int()
		if rInit.flags&(1<<11) == (1 << 11) {
			rInit.message = m.String()
		}
		if rInit.flags&(1<<2) == (1 << 2) {
			rInit.reply_markup = m.Object()
		}
		if rInit.flags&(1<<3) == (1 << 3) {
			rInit.entities = m.Vector()
		}
		r = rInit

	case crc_messages_editInlineBotMessage:
		rInit := TL_messages_editInlineBotMessage{}
		rInit.flags = m.Int()
		if rInit.flags&(1<<1) == (1 << 1) {
			rInit.no_webpage = true
		}
		rInit.id = m.Object()
		if rInit.flags&(1<<11) == (1 << 11) {
			rInit.message = m.String()
		}
		if rInit.flags&(1<<2) == (1 << 2) {
			rInit.reply_markup = m.Object()
		}
		if rInit.flags&(1<<3) == (1 << 3) {
			rInit.entities = m.Vector()
		}
		r = rInit

	case crc_messages_getBotCallbackAnswer:
		r = TL_messages_getBotCallbackAnswer{
			m.Object(),
			m.Int(),
			m.StringBytes(),
		}

	case crc_messages_setBotCallbackAnswer:
		rInit := TL_messages_setBotCallbackAnswer{}
		rInit.flags = m.Int()
		if rInit.flags&(1<<1) == (1 << 1) {
			rInit.alert = true
		}
		rInit.query_id = m.Long()
		if rInit.flags&(1<<0) == (1 << 0) {
			rInit.message = m.String()
		}
		if rInit.flags&(1<<2) == (1 << 2) {
			rInit.url = m.String()
		}
		r = rInit

	case crc_contacts_getTopPeers:
		rInit := TL_contacts_getTopPeers{}
		rInit.flags = m.Int()
		if rInit.flags&(1<<0) == (1 << 0) {
			rInit.correspondents = true
		}
		if rInit.flags&(1<<1) == (1 << 1) {
			rInit.bots_pm = true
		}
		if rInit.flags&(1<<2) == (1 << 2) {
			rInit.bots_inline = true
		}
		if rInit.flags&(1<<10) == (1 << 10) {
			rInit.groups = true
		}
		if rInit.flags&(1<<15) == (1 << 15) {
			rInit.channels = true
		}
		rInit.offset = m.Int()
		rInit.limit = m.Int()
		rInit.hash = m.Int()
		r = rInit

	case crc_contacts_resetTopPeerRating:
		r = TL_contacts_resetTopPeerRating{
			m.Object(),
			m.Object(),
		}

	case crc_messages_getPeerDialogs:
		r = TL_messages_getPeerDialogs{
			m.Vector(),
		}

	case crc_messages_saveDraft:
		rInit := TL_messages_saveDraft{}
		rInit.flags = m.Int()
		if rInit.flags&(1<<1) == (1 << 1) {
			rInit.no_webpage = true
		}
		if rInit.flags&(1<<0) == (1 << 0) {
			rInit.reply_to_msg_id = m.Int()
		}
		rInit.peer = m.Object()
		rInit.message = m.String()
		if rInit.flags&(1<<3) == (1 << 3) {
			rInit.entities = m.Vector()
		}
		r = rInit

	case crc_messages_getAllDrafts:
		r = TL_messages_getAllDrafts{}

	case crc_messages_getFeaturedStickers:
		r = TL_messages_getFeaturedStickers{
			m.Int(),
		}

	case crc_messages_readFeaturedStickers:
		r = TL_messages_readFeaturedStickers{}

	case crc_messages_getRecentStickers:
		r = TL_messages_getRecentStickers{
			m.Int(),
		}

	case crc_messages_saveRecentSticker:
		r = TL_messages_saveRecentSticker{
			m.Object(),
			m.Object(),
		}

	case crc_messages_clearRecentStickers:
		r = TL_messages_clearRecentStickers{}

	case crc_messages_getUnusedStickers:
		r = TL_messages_getUnusedStickers{
			m.Int(),
		}

	case crc_messages_getArchivedStickers:
		r = TL_messages_getArchivedStickers{
			m.Long(),
			m.Int(),
		}

	case crc_account_sendConfirmPhoneCode:
		rInit := TL_account_sendConfirmPhoneCode{}
		rInit.flags = m.Int()
		if rInit.flags&(1<<0) == (1 << 0) {
			rInit.allow_flashcall = true
		}
		rInit.hash = m.String()
		if rInit.flags&(1<<0) == (1 << 0) {
			rInit.current_number = m.Object()
		}
		r = rInit

	case crc_account_confirmPhone:
		r = TL_account_confirmPhone{
			m.String(),
			m.String(),
		}

	default:
		m.err = fmt.Errorf("Unknown constructor: %08x", constructor)
		return nil

	}

	if m.err != nil {
		return nil
	}

	return
}
