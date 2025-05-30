const WEEKDAYS = ['周日', '周一', '周二', '周三', '周四', '周五', '周六'];
const exercisePlan = {
    '周一': '慢跑 40 min',
    '周二': '慢跑 40 min',
    '周三': '跑休',
    '周四': '配速跑 40 min',
    '周五': '慢跑 40 min',
    '周六': '长跑 60 min / 10km',
    '周日': '跑休'
};

const TodayExercise = () => {
    const today = new Date();
    const weekday = WEEKDAYS[today.getDay()];
    const plan = exercisePlan[weekday];

    return (
        <div className="mb-4 bg-white p-4 rounded-xl shadow bg-gradient-to-r from-blue-100 to-green-100">
            <h2 className="font-semibold text-lg mb-2">🏃 今日运动安排</h2>
            <table className="w-full border text-sm">
                <thead>
                <tr className="bg-gray-100 text-center">
                    <th className="border px-2 py-1">时间</th>
                    <th className="border px-2 py-1">锻炼内容</th>
                </tr>
                </thead>
                <tbody>
                <tr className="text-center">
                    <td className="border px-2 py-1 font-semibold text-blue-500">{weekday}</td>
                    <td className="border px-2 py-1">{plan}</td>
                </tr>
                </tbody>
            </table>
        </div>
    );
};

export default TodayExercise;