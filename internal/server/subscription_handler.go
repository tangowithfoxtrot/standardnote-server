package server

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type subscription struct {
	filesServerUrl string
}

func (h *subscription) SubscriptionV1(c echo.Context) error {
	user := currentUser(c)

	// The official Standard Notes client has a race condition,
	// the features endpoint will only be called when delaying response...
	time.Sleep(1 * time.Second)

	return c.JSON(http.StatusOK, echo.Map{
		"meta": echo.Map{
			"auth": echo.Map{
				"userUuid": user.ID,
				"roles": []echo.Map{
					{
						"uuid": "8047edbb-a10a-4ff8-8d53-c2cae600a8e8",
						"name": "PRO_USER",
					},
					{
						"uuid": "8802d6a3-b97c-4b25-968a-8fb21c65c3a1",
						"name": "CORE_USER",
					},
				},
			},
			"server": echo.Map{
				"filesServerUrl": h.filesServerUrl,
			},
		},
		"data": echo.Map{
			"success": true,
			"user": echo.Map{
				"uuid":  user.ID,
				"email": user.Email,
			},
			"subscription": echo.Map{
				"uuid":             "d4a65722-4f02-11ed-b7e0-0242ac12000a",
				"planName":         "PRO_PLAN",
				"endsAt":           8640000000000000,
				"createdAt":        0,
				"updatedAt":        0,
				"cancelled":        0,
				"subscriptionId":   1,
				"subscriptionType": "",
			},
		},
	})
}

func (h *subscription) Features(c echo.Context) error {
	user := currentUser(c)

	return c.JSON(http.StatusOK, echo.Map{
		"meta": echo.Map{
			"auth": echo.Map{
				"userUuid": user.ID,
				"roles": []echo.Map{
					{
						"uuid": "8047edbb-a10a-4ff8-8d53-c2cae600a8e8",
						"name": "PRO_USER",
					},
					{
						"uuid": "8802d6a3-b97c-4b25-968a-8fb21c65c3a1",
						"name": "CORE_USER",
					},
				},
			},
			"server": echo.Map{
				"filesServerUrl": h.filesServerUrl,
			},
		},
		"data": echo.Map{
			"success":  true,
			"userUuid": user.ID,
			"features": []echo.Map{
				{
					"name":            "Focus",
					"identifier":      "org.standardnotes.theme-focus",
					"permission_name": "theme:focused",
					"version":         "1.2.6",
					"description":     "For when you need to go in.",
					"git_repo_url":    "https://github.com/standardnotes/focus-theme",
					"marketing_url":   "https://standardnotes.com/extensions/focused",
					"thumbnail_url":   "https://s3.amazonaws.com/standard-notes/screenshots/models/themes/focus-with-mobile.jpg",
					"dock_icon": echo.Map{
						"type":             "circle",
						"background_color": "#a464c2",
						"foreground_color": "#ffffff",
						"border_color":     "#a464c2",
					},
					"static_files": []string{
						"dist",
						"package.json",
					},
					"download_url": "https://github.com/standardnotes/focus-theme/releases/download/1.2.6/org.standardnotes.theme-focus.zip",
					"index_path":   "dist/dist.css",
					"content_type": "SN|Theme",
					"area":         "editor-editor",
					"expires_at":   8640000000000000,
					"role_name":    "PRO_USER",
				},
				{
					"name":       "Folders",
					"identifier": "org.standardnotes.folders",
					"version":    "1.3.8",
					"index_path": "index.html",
					"component_permissions": []echo.Map{
						{
							"name": "stream-items",
							"content_types": []string{
								"Tag",
								"SN|SmartTag",
							},
						},
					},
					"permission_name": "component:folders",
					"area":            "tags-list",
					"description":     "Create nested folders with easy drag and drop.",
					"git_repo_url":    "https://github.com/standardnotes/folders-component",
					"marketing_url":   "https://standardnotes.com/extensions/folders",
					"thumbnail_url":   "https://s3.amazonaws.com/standard-notes/screenshots/models/components/folders.jpg",
					"static_files": []string{
						"index.html",
						"dist",
						"package.json",
					},
					"download_url":    "https://github.com/standardnotes/folders-component/releases/download/1.3.8/org.standardnotes.folders.zip",
					"content_type":    "SN|Component",
					"interchangeable": true,
					"expires_at":      8640000000000000,
					"role_name":       "PRO_USER",
				},
				{
					"name":            "Focus Mode",
					"identifier":      "org.standardnotes.focus-mode",
					"permission_name": "app:focus-mode",
					"description":     "",
					"expires_at":      8640000000000000,
					"role_name":       "PRO_USER",
				},
				{
					"name":            "Secure Spreadsheets",
					"identifier":      "org.standardnotes.standard-sheets",
					"version":         "1.4.4",
					"note_type":       "spreadsheet",
					"file_type":       "json",
					"interchangeable": false,
					"permission_name": "editor:sheets",
					"description":     "A powerful spreadsheet editor with formatting and formula support. Not recommended for large data sets, as encryption of such data may decrease editor performance.",
					"marketing_url":   "",
					"git_repo_url":    "https://github.com/standardnotes/secure-spreadsheets",
					"thumbnail_url":   "https://s3.amazonaws.com/standard-notes/screenshots/models/editors/spreadsheets.png",
					"static_files": []string{
						"index.html",
						"dist",
						"package.json",
					},
					"download_url": "https://github.com/standardnotes/secure-spreadsheets/releases/download/1.4.4/org.standardnotes.standard-sheets.zip",
					"index_path":   "dist/index.html",
					"component_permissions": []echo.Map{
						{
							"name": "stream-context-item",
							"content_types": []string{
								"Note",
							},
						},
					},
					"content_type": "SN|Component",
					"area":         "editor-editor",
					"expires_at":   8640000000000000,
					"role_name":    "PRO_USER",
				},
				{
					"name":            "Autobiography",
					"identifier":      "org.standardnotes.theme-autobiography",
					"permission_name": "theme:autobiography",
					"version":         "1.0.2",
					"description":     "A theme for writers and readers.",
					"git_repo_url":    "https://github.com/standardnotes/autobiography-theme",
					"marketing_url":   "",
					"thumbnail_url":   "https://s3.amazonaws.com/standard-notes/screenshots/models/themes/autobiography.jpg",
					"dock_icon": echo.Map{
						"type":             "circle",
						"background_color": "#9D7441",
						"foreground_color": "#ECE4DB",
						"border_color":     "#9D7441",
					},
					"static_files": []string{
						"dist",
						"package.json",
					},
					"download_url": "https://github.com/standardnotes/autobiography-theme/releases/download/1.0.2/org.standardnotes.theme-autobiography.zip",
					"index_path":   "dist/dist.css",
					"content_type": "SN|Theme",
					"area":         "editor-editor",
					"expires_at":   8640000000000000,
					"role_name":    "PRO_USER",
				},
				{
					"name":              "Markdown Minimist",
					"identifier":        "org.standardnotes.minimal-markdown-editor",
					"note_type":         "markdown",
					"file_type":         "md",
					"index_path":        "index.html",
					"permission_name":   "editor:markdown-minimist",
					"version":           "1.3.9",
					"spellcheckControl": true,
					"description":       "A minimal Markdown editor with live rendering and in-text search via Ctrl/Cmd + F",
					"git_repo_url":      "https://github.com/standardnotes/minimal-markdown-editor",
					"marketing_url":     "https://standardnotes.com/extensions/minimal-markdown-editor",
					"thumbnail_url":     "https://s3.amazonaws.com/standard-notes/screenshots/models/editors/min-markdown.jpg",
					"static_files": []string{
						"index.html",
						"dist",
						"package.json",
					},
					"download_url": "https://github.com/standardnotes/minimal-markdown-editor/releases/download/1.3.9/org.standardnotes.minimal-markdown-editor.zip",
					"component_permissions": []echo.Map{
						{
							"name": "stream-context-item",
							"content_types": []string{
								"Note",
							},
						},
					},
					"content_type":    "SN|Component",
					"area":            "editor-editor",
					"interchangeable": true,
					"expires_at":      8640000000000000,
					"role_name":       "PRO_USER",
				},
				{
					"name":            "Tag Nesting",
					"identifier":      "org.standardnotes.tag-nesting",
					"permission_name": "app:tag-nesting",
					"description":     "Organize your tags into folders.",
					"expires_at":      8640000000000000,
					"role_name":       "PRO_USER",
				},
				{
					"name":            "Midnight",
					"identifier":      "org.standardnotes.theme-midnight",
					"permission_name": "theme:midnight",
					"version":         "1.2.5",
					"description":     "Elegant utilitarianism.",
					"git_repo_url":    "https://github.com/standardnotes/midnight-theme",
					"marketing_url":   "https://standardnotes.com/extensions/midnight",
					"thumbnail_url":   "https://s3.amazonaws.com/standard-notes/screenshots/models/themes/midnight-with-mobile.jpg",
					"dock_icon": echo.Map{
						"type":             "circle",
						"background_color": "#086DD6",
						"foreground_color": "#ffffff",
						"border_color":     "#086DD6",
					},
					"static_files": []string{
						"dist",
						"package.json",
					},
					"download_url": "https://github.com/standardnotes/midnight-theme/releases/download/1.2.5/org.standardnotes.theme-midnight.zip",
					"index_path":   "dist/dist.css",
					"content_type": "SN|Theme",
					"area":         "editor-editor",
					"expires_at":   8640000000000000,
					"role_name":    "PRO_USER",
				},
				{
					"name":            "Dynamic",
					"identifier":      "org.standardnotes.theme-dynamic",
					"permission_name": "theme:dynamic",
					"layerable":       true,
					"no_mobile":       true,
					"version":         "1.0.3",
					"description":     "A smart theme that minimizes the tags and notes panels when they are not in use.",
					"git_repo_url":    "https://github.com/standardnotes/dynamic-theme",
					"marketing_url":   "https://standardnotes.com/extensions/dynamic",
					"static_files": []string{
						"dist",
						"package.json",
					},
					"download_url": "https://github.com/standardnotes/dynamic-theme/releases/download/1.0.3/org.standardnotes.theme-dynamic.zip",
					"index_path":   "dist/dist.css",
					"content_type": "SN|Theme",
					"area":         "editor-editor",
					"expires_at":   8640000000000000,
					"role_name":    "PRO_USER",
				},
				{
					"name":              "Markdown Math",
					"identifier":        "org.standardnotes.fancy-markdown-editor",
					"version":           "1.3.6",
					"spellcheckControl": true,
					"permission_name":   "editor:markdown-math",
					"note_type":         "markdown",
					"file_type":         "md",
					"index_path":        "index.html",
					"description":       "A beautiful split-pane Markdown editor with synced-scroll, LaTeX support, and colorful syntax.",
					"git_repo_url":      "https://github.com/standardnotes/math-editor",
					"marketing_url":     "https://standardnotes.com/extensions/math-editor",
					"thumbnail_url":     "https://s3.amazonaws.com/standard-notes/screenshots/models/editors/fancy-markdown.jpg",
					"static_files": []string{
						"index.html",
						"dist",
						"package.json",
					},
					"download_url": "https://github.com/standardnotes/math-editor/releases/download/1.3.6/org.standardnotes.fancy-markdown-editor.zip",
					"component_permissions": []echo.Map{
						{
							"name": "stream-context-item",
							"content_types": []string{
								"Note",
							},
						},
					},
					"content_type":    "SN|Component",
					"area":            "editor-editor",
					"interchangeable": true,
					"expires_at":      8640000000000000,
					"role_name":       "PRO_USER",
				},
				{
					"identifier":      "org.standardnotes.files-25-gb",
					"permission_name": "server:files-25-gb",
					"expires_at":      8640000000000000,
					"role_name":       "PRO_USER",
				},
				{
					"name":            "Solarized Dark",
					"identifier":      "org.standardnotes.theme-solarized-dark",
					"permission_name": "theme:solarized-dark",
					"version":         "1.2.4",
					"description":     "The perfect theme for any time.",
					"git_repo_url":    "https://github.com/standardnotes/solarized-dark-theme",
					"marketing_url":   "https://standardnotes.com/extensions/solarized-dark",
					"thumbnail_url":   "https://s3.amazonaws.com/standard-notes/screenshots/models/themes/solarized-dark.jpg",
					"dock_icon": echo.Map{
						"type":             "circle",
						"background_color": "#2AA198",
						"foreground_color": "#ffffff",
						"border_color":     "#2AA198",
					},
					"static_files": []string{
						"dist",
						"package.json",
					},
					"download_url": "https://github.com/standardnotes/solarized-dark-theme/releases/download/1.2.4/org.standardnotes.theme-solarized-dark.zip",
					"index_path":   "dist/dist.css",
					"content_type": "SN|Theme",
					"area":         "editor-editor",
					"expires_at":   8640000000000000,
					"role_name":    "PRO_USER",
				},
				{
					"name":              "Markdown Pro",
					"identifier":        "org.standardnotes.advanced-markdown-editor",
					"version":           "1.4.2",
					"note_type":         "markdown",
					"file_type":         "md",
					"permission_name":   "editor:markdown-pro",
					"spellcheckControl": true,
					"description":       "A fully featured Markdown editor that supports live preview, a styling toolbar, and split pane support.",
					"git_repo_url":      "https://github.com/standardnotes/advanced-markdown-editor",
					"marketing_url":     "https://standardnotes.com/extensions/advanced-markdown",
					"thumbnail_url":     "https://s3.amazonaws.com/standard-notes/screenshots/models/editors/adv-markdown.jpg",
					"static_files": []string{
						"index.html",
						"dist",
						"package.json",
					},
					"download_url": "https://github.com/standardnotes/advanced-markdown-editor/releases/download/1.4.2/org.standardnotes.advanced-markdown-editor.zip",
					"index_path":   "dist/index.html",
					"component_permissions": []echo.Map{
						{
							"name": "stream-context-item",
							"content_types": []string{
								"Note",
							},
						},
					},
					"content_type":    "SN|Component",
					"area":            "editor-editor",
					"interchangeable": true,
					"expires_at":      8640000000000000,
					"role_name":       "PRO_USER",
				},
				{
					"name":            "Futura",
					"identifier":      "org.standardnotes.theme-futura",
					"permission_name": "theme:futura",
					"version":         "1.2.5",
					"description":     "Calm and relaxed. Take some time off.",
					"git_repo_url":    "https://github.com/standardnotes/futura-theme",
					"marketing_url":   "https://standardnotes.com/extensions/futura",
					"thumbnail_url":   "https://s3.amazonaws.com/standard-notes/screenshots/models/themes/futura-with-mobile.jpg",
					"dock_icon": echo.Map{
						"type":             "circle",
						"background_color": "#fca429",
						"foreground_color": "#ffffff",
						"border_color":     "#fca429",
					},
					"static_files": []string{
						"dist",
						"package.json",
					},
					"download_url": "https://github.com/standardnotes/futura-theme/releases/download/1.2.5/org.standardnotes.theme-futura.zip",
					"index_path":   "dist/dist.css",
					"content_type": "SN|Theme",
					"area":         "editor-editor",
					"expires_at":   8640000000000000,
					"role_name":    "PRO_USER",
				},
				{
					"name":            "Smart Filters",
					"identifier":      "org.standardnotes.smart-filters",
					"permission_name": "app:smart-filters",
					"description":     "Create smart filters for viewing notes matching specific criteria.",
					"expires_at":      8640000000000000,
					"role_name":       "PRO_USER",
				},
				{
					"name":            "Listed Custom Domain",
					"identifier":      "org.standardnotes.listed-custom-domain",
					"permission_name": "listed:custom-domain",
					"description":     "",
					"expires_at":      8640000000000000,
					"role_name":       "PRO_USER",
				},
				{
					"name":            "TokenVault",
					"note_type":       "authentication",
					"file_type":       "json",
					"interchangeable": false,
					"identifier":      "org.standardnotes.token-vault",
					"permission_name": "editor:token-vault",
					"version":         "2.0.9",
					"description":     "Encrypt and protect your 2FA secrets for all your internet accounts. TokenVault handles your 2FA secrets so that you never lose them again, or have to start over when you get a new device.",
					"marketing_url":   "",
					"git_repo_url":    "https://github.com/standardnotes/token-vault",
					"thumbnail_url":   "https://standard-notes.s3.amazonaws.com/screenshots/models/editors/token-vault.png",
					"static_files": []string{
						"index.html",
						"dist",
						"package.json",
					},
					"download_url": "https://github.com/standardnotes/token-vault/releases/download/2.0.9/org.standardnotes.token-vault.zip",
					"index_path":   "dist/index.html",
					"component_permissions": []echo.Map{
						{
							"name": "stream-context-item",
							"content_types": []string{
								"Note",
							},
						},
					},
					"content_type": "SN|Component",
					"area":         "editor-editor",
					"expires_at":   8640000000000000,
					"role_name":    "PRO_USER",
				},
				{
					"name":            "Encrypted files (coming soon)",
					"identifier":      "org.standardnotes.files",
					"permission_name": "app:files",
					"description":     "",
					"expires_at":      8640000000000000,
					"role_name":       "PRO_USER",
				},
				{
					"name":            "Titanium",
					"version":         "1.2.5",
					"identifier":      "org.standardnotes.theme-titanium",
					"permission_name": "theme:titanium",
					"description":     "Light on the eyes, heavy on the spirit.",
					"git_repo_url":    "https://github.com/standardnotes/titanium-theme",
					"marketing_url":   "https://standardnotes.com/extensions/titanium",
					"thumbnail_url":   "https://s3.amazonaws.com/standard-notes/screenshots/models/themes/titanium-with-mobile.jpg",
					"dock_icon": echo.Map{
						"type":             "circle",
						"background_color": "#6e2b9e",
						"foreground_color": "#ffffff",
						"border_color":     "#6e2b9e",
					},
					"static_files": []string{
						"dist",
						"package.json",
					},
					"download_url": "https://github.com/standardnotes/titanium-theme/releases/download/1.2.5/org.standardnotes.theme-titanium.zip",
					"index_path":   "dist/dist.css",
					"content_type": "SN|Theme",
					"area":         "editor-editor",
					"expires_at":   8640000000000000,
					"role_name":    "PRO_USER",
				},
				{
					"identifier":      "org.standardnotes.daily-gdrive-backup",
					"permission_name": "server:daily-gdrive-backup",
					"expires_at":      8640000000000000,
					"role_name":       "PRO_USER",
				},
				{
					"identifier":      "org.standardnotes.daily-onedrive-backup",
					"permission_name": "server:daily-onedrive-backup",
					"expires_at":      8640000000000000,
					"role_name":       "PRO_USER",
				},
				{
					"name":            "Two factor authentication",
					"identifier":      "org.standardnotes.two-factor-auth",
					"permission_name": "server:two-factor-auth",
					"expires_at":      8640000000000000,
					"role_name":       "PRO_USER",
				},
				{
					"name":              "Plus Editor",
					"note_type":         "rich-text",
					"file_type":         "html",
					"identifier":        "org.standardnotes.plus-editor",
					"permission_name":   "editor:plus",
					"version":           "1.6.1",
					"spellcheckControl": true,
					"description":       "From highlighting to custom font sizes and colors, to tables and lists, this editor is perfect for crafting any document.",
					"git_repo_url":      "https://github.com/standardnotes/plus-editor",
					"marketing_url":     "https://standardnotes.com/extensions/plus-editor",
					"thumbnail_url":     "https://s3.amazonaws.com/standard-notes/screenshots/models/editors/plus-editor.jpg",
					"static_files": []string{
						"index.html",
						"dist",
						"package.json",
					},
					"download_url": "https://github.com/standardnotes/plus-editor/releases/download/1.6.1/org.standardnotes.plus-editor.zip",
					"index_path":   "dist/index.html",
					"component_permissions": []echo.Map{
						{
							"name": "stream-context-item",
							"content_types": []string{
								"Note",
							},
						},
					},
					"content_type":    "SN|Component",
					"area":            "editor-editor",
					"interchangeable": true,
					"expires_at":      8640000000000000,
					"role_name":       "PRO_USER",
				},
				{
					"name":              "Task Editor",
					"identifier":        "org.standardnotes.simple-task-editor",
					"note_type":         "task",
					"version":           "1.3.9",
					"spellcheckControl": true,
					"file_type":         "md",
					"interchangeable":   false,
					"permission_name":   "editor:task-editor",
					"description":       "A great way to manage short-term and long-term to-do\"s. You can mark tasks as completed, change their order, and edit the text naturally in place.",
					"git_repo_url":      "https://github.com/standardnotes/simple-task-editor",
					"marketing_url":     "https://standardnotes.com/extensions/simple-task-editor",
					"thumbnail_url":     "https://s3.amazonaws.com/standard-notes/screenshots/models/editors/task-editor.jpg",
					"static_files": []string{
						"index.html",
						"dist",
						"package.json",
					},
					"download_url": "https://github.com/standardnotes/simple-task-editor/releases/download/1.3.9/org.standardnotes.simple-task-editor.zip",
					"index_path":   "dist/index.html",
					"component_permissions": []echo.Map{
						{
							"name": "stream-context-item",
							"content_types": []string{
								"Note",
							},
						},
					},
					"content_type": "SN|Component",
					"area":         "editor-editor",
					"expires_at":   8640000000000000,
					"role_name":    "PRO_USER",
				},
				{
					"name":            "Unlimited note history",
					"identifier":      "org.standardnotes.note-history-unlimited",
					"permission_name": "server:note-history-unlimited",
					"expires_at":      8640000000000000,
					"role_name":       "PRO_USER",
				},
				{
					"identifier":      "org.standardnotes.daily-dropbox-backup",
					"permission_name": "server:daily-dropbox-backup",
					"expires_at":      8640000000000000,
					"role_name":       "PRO_USER",
				},
				{
					"name":              "Code Editor",
					"version":           "1.3.10",
					"spellcheckControl": true,
					"identifier":        "org.standardnotes.code-editor",
					"permission_name":   "editor:code-editor",
					"note_type":         "code",
					"file_type":         "txt",
					"interchangeable":   true,
					"index_path":        "index.html",
					"static_files": []string{
						"index.html",
						"dist",
						"package.json",
						"vendor",
					},
					"description":   "Syntax highlighting and convenient keyboard shortcuts for over 120 programminglanguages. Ideal for code snippets and procedures.",
					"git_repo_url":  "https://github.com/standardnotes/code-editor",
					"marketing_url": "https://standardnotes.com/extensions/code-editor",
					"thumbnail_url": "https://s3.amazonaws.com/standard-notes/screenshots/models/editors/code.jpg",
					"download_url":  "https://github.com/standardnotes/code-editor/releases/download/1.3.10/org.standardnotes.code-editor.zip",
					"component_permissions": []echo.Map{
						{
							"name": "stream-context-item",
							"content_types": []string{
								"Note",
							},
						},
					},
					"content_type": "SN|Component",
					"area":         "editor-editor",
					"expires_at":   8640000000000000,
					"role_name":    "PRO_USER",
				},
				{
					"name":       "Bold Editor",
					"identifier": "org.standardnotes.bold-editor",
					"version":    "1.3.2",
					"note_type":  "rich-text",
					"file_type":  "html",
					"component_permissions": []echo.Map{
						{
							"name": "stream-context-item",
							"content_types": []string{
								"Note",
							},
						},
						{
							"name": "stream-items",
							"content_types": []string{
								"SN|FileSafe|Credentials",
								"SN|FileSafe|FileMetadata",
								"SN|FileSafe|Integration",
							},
						},
					},
					"spellcheckControl": true,
					"permission_name":   "editor:bold",
					"description":       "A simple and peaceful rich editor that helps you write and think clearly. Features FileSafe integration, so you can embed your encrypted images, videos, and audio recordings directly inline.",
					"marketing_url":     "",
					"thumbnail_url":     "https://s3.amazonaws.com/standard-notes/screenshots/models/editors/bold.jpg",
					"git_repo_url":      "https://github.com/standardnotes/bold-editor",
					"static_files": []string{
						"index.html",
						"dist",
						"package.json",
					},
					"download_url":    "https://github.com/standardnotes/bold-editor/releases/download/1.3.2/org.standardnotes.bold-editor.zip",
					"index_path":      "dist/index.html",
					"content_type":    "SN|Component",
					"area":            "editor-editor",
					"interchangeable": true,
					"expires_at":      8640000000000000,
					"role_name":       "PRO_USER",
				},
				{
					"name":              "Markdown Basic",
					"identifier":        "org.standardnotes.simple-markdown-editor",
					"note_type":         "markdown",
					"version":           "1.4.2",
					"spellcheckControl": true,
					"file_type":         "md",
					"permission_name":   "editor:markdown-basic",
					"description":       "A Markdown editor with dynamic split-pane preview.",
					"git_repo_url":      "https://github.com/standardnotes/markdown-basic",
					"marketing_url":     "https://standardnotes.com/extensions/simple-markdown-editor",
					"thumbnail_url":     "https://s3.amazonaws.com/standard-notes/screenshots/models/editors/simple-markdown.jpg",
					"static_files": []string{
						"index.html",
						"dist",
						"package.json",
					},
					"download_url": "https://github.com/standardnotes/markdown-basic/releases/download/1.4.2/org.standardnotes.simple-markdown-editor.zip",
					"index_path":   "dist/index.html",
					"component_permissions": []echo.Map{
						{
							"name": "stream-context-item",
							"content_types": []string{
								"Note",
							},
						},
					},
					"content_type":    "SN|Component",
					"area":            "editor-editor",
					"interchangeable": true,
					"expires_at":      8640000000000000,
					"role_name":       "PRO_USER",
				},
				{
					"name":            "Email backups",
					"identifier":      "org.standardnotes.daily-email-backup",
					"permission_name": "server:daily-email-backup",
					"expires_at":      8640000000000000,
					"role_name":       "PRO_USER",
				},
			},
		},
	})
}
