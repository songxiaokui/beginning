export default function CheckinForm({ checkin, setCheckin }) {
    const handleChange = (field: string) => {
        setCheckin({ ...checkin, [field]: !checkin[field] });
    };

    //@ts-ignore
    const checkedItems = Object.entries(checkin).filter(([_, value]) => value);
    //@ts-ignore
    const uncheckedItems = Object.entries(checkin).filter(([_, value]) => !value);

    const labelMap: Record<string, string> = {
        breakfast: "🥚 营养早餐",
        lunch: "🍱 午餐清淡",
        dinner: "🍲 低脂晚餐",
        exercise: "🏃 空腹有氧",
        sleep: "🌅 早睡早起",
    };

    const renderItem = (key: string, checked: boolean) => (
        <div
            key={key}
            onClick={() => handleChange(key)}
            className={`cursor-pointer p-3 rounded-xl shadow transition-all duration-200 flex items-center justify-between
                ${checked ? "bg-green-100 hover:bg-green-200" : "bg-gray-100 hover:bg-gray-200"}
            `}
        >
            <span className="text-gray-800 font-medium">{labelMap[key]}</span>
            <input
                type="checkbox"
                checked={checked}
                readOnly
                className="h-5 w-5 text-blue-500 focus:ring-blue-400 border-gray-300 rounded"
            />
        </div>
    );

    return (
        <div className="space-y-4 mb-4 p-4 rounded-xl bg-white shadow-md bg-gradient-to-r from-blue-50 to-green-50">
            {checkedItems.length > 0 && (
                <div className="space-y-2">
                    <h2 className="text-lg font-bold text-green-700">✅ 今日已打卡</h2>
                    {checkedItems.map(([key]) => renderItem(key, true))}
                </div>
            )}
            {uncheckedItems.length > 0 && (
                <div className="space-y-2 mt-4">
                    <h2 className="text-lg font-bold text-gray-700">🕒 今日待打卡</h2>
                    {uncheckedItems.map(([key]) => renderItem(key, false))}
                </div>
            )}
        </div>
    );
}
