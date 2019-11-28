import pandas as pd 

df = pd.read_excel('/Users/zhangxin/Desktop/ppdai/2drools/2drools/2drools_优选贷新客_zcx.xlsx')



class Data:
    df=df

# print(df)
# 初始化columns
def reanme_df():
    dic={}
    for i in df.columns:
        dic[i] = str(int(i.split(':')[1].strip(' '))+1)
    Data.df=df.rename(columns=dic)
    print(Data.df)

def write_mysql(df):
     for x in df.index:
        for iy, y in enumerate(df.columns):
            print('===========', x+1, iy+1, df.loc[x,y])

def all_due():
    df = Data.df.loc[3:22, [ '3',  '4' , '5','6','7']].reset_index()
    del df['index']
    print(df)
    write_mysql(df)


def due_3():
    df = Data.df.loc[27:47, [ '3',  '4' , '5','6','7']].reset_index()
    del df['index']
    print(df)
    write_mysql(df)

def due_6():   
    df = Data.df.loc[27:47, ['11',  '12' , '13','14','15']].reset_index()
    del df['index']
    print(df)
    write_mysql(df)

def due_12():   
    df = Data.df.loc[27:47, [ '19' , '20','21','22','23', ]].reset_index()
    del df['index']
    print(df)
    write_mysql(df)
if __name__ == "__main__":
    reanme_df()
    all_due()
    # due_6()
    # reanme_df()