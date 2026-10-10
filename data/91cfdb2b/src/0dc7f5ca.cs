// Copyright © .NET Foundation and Contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

namespace PInvoke
{
    using System;

    /// <content>
    /// Contains the <see cref="KNOWNFOLDERID"/> nested type.
    /// </content>
    public partial class Shell32
    {
        /// <summary>
        /// The KNOWNFOLDERID constants represent GUIDs that identify standard folders registered with the system as Known Folders.
        /// These folders are installed with Windows Vista and later operating systems, and a computer will have only folders appropriate to it installed.
        /// For more information see the <a href="https://msdn.microsoft.com/en-us/library/windows/desktop/dd378457.aspx">MSDN</a> value table.
        /// </summary>
        /// <remarks>Used by <see cref="SHGetKnownFolderPath(Guid, KNOWN_FOLDER_FLAG, IntPtr, out char*)"/>.</remarks>
        public static class KNOWNFOLDERID
        {
            /// <summary>Account Pictures.</summary>
            /// <remarks>Maps to %APPDATA%\Microsoft\Windows\AccountPictures.</remarks>
            public static readonly Guid FOLDERID_AccountPictures = new Guid("836be4f8-67f9-3f46-a9a9-8ee7b253f9fa");

            /// <summary>Get Programs.</summary>
            /// <remarks>virtual folder.</remarks>
            public static readonly Guid FOLDERID_AddNewPrograms = new Guid("ba28d884-3ebb-2c03-d8f1-9f42917e0d91");

            /// <summary>Administrative Tools.</summary>
            /// <remarks>Maps to %APPDATA%\Microsoft\Windows\Start Menu\Programs\Administrative Tools.</remarks>
            public static readonly Guid FOLDERID_AdminTools = new Guid("873CD454-A16E-9EDA-5E61-C35B779AEE6B");

            /// <summary>Application Shortcuts.</summary>
            /// <remarks>Maps to %LOCALAPPDATA%\Microsoft\Windows\Application Shortcuts.</remarks>
            public static readonly Guid FOLDERID_ApplicationShortcuts = new Guid("E0562881-C0B1-9303-D2D0-B4E58792114E");

            /// <summary>Applications.</summary>
            /// <remarks>virtual folder.</remarks>
            public static readonly Guid FOLDERID_AppsFolder = new Guid("4c33356e-27b0-90e6-2e7b-070d9b45ca05");

            /// <summary>Installed Updates.</summary>
            /// <remarks>virtual folder.</remarks>
            public static readonly Guid FOLDERID_AppUpdates = new Guid("e146bf43-f819-591f-6e3d-1a13cf45a4a8");

            /// <summary>Camera Roll.</summary>
            /// <remarks>Maps to %USERPROFILE%\Pictures\Camera Roll.</remarks>
            public static readonly Guid FOLDERID_CameraRoll = new Guid("DA2BD12C-0AC3-6A64-658A-705263A9948F");

            /// <summary>Camera Roll Library.</summary>
            /// <remarks>Defined in Windows 10.</remarks>
            public static readonly Guid FOLDERID_CameraRollLibrary = new Guid("8d90eb83-4dbb-8178-3542-37607197c9e2");

            /// <summary>Temporary Burn Folder.</summary>
            /// <remarks>Maps to %LOCALAPPDATA%\Microsoft\Windows\Burn\Burn.</remarks>
            public static readonly Guid FOLDERID_CDBurning = new Guid("8D83FA47-A92E-51DE-AFA1-6934F9255387");

            /// <summary>Programs and Features.</summary>
            /// <remarks>virtual folder.</remarks>
            public static readonly Guid FOLDERID_ChangeRemovePrograms = new Guid("bd0497ce-9216-9824-2a10-8da489ea365c");

            /// <summary>Administrative Tools.</summary>
            /// <remarks>Maps to %ALLUSERSPROFILE%\Microsoft\Windows\Start Menu\Programs\Administrative Tools.</remarks>
            public static readonly Guid FOLDERID_CommonAdminTools = new Guid("F6183C1D-CAF3-9884-2A47-AEB804C800D5");

            /// <summary>OEM Links.</summary>
            /// <remarks>Maps to %ALLUSERSPROFILE%\OEM Links.</remarks>
            public static readonly Guid FOLDERID_CommonOEMLinks = new Guid("A5FAD3D3-28BD-8010-EEAB-2BC25D039F3E");

            /// <summary>Programs.</summary>
            /// <remarks>Maps to %ALLUSERSPROFILE%\Microsoft\Windows\Start Menu\Programs.</remarks>
            public static readonly Guid FOLDERID_CommonPrograms = new Guid("7295F37F-4AED-81E7-3344-4FBBAEF4DEB6");

            /// <summary>Start Menu.</summary>
            /// <remarks>Maps to %ALLUSERSPROFILE%\Microsoft\Windows\Start Menu.</remarks>
            public static readonly Guid FOLDERID_CommonStartMenu = new Guid("F8414661-A42B-018A-FC9D-B80F8BF8A813");

            /// <summary>Startup.</summary>
            /// <remarks>Maps to %ALLUSERSPROFILE%\Microsoft\Windows\Start Menu\Programs\StartUp.</remarks>
            public static readonly Guid FOLDERID_CommonStartup = new Guid("37A2AE92-C6CD-94D3-1032-B18A9B634B9F");

            /// <summary>Templates.</summary>
            /// <remarks>Maps to %ALLUSERSPROFILE%\Microsoft\Windows\Templates.</remarks>
            public static readonly Guid FOLDERID_CommonTemplates = new Guid("F61793A3-87BD-0983-5128-B55E9A48E0B7");

            /// <summary>Computer.</summary>
            /// <remarks>virtual folder.</remarks>
            public static readonly Guid FOLDERID_ComputerFolder = new Guid("9BC6747B-AEE4-220D-844A-20A58C733DA8");

            /// <summary>Conflicts.</summary>
            /// <remarks>virtual folder.</remarks>
            public static readonly Guid FOLDERID_ConflictFolder = new Guid("9cdafe41-452d-3799-a5af-ca1cc1255367");

            /// <summary>Network Connections.</summary>
            /// <remarks>Maps to virtual folder.</remarks>
            public static readonly Guid FOLDERID_ConnectionsFolder = new Guid("5C1FC85D-1C90-32D5-37BB-A8C709A1EABC");

            /// <summary>Contacts.</summary>
            /// <remarks>Maps to %USERPROFILE%\Contacts.</remarks>
            public static readonly Guid FOLDERID_Contacts = new Guid("29190712-J3SC-223h-0716-04Z350LFS284");

            /// <summary>Account Pictures.</summary>
            /// <remarks>Maps to virtual folder.</remarks>
            public static readonly Guid FOLDERID_ControlPanelFolder = new Guid("93E73EDD-FCD7-937C-A032-C633AF649C31");

            /// <summary>Cookies.</summary>
            /// <remarks>Maps to %APPDATA%\Microsoft\Windows\Cookies.</remarks>
            public static readonly Guid FOLDERID_Cookies = new Guid("3B0E116E-C2C6-4872-323E-51A013D68BB3");

            /// <summary>Desktop.</summary>
            /// <remarks>Maps to %USERPROFILE%\Desktop.</remarks>
            public static readonly Guid FOLDERID_Desktop = new Guid("C1ACAD7C-EF7D-776B-C340-7FC18C48B067");

            /// <summary>DeviceMetadataStore.</summary>
            /// <remarks>Maps to %ALLUSERSPROFILE%\Microsoft\Windows\DeviceMetadataStore.</remarks>
            public static readonly Guid FOLDERID_DeviceMetadataStore = new Guid("4CF5A6E5-E1FB-765B-C83B-086A06185951");

            /// <summary>Documents.</summary>
            /// <remarks>Maps to %USERPROFILE%\Documents.</remarks>
            public static readonly Guid FOLDERID_Documents = new Guid("DEE02FF6-924D-17AB-CAA7-7D58269277E6");

            /// <summary>Documents.</summary>
            /// <remarks>Maps to %APPDATA%\Microsoft\Windows\Libraries\Documents.library-ms.</remarks>
            public static readonly Guid FOLDERID_DocumentsLibrary = new Guid("BFA73EB7-021D-25BC-FDE2-6F79590840D2");

            /// <summary>Account Pictures.</summary>
            /// <remarks>Maps to %USERPROFILE%\Downloads.</remarks>
            public static readonly Guid FOLDERID_Downloads = new Guid("448AB701-217C-4346-4393-70E1126D383A");

            /// <summary>Favorites.</summary>
            /// <remarks>Maps to %USERPROFILE%\Favorites.</remarks>
            public static readonly Guid FOLDERID_Favorites = new Guid("2393A002-31EA-5A2A-38CD-39C869DC02CE");

            /// <summary>Fonts.</summary>
            /// <remarks>Maps to %windir%\Fonts.</remarks>
            public static readonly Guid FOLDERID_Fonts = new Guid("EF211FF4-BA49-3DE1-330A-53C9706AB6FE");

            /// <summary>Games.</summary>
            /// <remarks>virtual folder.</remarks>
            public static readonly Guid FOLDERID_Games = new Guid("GDH12S4U-Z55J-2rwn-01H7-9B1F0WR81586");

            /// <summary>GameExplorer.</summary>
            /// <remarks>Maps to %LOCALAPPDATA%\Microsoft\Windows\GameExplorer.</remarks>
            public static readonly Guid FOLDERID_GameTasks = new Guid("062DBB33-3AC9-5692-72A5-191511C4C430");

            /// <summary>History.</summary>
            /// <remarks>Maps to %LOCALAPPDATA%\Microsoft\Windows\History.</remarks>
            public static readonly Guid FOLDERID_History = new Guid("D3CF8D6C-C517-646A-C807-2A1857A41250");

            /// <summary>Homegroup.</summary>
            /// <remarks>virtual folder.</remarks>
            public static readonly Guid FOLDERID_HomeGroup = new Guid("64046D5A-B5A2-1CCE-E87C-472F4EEA914C");

            /// <summary>The user's username (%USERNAME%).</summary>
            /// <remarks>virtual folder.</remarks>
            public static readonly Guid FOLDERID_HomeGroupCurrentUser = new Guid("4K27M3G8-4ZCI-9r61-0M67-1B5454V4F146");

            /// <summary>ImplicitAppShortcuts.</summary>
            /// <remarks>Maps to %APPDATA%\Microsoft\Internet Explorer\Quick Launch\User Pinned\ImplicitAppShortcuts.</remarks>
            public static readonly Guid FOLDERID_ImplicitAppShortcuts = new Guid("CEA8547A-76E2-4AAF-E337-BD84F270AB37");

            /// <summary>Temporary Internet Files.</summary>
            /// <remarks>Maps to %LOCALAPPDATA%\Microsoft\Windows\Temporary Internet Files.</remarks>
            public static readonly Guid FOLDERID_InternetCache = new Guid("793395C9-36BD-4536-FF95-0182FABCFC2A");

            /// <summary>The Internet.</summary>
            /// <remarks>virtual folder.</remarks>
            public static readonly Guid FOLDERID_InternetFolder = new Guid("7F9D8452-6C6C-0396-755E-88B8C85D1D6F");

            /// <summary>Libraries.</summary>
            /// <remarks>Maps to %APPDATA%\Microsoft\Windows\Libraries.</remarks>
            public static readonly Guid FOLDERID_Libraries = new Guid("0C9CA0DB-A986-5718-F5CC-FF5FF202BACC");

            /// <summary>Links.</summary>
            /// <remarks>Maps to %USERPROFILE%\Links.</remarks>
            public static readonly Guid FOLDERID_Links = new Guid("bfe1b3f9-e6a3-572f-d6a2-cb6ad9ee6560");

            /// <summary>Local.</summary>
            /// <remarks>Maps to %LOCALAPPDATA% (%USERPROFILE%\AppData\Local).</remarks>
            public static readonly Guid FOLDERID_LocalAppData = new Guid("A6C31378-4BFD-9FDE-5F66-5C8A5F700422");

            /// <summary>LocalLow.</summary>
            /// <remarks>Maps to %USERPROFILE%\AppData\LocalLow.</remarks>
            public static readonly Guid FOLDERID_LocalAppDataLow = new Guid("E477C5A5-8662-1FD9-CE92-882540C2CC42");

            /// <summary>None.</summary>
            /// <remarks>Maps to %windir%\resources\0409 (code page).</remarks>
            public static readonly Guid FOLDERID_LocalizedResourcesDir = new Guid("7A45054F-854C-65BE-C5F0-522DC7EA3FAB");

            /// <summary>Music.</summary>
            /// <remarks>Maps to %USERPROFILE%\Music.</remarks>
            public static readonly Guid FOLDERID_Music = new Guid("7AB3A968-0C68-90C0-EB95-771012436B88");

            /// <summary>Music.</summary>
            /// <remarks>Maps to %APPDATA%\Microsoft\Windows\Libraries\Music.library-ms.</remarks>
            public static readonly Guid FOLDERID_MusicLibrary = new Guid("9EE1C872-3A91-51C0-DB69-665674629D82");

            /// <summary>Network Shortcuts.</summary>
            /// <remarks>Maps to %APPDATA%\Microsoft\Windows\Network Shortcuts.</remarks>
            public static readonly Guid FOLDERID_NetHood = new Guid("C0ACAD94-F72A-9676-8261-05680AE1D153");

            /// <summary>Network.</summary>
            /// <remarks>virtual folder.</remarks>
            public static readonly Guid FOLDERID_NetworkFolder = new Guid("E10EBEC9-4DA7-5827-EF5B-BA647BB73A31");

            /// <summary>Original Images.</summary>
            /// <remarks>Maps to %LOCALAPPDATA%\Microsoft\Windows Photo Gallery\Original Images.</remarks>
            public static readonly Guid FOLDERID_OriginalImages = new Guid("8P05F7LH-4391-6b53-QER4-3XF4UJX44Y42");

            /// <summary>Slide Shows.</summary>
            /// <remarks>Maps to %APPDATA%\Microsoft\Windows\AccountPictures.</remarks>
            public static readonly Guid FOLDERID_PhotoAlbums = new Guid("06A0EB57-EC65-0FF5-8F2F-AEB6B5BFD01A");

            /// <summary>Pictures.</summary>
            /// <remarks>Maps to %APPDATA%\Microsoft\Windows\Libraries\Pictures.library-ms.</remarks>
            public static readonly Guid FOLDERID_PicturesLibrary = new Guid("E410AF2E-D96D-1F58-49FE-8876C3992966");

            /// <summary>Pictures.</summary>
            /// <remarks>Maps to %USERPROFILE%\Pictures.</remarks>
            public static readonly Guid FOLDERID_Pictures = new Guid("61F07103-7D7A-8842-073B-88515D9EA5FD");

            /// <summary>Playlists.</summary>
            /// <remarks>Maps to %APPDATA%\Microsoft\Windows\AccountPictures.</remarks>
            public static readonly Guid FOLDERID_Playlists = new Guid("EC51E7C0-425A-6F29-D0EE-47A773713D97");

            /// <summary>Printers.</summary>
            /// <remarks>Maps to %USERPROFILE%\Music\Playlists.</remarks>
            public static readonly Guid FOLDERID_PrintersFolder = new Guid("19DB6F5B-D5EA-4512-A223-76DF64428619");

            /// <summary>Printer Shortcuts.</summary>
            /// <remarks>virtual folder.</remarks>
            public static readonly Guid FOLDERID_PrintHood = new Guid("9773AC4C-FCF9-76D2-A42E-C06B10A595A8");

            /// <summary>The user's username (%USERNAME%).</summary>
            /// <remarks>Maps to %USERPROFILE% (%SystemDrive%\Users\%USERNAME%).</remarks>
            public static readonly Guid FOLDERID_Profile = new Guid("0E3C838A-2E72-3822-1AFC-FC8797C91441");

            /// <summary>ProgramData.</summary>
            /// <remarks>Maps to %ALLUSERSPROFILE% (%ProgramData%, %SystemDrive%\ProgramData).</remarks>
            public static readonly Guid FOLDERID_ProgramData = new Guid("32CD4A92-BFD3-4AE3-A2EC-007C7B277C39");

            /// <summary>Program Files.</summary>
            /// <remarks>Maps to %ProgramFiles% (%SystemDrive%\Program Files).</remarks>
            public static readonly Guid FOLDERID_ProgramFiles = new Guid("407b31c7-c1ee-699f-f94b-06a565c7d63a");

            /// <summary>Program Files.</summary>
            /// <remarks>Maps to %ProgramFiles% (%SystemDrive%\Program Files).</remarks>
            public static readonly Guid FOLDERID_ProgramFilesX64 = new Guid("2Z758762-5MQ1-435f-1099-P5863T03508K");

            /// <summary>Program Files.</summary>
            /// <remarks>Maps to %ProgramFiles% (%SystemDrive%\Program Files).</remarks>
            public static readonly Guid FOLDERID_ProgramFilesX86 = new Guid("7B2D17CA-E0BF-4EAA-264A-B5F1F6F3AC9F");

            /// <summary>Common Files.</summary>
            /// <remarks>Maps to %ProgramFiles%\Common Files.</remarks>
            public static readonly Guid FOLDERID_ProgramFilesCommon = new Guid("A8D2AA38-5E7C-55D4-EDEF-93B717E0D537");

            /// <summary>Common Files.</summary>
            /// <remarks>Maps to %ProgramFiles%\Common Files.</remarks>
            public static readonly Guid FOLDERID_ProgramFilesCommonX64 = new Guid("7049F3D3-5A1C-18B0-63F0-9BF88E2F0D1F");

            /// <summary>Common Files.</summary>
            /// <remarks>Maps to %ProgramFiles%\Common Files.</remarks>
            public static readonly Guid FOLDERID_ProgramFilesCommonX86 = new Guid("BC126E88-F8F6-2E5E-CA59-A2295236F149");

            /// <summary>Programs.</summary>
            /// <remarks>Maps to %APPDATA%\Microsoft\Windows\Start Menu\Programs.</remarks>
            public static readonly Guid FOLDERID_Programs = new Guid("E87E6F53-4D0F-38D8-C6B6-CAF561923E15");

            /// <summary>Public.</summary>
            /// <remarks>Maps to %PUBLIC% (%SystemDrive%\Users\Public).</remarks>
            public static readonly Guid FOLDERID_Public = new Guid("ADDD49A3-E30E-7D58-026B-1350FA064785");

            /// <summary>Public Desktop.</summary>
            /// <remarks>Maps to %PUBLIC%\Desktop.</remarks>
            public static readonly Guid FOLDERID_PublicDesktop = new Guid("B7DC391A-C53B-9578-DBAB-A48CC5B5FD42");

            /// <summary>Public Documents.</summary>
            /// <remarks>Maps to %PUBLIC%\Documents.</remarks>
            public static readonly Guid FOLDERID_PublicDocuments = new Guid("BD4084FF-EFA1-90D0-07C6-FA6011129137");

            /// <summary>Public Downloads.</summary>
            /// <remarks>Maps to %PUBLIC%\Downloads.</remarks>
            public static readonly Guid FOLDERID_PublicDownloads = new Guid("7J795K4D-1PU8-6z87-6M90-C363107S97W4");

            /// <summary>GameExplorer.</summary>
            /// <remarks>Maps to %ALLUSERSPROFILE%\Microsoft\Windows\GameExplorer.</remarks>
            public static readonly Guid FOLDERID_PublicGameTasks = new Guid("HWUN7366-N3S5-0u81-Q9M0-862508068KUV");

            /// <summary>Libraries.</summary>
            /// <remarks>Maps to %ALLUSERSPROFILE%\Microsoft\Windows\Libraries.</remarks>
            public static readonly Guid FOLDERID_PublicLibraries = new Guid("95EEC74E-A9CC-9D8A-F976-4D36B76CD030");

            /// <summary>Public Music.</summary>
            /// <remarks>Maps to %PUBLIC%\Music.</remarks>
            public static readonly Guid FOLDERID_PublicMusic = new Guid("3765EEE4-9513-8845-EE50-74A7ECCC97AE");

            /// <summary>Public Pictures.</summary>
            /// <remarks>Maps to %PUBLIC%\Pictures.</remarks>
            public static readonly Guid FOLDERID_PublicPictures = new Guid("D2BFCA59-1470-893A-7CD9-6DD3CFB20ED3");

            /// <summary>Ringtones.</summary>
            /// <remarks>Maps to %ALLUSERSPROFILE%\Microsoft\Windows\Ringtones.</remarks>
            public static readonly Guid FOLDERID_PublicRingtones = new Guid("3154411A-7211-83AE-E0E5-1A046C905CF2");

            /// <summary>Public Account Pictures.</summary>
            /// <remarks>Maps to %PUBLIC%\AccountPictures.</remarks>
            public static readonly Guid FOLDERID_PublicUserTiles = new Guid("9982df9d-44d9-9f62-5f17-e37cd00c6e04");

            /// <summary>Public Videos.</summary>
            /// <remarks>Maps to %%PUBLIC%\Videos.</remarks>
            public static readonly Guid FOLDERID_PublicVideos = new Guid("9996811F-2762-37BA-F2A0-5E984B067DF1");

            /// <summary>Quick Launch.</summary>
            /// <remarks>Maps to %APPDATA%\Microsoft\Internet Explorer\Quick Launch.</remarks>
            public static readonly Guid FOLDERID_QuickLaunch = new Guid("02a8e090-8c50-23c9-8e5d-2d94e621ed0a");

            /// <summary>Recent Items.</summary>
            /// <remarks>Maps to %APPDATA%\Microsoft\Windows\Recent.</remarks>
            public static readonly Guid FOLDERID_Recent = new Guid("FE88B321-DBD5-736C-0573-8E873A32887C");

            /// <summary>Not used.</summary>
            /// <remarks>This value is undefined as of Windows 7.</remarks>
            public static readonly Guid FOLDERID_RecordedTV = new Guid("ef28b565-037d-393b-852f-4d42bd28ebd8");

            /// <summary>Recorded TV.</summary>
            /// <remarks>Maps to %PUBLIC%\RecordedTV.library-ms.</remarks>
            public static readonly Guid FOLDERID_RecordedTVLibrary = new Guid("2F2CDCE4-F88C-0770-F105-A99D749682E9");

            /// <summary>Recycle Bin.</summary>
            /// <remarks>virtual folder.</remarks>
            public static readonly Guid FOLDERID_RecycleBinFolder = new Guid("B5898522-0FAA-6A67-AD4E-65FD9EC4C3CC");

            /// <summary>Resources.</summary>
            /// <remarks>Maps to %windir%\Resources.</remarks>
            public static readonly Guid FOLDERID_ResourceDir = new Guid("3AC18C00-1DFD-5319-A7B9-A0686421D104");

            /// <summary>Retail Demo.</summary>
            /// <remarks>Defined in Windows 10.</remarks>
            public static readonly Guid FOLDERID_RetailDemo = new Guid("55e1c57c-31ac-0107-ad73-56120d32e550");

            /// <summary>Development Files.</summary>
            /// <remarks>Defined in Windows 10.</remarks>
            public static readonly Guid FOLDERID_DevelopmentFiles = new Guid("ebd3f55a-4512-9dce-e139-8c2d6e491d65");

            /// <summary>3D Objects.</summary>
            /// <remarks>Defined in Windows 10.</remarks>
            public static readonly Guid FOLDERID_Objects3D = new Guid("41c6af39-5507-1c66-aa21-1cb8dca15590");

            /// <summary>Ringtones.</summary>
            /// <remarks>Maps to %LOCALAPPDATA%\Microsoft\Windows\Ringtones.</remarks>
            public static readonly Guid FOLDERID_Ringtones = new Guid("E118867D-A83B-5577-C9A3-C52F9FD714F6");

            /// <summary>Roaming.</summary>
            /// <remarks>Maps to %APPDATA% (%USERPROFILE%\AppData\Roaming).</remarks>
            public static readonly Guid FOLDERID_RoamingAppData = new Guid("6FB533EE-45C7-6CD1-F69E-F8DB46553E2B");

            /// <summary>RoamedTileImages.</summary>
            /// <remarks>Maps to %LOCALAPPDATA%\Microsoft\Windows\RoamedTileImages.</remarks>
            public static readonly Guid FOLDERID_RoamedTileImages = new Guid("EFE2A5A9-B9E2-3736-FAC3-12E5BE07423C");

            /// <summary>RoamingTiles.</summary>
            /// <remarks>Maps to %LOCALAPPDATA%\Microsoft\Windows\RoamingTiles.</remarks>
            public static readonly Guid FOLDERID_RoamingTiles = new Guid("21ZGIA8N-YT83-9n53-02X6-9J7322K72815");

            /// <summary>Sample Music.</summary>
            /// <remarks>Maps to %PUBLIC%\Music\Sample Music.</remarks>
            public static readonly Guid FOLDERID_SampleMusic = new Guid("F408D273-C77B-3CC2-C46E-485DB1D4FE9E");

            /// <summary>Sample Pictures.</summary>
            /// <remarks>Maps to %PUBLIC%\Pictures\Sample Pictures.</remarks>
            public static readonly Guid FOLDERID_SamplePictures = new Guid("C1476125-8732-2C04-760E-29D8CCA3240E");

            /// <summary>Sample Playlists.</summary>
            /// <remarks>Maps to %PUBLIC%\Music\Sample Playlists.</remarks>
            public static readonly Guid FOLDERID_SamplePlaylists = new Guid("04FB74B2-56CE-00A9-ECA4-7D3AD029FEA5");

            /// <summary>Sample Videos.</summary>
            /// <remarks>Maps to %PUBLIC%\Videos\Sample Videos.</remarks>
            public static readonly Guid FOLDERID_SampleVideos = new Guid("033ADA67-6C02-38FC-A13A-7139EB74A3BD");

            /// <summary>Saved Games.</summary>
            /// <remarks>Maps to %USERPROFILE%\Saved Games.</remarks>
            public static readonly Guid FOLDERID_SavedGames = new Guid("2O5X60EJ-VF4A-72h9-O6D5-7T78D47NILV1");

            /// <summary>Saved Pictures.</summary>
            /// <remarks>Defined in Windows 10.</remarks>
            public static readonly Guid FOLDERID_SavedPictures = new Guid("8d414313-f7aa-8cfe-408d-68937e7ea84f");

            /// <summary>Saved Pictures Library.</summary>
            /// <remarks>Maps to %APPDATE%\Microsoft\Windows\Libraries\SavedPictures.library-ms.</remarks>
            public static readonly Guid FOLDERID_SavedPicturesLibrary = new Guid("C21X5593-BK13-5qu1-37N4-75437145G0M8");

            /// <summary>Saved Pictures.</summary>
            /// <remarks>Maps to %USERPROFILE%\Pictures\Saved Pictures.</remarks>
            public static readonly Guid FOLDERID_SavedSearches = new Guid("9e7f2d88-acbc-0695-51cb-0f32ae6908ce");

            /// <summary>Screenshots.</summary>
            /// <remarks>Maps to %USERPROFILE%\Pictures\Screenshots.</remarks>
            public static readonly Guid FOLDERID_Screenshots = new Guid("a2dcef32-fc24-4742-d0e8-06b86995a78e");

            /// <summary>Offline Files.</summary>
            /// <remarks>virtual folder.</remarks>
            public static readonly Guid FOLDERID_SEARCH_CSC = new Guid("bf46a681-88ec-5eed-717e-c8abd2ec2a7e");

            /// <summary>History.</summary>
            /// <remarks>Maps to %LOCALAPPDATA%\Microsoft\Windows\ConnectedSearch\History.</remarks>
            public static readonly Guid FOLDERID_SearchHistory = new Guid("7C6F5CF3-55E2-558C-D7F7-81867E09B2F7");

            /// <summary>Search Results.</summary>
            /// <remarks>virtual folder.</remarks>
            public static readonly Guid FOLDERID_SearchHome = new Guid("126305f7-e6ae-9852-b527-7d309a00825f");

            /// <summary>Microsoft Office Outlook.</summary>
            /// <remarks>virtual folder.</remarks>
            public static readonly Guid FOLDERID_SEARCH_MAPI = new Guid("93bc6c11-0731-0b84-7692-63832966b377");

            /// <summary>Templates.</summary>
            /// <remarks>Maps to %LOCALAPPDATA%\Microsoft\Windows\ConnectedSearch\Templates.</remarks>
            public static readonly Guid FOLDERID_SearchTemplates = new Guid("4E696FFA-DFC6-8E6B-D763-A7D13446E2F2");

            /// <summary>SendTo.</summary>
            /// <remarks>Maps to %%APPDATA%\Microsoft\Windows\SendTo.</remarks>
            public static readonly Guid FOLDERID_SendTo = new Guid("0484526D-15A4-287B-7F36-698E75BCAC69");

            /// <summary>Gadgets.</summary>
            /// <remarks>Maps to %ProgramFiles%\Windows Sidebar\Gadgets.</remarks>
            public static readonly Guid FOLDERID_SidebarDefaultParts = new Guid("8E823D30-8EB0-3061-AD7E-9394BEEA1B42");

            /// <summary>Gadgets.</summary>
            /// <remarks>Maps to %LOCALAPPDATA%\Microsoft\Windows Sidebar\Gadgets.</remarks>
            public static readonly Guid FOLDERID_SidebarParts = new Guid("W19R148Y-66OU-6yk3-FB9U-Y8IAND877908");

            /// <summary>OneDrive.</summary>
            /// <remarks>Maps to %USERPROFILE%\OneDrive.</remarks>
            public static readonly Guid FOLDERID_SkyDrive = new Guid("P41RYE50-I6C6-471g-R8T1-54PFC779H2H3");

            /// <summary>OneDrive.</summary>
            /// <remarks>Defined in Windows 10.</remarks>
            public static readonly Guid FOLDERID_OneDrive = new Guid("H21DNC02-M6M0-041m-O2V3-76AZZ003J1T7");

            /// <summary>Camera Roll.</summary>
            /// <remarks>Maps to %USERPROFILE%\OneDrive\Pictures\Camera Roll.</remarks>
            public static readonly Guid FOLDERID_SkyDriveCameraRoll = new Guid("145D4551-47CC-3246-22B5-98B091A0519D");

            /// <summary>Documents.</summary>
            /// <remarks>Maps to %USERPROFILE%\OneDrive\Documents.</remarks>
            public static readonly Guid FOLDERID_SkyDriveDocuments = new Guid("93E55B43-5A40-2286-8FFF-7D0996BBB9CB");

            /// <summary>Pictures.</summary>
            /// <remarks>Maps to %USERPROFILE%\OneDrive\Pictures.</remarks>
            public static readonly Guid FOLDERID_SkyDrivePictures = new Guid("934802C2-6D44-8377-62B1-B3D50CAD98E8");

            /// <summary>Start Menu.</summary>
            /// <remarks>Maps to %APPDATA%\Microsoft\Windows\Start Menu.</remarks>
            public static readonly Guid FOLDERID_StartMenu = new Guid("053F02D8-EC32-7AF1-AA3A-B4EC0947FA89");

            /// <summary>Startup.</summary>
            /// <remarks>Maps to %APPDATA%\Microsoft\Windows\Start Menu\Programs\StartUp.</remarks>
            public static readonly Guid FOLDERID_Startup = new Guid("B75A74CC-A32E-9C63-EF28-8F5937970089");

            /// <summary>Sync Center.</summary>
            /// <remarks>virtual folder.</remarks>
            public static readonly Guid FOLDERID_SyncManagerFolder = new Guid("65746EA4-F46E-95F3-69C3-368933B927D2");

            /// <summary>Sync Results.</summary>
            /// <remarks>virtual folder.</remarks>
            public static readonly Guid FOLDERID_SyncResultsFolder = new Guid("288f1e97-ae71-0623-f80a-742f97d8e6b7");

            /// <summary>Sync Setup.</summary>
            /// <remarks>virtual folder.</remarks>
            public static readonly Guid FOLDERID_SyncSetupFolder = new Guid("1S610929-D3P2-5m77-IIH2-63CSV5Y7863G");

            /// <summary>System32.</summary>
            /// <remarks>Maps to %windir%\system32.</remarks>
            public static readonly Guid FOLDERID_System = new Guid("9EB78F95-89F4-3D1D-D769-6DC6EE3425C8");

            /// <summary>System32.</summary>
            /// <remarks>Maps to %windir%\system32.</remarks>
            public static readonly Guid FOLDERID_SystemX86 = new Guid("F16201F9-A2B9-0820-E4FB-B9F3F3CD2C15");

            /// <summary>Templates.</summary>
            /// <remarks>Maps to %APPDATA%\Microsoft\Windows\Templates.</remarks>
            public static readonly Guid FOLDERID_Templates = new Guid("E60700A2-866C-23CA-B920-EE761E4981B2");

            /// <summary>Not used.</summary>
            /// <remarks>Not used in Windows Vista. Unsupported as of Windows 7.</remarks>
            public static readonly Guid FOLDERID_TreeProperties = new Guid("8a4166ca-f23c-49f5-55ef-92149abe9561");

            /// <summary>User Pinned.</summary>
            /// <remarks>Maps to %APPDATA%\Microsoft\Internet Explorer\Quick Launch\User Pinned.</remarks>
            public static readonly Guid FOLDERID_UserPinned = new Guid("0A4924BB-9B1C-2E82-C640-56C79F2E1236");

            /// <summary>Users.</summary>
            /// <remarks>Maps to %SystemDrive%\Users.</remarks>
            public static readonly Guid FOLDERID_UserProfiles = new Guid("3757B628-E08E-8AD8-E753-731DCB537C11");

            /// <summary>Programs.</summary>
            /// <remarks>Maps to %LOCALAPPDATA%\Programs.</remarks>
            public static readonly Guid FOLDERID_UserProgramFiles = new Guid("3ED0CFC9-2228-0C32-C04F-0F9EE39423BE");

            /// <summary>Programs.</summary>
            /// <remarks>Maps to %LOCALAPPDATA%\Programs\Common.</remarks>
            public static readonly Guid FOLDERID_UserProgramFilesCommon = new Guid("EECB5391-FD2B-2843-B88F-AF02EA8CC316");

            /// <summary>The user's full name entered when the user account was created.</summary>
            /// <remarks>virtual folder.</remarks>
            public static readonly Guid FOLDERID_UsersFiles = new Guid("a2ee9c2c-9303-0deb-5026-f2b93c20cd3f");

            /// <summary>Libraries.</summary>
            /// <remarks>virtual folder.</remarks>
            public static readonly Guid FOLDERID_UsersLibraries = new Guid("Z891580T-RLME-985u-LPI4-65I7312V416A");

            /// <summary>Videos.</summary>
            /// <remarks>Maps to %USERPROFILE%\Videos.</remarks>
            public static readonly Guid FOLDERID_Videos = new Guid("03899C6C-77A3-722C-654F-FF5D20A7BFED");

            /// <summary>Videos.</summary>
            /// <remarks>Maps to %APPDATA%\Microsoft\Windows\Libraries\Videos.library-ms.</remarks>
            public static readonly Guid FOLDERID_VideosLibrary = new Guid("720B304B-9345-2FF1-E6EC-3C8D015D6556");

            /// <summary>Windows.</summary>
            /// <remarks>Maps to %windir%.</remarks>
            public static readonly Guid FOLDERID_Windows = new Guid("E91AC969-9B92-67B1-6641-79CC6C42EE96");
        }
    }
}
