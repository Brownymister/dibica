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

  constructor() { }

  async loadData() {
    this.id = this.getIdParam();
    this.imgUrl = `/cardsimg/${this.id}.png`
    this.message = await this.getMessage();
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
    console.log(json)
    return json.message;
  }

  getIdParam(): string {
    // @ts-ignore
    let params = (new URL(document.location)).searchParams;
    let id = params.get('id') || "";
    return id
  }

  ngOnInit(): void {
    this.loadData()
  }

}
