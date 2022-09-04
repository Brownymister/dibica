import { Component, OnInit } from '@angular/core';

interface Alert {
  type: string;
  message: string;
}

@Component({
  selector: 'app-create-page',
  templateUrl: './create-page.component.html',
  styleUrls: ['./create-page.component.css'],
})
export class CreatePageComponent implements OnInit {
  
  constructor() {}
  
  id: string ="";
  imageURl: string = "";
  error: string = "";
  displayError: boolean = false;
  LinkAlertMessage: string = "";
  LinkAlerType: string = "dark";
  
  SetMessage(msg: string) {
    // @ts-ignore
    document.getElementById("message").value = msg;
  }

  copyToCb() {
    navigator.clipboard.writeText("https://dibica.herokuapp.com/card?id="+this.id);
  }

  async SendCardInformation() {
    this.error = "";
    this.displayError = false;
    const response = await fetch("/cards/create", {
      method: 'POST', 
      mode: 'cors', 
      cache: 'no-cache',
      credentials: 'same-origin',
      headers: {
      'Content-Type': 'application/json'
      },
      redirect: 'follow', 
      referrerPolicy: 'no-referrer',
      body: JSON.stringify({
          // @ts-ignore
          "name":document.getElementById("firstName").value,
          // @ts-ignore
          "message":document.getElementById("message").value,
          "template":"bd1",
      })
    });

    if (response.status == 200) {
      var json = await response.json()
      // @ts-ignore
      console.log(json);
      this.imageURl = json.link;
      // @ts-ignore
      document.getElementById("previewImg").src = json.link;
      // @ts-ignore
      document.getElementById("clipboard").style.display = "block";
      console.log("https://dibica.herokuapp.com"+this.imageURl);
      var id = this.imageURl.replace(".png", "");
      this.id = id.replace("/cardsimg/", "");
      // @ts-ignore
      document.getElementById("in01").value = "https://dibica.herokuapp.com/card?id="+this.id;
      // @ts-ignore
      this.LinkAlertMessage = `Card to ${document.getElementById("firstName").value}`;
    } else {
      this.displayError = true;
      this.error = "Name and Message needed!"
    }
  }

  ngOnInit(): void {
  }

}
