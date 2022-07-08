"use strict";(self.webpackChunktrainly_docs=self.webpackChunktrainly_docs||[]).push([[155],{9613:function(e,t,n){n.d(t,{Zo:function(){return d},kt:function(){return m}});var r=n(9496);function o(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function a(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function i(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?a(Object(n),!0).forEach((function(t){o(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):a(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function c(e,t){if(null==e)return{};var n,r,o=function(e,t){if(null==e)return{};var n,r,o={},a=Object.keys(e);for(r=0;r<a.length;r++)n=a[r],t.indexOf(n)>=0||(o[n]=e[n]);return o}(e,t);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);for(r=0;r<a.length;r++)n=a[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(o[n]=e[n])}return o}var s=r.createContext({}),l=function(e){var t=r.useContext(s),n=t;return e&&(n="function"==typeof e?e(t):i(i({},t),e)),n},d=function(e){var t=l(e.components);return r.createElement(s.Provider,{value:t},e.children)},u={inlineCode:"code",wrapper:function(e){var t=e.children;return r.createElement(r.Fragment,{},t)}},p=r.forwardRef((function(e,t){var n=e.components,o=e.mdxType,a=e.originalType,s=e.parentName,d=c(e,["components","mdxType","originalType","parentName"]),p=l(n),m=o,h=p["".concat(s,".").concat(m)]||p[m]||u[m]||a;return n?r.createElement(h,i(i({ref:t},d),{},{components:n})):r.createElement(h,i({ref:t},d))}));function m(e,t){var n=arguments,o=t&&t.mdxType;if("string"==typeof e||o){var a=n.length,i=new Array(a);i[0]=p;var c={};for(var s in t)hasOwnProperty.call(t,s)&&(c[s]=t[s]);c.originalType=e,c.mdxType="string"==typeof e?e:o,i[1]=c;for(var l=2;l<a;l++)i[l]=n[l];return r.createElement.apply(null,i)}return r.createElement.apply(null,n)}p.displayName="MDXCreateElement"},2110:function(e,t,n){n.r(t),n.d(t,{assets:function(){return d},contentTitle:function(){return s},default:function(){return m},frontMatter:function(){return c},metadata:function(){return l},toc:function(){return u}});var r=n(2848),o=n(9213),a=(n(9496),n(9613)),i=["components"],c={id:"Project_Firmware",title:"Firmware / Embedded"},s=void 0,l={unversionedId:"Project/Project_Firmware",id:"Project/Project_Firmware",title:"Firmware / Embedded",description:"Collect and sending data",source:"@site/docs/Project/Firmware.md",sourceDirName:"Project",slug:"/Project/Project_Firmware",permalink:"/ProjectWork-DIQU-Group1-2022/docs/Project/Project_Firmware",draft:!1,editUrl:"https://github.com/facebook/docusaurus/tree/main/packages/create-docusaurus/templates/shared/docs/Project/Firmware.md",tags:[],version:"current",frontMatter:{id:"Project_Firmware",title:"Firmware / Embedded"},sidebar:"tutorialSidebar",previous:{title:"Group Members",permalink:"/ProjectWork-DIQU-Group1-2022/docs/Introduction/Group_Members"},next:{title:"Gateway",permalink:"/ProjectWork-DIQU-Group1-2022/docs/Project/Project_Gateway"}},d={},u=[{value:"Collect and sending data",id:"collect-and-sending-data",level:2},{value:"Gallery",id:"gallery",level:2},{value:"Behind the scene",id:"behind-the-scene",level:2}],p={toc:u};function m(e){var t=e.components,c=(0,o.Z)(e,i);return(0,a.kt)("wrapper",(0,r.Z)({},p,c,{components:t,mdxType:"MDXLayout"}),(0,a.kt)("h2",{id:"collect-and-sending-data"},"Collect and sending data"),(0,a.kt)("p",null,"We use PicSimLab to simulate the electronic board that due to technical problems we could not get working."),(0,a.kt)("h2",{id:"gallery"},"Gallery"),(0,a.kt)("p",null,"The following picture shows the main screen of the LCD display."),(0,a.kt)("p",null,(0,a.kt)("img",{src:n(5426).Z,width:"1272",height:"826"})),(0,a.kt)("p",null,"RB3 button opens the doors."),(0,a.kt)("p",null,(0,a.kt)("img",{src:n(3799).Z,width:"1262",height:"825"})),(0,a.kt)("p",null,"RB4 button closes the doors."),(0,a.kt)("p",null,(0,a.kt)("img",{src:n(2929).Z,width:"1266",height:"816"})),(0,a.kt)("p",null,"RB5 button sets the bathroom as busy."),(0,a.kt)("p",null,(0,a.kt)("img",{src:n(2650).Z,width:"1268",height:"817"})),(0,a.kt)("h2",{id:"behind-the-scene"},"Behind the scene"),(0,a.kt)("p",null,"In this section we are going to give a brief explanation about the used code."),(0,a.kt)("p",null,"This function has the aim to read the data (like temperature) and return a value between 0 and 1024."),(0,a.kt)("p",null,"The first step is to set the analog-channel-corresponding-bit register.\nBy doing this, we are going to be able to read analog data throught the analog channel AN2."),(0,a.kt)("p",null,"ADCON0 register is used to manage the conversion process. First we select the channel, then wait 20ms for the\ncapacitor to charge, set the GO bit to start reading the data, wait until it's resetted (conversion ended) and\nreturn the data from the function."),(0,a.kt)("p",null,(0,a.kt)("img",{src:n(2738).Z,width:"927",height:"303"})),(0,a.kt)("p",null,"This function converts int values to string values. Used to display data on LCD display."),(0,a.kt)("p",null,(0,a.kt)("img",{src:n(3453).Z,width:"447",height:"175"})),(0,a.kt)("p",null,"Send data or command to LCD display."),(0,a.kt)("p",null,(0,a.kt)("img",{src:n(7500).Z,width:"457",height:"383"})),(0,a.kt)("p",null,"Initialization of the LCD display."),(0,a.kt)("p",null,(0,a.kt)("img",{src:n(2298).Z,width:"388",height:"390"})))}m.isMDXComponent=!0},5426:function(e,t,n){t.Z=n.p+"assets/images/1Embedded-5c6b1e6320509bcb8a56567fb78207c8.jpg"},3799:function(e,t,n){t.Z=n.p+"assets/images/2Embedded-3b30002085419816f2822ee2348116b2.jpg"},2929:function(e,t,n){t.Z=n.p+"assets/images/3Emdedded-b4416ae4deed33213c63044701b28f72.jpg"},2650:function(e,t,n){t.Z=n.p+"assets/images/4Embedded-60653f9cb1e33ad25924e5f0f5728b2a.jpg"},3453:function(e,t,n){t.Z=n.p+"assets/images/ConversioneIntString-67c536c87c4367b3c33b829acd0d8503.jpg"},2298:function(e,t,n){t.Z=n.p+"assets/images/InizializzazioneLCD-c84004003f28710bea438e0be46cc77b.jpg"},7500:function(e,t,n){t.Z=n.p+"assets/images/InvioDatoComandoLCD-08a4f0c1d23d97f0901e98c1d36ed095.jpg"},2738:function(e,t,n){t.Z=n.p+"assets/images/letturaDatoAnalogico-307b40b3407d7b9120b6c54acdeea84e.jpg"}}]);