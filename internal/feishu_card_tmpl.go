package internal

func PushFeishuCardTmpl() string {
	return `
		{
		  "config": {
			"wide_screen_mode": true
		  },
		  "elements": [
			{
			  "tag": "div",
			  "fields": [
				{
				  "is_short": true,
				  "text": {
					"tag": "lark_md",
					"content": "**项目：**\n{{.projectName}}"
				  }
				},
				{
				  "is_short": true,
				  "text": {
					"tag": "lark_md",
					"content": "**推送用户：**\n{{.userName}}"
				  }
				}
			  ]
			},
			{
			  "tag": "div",
			  "fields": [
				{
				  "is_short": true,
				  "text": {
					"tag": "lark_md",
					"content": "**Ref：**\n{{.ref}}"
				  }
				},
				{
				  "is_short": true,
				  "text": {
					"tag": "lark_md",
					"content": "**项目地址：**\n{{.webUrl}}"
				  }
				}
			  ]
			},
			{
			  "tag": "hr"
			},
			{
			  "tag": "note",
			  "elements": [{
				"content": "{{.commit}}",
				"tag": "plain_text"
			  }]
			}
		  ],
		  "header": {
			"template": "{{.headerColor}}",
			"title": {
			  "content": "{{.title}}",
			  "tag": "plain_text"
			}
		  }
		}
		`
}

func MergeRequestFeishuCardTmpl() string {
	return `{
		  "config": {
			"wide_screen_mode": true
		  },
		  "elements": [
			{
			  "tag": "div",
			  "fields": [
				{
				  "is_short": true,
				  "text": {
					"tag": "lark_md",
					"content": "**项目：**\n{{.projectName}}"
				  }
				},
				{
				  "is_short": true,
				  "text": {
					"tag": "lark_md",
					"content": "**合并请求用户：**\n{{.userName}}"
				  }
				}
			  ]
			},
			{
			  "tag": "div",
			  "fields": [
				{
				  "is_short": true,
				  "text": {
					"tag": "lark_md",
					"content": "**source branch：**\n{{.sourceBranch}}"
				  }
				},
				{
				  "is_short": true,
				  "text": {
					"tag": "lark_md",
					"content": "**target branch：**\n{{.targetBranch}}"
				  }
				}
			  ]
			},
			{
			  "tag": "div",
			  "fields": [
				{
				  "is_short": true,
				  "text": {
					"tag": "lark_md",
					"content": "**项目地址：**\n{{.webUrl}}"
				  }
				},
				{
				  "is_short": true,
				  "text": {
					"tag": "lark_md",
					"content": " "
				  }
				}
			  ]
			},
			{
				"tag": "div",
				"fields": [
				  {
					"is_short": true,
					"text": {
					  "tag": "lark_md",
					  "content": "**合并标题：**{{.mergeTitle}}"
					}
				  },
				  {
					"is_short": true,
					"text": {
					  "tag": "lark_md",
					  "content": " "
					}
				  }
				]
			},
			{
				"tag": "div",
				"text": {
					"content":"<at id=7337884226983575556>自增</at><at id=7337660809616687106>永军</at><at id=7337660809616687106>永军</at><at id=7337661571016949762>明杰</at>",
					"tag":"lark_md"
				}
			}
		  ],
		  "header": {
			"template": "{{.headerColor}}",
			"title": {
			  "content": "{{.title}}",
			  "tag": "plain_text"
			}
		  }
		}`
}
