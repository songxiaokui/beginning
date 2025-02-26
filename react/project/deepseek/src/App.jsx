import { useState, useEffect } from 'react';
import ChatWindow from './ChatWindow';
import ChatList from './ChatList';
import './styles.css';

function App() {
    const [chats, setChats] = useState([]); // 主题列表
    const [selectedChat, setSelectedChat] = useState(null); // 当前选中的聊天主题

    // 加载本地存储的聊天主题
    useEffect(() => {
        const storedChats = JSON.parse(localStorage.getItem('chats')) || [];
        setChats(storedChats);
    }, []);

    // 创建新聊天主题
    const createNewChat = () => {
        const newChat = {
            id: Date.now(), // 唯一 ID
            title: `会话 ${chats.length + 1}`,
            messages: []
        };
        const updatedChats = [newChat, ...chats];
        setChats(updatedChats);
        setSelectedChat(newChat);
        localStorage.setItem('chats', JSON.stringify(updatedChats));
    };

    // 更新聊天记录
    const updateChatMessages = (chatId, newMessages) => {
        const updatedChats = chats.map(chat =>
            chat.id === chatId ? { ...chat, messages: newMessages } : chat
        );
        setChats(updatedChats);
        localStorage.setItem('chats', JSON.stringify(updatedChats));
    };

    return (
        <div className="app-container">
            <ChatList chats={chats} onSelectChat={setSelectedChat} onNewChat={createNewChat} />
            {selectedChat ? (
                <ChatWindow chat={selectedChat} updateMessages={updateChatMessages} />
            ) : (
                <div className="empty-chat">请选择或创建一个新聊天</div>
            )}
        </div>
    );
}

export default App;
