-- =========================================
-- Create table template
-- =========================================
USE dbtest
GO

IF OBJECT_ID('dbo.Users', 'U') IS NOT NULL
  DROP TABLE dbo.Users
GO

CREATE TABLE dbo.Users
(
	Id int not null identity(1,1),
	[Name] varchar(250) not null,
	Age  int not null,
	CreateAt datetime not null,
	UpdateAt datetime null,
	DelteAt datetime null
    CONSTRAINT PK_Users_ID PRIMARY KEY (Id)
)
GO


insert into users select 'Juan Perez',25,GETDATE(),null,null;

insert into users select 'Gian Perez',25,GETDATE(),null,null;

insert into users select 'Juan Mauel Perez',25,GETDATE(),null,null;

select * from Users

