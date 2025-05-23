function fetchAndDisplayArticle() {
    // 显示加载动画
    document.getElementById('loading').style.display = 'block';
    document.getElementById('article-details').style.display = 'none';
    document.getElementById('error-message').style.display = 'none';

    console.log('开始获取随机物品...');

    // 使用绝对路径访问API
    fetch(window.location.origin + '/random-article')
        .then(response => response.json())
        .then(article => {
            // 显示物品的空间信息
            console.log('Article space:', article.space.width + 'x' + article.space.height);
            document.getElementById('space').innerText = `空间: ${article.space.width}x${article.space.height}`;

            // 根据物品的品质来决定延迟时间
            let delay;
            switch (article.quality) {
                case "红": // Red
                    delay = 2000;
                    break;
                case "金": // Gold
                    delay = 1500;
                    break;
                case "紫": // Purple
                    delay = 1000;
                    break;
                case "蓝": // Blue
                    delay = 500;
                    break;
                case "绿": // Green
                    delay = 300;
                    break;
                case "白": // White
                    delay = 200;
                    break;
                default:
                    delay = 0;
            }

            // 延迟显示物品名称和其他信息
            setTimeout(() => {
                document.getElementById('loading').style.display = 'none';
                document.getElementById('article-details').style.display = 'block';
                displayArticleDetails(article);
            }, delay);
        })
        .catch(error => {
            console.error('Error fetching article:', error);
            document.getElementById('loading').style.display = 'none';
            document.getElementById('error-message').innerText = '获取物品失败，请重试';
            document.getElementById('error-message').style.display = 'block';
        });
}

function displayArticleDetails(article) {
    // 显示物品的详细信息
    document.getElementById('name').innerText = `名称: ${article.name}`;

    // 设置品质文本和颜色
    const qualityElement = document.getElementById('quality');
    qualityElement.innerText = `品质: ${article.quality}`;
    qualityElement.style.color = article.color;

    document.getElementById('price').innerText = `价格: ${article.price}`;
    document.getElementById('type').innerText = `类型: ${article.type}`;

    // 添加动画效果
    const detailsElement = document.getElementById('article-details');
    detailsElement.classList.add('reveal-animation');
    setTimeout(() => {
        detailsElement.classList.remove('reveal-animation');
    }, 1000);
}

// 当页面加载完成时设置事件监听器
document.addEventListener('DOMContentLoaded', function () {
    // 当用户点击搜索按钮时调用fetchAndDisplayArticle函数
    document.getElementById('search-button').addEventListener('click', fetchAndDisplayArticle);
});