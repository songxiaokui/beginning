import Clock from "./Clock/Clock";
import CheckinForm from "./CheckinFrom/CheckinForm";
import TodayBreakfast from "./TodayBreakfast/TodayBreakfast";
import TodayExercise from "./TodayExercise/TodayExercise";
// @ts-ignore
import backgroundImage from "../assets/bg.jpg";
import { useState, useEffect } from "react";
import axios from 'axios';
import SubmitButton from "./Submit/SubmitButton";
import {Toaster} from "react-hot-toast";
import ChatPop from "./ChatAI/ChatPop"
import NutritionAssistant from "./ChatAI/NutritionAssistant";


export default function CheckinPage() {
    const [checkin, setCheckin] = useState({
        breakfast: false,
        lunch: false,
        dinner: false,
        exercise: false,
        sleep: false,
    });

    const [loading, setLoading] = useState(true);
    const [showChat, setShowChat] = useState(false);

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

    return (
        <div className="w-screen min-h-screen bg-gradient-to-b from-white to-blue-50 flex items-center justify-center px-4 py-8">
            <Toaster position="top-center" reverseOrder={false} />
            <div className="relative z-10 w-full max-w-md p-6 bg-white/80 rounded-xl shadow bg-gradient-to-r from-blue-50 to-green-50">
                <h1 className="text-2xl font-bold text-center mb-4">减脂养成打卡</h1>
                <Clock />
                {!loading && <CheckinForm checkin={checkin} setCheckin={setCheckin} />}
                <TodayBreakfast />
                <TodayExercise />
                <SubmitButton checkin={checkin}/>
            </div>
            <ChatPop onClick={() => setShowChat(!showChat)} />
            {showChat && <NutritionAssistant checkin={checkin} />}
        </div>
    );
}