import "./style.css"

const FooterContent = () =>{
    return (
        <footer className="footer-classic" role="contentinfo">
            <div className="footer-classic__container">
                <div className="footer-classic__bottom">
                    <p className="footer-classic__copyright">
                        2025 Магнитогорский завод микроэлектроники. <a href="#">Контакты</a> | <a href="#">О нас</a>
                    </p>
                    <nav className="footer-classic__social" aria-label="Социальные сети">
                        <ul>
                            <li><a href="https://vk.com" aria-label="Наша группа VK">
                                <img alt ="vk" src ="https://static.tildacdn.com/tild6433-3539-4161-a463-633934636231/Group_354.svg"></img>
                            </a></li>
                            <li><a href="https://web.telegram.org/" aria-label="Наш канал в телеграме">
                                <img alt ="tg" src ="https://static.tildacdn.com/tild6634-3434-4734-b034-333430386461/Group.svg"></img>
                            </a></li>
                            <li><a href="https://www.youtube.com/" aria-label="Наш ютуб-канал">
                                <img alt = "Youtube" src ="https://static.tildacdn.com/tild6130-6136-4138-b837-656434333730/cdnlogocom_youtube_1.svg"></img>
                            </a></li>
                        </ul>
                    </nav>

                </div>
            </div>
        </footer>
    )
}

export default FooterContent;