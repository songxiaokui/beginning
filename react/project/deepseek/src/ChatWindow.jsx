import { useState, useEffect, useRef } from 'react';
import ReactMarkdown from 'react-markdown';
import remarkGfm from 'remark-gfm';
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
            setMessages(prev => [...prev, { role: 'assistant', content: '' }]); // 先插入空的 AI 消息
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

                    let eventData = {};
                    const lines = chunk.split('\n');
                    for (const line of lines) {
                        if (line.startsWith('event:')) {
                            eventData.event = line.replace('event:', '').trim();
                        } else if (line.startsWith('data:')) {
                            eventData.data = line.replace('data:', '').trim();
                        }
                    }

                    if (eventData.event === 'message' && eventData.data) {
                        let newData = eventData.data.trim();
                        if (!newData || newData.startsWith('<think>') || newData.startsWith('</think>') || newData === 'stream ended') {
                            continue;
                        }

                        setMessages(prev => {
                            const newMessages = [...prev];
                            const lastIndex = newMessages.length - 1;
                            if (newMessages[lastIndex]?.role === 'assistant') {
                                if (!newMessages[lastIndex].content.endsWith(newData)) {
                                    newMessages[lastIndex].content += newData;
                                }
                            }
                            return [...newMessages];
                        });
                    }
                }
            }
        } catch (err) {
            console.error('请求失败:', err);
            setMessages(prev => [...prev, { role: 'error', content: '服务器繁忙，请稍后重试。' }]);
        } finally {
            setLoading(false);
        }
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        if (!input.trim() || loading) return;
        setMessages(prev => [...prev, { role: 'user', content: input }]);
        updateMessages(chat.id, [...messages, { role: 'user', content: input }]);
        setInput('');
        await handleStream(input);
    };

    return (
        <div className="chat-window">
            <div className="messages">
                {messages.map((msg, index) => (
                    <div key={index} className={`bubble-row ${msg.role}`}>
                        <div className="avatar">
                            <img src={msg.role === 'user' ? "/user-avatar.png" : "/ai-avatar.png"} alt={msg.role} />
                        </div>
                        <div className="bubble">
                            {msg.role === 'error' ? (
                                <div className="error-message">⚠ {msg.content}</div>
                            ) : (
                                <ReactMarkdown
                                    remarkPlugins={[remarkGfm]}
                                    components={{
                                        a: ({ node, ...props }) => <a {...props} target="_blank" rel="noopener noreferrer" />,
                                        pre: ({ children }) => <pre className="md-pre">{children}</pre>,
                                        code: ({ children }) => <code className="md-code">{children}</code>,
                                        p: ({ children }) => <p className="md-p">{children}</p>,
                                        strong: ({ children }) => <strong className="md-strong">{children}</strong>,
                                        em: ({ children }) => <em className="md-em">{children}</em>,
                                        ul: ({ children }) => <ul className="md-ul">{children}</ul>,
                                        ol: ({ children }) => <ol className="md-ol">{children}</ol>,
                                        h1: ({ children }) => <h1 className="md-h1">{children}</h1>,
                                        h2: ({ children }) => <h2 className="md-h2">{children}</h2>,
                                        h3: ({ children }) => <h3 className="md-h3">{children}</h3>,
                                    }}
                                >
                                    {msg.content || ' '}
                                </ReactMarkdown>
                            )}
                        </div>
                    </div>
                ))}
                <div ref={messagesEndRef} />
            </div>

            <div className="input-box">
                <form onSubmit={handleSubmit}>
                    <input
                        value={input}
                        onChange={e => setInput(e.target.value)}
                        placeholder="输入你的问题..."
                        disabled={loading}
                    />
                    <button type="submit" disabled={loading}>
                        {loading ? '发送中...' : '发送'}
                    </button>
                </form>
            </div>
        </div>
    );
}

export default ChatWindow;
