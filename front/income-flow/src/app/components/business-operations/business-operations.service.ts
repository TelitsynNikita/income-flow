import {inject, Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {map, Observable} from 'rxjs';

interface Contractor {
  name: string;
  inn: string;
  type: string;
}

export interface BusinessOperation {
  id: number;
  date: string;
  operation_type: string;
  operation_id: number;
  goods_count: number;
  goods_id: number;
  section_id: number;
  contractors_id: number;
  contract: Contractor;
}

export interface BusinessOperation2 {
  id: number;
  date: string;
  operation_type: string;
  operation_id: number;
  goods_count: number;
  goods_name: number;
  contractor_name: string;
}

interface Result {
  res: BusinessOperation2[]
}

@Injectable({providedIn: 'root'})
export class BusinessOperationsService {
  private http = inject(HttpClient);
  private url = 'http://localhost:8080'
  private url8081 = 'http://localhost:8081'


  getOperations(): Observable<BusinessOperation[]> {
    return this.http.get<BusinessOperation[]>(`${this.url}/operations/get_business_operations`);
  }

  getOperations2(): Observable<Map<string, BusinessOperation2[]>> {
    return this.http.get<Result>(`${this.url8081}/operations/get-operation`).pipe(
      map(item => {
        return item.res
      }),
      map(res => {
        let uniqRes = new Map<number, BusinessOperation2>();


        for (let item of res) {
          uniqRes.set(item.operation_id, item)
        }

        let uniqBody = new Map<string, BusinessOperation2[]>();

        uniqRes.forEach((value) => {
          let data = uniqBody.get(new Date(value.date).toLocaleDateString());

          if (!data) {
            uniqBody.set(new Date(value.date).toLocaleDateString(), [{
              ...value,
              date: new Date(value.date).toLocaleTimeString(),
            }])
          } else {
            uniqBody.set(new Date(value.date).toLocaleDateString(), [...data, {
              ...value,
              date: new Date(value.date).toLocaleTimeString(),
            }])
          }
        })

        return uniqBody;
      })
    );
  }
}
