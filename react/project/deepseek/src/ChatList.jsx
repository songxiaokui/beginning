import './styles.css';

function ChatList({ chats, onSelectChat, onNewChat }) {
    return (
        <div className="chat-list">
            <button onClick={onNewChat} className="new-chat-btn">+ 新建聊天</button>
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
