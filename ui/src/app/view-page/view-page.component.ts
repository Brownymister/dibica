import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-view-page',
  templateUrl: './view-page.component.html',
  styleUrls: ['./view-page.component.css']
})
export class ViewPageComponent implements OnInit {

  imgUrl: string = "";
  id: string = "";
  message: any = "";
  createDate: string = "";

  constructor() { }

  async loadData() {
    this.id = this.getIdParam();
    this.imgUrl = `/cardsimg/${this.id}.png`
    this.message = await this.getMessage();
    this.createDate = await this.getCreateDate();
    // Get year, month, and day part from the date

    if (this.message != "") { 
      this.AddToLocalStore();
    }
  }

  async getMessage(): Promise<string> {
    const response = await fetch(`/cards/message?id=${this.id}`, {
      method: 'GET', 
      mode: 'cors', 
      cache: 'no-cache',
      credentials: 'same-origin',
      headers: {
      'Content-Type': 'application/json'
      },
      redirect: 'follow', 
      referrerPolicy: 'no-referrer',
    });
    var json = await response.json();
    return json.message;
  }

  async getCreateDate():Promise<string> {
    const response = await fetch(`/cards/createDate?id=${this.id}`, {
      method: 'GET', 
      mode: 'cors', 
      cache: 'no-cache',
      credentials: 'same-origin',
      headers: {
      'Content-Type': 'application/json'
      },
      redirect: 'follow', 
      referrerPolicy: 'no-referrer',
    });
    var json = await response.json();
    return json.date;
  }

  getIdParam(): string {
    // @ts-ignore
    let params = (new URL(document.location)).searchParams;
    let id = params.get('id') || "";
    return id
  }

  AddToLocalStore() {
    const card = {
      id: this.id,
      imgUrl: this.imgUrl,
      message: this.message,
      createDate: this.createDate,
    }
    
    if (window.localStorage.getItem("collection") != null) {
      var allreadyadded = this.AllReadAddedToCollection()
      console.log("all "+ allreadyadded)
      if (!allreadyadded) {
        console.log(window.localStorage.getItem("collection")+","+JSON.stringify(card))
        window.localStorage.setItem("collection",window.localStorage.getItem("collection")+","+JSON.stringify(card));
      }
    }else {
      console.log("2")
      window.localStorage.setItem("collection",JSON.stringify(card));
    }
  }

  AllReadAddedToCollection(): boolean {
    // @ts-ignore
    var json = JSON.parse("["+window.localStorage.getItem("collection")+"]")
    var found: boolean = false;
    // @ts-ignore
    json.forEach(element => {
      console.log(element.id == this.id)
      if(element.id == this.id) {
        found = true;
      }
    });
    return found
  }

  ngOnInit(): void {
    this.loadData()
  }

}
