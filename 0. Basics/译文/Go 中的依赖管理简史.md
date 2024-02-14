原文：[Go 中的依赖管理简史](https://www.bytesizego.com/blog/history-of-dependency-management-go)

---




ByteSize Tip #7: A Brief History of Dependency Management in Go
















.messaging-header{background-color:#2C41DD;color:#EAF2FA;fill:#EAF2FA}.conversation-bubble-background{background:#2C41DD}.unread .conversation-bubble-background svg{fill:#2C41DD}.conversation-bubble-icon{fill:#EAF2FA}.conversation-bubble{position:fixed;width:56px;height:56px;bottom:20px;right:20px;cursor:pointer;z-index:101}.conversation-bubble-background{position:relative;z-index:103;width:56px;height:56px;border-radius:50%}.conversation-bubble-background svg{display:none}.conversation-bubble-shadow{position:absolute;z-index:102;top:-8px;left:-16px;width:88px;height:88px;background-size:88px 88px}.conversation-bubble-icon{position:absolute;top:12px;left:12px;width:32px;height:32px;z-index:111}.conversation-dismiss-icon{fill:#EAF2FA}.conversation-dismiss-icon path{stroke:#EAF2FA}.unread .conversation-bubble-background{background:none}.unread .conversation-bubble-background svg{display:block;width:56px;height:56px}.unread:after{position:absolute;z-index:103;left:43px;top:3px;width:10px;height:10px;content:"";background:#E24646;border-radius:50%}.messaging-chat-window{width:380px;height:480px;position:fixed;right:20px;bottom:100px;box-sizing:border-box;box-shadow:0 0 0 1px rgba(0,0,0,0.15),0px 7px 40px 2px rgba(148,149,150,0.3);background:white;display:flex;flex-direction:column;justify-content:space-between;transition:0.3s ease-in-out;border-radius:12px;z-index:112}@media screen and (max-width: 420px),screen and (max-height: 590px){.messaging-chat-window{width:100%;height:100%;top:0;bottom:0;right:0;border-radius:0;z-index:1143}}.messaging-message-list ::-webkit-scrollbar{display:none;height:344px;min-height:200px;overflow-y:scroll;background-color:white;background-size:100%;overflow:hidden;position:relative;z-index:113}@media screen and (max-width: 420px),screen and (max-height: 590px){.messaging-message-list{overflow-x:hidden;height:100%}}.messaging-header{height:100%;max-height:70px;border-top-left-radius:12px;border-top-right-radius:12px;padding:16px;position:relative;box-sizing:border-box;display:flex;align-items:center;z-index:114;-webkit-box-shadow:0px 2px 4px rgba(41,37,51,0.1);box-shadow:0px 2px 4px rgba(41,37,51,0.1)}@media screen and (max-width: 420px),screen and (max-height: 590px){.messaging-header{border-top-left-radius:0;border-top-right-radius:0}}.messaging-header--img{border-radius:50%}.messaging-header--user{display:grid;width:100%}.messaging-header--user-name{border-radius:5px;font-size:16px;font-weight:bold;line-height:20px;overflow:hidden;text-overflow:ellipsis;white-space:nowrap}.messaging-header--user-status{align-self:center;border-radius:5px;font-size:14px;line-height:1}.messaging-header--dismiss{align-self:center;display:none}@media screen and (max-width: 420px),screen and (max-height: 590px){.messaging-header--dismiss{display:block}}.messaging-user-input--container{display:flex;align-items:flex-start;-webkit-box-shadow:0px -2px 4px rgba(41,37,51,0.1);box-shadow:0px -2px 4px rgba(41,37,51,0.1);z-index:114}.messaging-user-input--text{width:100%;line-height:1.5em;resize:none;border:none;outline:none;border-bottom-left-radius:12px;border-bottom-right-radius:12px;box-sizing:border-box;padding:16px;font-size:16px;font-weight:400;white-space:pre-wrap;word-wrap:break-word;color:#565867;-webkit-font-smoothing:antialiased;overflow:scroll;bottom:0;overflow-x:hidden;overflow-y:auto}.messaging-user-input--send-btn{height:35px;min-width:35px;border-radius:50%;display:flex;align-items:center;justify-content:center;cursor:pointer;margin:1rem 1rem 0 0;background-color:#2C41DD}.messaging-user-input--send-btn svg path{stroke:#EAF2FA}.message{background-color:#EDF0F2;border-radius:12px 12px 12px 4px;white-space:pre-wrap;word-wrap:break-word;font-size:14px}.message.customer{color:#EAF2FA;background-color:#2C41DD;border-radius:12px 12px 4px 12px}.message.bg-danger{background-color:#8A1C1C;color:white}.dots{align-items:center;display:flex;gap:.25rem}.typing{padding:1rem}.typing .dot{background:#6E7577}.dot{display:block;border-radius:.5rem;height:0.50rem;width:0.50rem;animation:typingAnimation 1s infinite ease-in-out}@keyframes typingAnimation{0%{transform:translateY(0px)}28%{transform:translateY(-0.25rem)}44%{transform:translateY(0px)}}.dot:nth-child(1){animation-delay:200ms}.dot:nth-child(2){animation-delay:300ms}.dot:nth-child(3){animation-delay:400ms}.text-sentence::first-letter{text-transform:uppercase}.visitor-email-capture-btn{background-color:#2C41DD}.visitor-email-capture-btn svg path{stroke:#EAF2FA}.messaging-chat{font-family:var(--system-font)}.messaging-chat .pt2{padding-top:6.25px !important}.messaging-chat .pb2{padding-bottom:6.25px !important}.messaging-chat .mb3{margin-bottom:9.375px !important}.messaging-chat .pl3{padding-left:9.375px !important}.messaging-chat .pr3{padding-right:9.375px !important}.messaging-chat .pl4{padding-left:12.5px !important}.messaging-chat .pr4{padding-right:12.5px !important}.messaging-chat .mb1{margin-bottom:3.125px !important}.messaging-chat .mr2{margin-right:6.25px !important}.messaging-chat .text-gray{color:#8599a9}.messaging-chat .text-muted{color:#a09ea2 !important}.btn-primary{color:#fff;background-color:#2C41DD;border-color:rgba(0,0,0,0)}.btn-primary:focus,.btn-primary.focus{color:#fff;background-color:#1e30b8;border-color:rgba(0,0,0,0)}.btn-primary:hover{color:#fff;background-color:#1e30b8;border-color:rgba(0,0,0,0)}.btn-primary:active,.btn-primary.active,.open>.btn-primary.dropdown-toggle{color:#fff;background-color:#1e30b8;background-image:none;border-color:rgba(0,0,0,0)}.btn-primary:active:hover,.btn-primary:active:focus,.btn-primary.focus:active,.btn-primary.active:hover,.btn-primary.active:focus,.btn-primary.active.focus,.open>.btn-primary.dropdown-toggle:hover,.open>.btn-primary.dropdown-toggle:focus,.open>.btn-primary.dropdown-toggle.focus{color:#fff;background-color:#19289a;border-color:rgba(0,0,0,0)}.btn-primary.disabled:hover,.btn-primary.disabled:focus,.btn-primary.disabled.focus,.btn-primary[disabled]:hover,.btn-primary[disabled]:focus,.btn-primary.focus[disabled],fieldset[disabled] .btn-primary:hover,fieldset[disabled] .btn-primary:focus,fieldset[disabled] .btn-primary.focus{background-color:#2C41DD;border-color:rgba(0,0,0,0)}.btn-primary .badge{color:#2C41DD;background-color:#fff}a{color:#2C41DD}a:hover,a:focus,a:active{color:#2136ce}.text-primary,a.text-muted:hover{color:#2C41DD}.bg-primary,.bg-confetti-primary,.user-course-pager .btn:hover{background:#2C41DD}.product-card-image.bg-primary:hover{background-color:#2136ce}.navbar-default .navbar-nav>li>a:hover,.navbar-default .navbar-nav>li>a:focus,.navbar-default .navbar-nav>li>a:active,.navbar-default .navbar-nav>.active>a,.navbar-default .navbar-nav>.active>a:hover,.navbar-default .navbar-nav>.active>a:focus,.navbar-default .navbar-nav>.open>a,.navbar-default .navbar-nav>.open>a:hover,.navbar-default .navbar-nav>.open>a:focus,.nav-tabs>li>a:hover{color:#2C41DD}#user-site-course-sidebar .block-link a.active,#user-site-course-sidebar .block-link a:hover{border-color:#2C41DD}.form-control.focused,.form-control:focus{border-color:#2C41DD}.nav-tabs>li>a:hover,.nav-tabs>li>a:active,.nav-tabs>li>a:focus{color:#2C41DD}.nav-tabs>li.active>a,.nav-tabs>li.active>a:hover,.nav-tabs>li.active>a:active,.nav-tabs>li.active>a:focus{border-color:#2C41DD}svg.stroke-primary \*{stroke:#2C41DD}svg.fill-primary \*{fill:#2C41DD}.comment.creator-comment img{box-shadow:0 0 0 2px white,0 0 0 4px #2C41DD}code{color:#1e30b8}.quiz-question .quiz-answer-list .quiz-answer-list-item.unanswered.chosen:before{background:#2C41DD}.quiz-question .quiz-answer-list .quiz-answer-list-item.unanswered.chosen:before,.quiz-question .quiz-answer-list .quiz-answer-list-item.unanswered.unanswered:hover:before{border-color:#2C41DD}



 :root {
 font-size: 18px
 }
 





 window.dataLayer = window.dataLayer || [];
 function gtag(){dataLayer.push(arguments);}
 gtag('js', new Date());
 







 Podia.Customer = null;























 window.Conversation = {
 connectOnLoad: 'false',
 creator: '{\"first\_name\":\"Matt\",\"name\":\"Matt Boyle\",\"imageUrl\":\"https://www.gravatar.com/avatar/a419889cc76057c29071281ab30caf22/?default=https%3A%2F%2Fd228am55mqbj0t.cloudfront.net%2Fdefaults%2Fpurple-MB.png\\u0026size=64\",\"availability\":\"away\",\"messagingEnabled\":true}',
 customDomain: 'true',
 isVisitor: 'true',
 messagingEnabled: 'true',
 locales: '{\"connectivity\":{\"error\":\"Something\'s not right. We\'ll be right back.\",\"success\":\"Connecting...\"},\"email\_capture\":{\"get\_notified\":\"Get notified for responses\",\"notify\_new\_messages\":\"We will notify you of new messages.\",\"subscribed\_message\":\"You\'ll be subscribed to future unread message notifications.\",\"your\_email\":\"Your email\",\"you\_got\_mail\":\"You\'ve got mail!\"},\"empty\_container\":{\"subtitle\":\"Start messaging below\",\"title\":\"No messages yet\"},\"input\_placeholder\":\"Message %{name}\",\"send\_message\_error\":\"Could not send.\",\"user\_status\":{\"away\":\"Away\",\"online\":\"Online\",\"unknown\":\"Unknown\"}}'
 }
 window.clickToastEvent = document.createEvent('Event');
 window.clickToastEvent.initEvent('podia.messagingToastClick', true, true);




 
\_linkedin\_partner\_id = "5441602";
window.\_linkedin\_data\_partner\_ids = window.\_linkedin\_data\_partner\_ids || [];
window.\_linkedin\_data\_partner\_ids.push(\_linkedin\_partner\_id);


(function(l) {
if (!l){window.lintrk = function(a,b){window.lintrk.q.push([a,b])};
window.lintrk.q=[]}
var s = document.getElementsByTagName("script")[0];
var b = document.createElement("script");
b.type = "text/javascript";b.async = true;
b.src = "https://snap.licdn.com/li.lms-analytics/insight.min.js";
s.parentNode.insertBefore(b, s);})(window.lintrk);


![](https://px.ads.linkedin.com/collect/?pid=5441602&fmt=gif)




