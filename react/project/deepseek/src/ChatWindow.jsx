import { useState, useEffect, useRef } from 'react';
import './styles.css';

function ChatWindow({ chat, updateMessages }) {
    const [messages, setMessages] = useState(chat.messages || []);
    const [input, setInput] = useState('');
    const [loading, setLoading] = useState(false);
    const controllerRef = useRef(null);
    const messagesEndRef = useRef(null);

    useEffect(() => {
        messagesEndRef.current?.scrollIntoView({ behavior: "smooth" });
    }, [messages]);

    useEffect(() => {
        setMessages(chat.messages);
    }, [chat]);

    const handleStream = async (message) => {
        try {
            setLoading(true);
            const assistantMessage = { role: 'assistant', content: '' };
            setMessages(prev => [...prev, assistantMessage]); // 先创建一个空的 AI 消息
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

                // 解析 SSE 数据
                while (buffer.includes('\n\n')) {
                    const splitIndex = buffer.indexOf('\n\n');
                    const chunk = buffer.slice(0, splitIndex);
                    buffer = buffer.slice(splitIndex + 2);

                    let eventData = {};
                    const lines = chunk.split('\n');

                    for (const line of lines) {
                        if (line.startsWith('event:')) {
                            eventData.event = line.replace('event:', '').trim();
                        } else if (line.startsWith('data:')) {
                            eventData.data = line.replace('data:', '').trim();
                        }
                    }

                    // 只处理 message 事件
                    if (eventData.event === 'message' && eventData.data) {
                        let newData = eventData.data.trim();

                        // 跳过无用信息
                        if (!newData || newData.startsWith('<think>') || newData.startsWith('</think>') || newData === 'stream ended') {
                            continue;
                        }

                        // 确保内容不重复
                        setMessages(prev => {
                            const newMessages = [...prev];
                            const lastMessageIndex = newMessages.length - 1;

                            if (newMessages[lastMessageIndex]?.role === 'assistant') {
                                if (!newMessages[lastMessageIndex].content.endsWith(newData)) {
                                    newMessages[lastMessageIndex].content += newData;
                                }
                            }

                            return [...newMessages]; // 触发 React 重新渲染
                        });
                    }
                }
            }
        } catch (err) {
            console.error('请求失败:', err);
            setMessages(prev => [...prev, { role: 'error', content: '请求失败' }]);
        } finally {
            setLoading(false);
        }
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        if (!input.trim() || loading) return;

        const newMessages = [...messages, { role: 'user', content: input }];
        setMessages(newMessages);
        updateMessages(chat.id, newMessages);
        setInput('');
        await handleStream(input);
    };

    return (
        <div className="chat-container">
            <div className="messages">
                {messages.map((msg, index) => (
                    <div key={index} className={`message-wrapper ${msg.role}`}>
                        {/* 头像部分，用户在右侧，AI在左侧 */}
                        <div className="avatar">
                            <img
                                src={msg.role === 'user' ? "/user-avatar.png" : "/ai-avatar.png"}
                                alt={msg.role}
                            />
                        </div>
                        {/* 消息气泡 */}
                        <div className={`message ${msg.role}`}>
                            <div className="content">{msg.content || '思考中...'}</div>
                        </div>
                    </div>
                ))}
                <div ref={messagesEndRef} />
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
            </form>
        </div>
    );
}

export default ChatWindow;
