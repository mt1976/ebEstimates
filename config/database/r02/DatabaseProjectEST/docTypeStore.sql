CREATE TABLE [dbo].[docTypeStore] (
    [_id]          INT            IDENTITY (1, 1) NOT NULL,
    [docTypeID]    NVARCHAR (MAX) NOT NULL,
    [code]         NVARCHAR (MAX) NOT NULL,
    [name]         NVARCHAR (MAX) NOT NULL,
    [_created]     NVARCHAR (MAX) NULL,
    [_createdBy]   NVARCHAR (MAX) NULL,
    [_createdHost] NVARCHAR (MAX) NULL,
    [_updated]     NVARCHAR (MAX) NULL,
    [_updatedBy]   NVARCHAR (MAX) NULL,
    [_updatedHost] NVARCHAR (MAX) NULL,
    [_deleted]     NVARCHAR (MAX) NULL,
    [_deletedBy]   NVARCHAR (MAX) NULL,
    [_deletedHost] NVARCHAR (MAX) NULL,
    [_dbVersion]   NVARCHAR (MAX) NULL,
    [comments]     NVARCHAR (MAX) NULL,
    CONSTRAINT [PK_docTypeStore] PRIMARY KEY CLUSTERED ([_id] ASC)
);


GO

EXECUTE sp_addextendedproperty @name = N'MS_Description', @value = 'system field', @level0type = N'SCHEMA', @level0name = N'dbo', @level1type = N'TABLE', @level1name = N'docTypeStore', @level2type = N'COLUMN', @level2name = N'_deletedBy';


GO

EXECUTE sp_addextendedproperty @name = N'MS_Description', @value = 'system field', @level0type = N'SCHEMA', @level0name = N'dbo', @level1type = N'TABLE', @level1name = N'docTypeStore', @level2type = N'COLUMN', @level2name = N'_updatedBy';


GO

EXECUTE sp_addextendedproperty @name = N'MS_Description', @value = 'Document Types "RSC/NC"', @level0type = N'SCHEMA', @level0name = N'dbo', @level1type = N'TABLE', @level1name = N'docTypeStore';


GO

EXECUTE sp_addextendedproperty @name = N'MS_Description', @value = 'system field', @level0type = N'SCHEMA', @level0name = N'dbo', @level1type = N'TABLE', @level1name = N'docTypeStore', @level2type = N'COLUMN', @level2name = N'_createdBy';


GO

EXECUTE sp_addextendedproperty @name = N'MS_Description', @value = 'system field', @level0type = N'SCHEMA', @level0name = N'dbo', @level1type = N'TABLE', @level1name = N'docTypeStore', @level2type = N'COLUMN', @level2name = N'_deleted';


GO

EXECUTE sp_addextendedproperty @name = N'MS_Description', @value = 'system field', @level0type = N'SCHEMA', @level0name = N'dbo', @level1type = N'TABLE', @level1name = N'docTypeStore', @level2type = N'COLUMN', @level2name = N'_updated';


GO

EXECUTE sp_addextendedproperty @name = N'MS_Description', @value = 'system field', @level0type = N'SCHEMA', @level0name = N'dbo', @level1type = N'TABLE', @level1name = N'docTypeStore', @level2type = N'COLUMN', @level2name = N'_deletedHost';


GO

EXECUTE sp_addextendedproperty @name = N'MS_Description', @value = 'system field', @level0type = N'SCHEMA', @level0name = N'dbo', @level1type = N'TABLE', @level1name = N'docTypeStore', @level2type = N'COLUMN', @level2name = N'_created';


GO

EXECUTE sp_addextendedproperty @name = N'MS_Description', @value = 'system field', @level0type = N'SCHEMA', @level0name = N'dbo', @level1type = N'TABLE', @level1name = N'docTypeStore', @level2type = N'COLUMN', @level2name = N'_updatedHost';


GO

EXECUTE sp_addextendedproperty @name = N'MS_Description', @value = 'system field', @level0type = N'SCHEMA', @level0name = N'dbo', @level1type = N'TABLE', @level1name = N'docTypeStore', @level2type = N'COLUMN', @level2name = N'_createdHost';


GO

EXECUTE sp_addextendedproperty @name = N'MS_Description', @value = 'system field', @level0type = N'SCHEMA', @level0name = N'dbo', @level1type = N'TABLE', @level1name = N'docTypeStore', @level2type = N'COLUMN', @level2name = N'_id';


GO

