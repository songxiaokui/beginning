import { useState, useRef, useEffect } from 'react';
import './App.css';

function App() {
    const [messages, setMessages] = useState([]);
    const [input, setInput] = useState('');
    const [loading, setLoading] = useState(false);
    const controllerRef = useRef(null);
    const messagesEndRef = useRef(null);

    // 🔹 自动滚动到底部
    useEffect(() => {
        messagesEndRef.current?.scrollIntoView({ behavior: "smooth" });
    }, [messages]);

    // 处理流式响应
    const handleStream = async (message) => {
        try {
            setLoading(true);
            setMessages(prev => [...prev, { role: 'assistant', content: '' }]); // 先添加一个空的 AI 响应
            controllerRef.current = new AbortController();

            const response = await fetch('http://localhost:8080/api/chat/stream', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ message, stream: true }),
                signal: controllerRef.current.signal
            });

            const reader = response.body.getReader();
            const decoder = new TextDecoder();
            let buffer = '';

            while (true) {
                const { done, value } = await reader.read();
                if (done) break;

                buffer += decoder.decode(value, { stream: true });

                while (buffer.includes('\n\n')) {
                    const splitIndex = buffer.indexOf('\n\n');
                    const chunk = buffer.slice(0, splitIndex);
                    buffer = buffer.slice(splitIndex + 2);

                    const lines = chunk.split('\n');
                    let eventData = {};

                    for (const line of lines) {
                        if (line.startsWith('event:')) {
                            eventData.event = line.replace('event:', '').trim();
                        } else if (line.startsWith('data:')) {
                            eventData.data = line.replace('data:', '').trim();
                        }
                    }

                    if (eventData.event === 'message') {
                        setMessages(prev => {
                            const newMessages = [...prev];
                            const lastMessage = newMessages[newMessages.length - 1];

                            if (lastMessage?.role === 'assistant') {
                                newMessages[newMessages.length - 1] = {
                                    ...lastMessage,
                                    content: lastMessage.content + eventData.data
                                };
                            }
                            return newMessages;
                        });
                    }
                }
            }
        } catch (err) {
            if (err.name !== 'AbortError') {
                console.error('请求失败:', err);
                setMessages(prev => [...prev, { role: 'error', content: '请求失败' }]);
            }
        } finally {
            setLoading(false);
        }
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        if (!input.trim() || loading) return;

        setMessages(prev => [...prev, { role: 'user', content: input }]);
        await handleStream(input);
        setInput('');
    };

    const stopGenerating = () => {
        if (controllerRef.current) {
            controllerRef.current.abort();
            setLoading(false);
        }
    };

    return (
        <div className="chat-container">
            <div className="messages">
                {messages.map((msg, index) => (
                    <div key={index} className={`message ${msg.role}`}>
                        <div className="role">{msg.role === 'user' ? '你' : 'AI助手'}</div>
                        <div className="content">{msg.content || <span className="placeholder-text">思考中...</span>}</div>
                    </div>
                ))}
                {loading && <div className="loading">思考中...</div>}
                <div ref={messagesEndRef} /> {/* 滚动到底部 */}
            </div>

            <form onSubmit={handleSubmit} className="input-area">
                <input
                    value={input}
                    onChange={(e) => setInput(e.target.value)}
                    placeholder="输入你的问题..."
                    disabled={loading}
                />
                <button type="submit" disabled={loading}>
                    {loading ? '发送中...' : '发送'}
                </button>
                {loading && (
                    <button type="button" onClick={stopGenerating} className="stop-btn">
                        停止
                    </button>
                )}
            </form>
        </div>
    );
}

export default App;
