import Clock from "./Clock/Clock";
import CheckinForm from "./CheckinFrom/CheckinForm";
import TodayBreakfast from "./TodayBreakfast/TodayBreakfast";
import TodayExercise from "./TodayExercise/TodayExercise";
// @ts-ignore
import backgroundImage from "../assets/bg.jpg";
import {useState, useEffect} from "react";
import axios from 'axios';


export default function CheckinPage() {
    const [checkin, setCheckin] = useState({
        breakfast: false,
        lunch: false,
        dinner: false,
        exercise: false,
        sleep: false,
    });

    const [loading, setLoading] = useState(true);
    const [submitStatus, setSubmitStatus] = useState<string | null>(null);

    useEffect(() => {
        // @ts-ignore
        const fetchCheckin = async () => {
            try {
                const today = new Date().toISOString().slice(0, 10);
                const res = await axios.get(`/api/checkin?date=${today}`);
                if (res.data) {
                    setCheckin(res.data);
                }
            } catch (err) {
                console.log("No checkin data yet.");
            } finally {
                setLoading(false);
            }
        };
        fetchCheckin();
    }, []);

    // @ts-ignore
    const handleSubmit = async () => {
        const today = new Date().toISOString().slice(0, 10);
        try {
            await axios.post("/api/checkin", {date: today, config: checkin});
            setSubmitStatus("success");
        } catch (err) {
            console.error(err);
            setSubmitStatus("error");
        }
    };

    return (
        <div
            className="min-h-screen bg-gradient-to-b from-white to-blue-50 flex items-center justify-center bg-gradient-to-r from-blue-100 to-green-100"
        >
            <div className="absolute inset-0 bg-white/10 backdrop-blur-sm z-0"></div>

            <div className="relative z-10 w-full max-w-md p-6 bg-white/80 rounded-xl shadow bg-gradient-to-r from-blue-80 to-green-50">
                <h1 className="text-2xl font-bold text-center mb-4">减脂养成打卡</h1>
                <Clock/>
                {!loading && <CheckinForm checkin={checkin} setCheckin={setCheckin}/>}
                <TodayBreakfast/>
                <TodayExercise/>
                <button
                    className="w-full bg-blue-500 text-white py-2 rounded hover:bg-blue-600 mt-4"
                    onClick={handleSubmit}
                >更新打卡
                </button>
                {submitStatus === "success" && (
                    <p className="text-green-600 text-center mt-2">提交成功！</p>
                )}
                {submitStatus === "error" && (
                    <p className="text-red-600 text-center mt-2">提交失败，请重试。</p>
                )}
            </div>
        </div>
    );
}