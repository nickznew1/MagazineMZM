import "./style.css"

const LoadingPage = ({ message = "Загрузка контента...", variant = "spinner" }) => {
    return (
        <div className="loading-page">
            <div className="loading-content">
                <div className="spinner-container">
                    <svg className="spinner" viewBox="0 0 50 50">
                        <circle
                            className="spinner-track"
                            cx="25"
                            cy="25"
                            r="20"
                            fill="none"
                            strokeWidth="5"
                        />
                        <circle
                            className="spinner-circle"
                            cx="25"
                            cy="25"
                            r="20"
                            fill="none"
                            strokeWidth="5"
                            strokeLinecap="round"
                        />
                    </svg>
                </div>

                <p className="loading-text">{message}</p>
            </div>
        </div>
    );
};


export default LoadingPage;