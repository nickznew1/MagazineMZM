\restrict i7WLMdBQ4KHSx6XkMa1OTSLfRiyEu2wZ7OcQSb49Qy8Bfp6gVThUdG7D9CQMlSn

-- Dumped from database version 14.23
-- Dumped by pg_dump version 14.23

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Data for Name: customer; Type: TABLE DATA; Schema: public; Owner: root
--

INSERT INTO public.customer (id, login, password, email, user_role, registration_date) VALUES (2, 'AdminUser', '$2a$12$Ve.w2S.3WxHYObTfeFiP8uQjku5rkeoJvLOSeh8btunIFlL2nwot6', 'testUser@gmail.com', 'admin', '2026-07-01 07:37:26.070058');
INSERT INTO public.customer (id, login, password, email, user_role, registration_date) VALUES (3, 'PrimaryUser', '$2a$12$h4EUVhThodORLKEqoX7Ske6GTLBqVEDiCNXbcU2OI2wdbeqX.umX.', 'PrimaryUser@gmail.com', 'ordinary', '2026-07-01 07:38:23.789422');
INSERT INTO public.customer (id, login, password, email, user_role, registration_date) VALUES (4, 'TestUser123', '$2a$12$C5cchaE1FMM9sh7Rcw3Mdu0ODSrkpJVoY3TvY/HcJTZQNeHpkvhF.', 'nickztest@gmail.com', 'ordinary', '2026-08-03 14:17:36.295132');


--
-- Data for Name: customer_delivery_info; Type: TABLE DATA; Schema: public; Owner: root
--

INSERT INTO public.customer_delivery_info (id, phone_number, city, address) VALUES (4, '89089081418', 'Магнитогорск', 'Доменщиков');


--
-- Data for Name: item; Type: TABLE DATA; Schema: public; Owner: root
--

INSERT INTO public.item (id, name, price, item_type, secondary_type, item_picture, item_description, item_short_description, article, visible) VALUES (1, 'МЗМ-ДД1 Дифференциального давления', 30000, 'Датчик давления', 'Дифференциальное давление', 'Датчик дифференциального давления.png', 'Данные приборы измеряют давление в трубах, сосудах, аппаратах, паровых и водогрейных котлах, в качестве точки отсчета при этом используется давления равного абсолютному нулю. Обеспечивается такое измерение наличием специальной камеры, из которой при изготовлении прибора откачивается воздух. Эта камера располагается в полости сенсора, а с другой стороны на чувствительный элемент воздействует давление измеряемой среды. Электронный блок производит расчет и выводит полученное значение на дисплей или передает с помощью выходных сигналов.

Чувствительным элементом служит монокристаллическая кремниевая мембрана, на которой расположен мост Уитстона, плечи моста - пьезорезисторы. Для защиты сенсора от воздействия агрессивной измеряемой и окружающих сред, в отдельных модификациях предусмотрены разделительные мембраны и заполняющая жидкость. При этом есть возможность выбрать тип заполняющей жидкости и материалы мембраны. В руководстве по эксплуатации прописаны все варианты изготовления, в том числе электронного блока и корпуса приборов с учетом условий эксплуатации.', 'Исполнения: штуцерное, с открытой разделительной мембранной или выносной разделительной мембраной, с традиционным фланцевым присоединением по стандартам EN 61518: IEC 61518. Диапазон настраиваемого верхнего предела измерения (ВПИ) 500 Па до 40 МПа. Перенастройка диапазона измерения 100:1, температура среды до 230 ºC, самодиагностика NAMUR NE107, соответствие требованиям функциональной безопасности SIL2.', 9823, true);
INSERT INTO public.item (id, name, price, item_type, secondary_type, item_picture, item_description, item_short_description, article, visible) VALUES (3, 'МЗМ-ИД-1 избыточного давления', 35000, 'Датчик давления', 'Избыточное давление', 'Датчик избыточного давления.png', 'Датчики давления 4-20 мА могут использоваться в системах автоматического контроля, регулирования и управления технологическими процессами в черной и цветной металлургии, нефтепереработке и газопереработке, добыча нефти газа, химической отрасли и пищевой отрасли. Также датчики давления ЭМИС-БАР успешно применяются на рудниках и морских платформах и судах.

Представляют собой сочетание преобразователей избыточного и вакуумметрического давлений, т.е. измеряют как давление, так и разрежение.

Датчики МЗМ для измерения избыточного давления применяются в системах коммерческого учета для компенсации давления измеряемой среды при приведении значения расхода к стандартным или нормальным условиям.

В энергетике данный тип приборов используется для измерения давления пара, воды, газа и других рабочих сред в котлах, турбинах, теплосчетчиках и другом оборудовании;

В химической промышленности преобразователи избыточного давления используются для контроля давления в реакторах, колоннах, емкостях и других аппаратах, а также для измерения скорости потока и расхода химических реагентов;

В вакуумной технике используются для контроля давления в вакуумных насосах, камерах, установках и других системах.', 'Исполнения: штуцерное, с открытой разделительной мембранной или выносной разделительной мембраной, с традиционным фланцевым присоединением по стандартам EN 61518: IEC 61518. Диапазон настраиваемого верхнего предела измерения (ВПИ) от 100 Па до 70 МПа. Перенастройка диапазона измерения 100:1, температура среды до 700 ºC, самодиагностика NAMUR NE107, соответствие требованиям функциональной безопасности SIL2.', 9823, true);
INSERT INTO public.item (id, name, price, item_type, secondary_type, item_picture, item_description, item_short_description, article, visible) VALUES (4, 'МЗМ-ГД-1 гидростатического давления', 60000, 'Датчик давления', 'Гидростатическое давление', 'Датчик гидростатического давления.png', 'Датчики гидростатического давления также получили наименование уровнемеры, поскольку они способны производить учет объема жидкости в емкости. Измерение проводится при помощи столба жидкости на плюсовую мембрану и, при необходимости, измерения минусовой полостью под куполом емкости, для исключения влияния насыщенного пара.

Устройства имеют дополнительный сертификат на эксплуатацию в среде сероводорода (ГОСТ Р 53679 и 53676), уровень полноты безопасности SIL-2 (ГОСТ Р МЭК 61508) SIL, заключение по санитарно-гигиенической экспертизе, что позволяет использовать их в пищевой промышленности.

Имеется дополнительная опция – исполнение с радиатором между корпусом датчика и разделительной мембраной для работы при температуре до 200°С. При спецзаказе делается внешняя защитная обработка, если окружающая среда обладает высокой коррозионной активностью. Перечень всех дополнительных опций и специсполнений прописан в руководстве по эксплуатации.', 'Варианты присоединения датчика гидростатического давления к процессу: с мембранным фланцевым разделителем стандартной и тубусной конструкции. До 3 МПа.', 9823, true);
INSERT INTO public.item (id, name, price, item_type, secondary_type, item_picture, item_description, item_short_description, article, visible) VALUES (5, 'МЗМ-ВА-1 вихреакустического расхода', 90000, 'Датчик расхода', 'Вихреакустический расход', 'вихреакустический.png', 'Принцип его работы заключается в следующем: электронный блок формирует на датчике-излучателе высокочастотный ультразвуковой сигнал, который проходя через вихревую дорожку, модулируется по фазе и поступает на датчик пьезоприемника. Такой метод обеспечивает виброустойчивость и расширяет динамический диапазон (до значения 1:100).

Также для расходомера внедрена автоматическая подстройка уровня сигналов, позволяющая компенсировать загрязнения сенсоров и их деградацию по мере старения прибора.

В базовой версии электронный блок поставляется с частотно/импульсным выходом и цифровым выходным сигналом Modbus с интерфейсом RS-485. Однако, по желанию заказчика, приборы могут быть изготовлены с четырехстрочным OLED-индикатором, аналоговым выходным сигналом и цифровым протоколом HART.

Взрывозащищенное исполнение и широкий температурный диапазон (от -60 до +70 ºC) позволяют эксплуатировать расходомер в экстремальных условиях.

Он предназначен для измерения объема и объемного расхода жидкостей в трубопроводах систем ППД (поддержания пластового давления), а также других жидкостей при высоком давлении.', 'Преобразователи расхода могут использоваться в составе автоматических систем управления и контроля, локальных схемах автоматизации с использованием частотно-импульсного сигнала, токового сигнала и цифрового сигнала ModBus (RS485) и HART.', 123456, true);
INSERT INTO public.item (id, name, price, item_type, secondary_type, item_picture, item_description, item_short_description, article, visible) VALUES (7, 'МЗМ-ВР-2 вихревого расхода', 89000, 'Датчик расхода', 'Вихревой расход', 'Датчик вихревой расход.png', 'Счетчик не требует периодической калибровки, а диагностика и замена узлов производится без демонтажа. Доставляют счетчики только после прохождения обязательного пролива на поверочном стенде. Удаленная передача данных, настройка, поверка приборов (через RS-485 на базе протокола Modbus RTU) позволяют снижать расходы на обслуживание. Счетчик имеет широкий динамический диапазон измерений; важным преимуществом является наличие конструктивного исполнения с коническими переходами.

Высокая точность измерений позволяет использовать вихревой счетчик газа и пара для коммерческого учета в составе теплосчетчиков и счетчиков пара', 'При покупке любого прибора мы предоставляем вам поддержку при проектировании, разработке комплексов и измерительных систем, индивидуальном подборе. В зависимости от проекта инженеры компании предоставят стандартное либо уникальное индивидуальное исполнение расходомера. Приборы выпускаются в нескольких модификациях, а значит это универсальные счетчики, которые можно использовать в различных сферах промышленности и хозяйства.', 123456, true);
INSERT INTO public.item (id, name, price, item_type, secondary_type, item_picture, item_description, item_short_description, article, visible) VALUES (10, 'МЗМ-ВУ-1 волноводный', 123221, 'Датчик уровня', 'Волноводные', 'Волновой уровнемер.png', 'Уровнемеры волноводные МЗМ-ВУ-1 предназначены для измерений уровня жидкости, в том числе сжиженных газов, и сыпучих материалов при атмосферном и избыточном давлении.

Принцип действия волноводных уровнемеров основан на излучении электромагнитных импульсов, распространяющихся вдоль зонда в сторону измеряемой среды. При достижении среды с другой диэлектрической проницаемостью часть импульсов отражается и передаётся обратно.

На основании временной задержки между излученным и принятым сигналом рассчитывается значение уровня измеряемой среды. Аналогичным образом измеряется расстояние между сенсором и границей раздела двух жидких сред с различными коэффициентами диэлектрической проницаемости.', 'Волноводные уровнемеры МЗМ-ВУ-1 предназначены для измерения и контроля уровня и границы раздела фаз жидких и сыпучих сред.', 9999, true);
INSERT INTO public.item (id, name, price, item_type, secondary_type, item_picture, item_description, item_short_description, article, visible) VALUES (9, 'МЗМ-РУ-1 радарный', 55555, 'Датчик уровня', 'Радарные', 'УРОВНЕМЕР РАДАРНЫЙ.png', 'Принцип действия уровнемеров основан на излучении антенной уровнемера непрерывного частотно-модулированного сигнала, который, отражаясь от поверхности измеряемой среды, принимается антенной уровнемера с временной задержкой. Используя разность частот излучаемого и принимаемого сигналов, вычисляется значение уровня измеряемой среды.', 'Уровнемер радарный МЗМ-РУ-1 предназначен для измерения уровня жидкости (в том числе сжиженных газов) и сыпучих материалов при атмосферном и избыточном давлении.', 231231231, true);
INSERT INTO public.item (id, name, price, item_type, secondary_type, item_picture, item_description, item_short_description, article, visible) VALUES (8, 'МЗМ-ЭР-1 электромагнитного расхода', 99999, 'Датчик расхода', 'Электромагнитный расход', 'Электромагнитный расход.png', 'Электромагнитные расходомеры-счетчики ЭМИС-МАГ 270 предназначены для измерений объемного расхода электропроводных жидкостей в прямом и обратном направлении потока. В том числе агрессивных, двухкомпонентных и загрязненных жидкостей с твердыми включениями.

Применяются в нефтяной, химической, нефтехимической, металлургической, пищевой промышленности и в коммунальном хозяйстве', 'Измерение расхода электропроводных жидкостей, в том числе загрязненных и агрессивных сред электромагнитным расходомером.', 1232311, true);
INSERT INTO public.item (id, name, price, item_type, secondary_type, item_picture, item_description, item_short_description, article, visible) VALUES (6, 'МЗМ-ВР-1 высокого давления', 50000, 'Датчик расхода', 'Вихревой расход', 'Вихревой расход высокого давления.png', 'Принцип его работы заключается в следующем: электронный блок формирует на датчике-излучателе высокочастотный ультразвуковой сигнал, который проходя через вихревую дорожку, модулируется по фазе и поступает на датчик пьезоприемника. Такой метод обеспечивает виброустойчивость и расширяет динамический диапазон (до значения 1:100).

Также для расходомера внедрена автоматическая подстройка уровня сигналов, позволяющая компенсировать загрязнения сенсоров и их деградацию по мере старения прибора.

В базовой версии электронный блок поставляется с частотно/импульсным выходом и цифровым выходным сигналом Modbus с интерфейсом RS-485. Однако, по желанию заказчика, приборы могут быть изготовлены с четырехстрочным OLED-индикатором, аналоговым выходным сигналом и цифровым протоколом HART.

Взрывозащищенное исполнение и широкий температурный диапазон (от -60 до +70 ºC) позволяют эксплуатировать расходомер в экстремальных условиях.

Он предназначен для измерения объема и объемного расхода жидкостей в трубопроводах систем ППД (поддержания пластового давления), а также других жидкостей при высоком давлении.', 'Преобразователи расхода могут использоваться в составе автоматических систем управления и контроля, локальных схемах автоматизации с использованием частотно-импульсного сигнала, токового сигнала и цифрового сигнала ModBus (RS485) и HART.', 123456, true);


--
-- Data for Name: customer_item; Type: TABLE DATA; Schema: public; Owner: root
--

INSERT INTO public.customer_item (item_id, count, customer_id, props) VALUES (5, 4, 3, '{"Взрывозащита": "1 Exd IIC T5 Gb X", "Погрешность, %": "0...+100", "Пылевлагозащита": "IP 66/68", "Спец. исполнение": "Сероводородное", "Выходные сигналы": " Цифровой: RS-485 с протоколом Modbus RTU", "Измеряемая среда": "Жидкость", "Напряжение питания, В": " ±1.5", "Межповерочный интервал": "4 года", "Диаметр условного прохода, мм": " 100", "Температура измеряемой среды, °С": "-60...+70", "Температура окружающей среды, °С": "-60...+70"}');


--
-- Data for Name: customer_personal_info; Type: TABLE DATA; Schema: public; Owner: root
--

INSERT INTO public.customer_personal_info (id, company, first_name, second_name) VALUES (4, 'ММКММкМММ', 'Никита', 'Слуквенко');


--
-- Data for Name: item_properties; Type: TABLE DATA; Schema: public; Owner: root
--

INSERT INTO public.item_properties (id, name) VALUES (1, 'Измеряемая среда');
INSERT INTO public.item_properties (id, name) VALUES (2, 'Спец. исполнение');
INSERT INTO public.item_properties (id, name) VALUES (3, 'Взрывозащита');
INSERT INTO public.item_properties (id, name) VALUES (4, 'Выходные сигналы');
INSERT INTO public.item_properties (id, name) VALUES (5, 'Долговременная стабильность');
INSERT INTO public.item_properties (id, name) VALUES (6, 'Рабочая температура, °С');
INSERT INTO public.item_properties (id, name) VALUES (7, 'Диапазон измеряемых давлений (мин, макс), МПа');
INSERT INTO public.item_properties (id, name) VALUES (8, 'Напряжение питания, В');
INSERT INTO public.item_properties (id, name) VALUES (9, 'Погрешность, %');
INSERT INTO public.item_properties (id, name) VALUES (10, 'Температура измеряемой среды, °С');
INSERT INTO public.item_properties (id, name) VALUES (11, 'Температура окружающей среды, °С');
INSERT INTO public.item_properties (id, name) VALUES (12, 'Диапазон перенастройки');
INSERT INTO public.item_properties (id, name) VALUES (13, 'Материал мембраны процесса');
INSERT INTO public.item_properties (id, name) VALUES (14, 'Заполняющая жидкость');
INSERT INTO public.item_properties (id, name) VALUES (15, 'Материал корпуса электронного блока');
INSERT INTO public.item_properties (id, name) VALUES (16, 'Межповерочный интервал');
INSERT INTO public.item_properties (id, name) VALUES (17, 'Пылевлагозащита');
INSERT INTO public.item_properties (id, name) VALUES (18, 'Технологическое присоединение');
INSERT INTO public.item_properties (id, name) VALUES (19, 'Вибростойкость');
INSERT INTO public.item_properties (id, name) VALUES (20, 'Диаметр условного прохода, мм');
INSERT INTO public.item_properties (id, name) VALUES (21, 'Типоразмер');
INSERT INTO public.item_properties (id, name) VALUES (22, 'Частота');
INSERT INTO public.item_properties (id, name) VALUES (23, 'Диэлектрическая проницаемость среды');
INSERT INTO public.item_properties (id, name) VALUES (24, 'Кабельные вводы');


--
-- Data for Name: item_properties_values; Type: TABLE DATA; Schema: public; Owner: root
--

INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (1, 1, 'Газ,Жидкость,Пар');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (1, 2, 'Водородное,Высокотемпературное,Кислородное,Криогенное,Морское,Пищевое,Рудничное,Сероводородное,Хлорное');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (1, 3, '1Ex d IIC T6…T4 Gb X,1Ex d ia IIC T6…T4 Gb X,0Ex ia IIB T6…T4 Ga X,0Ex ia IIC T6…T4 Ga X,PO Ex ia I Ma X,PB Ex d I Mb X,PB Ex d ia I Mb X');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (1, 4, '4-20мА + HART с наличием DD');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (1, 5, 'Не более 0,1% в течение 10 лет');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (1, 6, '-120...+700');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (1, 7, '0..40');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (1, 8, '10.5...45');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (1, 9, '±0.04,0.065,0.1,0.2 ');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (1, 10, '-90...+230 (с разделителем сред)');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (1, 11, '-60..+85');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (1, 12, 'до 100:1');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (1, 13, 'Нержавеющая сталь 316L,сплав Хастеллой,Тантал,Монель,Никель,316L с золотым напылением');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (1, 14, 'Силиконовое масло, Инертное масло');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (1, 15, 'Алюминий,Алюминий с защитным слоем,Нержавеющая сталь');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (1, 16, '6 лет');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (1, 17, 'IP65,IP66,IP67,IP68');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (1, 18, 'М20х1.5,М44х1.25,G1/2 наружная резьба,1” безрезьбовое присоединение,1/2NPT наружная/внутренняя резьба для датчиков штуцерного исполнения,1/4NPT фланцевое присоединение для датчиков фланцевого исполнения.');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (1, 19, 'V2,G2 по ГОСТ Р 52931-2008');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (3, 1, 'Газ, Жидкость, Пар');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (3, 2, 'Водородное, Высокотемпературное, Кислородное, Криогенное, Морское, Пищевое, Рудничное, Сероводородное, Хлорное');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (3, 3, '1Ex d IIC T6…T4 Gb X, 1Ex d ia IIC T6…T4 Gb X, 0Ex ia IIB T6…T4 Ga X, 0Ex ia IIC T6…T4 Ga X, PO Ex ia I Ma X, PB Ex d I Mb X, PB Ex d ia I Mb X');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (3, 4, '4-20 мА + HART с наличием DD');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (3, 5, 'Не более 0,1% в течение 10 лет');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (3, 6, '-40...+120');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (3, 7, '0..69');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (3, 8, '10.5...45 ');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (3, 9, '±0.065, 0.1, 0.2 ');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (3, 10, '-90...+700 (с разделителем сред)');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (3, 11, '-60...+85 с сохран. взрывозащиты');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (3, 12, 'до 100:1');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (3, 13, 'Нержавеющая сталь 316L, сплав Хастеллой, Тантал, Монель, Никель, 316L с золотым напылением');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (3, 14, 'Силиконовое масло, Инертное масло');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (3, 15, 'Алюминий, Алюминий с защитным слоем, Нержавеющая сталь');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (3, 16, '6 лет');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (3, 17, 'IP65, IP66, IP67, IP68');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (3, 18, 'М20х1.5, М44х1.25, G1/2 наружная резьба, 1” безрезьбовое присоединение. 1/2NPT наружная/внутренняя резьба для датчиков штуцерного исполнения, 1/4NPT фланцевое присоединение для датчиков фланцевого исполнения');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (3, 19, 'G2 по ГОСТ Р 52931-2008');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (4, 1, 'Газ, Жидкость, Пар');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (4, 2, 'Водородное, Высокотемпературное, Кислородное, Криогенное, Морское, Пищевое, Рудничное, Сероводородное, Хлорное');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (4, 3, '1Ex d IIC T6…T4 Gb X, 1Ex d ia IIC T6…T4 Gb X, 0Ex ia IIB T6…T4 Ga X, 0Ex ia IIC T6…T4 Ga X, PO Ex ia I Ma X, PB Ex d I Mb X, PB Ex d ia I Mb X');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (4, 4, '4-20 мА + HART с наличием DD');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (4, 5, 'Не более 0.1% в течение 10 лет');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (4, 6, '-40...+120');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (4, 7, '0..69');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (4, 8, '10.5...45');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (4, 9, '±0.065, 0.1, 0.2 ');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (4, 10, '60..+85');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (4, 11, '-60..+85');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (4, 12, 'до 100:1');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (4, 13, 'Нержавеющая сталь 316L, сплав Хастеллой, Тантал, Монель, Никель, 316L с золотым напылением');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (4, 14, 'Силиконовое масло, Инертное масло');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (4, 15, 'Алюминий, Алюминий с защитным слоем, Нержавеющая сталь');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (4, 16, '6 лет');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (4, 17, 'IP65,IP66,IP67,IP68');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (4, 18, 'М20х1.5, М44х1.25, G1/2 наружная резьба, 1” безрезьбовое присоединение, 1/2NPT наружная/внутренняя резьба для датчиков штуцерного исполнения, 1/4NPT фланцевое присоединение для датчиков фланцевого исполнения.');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (4, 19, 'G2 по ГОСТ Р 52931-2008');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (5, 1, 'Жидкость');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (5, 2, 'Сероводородное');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (5, 3, '1 Exd IIC T5 Gb X');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (5, 4, 'Частотно-импульсный, Аналоговый: токовый 4…20мА, Цифровой: RS-485 с протоколом Modbus RTU, HART, USB (технологический)');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (5, 8, '±1.0, ±1.5, ±3.0');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (5, 9, '0...+100');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (5, 10, '-60...+70');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (5, 11, '-60...+70');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (5, 16, '4 года');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (5, 17, 'IP 66/68');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (5, 20, '50, 80, 100, 150');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (7, 1, 'Газ, Жидкость, Пар');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (7, 2, 'Водородное, Высокотемпературное, Кислородное, Криогенное, Морское, Пищевое, Рудничное, Сероводородное, Хлорное');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (7, 3, '1 Exd IIC (T1-T6) Gb X, 1 Exib IIB/IIС (T1-T6) Gb X, 1 Exia IIB/IIС (T1-T6) Gb X, 0 Exia IIB/IIC (T1-T6) Gb Х, РВ Exib IМb Х, РО Exia IMa Х, РВ ExdI Mb Х');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (7, 4, 'Частотно-импульсный, Аналоговый, токовый 4…20мА, Цифровой: RS-485 с протоколом Modbus RTU, HART, USB (технологический)');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (7, 8, '12-30');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (7, 9, '±0,5,±0,7');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (7, 10, '-200...+450');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (7, 11, '-60...+70');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (7, 16, '5 лет');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (7, 17, 'IP 66/68, IP 66 (уровня РВ; РВИ; РО; РО-РВ)');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (7, 20, '15...300');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (8, 1, 'Жидкость');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (8, 2, 'Объём');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (8, 3, '1Ex db IIC T6…T3 Gb X, 1Ex db [ia] IIC Т6…Т3 Gb X, РВ Ex db I Mb X, Ex tb IIIC T80°C...T180°C Db X');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (8, 4, 'Частотно-импульсный, Аналоговый: токовый 4…20мА, Цифровой: RS-485 с протоколом Modbus RTU, HART, Сигнал тревоги');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (8, 8, '24 В постоянного тока, 220 В переменного тока');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (8, 9, '±0.5');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (8, 10, '-40...+180');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (8, 11, '-60...+70');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (8, 16, '5 лет');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (8, 17, 'IP 65, IP 66, IP 67, IP 66/67');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (8, 20, '15...600');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (9, 1, 'Жидкость, Сжиженные газы, Сыпучие');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (9, 3, '0Ex ia IIC T6…T1 Ga X, Ex ia IIIC T80°C…T450°C, 0Ex ia IIB T6…T1 Ga X, Ex ia IIIB T80°C…T450°C, 1Ex db IIC T6…T1 Gb X, Ex tb IIIC T80°C...T450°C, 1Ex db ia IIC T6...T1 Gb X.');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (9, 4, 'Аналоговый 4-20 мА / цифровой HART v7');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (9, 8, '24 В постоянного, 220 В переменного');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (9, 10, 'от -60 до +450');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (9, 11, 'от -60…+85, -70 °С до +85 °С с термочехлом');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (9, 17, 'IP66,IP67,IP68');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (9, 15, 'Корпус: Алюминий / нержавеющая сталь. Антенна: Стали – 304/316; фторопласт – PTFE (в зависимости от типа антенны)');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (9, 18, 'Фланцевое присоединение от Ду50, фланцы ГОСТ/EN/ASME, Резьбовое присоединение от 1.5 дюймов, тип резьбы – G и NPT');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (9, 22, '6 ГГц – конические/стержневые, 26 ГГц – конические/параболические/противокоррозионные, 80 ГГц – линзовые антенны');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (9, 23, 'от 1.4 (в зависимости от типа антенны и измеряемой среды)');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (10, 1, 'Жидкость, Сжиженные газы, Сыпучие');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (10, 3, '0Ex ia IIC T6…T1, Ga X 0Ex ia IIB T6…T1 Ga X,, 1Ex db IIC T6…T1 Gb X, 1Ex db ia IIC T6…T1 Gb X, Ex ia IIIC T80°….T450°C Da X, Ex ia IIIB T80°….T450°C Da X, Ex tb IIIC T80°….T450°C Db X');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (10, 4, 'Аналоговый: 4-20 мА Цифровой: HART v7 дополнительный аналоговый выход 4-20 мА:');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (10, 8, 'постоянного тока от 18 до 30, переменного тока от 187 до 242 частота переменного тока 50±1 Гц');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (10, 10, 'от -196 до +445');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (10, 11, 'от -60…+80');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (10, 17, 'IP66/IP67, IP66/IP68');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (10, 15, '');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (10, 18, 'Антенна: Стали – 304/316/12Х18Н10Т; фторопласт – PTFE; Корпус: Алюминий / нержавеющая сталь');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (10, 22, '~ 1 ГГц');
INSERT INTO public.item_properties_values (item_id, property_id, value) VALUES (10, 23, 'от 1,4');


--
-- Data for Name: item_spec_files; Type: TABLE DATA; Schema: public; Owner: root
--

INSERT INTO public.item_spec_files (id, name, link, picture) VALUES (1, 'Руководство', 'Документация абсолютное давление.pdf', 'logos/pdf_icon.png');
INSERT INTO public.item_spec_files (id, name, link, picture) VALUES (3, 'Руководство', 'Документация избыточное давление.pdf', 'logos/pdf_icon.png');
INSERT INTO public.item_spec_files (id, name, link, picture) VALUES (4, 'Руководство', 'Документация гидростатическое давление.pdf', 'logos/pdf_icon.png');
INSERT INTO public.item_spec_files (id, name, link, picture) VALUES (5, 'Руководство', 'вихреакустический.pdf', 'logos/pdf_icon.png');
INSERT INTO public.item_spec_files (id, name, link, picture) VALUES (6, 'Руководство', 'Документация вихревой высокого давления.pdf', 'logos/pdf_icon.png');
INSERT INTO public.item_spec_files (id, name, link, picture) VALUES (7, 'Руководство', 'Документация вихревой расходомер.pdf', 'logos/pdf_icon.png');
INSERT INTO public.item_spec_files (id, name, link, picture) VALUES (8, 'Руководство', 'Электромагнитный расхо.pdf', 'logos/pdf_icon.png');
INSERT INTO public.item_spec_files (id, name, link, picture) VALUES (9, 'Руководство', 'Уровнемер радарный.pdf', 'logos/pdf_icon.png');
INSERT INTO public.item_spec_files (id, name, link, picture) VALUES (10, 'Руководство', 'Уровнемер волноводный.pdf', 'logos/pdf_icon.png');


--
-- Data for Name: schema_migrations; Type: TABLE DATA; Schema: public; Owner: root
--

INSERT INTO public.schema_migrations (version, dirty) VALUES (3, false);


--
-- Data for Name: spec; Type: TABLE DATA; Schema: public; Owner: root
--



--
-- Data for Name: user_applications; Type: TABLE DATA; Schema: public; Owner: root
--

INSERT INTO public.user_applications (id, user_id, email, first_name, second_name, login, phone_number, company, address, city, order_date, items, order_status) VALUES (1, 2, 'testUser@gmail.com', 'НИКИТОС', 'НИКИТОС', 'AdminUser', '89094882121', 'mmk', 'доменщиков', 'магнитка', '2026-07-02', '[{"id": 2, "name": "МЗМ-ДД1 Дифференциального давления", "count": 4, "price": 30000, "props": {"Взрывозащита": "1Ex d ia IIC T6…T4 Gb X", "Погрешность, %": "0.065", "Вибростойкость": "V2", "Пылевлагозащита": "IP65", "Спец. исполнение": "Высокотемпературное", "Выходные сигналы": "4-20мА + HART с наличием DD", "Измеряемая среда": "Жидкость", "Заполняющая жидкость": "Силиконовое масло", "Напряжение питания, В": "10.5...45", "Диапазон перенастройки": "до 100:1", "Межповерочный интервал": "6 лет", "Рабочая температура, °С": "-120...+700", "Материал мембраны процесса": "Нержавеющая сталь 316L", "Долговременная стабильность": "Не более 0", "Технологическое присоединение": "М44х1.25", "Температура измеряемой среды, °С": "-90...+230 (с разделителем сред)", "Температура окружающей среды, °С": "-60..+85", "Материал корпуса электронного блока": "Алюминий", "Диапазон измеряемых давлений (мин, макс), МПа": "0..40"}, "item_id": 1, "item_type": "Датчик давления", "item_article": "9823", "item_picture": "Датчик дифференциального давления.png", "item_spec_id": 0, "item_description": "Данные приборы измеряют давление в трубах, сосудах, аппаратах, паровых и водогрейных котлах, в качестве точки отсчета при этом используется давления равного абсолютному нулю. Обеспечивается такое измерение наличием специальной камеры, из которой при изготовлении прибора откачивается воздух. Эта камера располагается в полости сенсора, а с другой стороны на чувствительный элемент воздействует давление измеряемой среды. Электронный блок производит расчет и выводит полученное значение на дисплей или передает с помощью выходных сигналов.\r\n\r\nЧувствительным элементом служит монокристаллическая кремниевая мембрана, на которой расположен мост Уитстона, плечи моста - пьезорезисторы. Для защиты сенсора от воздействия агрессивной измеряемой и окружающих сред, в отдельных модификациях предусмотрены разделительные мембраны и заполняющая жидкость. При этом есть возможность выбрать тип заполняющей жидкости и материалы мембраны. В руководстве по эксплуатации прописаны все варианты изготовления, в том числе электронного блока и корпуса приборов с учетом условий эксплуатации.", "item_secondary_type": "Дифференциальное давление"}]', 'Отменен');


--
-- Name: customer_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.customer_id_seq', 5, true);


--
-- Name: item_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.item_id_seq', 11, true);


--
-- Name: item_properties_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.item_properties_id_seq', 24, true);


--
-- Name: spec_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.spec_id_seq', 1, false);


--
-- Name: user_applications_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.user_applications_id_seq', 1, true);


--
-- PostgreSQL database dump complete
--

\unrestrict i7WLMdBQ4KHSx6XkMa1OTSLfRiyEu2wZ7OcQSb49Qy8Bfp6gVThUdG7D9CQMlSn