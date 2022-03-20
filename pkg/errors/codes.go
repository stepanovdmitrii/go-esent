package errors

// ErrorCode ...
type ErrorCode int64

const (
	JET_wrnNyi                                       ErrorCode = -1
	JET_errRfsFailure                                ErrorCode = -100
	JET_errRfsNotArmed                               ErrorCode = -101
	JET_errFileClose                                 ErrorCode = -102
	JET_errOutOfThreads                              ErrorCode = -103
	JET_errTooManyIO                                 ErrorCode = -105
	JET_errTaskDropped                               ErrorCode = -106
	JET_errInternalError                             ErrorCode = -107
	JET_errDatabaseBufferDependenciesCorrupted       ErrorCode = -255
	JET_errPreviousVersion                           ErrorCode = -322
	JET_errPageBoundary                              ErrorCode = -323
	JET_errKeyBoundary                               ErrorCode = -324
	JET_errBadPageLink                               ErrorCode = -327
	JET_errBadBookmark                               ErrorCode = -328
	JET_errNTSystemCallFailed                        ErrorCode = -334
	JET_errBadParentPageLink                         ErrorCode = -338
	JET_errSPAvailExtCacheOutOfSync                  ErrorCode = -340
	JET_errSPAvailExtCorrupted                       ErrorCode = -341
	JET_errSPAvailExtCacheOutOfMemory                ErrorCode = -342
	JET_errSPOwnExtCorrupted                         ErrorCode = -343
	JET_errDbTimeCorrupted                           ErrorCode = -344
	JET_errKeyTruncated                              ErrorCode = -346
	JET_errKeyTooBig                                 ErrorCode = -408
	JET_errInvalidLoggedOperation                    ErrorCode = -500
	JET_errLogFileCorrupt                            ErrorCode = -501
	JET_errNoBackupDirectory                         ErrorCode = -503
	JET_errBackupDirectoryNotEmpty                   ErrorCode = -504
	JET_errBackupInProgress                          ErrorCode = -505
	JET_errRestoreInProgress                         ErrorCode = -506
	JET_errMissingPreviousLogFile                    ErrorCode = -509
	JET_errLogWriteFail                              ErrorCode = -510
	JET_errLogDisabledDueToRecoveryFailure           ErrorCode = -511
	JET_errCannotLogDuringRecoveryRedo               ErrorCode = -512
	JET_errLogGenerationMismatch                     ErrorCode = -513
	JET_errBadLogVersion                             ErrorCode = -514
	JET_errInvalidLogSequence                        ErrorCode = -515
	JET_errLoggingDisabled                           ErrorCode = -516
	JET_errLogBufferTooSmall                         ErrorCode = -517
	JET_errLogSequenceEnd                            ErrorCode = -519
	JET_errNoBackup                                  ErrorCode = -520
	JET_errInvalidBackupSequence                     ErrorCode = -521
	JET_errBackupNotAllowedYet                       ErrorCode = -523
	JET_errDeleteBackupFileFail                      ErrorCode = -524
	JET_errMakeBackupDirectoryFail                   ErrorCode = -525
	JET_errInvalidBackup                             ErrorCode = -526
	JET_errRecoveredWithErrors                       ErrorCode = -527
	JET_errMissingLogFile                            ErrorCode = -528
	JET_errLogDiskFull                               ErrorCode = -529
	JET_errBadLogSignature                           ErrorCode = -530
	JET_errBadDbSignature                            ErrorCode = -531
	JET_errBadCheckpointSignature                    ErrorCode = -532
	JET_errCheckpointCorrupt                         ErrorCode = -533
	JET_errMissingPatchPage                          ErrorCode = -534
	JET_errBadPatchPage                              ErrorCode = -535
	JET_errRedoAbruptEnded                           ErrorCode = -536
	JET_errBadSLVSignature                           ErrorCode = -537
	JET_errPatchFileMissing                          ErrorCode = -538
	JET_errDatabaseLogSetMismatch                    ErrorCode = -539
	JET_errDatabaseStreamingFileMismatch             ErrorCode = -540
	JET_errLogFileSizeMismatch                       ErrorCode = -541
	JET_errCheckpointFileNotFound                    ErrorCode = -542
	JET_errRequiredLogFilesMissing                   ErrorCode = -543
	JET_errSoftRecoveryOnBackupDatabase              ErrorCode = -544
	JET_errLogFileSizeMismatchDatabasesConsistent    ErrorCode = -545
	JET_errLogSectorSizeMismatch                     ErrorCode = -546
	JET_errLogSectorSizeMismatchDatabasesConsistent  ErrorCode = -547
	JET_errLogSequenceEndDatabasesConsistent         ErrorCode = -548
	JET_errStreamingDataNotLogged                    ErrorCode = -549
	JET_errDatabaseDirtyShutdown                     ErrorCode = -550
	JET_errDatabaseInconsistent                      ErrorCode = -550
	JET_errConsistentTimeMismatch                    ErrorCode = -551
	JET_errDatabasePatchFileMismatch                 ErrorCode = -552
	JET_errEndingRestoreLogTooLow                    ErrorCode = -553
	JET_errStartingRestoreLogTooHigh                 ErrorCode = -554
	JET_errGivenLogFileHasBadSignature               ErrorCode = -555
	JET_errGivenLogFileIsNotContiguous               ErrorCode = -556
	JET_errMissingRestoreLogFiles                    ErrorCode = -557
	JET_errMissingFullBackup                         ErrorCode = -560
	JET_errBadBackupDatabaseSize                     ErrorCode = -561
	JET_errDatabaseAlreadyUpgraded                   ErrorCode = -562
	JET_errDatabaseIncompleteUpgrade                 ErrorCode = -563
	JET_errMissingCurrentLogFiles                    ErrorCode = -565
	JET_errDbTimeTooOld                              ErrorCode = -566
	JET_errDbTimeTooNew                              ErrorCode = -567
	JET_errMissingFileToBackup                       ErrorCode = -569
	JET_errLogTornWriteDuringHardRestore             ErrorCode = -570
	JET_errLogTornWriteDuringHardRecovery            ErrorCode = -571
	JET_errLogCorruptDuringHardRestore               ErrorCode = -573
	JET_errLogCorruptDuringHardRecovery              ErrorCode = -574
	JET_errMustDisableLoggingForDbUpgrade            ErrorCode = -575
	JET_errBadRestoreTargetInstance                  ErrorCode = -577
	JET_errRecoveredWithoutUndo                      ErrorCode = -579
	JET_errDatabasesNotFromSameSnapshot              ErrorCode = -580
	JET_errSoftRecoveryOnSnapshot                    ErrorCode = -581
	JET_errUnicodeTranslationBufferTooSmall          ErrorCode = -601
	JET_errUnicodeTranslationFail                    ErrorCode = -602
	JET_errUnicodeNormalizationNotSupported          ErrorCode = -603
	JET_errExistingLogFileHasBadSignature            ErrorCode = -610
	JET_errExistingLogFileIsNotContiguous            ErrorCode = -611
	JET_errLogReadVerifyFailure                      ErrorCode = -612
	JET_errSLVReadVerifyFailure                      ErrorCode = -613
	JET_errCheckpointDepthTooDeep                    ErrorCode = -614
	JET_errRestoreOfNonBackupDatabase                ErrorCode = -615
	JET_errInvalidGrbit                              ErrorCode = -900
	JET_errTermInProgress                            ErrorCode = -1000
	JET_errFeatureNotAvailable                       ErrorCode = -1001
	JET_errInvalidName                               ErrorCode = -1002
	JET_errInvalidParameter                          ErrorCode = -1003
	JET_errDatabaseFileReadOnly                      ErrorCode = -1008
	JET_errInvalidDatabaseId                         ErrorCode = -1010
	JET_errOutOfMemory                               ErrorCode = -1011
	JET_errOutOfDatabaseSpace                        ErrorCode = -1012
	JET_errOutOfCursors                              ErrorCode = -1013
	JET_errOutOfBuffers                              ErrorCode = -1014
	JET_errTooManyIndexes                            ErrorCode = -1015
	JET_errTooManyKeys                               ErrorCode = -1016
	JET_errRecordDeleted                             ErrorCode = -1017
	JET_errReadVerifyFailure                         ErrorCode = -1018
	JET_errPageNotInitialized                        ErrorCode = -1019
	JET_errOutOfFileHandles                          ErrorCode = -1020
	JET_errDiskIO                                    ErrorCode = -1022
	JET_errInvalidPath                               ErrorCode = -1023
	JET_errInvalidSystemPath                         ErrorCode = -1024
	JET_errInvalidLogDirectory                       ErrorCode = -1025
	JET_errRecordTooBig                              ErrorCode = -1026
	JET_errTooManyOpenDatabases                      ErrorCode = -1027
	JET_errInvalidDatabase                           ErrorCode = -1028
	JET_errNotInitialized                            ErrorCode = -1029
	JET_errAlreadyInitialized                        ErrorCode = -1030
	JET_errInitInProgress                            ErrorCode = -1031
	JET_errFileAccessDenied                          ErrorCode = -1032
	JET_errBufferTooSmall                            ErrorCode = -1038
	JET_errTooManyColumns                            ErrorCode = -1040
	JET_errContainerNotEmpty                         ErrorCode = -1043
	JET_errInvalidFilename                           ErrorCode = -1044
	JET_errInvalidBookmark                           ErrorCode = -1045
	JET_errColumnInUse                               ErrorCode = -1046
	JET_errInvalidBufferSize                         ErrorCode = -1047
	JET_errColumnNotUpdatable                        ErrorCode = -1048
	JET_errIndexInUse                                ErrorCode = -1051
	JET_errLinkNotSupported                          ErrorCode = -1052
	JET_errNullKeyDisallowed                         ErrorCode = -1053
	JET_errNotInTransaction                          ErrorCode = -1054
	JET_errTooManyActiveUsers                        ErrorCode = -1059
	JET_errInvalidCountry                            ErrorCode = -1061
	JET_errInvalidLanguageId                         ErrorCode = -1062
	JET_errInvalidCodePage                           ErrorCode = -1063
	JET_errInvalidLCMapStringFlags                   ErrorCode = -1064
	JET_errVersionStoreEntryTooBig                   ErrorCode = -1065
	JET_errVersionStoreOutOfMemoryAndCleanupTimedOut ErrorCode = -1066
	JET_errVersionStoreOutOfMemory                   ErrorCode = -1069
	JET_errCannotIndex                               ErrorCode = -1071
	JET_errRecordNotDeleted                          ErrorCode = -1072
	JET_errTooManyMempoolEntries                     ErrorCode = -1073
	JET_errOutOfObjectIDs                            ErrorCode = -1074
	JET_errOutOfLongValueIDs                         ErrorCode = -1075
	JET_errOutOfAutoincrementValues                  ErrorCode = -1076
	JET_errOutOfDbtimeValues                         ErrorCode = -1077
	JET_errOutOfSequentialIndexValues                ErrorCode = -1078
	JET_errRunningInOneInstanceMode                  ErrorCode = -1080
	JET_errRunningInMultiInstanceMode                ErrorCode = -1081
	JET_errSystemParamsAlreadySet                    ErrorCode = -1082
	JET_errSystemPathInUse                           ErrorCode = -1083
	JET_errLogFilePathInUse                          ErrorCode = -1084
	JET_errTempPathInUse                             ErrorCode = -1085
	JET_errInstanceNameInUse                         ErrorCode = -1086
	JET_errInstanceUnavailable                       ErrorCode = -1090
	JET_errDatabaseUnavailable                       ErrorCode = -1091
	JET_errInstanceUnavailableDueToFatalLogDiskFull  ErrorCode = -1092
	JET_errOutOfSessions                             ErrorCode = -1101
	JET_errWriteConflict                             ErrorCode = -1102
	JET_errTransTooDeep                              ErrorCode = -1103
	JET_errInvalidSesid                              ErrorCode = -1104
	JET_errWriteConflictPrimaryIndex                 ErrorCode = -1105
	JET_errInTransaction                             ErrorCode = -1108
	JET_errRollbackRequired                          ErrorCode = -1109
	JET_errTransReadOnly                             ErrorCode = -1110
	JET_errSessionWriteConflict                      ErrorCode = -1111
	JET_errRecordTooBigForBackwardCompatibility      ErrorCode = -1112
	JET_errCannotMaterializeForwardOnlySort          ErrorCode = -1113
	JET_errSesidTableIdMismatch                      ErrorCode = -1114
	JET_errInvalidInstance                           ErrorCode = -1115
	JET_errReadLostFlushVerifyFailure                ErrorCode = -1119
	JET_errDatabaseDuplicate                         ErrorCode = -1201
	JET_errDatabaseInUse                             ErrorCode = -1202
	JET_errDatabaseNotFound                          ErrorCode = -1203
	JET_errDatabaseInvalidName                       ErrorCode = -1204
	JET_errDatabaseInvalidPages                      ErrorCode = -1205
	JET_errDatabaseCorrupted                         ErrorCode = -1206
	JET_errDatabaseLocked                            ErrorCode = -1207
	JET_errCannotDisableVersioning                   ErrorCode = -1208
	JET_errInvalidDatabaseVersion                    ErrorCode = -1209
	JET_errDatabase200Format                         ErrorCode = -1210
	JET_errDatabase400Format                         ErrorCode = -1211
	JET_errDatabase500Format                         ErrorCode = -1212
	JET_errPageSizeMismatch                          ErrorCode = -1213
	JET_errTooManyInstances                          ErrorCode = -1214
	JET_errDatabaseSharingViolation                  ErrorCode = -1215
	JET_errAttachedDatabaseMismatch                  ErrorCode = -1216
	JET_errDatabaseInvalidPath                       ErrorCode = -1217
	JET_errDatabaseIdInUse                           ErrorCode = -1218
	JET_errForceDetachNotAllowed                     ErrorCode = -1219
	JET_errCatalogCorrupted                          ErrorCode = -1220
	JET_errPartiallyAttachedDB                       ErrorCode = -1221
	JET_errDatabaseSignInUse                         ErrorCode = -1222
	JET_errDatabaseCorruptedNoRepair                 ErrorCode = -1224
	JET_errInvalidCreateDbVersion                    ErrorCode = -1225
	JET_errTableLocked                               ErrorCode = -1302
	JET_errTableDuplicate                            ErrorCode = -1303
	JET_errTableInUse                                ErrorCode = -1304
	JET_errObjectNotFound                            ErrorCode = -1305
	JET_errDensityInvalid                            ErrorCode = -1307
	JET_errTableNotEmpty                             ErrorCode = -1308
	JET_errInvalidTableId                            ErrorCode = -1310
	JET_errTooManyOpenTables                         ErrorCode = -1311
	JET_errIllegalOperation                          ErrorCode = -1312
	JET_errTooManyOpenTablesAndCleanupTimedOut       ErrorCode = -1313
	JET_errObjectDuplicate                           ErrorCode = -1314
	JET_errInvalidObject                             ErrorCode = -1316
	JET_errCannotDeleteTempTable                     ErrorCode = -1317
	JET_errCannotDeleteSystemTable                   ErrorCode = -1318
	JET_errCannotDeleteTemplateTable                 ErrorCode = -1319
	JET_errExclusiveTableLockRequired                ErrorCode = -1322
	JET_errFixedDDL                                  ErrorCode = -1323
	JET_errFixedInheritedDDL                         ErrorCode = -1324
	JET_errCannotNestDDL                             ErrorCode = -1325
	JET_errDDLNotInheritable                         ErrorCode = -1326
	JET_errInvalidSettings                           ErrorCode = -1328
	JET_errClientRequestToStopJetService             ErrorCode = -1329
	JET_errCannotAddFixedVarColumnToDerivedTable     ErrorCode = -1330
	JET_errIndexCantBuild                            ErrorCode = -1401
	JET_errIndexHasPrimary                           ErrorCode = -1402
	JET_errIndexDuplicate                            ErrorCode = -1403
	JET_errIndexNotFound                             ErrorCode = -1404
	JET_errIndexMustStay                             ErrorCode = -1405
	JET_errIndexInvalidDef                           ErrorCode = -1406
	JET_errInvalidCreateIndex                        ErrorCode = -1409
	JET_errTooManyOpenIndexes                        ErrorCode = -1410
	JET_errMultiValuedIndexViolation                 ErrorCode = -1411
	JET_errIndexBuildCorrupted                       ErrorCode = -1412
	JET_errPrimaryIndexCorrupted                     ErrorCode = -1413
	JET_errSecondaryIndexCorrupted                   ErrorCode = -1414
	JET_errInvalidIndexId                            ErrorCode = -1416
	JET_errIndexTuplesSecondaryIndexOnly             ErrorCode = -1430
	JET_errIndexTuplesTooManyColumns                 ErrorCode = -1431
	JET_errIndexTuplesNonUniqueOnly                  ErrorCode = -1432
	JET_errIndexTuplesTextBinaryColumnsOnly          ErrorCode = -1433
	JET_errIndexTuplesVarSegMacNotAllowed            ErrorCode = -1434
	JET_errIndexTuplesInvalidLimits                  ErrorCode = -1435
	JET_errIndexTuplesCannotRetrieveFromIndex        ErrorCode = -1436
	JET_errIndexTuplesKeyTooSmall                    ErrorCode = -1437
	JET_errColumnLong                                ErrorCode = -1501
	JET_errColumnNoChunk                             ErrorCode = -1502
	JET_errColumnDoesNotFit                          ErrorCode = -1503
	JET_errNullInvalid                               ErrorCode = -1504
	JET_errColumnIllegalNull                         ErrorCode = -1504
	JET_errColumnIndexed                             ErrorCode = -1505
	JET_errColumnTooBig                              ErrorCode = -1506
	JET_errColumnNotFound                            ErrorCode = -1507
	JET_errColumnDuplicate                           ErrorCode = -1508
	JET_errMultiValuedColumnMustBeTagged             ErrorCode = -1509
	JET_errColumnRedundant                           ErrorCode = -1510
	JET_errInvalidColumnType                         ErrorCode = -1511
	JET_errTaggedNotNULL                             ErrorCode = -1514
	JET_errNoCurrentIndex                            ErrorCode = -1515
	JET_errKeyIsMade                                 ErrorCode = -1516
	JET_errBadColumnId                               ErrorCode = -1517
	JET_errBadItagSequence                           ErrorCode = -1518
	JET_errColumnInRelationship                      ErrorCode = -1519
	JET_errCannotBeTagged                            ErrorCode = -1521
	JET_errDefaultValueTooBig                        ErrorCode = -1524
	JET_errMultiValuedDuplicate                      ErrorCode = -1525
	JET_errLVCorrupted                               ErrorCode = -1526
	JET_errMultiValuedDuplicateAfterTruncation       ErrorCode = -1528
	JET_errDerivedColumnCorruption                   ErrorCode = -1529
	JET_errInvalidPlaceholderColumn                  ErrorCode = -1530
	JET_errRecordNotFound                            ErrorCode = -1601
	JET_errRecordNoCopy                              ErrorCode = -1602
	JET_errNoCurrentRecord                           ErrorCode = -1603
	JET_errRecordPrimaryChanged                      ErrorCode = -1604
	JET_errKeyDuplicate                              ErrorCode = -1605
	JET_errAlreadyPrepared                           ErrorCode = -1607
	JET_errKeyNotMade                                ErrorCode = -1608
	JET_errUpdateNotPrepared                         ErrorCode = -1609
	JET_errDataHasChanged                            ErrorCode = -1611
	JET_errLanguageNotSupported                      ErrorCode = -1619
	JET_errTooManySorts                              ErrorCode = -1701
	JET_errInvalidOnSort                             ErrorCode = -1702
	JET_errTempFileOpenError                         ErrorCode = -1803
	JET_errTooManyAttachedDatabases                  ErrorCode = -1805
	JET_errDiskFull                                  ErrorCode = -1808
	JET_errPermissionDenied                          ErrorCode = -1809
	JET_errFileNotFound                              ErrorCode = -1811
	JET_errFileInvalidType                           ErrorCode = -1812
	JET_errAfterInitialization                       ErrorCode = -1850
	JET_errLogCorrupted                              ErrorCode = -1852
	JET_errInvalidOperation                          ErrorCode = -1906
	JET_errAccessDenied                              ErrorCode = -1907
	JET_errTooManySplits                             ErrorCode = -1909
	JET_errSessionSharingViolation                   ErrorCode = -1910
	JET_errEntryPointNotFound                        ErrorCode = -1911
	JET_errSessionContextAlreadySet                  ErrorCode = -1912
	JET_errSessionContextNotSetByThisThread          ErrorCode = -1913
	JET_errSessionInUse                              ErrorCode = -1914
	JET_errRecordFormatConversionFailed              ErrorCode = -1915
	JET_errOneDatabasePerSession                     ErrorCode = -1916
	JET_errRollbackError                             ErrorCode = -1917
	JET_errCallbackFailed                            ErrorCode = -2101
	JET_errCallbackNotResolved                       ErrorCode = -2102
	JET_errOSSnapshotInvalidSequence                 ErrorCode = -2401
	JET_errOSSnapshotTimeOut                         ErrorCode = -2402
	JET_errOSSnapshotNotAllowed                      ErrorCode = -2403
	JET_errOSSnapshotInvalidSnapId                   ErrorCode = -2404
	JET_errLSCallbackNotSpecified                    ErrorCode = -3000
	JET_errLSAlreadySet                              ErrorCode = -3001
	JET_errLSNotSet                                  ErrorCode = -3002
	JET_errFileIOSparse                              ErrorCode = -4000
	JET_errFileIOBeyondEOF                           ErrorCode = -4001
	JET_errFileIOAbort                               ErrorCode = -4002
	JET_errFileIORetry                               ErrorCode = -4003
	JET_errFileIOFail                                ErrorCode = -4004
	JET_errFileCompressed                            ErrorCode = -4005
)
