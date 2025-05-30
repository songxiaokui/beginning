
export default function ChatPop({onClick}) {
    return (
        <button
            onClick={onClick}
            className="fixed bottom-6 right-6 bg-green-500 hover:bg-green-600 text-white p-4 rounded-full shadow-lg z-50"
        >
            💬
        </button>
    )
}