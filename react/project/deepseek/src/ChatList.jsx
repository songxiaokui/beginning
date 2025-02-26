import './styles.css';

function ChatList({ chats, onSelectChat, onNewChat }) {
    return (
        <div className="chat-list">
            <button onClick={onNewChat} className="new-chat-btn">+ 开启新对话</button>
            <ul>
                {chats.map(chat => (
                    <li key={chat.id} onClick={() => onSelectChat(chat)}>
                        {chat.title}
                    </li>
                ))}
            </ul>
        </div>
    );
}

export default ChatList;
