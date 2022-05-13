# Article
class Article(object):
    """
    init(
        id: id
        user_id: 用户id
        content: 正文
        create_time: 创建时间
        update_time: 更新时间
    )
    """
    def __init__(self, id, user_id, content, create_time, update_time):
        self.id = id
        self.user_id = user_id
        self.content = content
        self.create_time = create_time
        self.update_time = update_time

# TEST
article = Article(1, 2, "hello world", "2022-05-01 12:00:00", "2022-05-01 12:00:00")
article.content = "just test"
print(article.user_id)
print(article.content)
