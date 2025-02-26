import { useState, useEffect } from 'react';
import ChatWindow from './ChatWindow';
import ChatList from './ChatList';
import './App.css';

function App() {
    const [chats, setChats] = useState([]);
    const [selectedChat, setSelectedChat] = useState(null);

    useEffect(() => {
        const storedChats = JSON.parse(localStorage.getItem('chats')) || [];
        setChats(storedChats);
    }, []);

    const createNewChat = () => {
        const newChat = {
            id: Date.now(),
            title: `会话 ${chats.length + 1}`,
            messages: []
        };
        setChats([newChat, ...chats]);
        setSelectedChat(newChat);
        localStorage.setItem('chats', JSON.stringify([newChat, ...chats]));
    };

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
