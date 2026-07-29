import {AnimatePresence, motion} from "framer-motion";
import {memo, useState} from "react";
import "./style.css"
import {config} from "../../../config";

const ProfileApplications = ({profileApplications}) => {
    const profile = profileApplications

    const date = (item) =>{
        return new Date(item).toLocaleDateString('ru-RU');
    }


    const [activeSpec, setActiveSpec] = useState(false)
    const handleSpec =(id) =>{
        setActiveSpec(activeSpec === id ? null : id)
    }
    return (
       <div className= "user-profile-applications">
           {[...profile.user.itemData]
                   .sort((a, b) => b.id - a.id)
                   .map((item,index) => (
                       <div className = "user-profile-items-wrapper"  key={index}>
                           <>
                               <h1 id= {item.id}>Дата заявки: {date(item.order_date)}</h1>
                               {item.items.map((item) =>(
                                   <div className = "user-profile-item" key={item.id}>
                                       <img alt = {item.name} src ={`${config.imgServerUrl}${item.item_picture}`}></img>
                                       <h3>{item.name}</h3>
                                       <span>Количество: {item.count}</span>
                                       <button className = {activeSpec === item.item_spec_id ? "active" : ""} onClick = {()=>handleSpec(item.item_spec_id)}>Спецификация
                                           <svg className = "filter-dropdown-icon" viewBox="0 0 24 24" width="16" height="16" stroke="currentColor"
                                                strokeWidth="2" fill="none" strokeLinecap="round" strokeLinejoin="round">
                                               <polyline points="6 9 12 15 18 9"></polyline>
                                           </svg>
                                       </button>
                                       <AnimatePresence>
                                           {activeSpec === item.item_spec_id &&
                                               <motion.div
                                                   initial = {{ height:0, opacity:0}}
                                                   animate ={{height:"auto", opacity:1}}
                                                   exit={{height:0, opacity:0}}
                                                   transition={{duration: 0.3}}
                                                   className = "user-profile_item__specs-wrapper">
                                                   <div className = "user-profile_item__specs">
                                                       {Object.entries(item.props).map(([key,value]) => (
                                                           <div key = {item.id}>
                                                               <p key ={key} id ="key">{key}:</p>
                                                               <p key ={value} id ="value">{value}</p>
                                                           </div>
                                                       ))}
                                                   </div>
                                               </motion.div>
                                           }
                                       </AnimatePresence>
                                   </div>
                               ))}
                           </>
                           <div className="user-items-status">
                               <h1>Статус заявки: {item.application_status}</h1>
                           </div>
                       </div>
                   ))}
       </div>
    )
}

export default memo(ProfileApplications);