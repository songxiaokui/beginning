import { useState, useEffect } from "react";

export default function Clock() {
    const [now, setNow] = useState(new Date());

    const getWeekday = (date: Date) => {
        const days = ["星期日", "星期一", "星期二", "星期三", "星期四", "星期五", "星期六"];
        return days[date.getDay()];
    };

    useEffect(() => {
        const timer = setInterval(() => {
            setNow(new Date());
        }, 1000);

        return () => clearInterval(timer);
    }, []);

    return (
        <div className="flex flex-row items-center gap-4 mb-4 text-sm text-gray-700 justify-center">
            {/*@ts-ignore*/}
            <span>{`${now.getFullYear()}-${String(now.getMonth() + 1).padStart(2, '0')}-${String(now.getDate()).padStart(2, '0')}`}</span>
            <span>{now.toLocaleTimeString()}</span>
            <span>{getWeekday(now)}</span>
        </div>
    );
}
