export default function CheckinForm({checkin, setCheckin}) {
    const handleChange = (field: string) => {
        setCheckin({...checkin, [field]: !checkin[field]});
    };

    // @ts-ignore
    const checkedItems = Object.entries(checkin).filter(([_, value]) => value);
    // @ts-ignore
    const uncheckedItems = Object.entries(checkin).filter(([_, value]) => !value);

    const labelMap: Record<string, string> = {
        breakfast: '营养早餐',
        lunch: '午餐',
        dinner: '低脂晚餐',
        exercise: '空腹有氧',
        sleep: '早起',
    };

    return (
        <div className="space-y-4 mb-4 bg-white p-4 rounded-xl shadow bg-gradient-to-r from-blue-100 to-green-100">
            {checkedItems.length > 0 && (
                <div>
                    <h2 className="text-lg font-semibold mb-2">📋 今日已打卡</h2>
                    {checkedItems.map(([key], i) => (
                        <label key={i} className="flex items-center space-x-2 text-gray-800">
                            <input
                                type="checkbox"
                                className="h-5 w-5 text-blue-500 focus:ring-blue-400 border-gray-300 rounded"
                                checked={checkin[key as keyof typeof checkin]}
                                onChange={() => handleChange(key)}
                            />
                            <span>{labelMap[key]}</span>
                        </label>
                    ))}
                </div>
            )}
            {uncheckedItems.length > 0 && (
                <div>
                    <h2 className="text-lg font-semibold mb-2">📋 今日未打卡</h2>
                    {uncheckedItems.map(([key], i) => (
                        <label key={i} className="flex items-center space-x-2 text-gray-800">
                            <input
                                type="checkbox"
                                className="h-5 w-5 text-blue-500 focus:ring-blue-400 border-gray-300 rounded"
                                checked={checkin[key as keyof typeof checkin]}
                                onChange={() => handleChange(key)}
                            />
                            <span>{labelMap[key]}</span>
                        </label>
                    ))}
                </div>
            )}
        </div>
    );
};
