package publishing

func formatGetArticleResponse(res GetArticleResponseModel) map[string]interface{} {
	return map[string]interface{}{
		"data": map[string]interface{}{
			"article": map[string]interface{}{
				"id":    res.Article.ID,
				"title": res.Article.Title,
				"text":  res.Article.Text,
			},
		},
	}
}

func formatCreateArticleResponse(res CreateArticleResponseModel) map[string]interface{} {
	return map[string]interface{}{
		"data": map[string]interface{}{
			"id": res.ID,
		},
	}
}
