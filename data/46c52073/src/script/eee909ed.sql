-- Add Test Member
INSERT INTO users.UserRegistrations VALUES 
(
	'4CDABBFE-BE20-50F9-D170-0FD12D75D624',
	'testMember@mail.com',
	'testMember@mail.com',
	'ANO7TKjxh/Dom6LG0dyoQfJciLca+e1itHQ6BZMQYs+aMbKL9MjCv6bq4gy4+MRY0w==', -- testMemberPass
	'John',
	'Doe',
	'John Doe',
	'Confirmed',
	GETDATE(),
	GETDATE()
)

INSERT INTO users.Users VALUES
(
	'0FEEDDCB-AA05-66C2-B774-8CC80F29D427',
	'testMember@mail.com',
	'testMember@mail.com',
	'ANO7TKjxh/Dom6LG0dyoQfJciLca+e1itHQ6BZMQYs+aMbKL9MjCv6bq4gy4+MRY0w==', -- testMemberPass
	1,
	'John',
	'Doe',
	'John Doe'
)

INSERT INTO meetings.Members VALUES
(
	'4EBEFACD-BC29-35B5-B349-8DC91A37E563',
	'testMember@mail.com',
	'testMember@mail.com',
	'John',
	'Doe',
	'John Doe'
)

INSERT INTO payments.Payers VALUES
(
	'0DBEBEEC-EA52-49A5-A995-7BC27B55B615',
	'testMember@mail.com',
	'testMember@mail.com',
	'John',
	'Doe',
	'John Doe'
)

INSERT INTO users.UserRoles VALUES
('6BEFAADC-FE30-71F1-E701-6DE56A01E722', 'Member')

-- Add Test Administrator
INSERT INTO users.UserRegistrations VALUES 
(
	'6185441B-6E2E-3D88-1672-7EABB1A7A77A',
	'testAdmin@mail.com',
	'testAdmin@mail.com',
	'AK0qplH5peUHwnCVuzW9zy0JGZTTG6/Ji88twX+nw9JdTUwqa3Wol1K4m5aCG9pE2A==', -- testAdminPass
	'Jane',
	'Doe',
	'Jane Doe',
	'Confirmed',
	GETDATE(),
	GETDATE()
)

INSERT INTO users.Users VALUES
(
	'8453105E-7C0E-7C27-1918-3DBCB4C3B14D',
	'testAdmin@mail.com',
	'testAdmin@mail.com',
	'AK0qplH5peUHwnCVuzW9zy0JGZTTG6/Ji88twX+nw9JdTUwqa3Wol1K4m5aCG9pE2A==', -- testAdminPass
	1,
	'Jane',
	'Doe',
	'Jane Doe'
)

INSERT INTO users.UserRoles VALUES
('4468259C-7E6C-9C81-2442-8FBDA3E7C73A', 'Administrator')

-- Roles to Permissions

INSERT INTO users.[Permissions] ([Code], [Name]) VALUES
	-- Meetings
	('GetMeetingGroupProposals', 'GetMeetingGroupProposals'),
	('ProposeMeetingGroup', 'ProposeMeetingGroup'),
	('CreateNewMeeting','CreateNewMeeting'),
	('EditMeeting','EditMeeting'),
	('AddMeetingAttendee','AddMeetingAttendee'),
	('RemoveMeetingAttendee','RemoveMeetingAttendee'),
	('AddNotAttendee','AddNotAttendee'),
	('ChangeNotAttendeeDecision','ChangeNotAttendeeDecision'),
	('SignUpMemberToWaitlist','SignUpMemberToWaitlist'),
	('SignOffMemberFromWaitlist','SignOffMemberFromWaitlist'),
	('SetMeetingHostRole','SetMeetingHostRole'),
	('SetMeetingAttendeeRole','SetMeetingAttendeeRole'),
	('CancelMeeting','CancelMeeting'),
	('GetAllMeetingGroups','GetAllMeetingGroups'),
	('EditMeetingGroupGeneralAttributes','EditMeetingGroupGeneralAttributes'),
	('JoinToGroup','JoinToGroup'),
	('LeaveMeetingGroup','LeaveMeetingGroup'),
	('AddMeetingComment','AddMeetingComment'),
	('EditMeetingComment','EditMeetingComment'),
	('RemoveMeetingComment','RemoveMeetingComment'),
	('AddMeetingCommentReply','AddMeetingCommentReply'),
	('LikeMeetingComment','LikeMeetingComment'),
	('UnlikeMeetingComment','UnlikeMeetingComment'),
	('EnableMeetingCommenting','EnableMeetingCommenting'),
	('DisableMeetingCommenting','DisableMeetingCommenting'),
	('MyMeetingGroupsView','MyMeetingGroupsView'),
	('AllMeetingGroupsView','AllMeetingGroupsView'),
	('SubscriptionView','SubscriptionView'),
	('EmailsView','EmailsView'),
	('MyMeetingsView','MyMeetingsView'),
	('GetAuthenticatedMemberMeetings','GetAuthenticatedMemberMeetings'),

	-- Administration
	('AcceptMeetingGroupProposal','AcceptMeetingGroupProposal'),
	('AdministrationsView','AdministrationsView'),

	-- Payments
	('RegisterPayment','RegisterPayment'),
	('BuySubscription','BuySubscription'),
	('RenewSubscription','RenewSubscription'),
	('CreatePriceListItem','CreatePriceListItem'),
	('ActivatePriceListItem','ActivatePriceListItem'),
	('DeactivatePriceListItem','DeactivatePriceListItem'),
	('ChangePriceListItemAttributes','ChangePriceListItemAttributes'),
	('GetAuthenticatedPayerSubscription','GetAuthenticatedPayerSubscription'),
	('GetPriceListItem','GetPriceListItem')

-- Meetings
INSERT INTO users.RolesToPermissions VALUES ('Member', 'GetMeetingGroupProposals')
INSERT INTO users.RolesToPermissions VALUES ('Member', 'ProposeMeetingGroup')
INSERT INTO users.RolesToPermissions VALUES ('Member', 'CreateNewMeeting')
INSERT INTO users.RolesToPermissions VALUES ('Member', 'EditMeeting')
INSERT INTO users.RolesToPermissions VALUES ('Member', 'AddMeetingAttendee')
INSERT INTO users.RolesToPermissions VALUES ('Member', 'RemoveMeetingAttendee')
INSERT INTO users.RolesToPermissions VALUES ('Member', 'AddNotAttendee')
INSERT INTO users.RolesToPermissions VALUES ('Member', 'ChangeNotAttendeeDecision')
INSERT INTO users.RolesToPermissions VALUES ('Member', 'SignUpMemberToWaitlist')
INSERT INTO users.RolesToPermissions VALUES ('Member', 'SignOffMemberFromWaitlist')
INSERT INTO users.RolesToPermissions VALUES ('Member', 'SetMeetingHostRole')
INSERT INTO users.RolesToPermissions VALUES ('Member', 'SetMeetingAttendeeRole')
INSERT INTO users.RolesToPermissions VALUES ('Member', 'CancelMeeting')
INSERT INTO users.RolesToPermissions VALUES ('Member', 'GetAllMeetingGroups')
INSERT INTO users.RolesToPermissions VALUES ('Member', 'EditMeetingGroupGeneralAttributes')
INSERT INTO users.RolesToPermissions VALUES ('Member', 'JoinToGroup')
INSERT INTO users.RolesToPermissions VALUES ('Member', 'LeaveMeetingGroup')
INSERT INTO users.RolesToPermissions VALUES ('Member', 'AddMeetingComment')
INSERT INTO users.RolesToPermissions VALUES ('Member', 'EditMeetingComment')
INSERT INTO users.RolesToPermissions VALUES ('Member', 'RemoveMeetingComment')
INSERT INTO users.RolesToPermissions VALUES ('Member', 'AddMeetingCommentReply')
INSERT INTO users.RolesToPermissions VALUES ('Member', 'LikeMeetingComment')
INSERT INTO users.RolesToPermissions VALUES ('Member', 'UnlikeMeetingComment')
INSERT INTO users.RolesToPermissions VALUES ('Member', 'GetAuthenticatedMemberMeetingGroups')
INSERT INTO users.RolesToPermissions VALUES ('Member', 'GetMeetingGroupDetails')
INSERT INTO users.RolesToPermissions VALUES ('Member', 'GetMeetingDetails')
INSERT INTO users.RolesToPermissions VALUES ('Member', 'GetMeetingAttendees')
INSERT INTO users.RolesToPermissions VALUES ('Member', 'MyMeetingsGroupsView')
INSERT INTO users.RolesToPermissions VALUES ('Member', 'SubscriptionView')
INSERT INTO users.RolesToPermissions VALUES ('Member', 'EmailsView')
INSERT INTO users.RolesToPermissions VALUES ('Member', 'AllMeetingGroupsView')
INSERT INTO users.RolesToPermissions VALUES ('Member', 'MyMeetingsView')
INSERT INTO users.RolesToPermissions VALUES ('Member', 'GetAuthenticatedMemberMeetings')

-- Administration
INSERT INTO users.RolesToPermissions VALUES ('Administrator', 'AcceptMeetingGroupProposal')
INSERT INTO users.RolesToPermissions VALUES ('Administrator', 'AdministrationsView')

-- Payments
INSERT INTO users.RolesToPermissions VALUES ('Member', 'RegisterPayment')
INSERT INTO users.RolesToPermissions VALUES ('Member', 'BuySubscription')
INSERT INTO users.RolesToPermissions VALUES ('Member', 'RenewSubscription')
INSERT INTO users.RolesToPermissions VALUES ('Member', 'GetAuthenticatedPayerSubscription')
INSERT INTO users.RolesToPermissions VALUES ('Member', 'GetPriceListItem')
INSERT INTO users.RolesToPermissions VALUES ('Administrator', 'CreatePriceListItem')
INSERT INTO users.RolesToPermissions VALUES ('Administrator', 'ActivatePriceListItem')
INSERT INTO users.RolesToPermissions VALUES ('Administrator', 'DeactivatePriceListItem')
INSERT INTO users.RolesToPermissions VALUES ('Administrator', 'ChangePriceListItemAttributes')
INSERT INTO users.RolesToPermissions VALUES ('Administrator', 'GetPriceListItem')