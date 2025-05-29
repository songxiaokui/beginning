import { useState } from "react"

export default function App() {
    const [messages, setMessages] = useState([
        { role: "user", content: "你好！" },
        { role: "ai", content: "🤖 回应内容（待接入模型）" },
    ])
    const [input, setInput] = useState("")

    const handleSend = () => {
        if (!input.trim()) return

        setMessages([
            ...messages,
            { role: "user", content: input },
            { role: "ai", content: "🤖 回应内容（待接入模型）" },
        ])
        setInput("")
    }

    return (
        <div className="h-screen flex flex-col bg-gray-100">
            {/* 消息区 */}
            <div className="flex-1 p-4 space-y-4 overflow-auto">
                {messages.map((msg, idx) => (
                    <div
                        key={idx}
                        className={`flex items-start gap-2 p-2 ${
                            msg.role === "ai" ? "justify-end" : "justify-start"
                        }`}
                    >
                        {msg.role === "user" && (
                            <img
                                src="/user-avatar.png"
                                alt="User"
                                className="w-6 h-6 rounded-full border"
                            />
                        )}

                        <div
                            className={`max-w-xs p-2 rounded-lg ${
                                msg.role === "ai"
                                    ? "bg-blue-200 text-left"
                                    : "bg-green-100 text-left"
                            }`}
                        >
                            {msg.content}
                        </div>

                        {msg.role === "ai" && (
                            <img
                                src="/bot-avatar.png"
                                alt="AI"
                                className="w-10 h-10 rounded-full border"
                            />
                        )}
                    </div>
                ))}

            </div>

            {/* 输入区 */}
            <div className="flex items-center p-4 border-t bg-white">
                <input
                    value={input}
                    onChange={(e) => setInput(e.target.value)}
                    onKeyDown={(e) => e.key === "Enter" && handleSend()}
                    placeholder="请输入内容..."
                    className="flex-1 px-4 py-2 border rounded"
                />
                <button
                    onClick={handleSend}
                    className="ml-2 bg-blue-500 text-white px-4 py-2 rounded"
                >
                    发送
                </button>
            </div>
        </div>
    )
}
