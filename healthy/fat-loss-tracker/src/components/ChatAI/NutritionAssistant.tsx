import { useState } from "react";

type Message = {
    role: "user" | "assistant";
    content: string;
};

export default function NutritionAssistant({ checkin }: { checkin: any }) {
    const [messages, setMessages] = useState<Message[]>([]);
    const [input, setInput] = useState("");
    //@ts-ignore
    const handleSend = async () => {
        if (!input.trim()) return;

        const userMessage: Message = { role: "user", content: input };
        setMessages((prev) => [...prev, userMessage]);

        try {
            const res = await fetch("/api/chat", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({ message: input, checkin }),
            });

            if (!res.ok || !res.body) {
                throw new Error(`请求失败，状态码: ${res.status}`);
            }

            const reader = res.body.getReader();
            const decoder = new TextDecoder("utf-8");
            let fullText = "";
            const assistantMsg: Message = { role: "assistant", content: "" };
            setMessages((prev) => [...prev, assistantMsg]);

            while (true) {
                const { value, done } = await reader.read();
                if (done) break;
                const chunk = decoder.decode(value, { stream: true });
                const content = chunk.replace(/^data:\s*/, "").trim();
                fullText += content;
                setMessages((prev) => {
                    const updated = [...prev];
                    updated[updated.length - 1] = { role: "assistant", content: fullText };
                    return updated;
                });
            }
        } catch (error) {
            console.error("Chat 发送失败:", error);
            setMessages((prev) => [
                ...prev,
                { role: "assistant", content: "🤖：服务器开小差啦，请稍后再试~ 🚧" },
            ]);
        } finally {
            setInput("");
        }
    };

    return (
        <div className="fixed bottom-20 right-6 w-80 bg-white border shadow-lg rounded-lg z-50 p-4 flex flex-col">
            <h2 className="font-bold text-lg mb-2">🧠 营养问答助手</h2>
            <div className="flex-1 overflow-y-auto mb-2 max-h-64 text-sm space-y-2">
                {messages.map((msg, i) => (
                    <div
                        key={i}
                        className={`p-2 rounded ${
                            msg.role === "user" ? "bg-blue-100 text-right" : "bg-gray-100 text-left"
                        }`}
                    >
                        {msg.content}
                    </div>
                ))}
            </div>
            <div className="flex">
                <input
                    value={input}
                    onChange={(e) => setInput(e.target.value)}
                    placeholder="请输入你的问题..."
                    className="flex-1 border rounded-l px-2 py-1 text-sm"
                />
                <button
                    onClick={handleSend}
                    className="bg-blue-500 text-white px-3 rounded-r"
                >
                    发送
                </button>
            </div>
        </div>
    );
}
