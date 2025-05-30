const WEEKDAYS = ['周日', '周一', '周二', '周三', '周四', '周五', '周六'];
const breakfastPlan = {
    '周一': {
        main: '蒸红薯 1 个',
        protein: '蒸鸡蛋 2 个',
        veggie: '水黄瓜 1 根',
        comment: '黄瓜切片，使用橄榄油，酱油、醋将黄瓜和鸡蛋一起搅拌，避免蛋黄浪费，口感更佳，醋可以调理胃酸，避免血糖峰值产生，按黄瓜->蛋->红薯饮食顺序'
    },
    '周二': {
        main: '蒸玉米 1/2 根',
        protein: '蒸鸡蛋 2 个',
        veggie: '苹果 1/2 个 + 坚果',
        comment: '加点干果（巴旦木、核桃、腰果、红色那个果子）'
    },
    '周三': {
        main: '蒸土豆 1 个',
        protein: '蒸牛肉 150g',
        veggie: '小番茄',
        comment: '牛肉切片，使用小苏打处理，然后洗干净，加点酱油、料酒、橄榄油直接蒸熟'
    },
    '周四': {
        main: '蒸红薯 1 个',
        protein: '蒸鸡蛋 2 个',
        veggie: '西兰花',
        comment: '焯水，使用橄榄油，酱油、醋将西兰花和鸡蛋一起搅拌，避免蛋黄浪费，口感更佳，醋可以调理胃酸，避免血糖峰值产生，按蔬菜->蛋->红薯饮食顺序'
    },
    '周五': {
        main: '蒸玉米 1/2 根',
        protein: '蒸鸡蛋 2 个',
        veggie: '水黄瓜 1 根',
        comment: '黄瓜切片，使用橄榄油，酱油、醋将黄瓜和鸡蛋一起搅拌，避免蛋黄浪费，口感更佳，醋可以调理胃酸，避免血糖峰值产生，按黄瓜->蛋->红薯饮食顺序'
    },
    '周六': {main: '拌面 1 份', protein: '蒸鸡蛋 1 个', veggie: '橙子/苹果', comment: '灵活调整'},
    '周日': {
        main: '手擀面条',
        protein: '牛肉/虾',
        veggie: '西红柿/青菜（灵活）',
        comment: '水煮菜，加西红柿、青菜、蛋白质、面条一锅煮。根据前一日摄入与体重浮动灵活调整'
    }
};

const TodayBreakfast = () => {
    const today = new Date();
    const weekday = WEEKDAYS[today.getDay()];
    const plan = breakfastPlan[weekday];

    return (
        <div className="mb-4 bg-white p-4 rounded-xl shadow">
            <h2 className="font-semibold text-lg mb-2">🍠 今日早餐推荐</h2>
            <table className="w-full border text-sm">
                <thead>
                <tr className="bg-gray-100 text-center">
                    <th className="border px-2 py-1">时间</th>
                    <th className="border px-2 py-1">主食</th>
                    <th className="border px-2 py-1">蛋白质</th>
                    <th className="border px-2 py-1">蔬果</th>
                </tr>
                </thead>
                <tbody>
                <tr className="text-center">
                    <td className="border px-2 py-1 font-semibold text-blue-500">{weekday}</td>
                    <td className="border px-2 py-1">{plan.main}</td>
                    <td className="border px-2 py-1">{plan.protein}</td>
                    <td className="border px-2 py-1">{plan.veggie}</td>
                </tr>

                </tbody>
            </table>
            <div className="text-center">
                <p className="border px-2 py-1 bg-gradient-to-r from-blue-100 to-green-100">{plan.comment}</p>
            </div>
        </div>
    );
};

export default TodayBreakfast;