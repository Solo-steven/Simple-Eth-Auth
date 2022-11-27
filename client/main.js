let state = "initial";
const textBlock = document.getElementById("state");
const button = document.getElementById("button");

const handleOnClickButton = async () => {
    if(state === "pending") {
        return;
    }
    if(!window.ethereum) {
        textBlock.textContent = "No Provider";
        return;
    }
    try {
        const accounts = await window.ethereum.request({method: "eth_requestAccounts"});
        const nonce = "Example `personal_sign` message";
        const _signNonce = await window.ethereum.request({method: "personal_sign", params: [nonce, accounts[0]]});
        console.log("Account Name:",accounts);
        console.log("Message:", nonce);
        console.log("Singed Data", _signNonce);
        state = "fullfill";
        textBlock.textContent = "Ok";
        button.disabled = true;
    } catch(e) {
        state = "rejected";
        textBlock.textContent = "Fail";
    }
}


function main() {
    textBlock.textContent = "Weclome";
    button.addEventListener('click', handleOnClickButton);
}
main();
