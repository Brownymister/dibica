import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-collection',
  templateUrl: './collection.component.html',
  styleUrls: ['./collection.component.css']
})
export class CollectionComponent implements OnInit {

  constructor() { }

  collection: any = [{"id":"4dd03b89-56d6-4841-8102-58d58a4560e7","imgUrl":"/cardsimg/4dd03b89-56d6-4841-8102-58d58a4560e7.png","message":"Happy Birthday!"}];

  readLocalstorage() {
    var json = JSON.parse("["+window.localStorage.getItem("collection")+"]")
    console.log(json)
    this.collection = json;
  }

  open(id:string) {
    window.open("/card?id="+id)
  } 

  ngOnInit(): void {
    this.readLocalstorage();
  }

}
