import axios from "axios";
import { Toaster, toast } from 'react-hot-toast';

export default  function SubmitButton({checkin}) {

    // @ts-ignore
    const handleSubmit = async () => {
        const today = new Date().toISOString().slice(0, 10);
        try {
            await axios.post("/api/checkin", { date: today, config: checkin });
            toast.success("✅ 提交成功！", { duration: 2000 });
        } catch (err) {
            console.error(err);
            toast.error("❌ 提交失败，请重试！", { duration: 2000 });
        }
    };

    return (
        <>
            <button
                className="w-full bg-blue-500 text-white py-2 rounded hover:bg-blue-600 mt-4"
                onClick={handleSubmit}
            >
                更新打卡
            </button>
        </>
    )
}