import { inject, Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

export interface Good {
  id: number;
  name: string;
  description: string;
  volume: number;
}

@Injectable( {providedIn: 'root'} )
export class GoodsService {
  private http = inject(HttpClient);
  private url = 'http://localhost:8080'


  createGoods(name: string, description: string, volume: number) {
    return this.http.post(`${this.url}/goods/create`, {
      name,
      description,
      volume
    });
  }

  getAllGoods(): Observable<Good[]> {
    return this.http.get<Good[]>(`${this.url}/goods/get_all`);
  }

  createSection(volume: number) {
    return this.http.post(`${this.url}/sections/create`, {
      volume
    })
  }

  createContractor(name: string, inn: string, type: string) {
    return this.http.post(`${this.url}/contractors/create`, {
      name,
      inn,
      type
    })
  }
}
